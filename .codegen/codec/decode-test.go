package codec

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"go/printer"
	"go/token"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"

	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
	"codegen/model"
)

func decodeTest() {
	outfile := filepath.Join("codec", "decode", "generated_test.go")
	decl := buildDecodeTest()

	// .. convert dst to ast
	fset, file, err := decorator.RestoreFile(decl)
	if err != nil {
		log.Fatalf("error converting dst to ast (%v)", err)
	}

	// ... pretty print
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, file); err != nil {
		log.Fatalf("error pretty-printing generated code (%v)", err)
	}

	// ... header comment
	if f, err := os.Create(outfile); err != nil {
		log.Fatalf("error creating file %s (%v)", outfile, err)
	} else {
		defer f.Close()

		writeln(f, "// generated code - ** DO NOT EDIT **")
		writeln(f, "")
		writeln(f, buf.String())

		f.Close()
	}
}

func buildDecodeTest() *dst.File {
	imports := &dst.GenDecl{
		Tok: token.IMPORT,
		Specs: []dst.Spec{
			&dst.ImportSpec{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"net/netip"`,
				},
			},

			&dst.ImportSpec{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"reflect"`,
				},
			},

			&dst.ImportSpec{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"testing"`,
				},
			},
			&dst.ImportSpec{
				Path: &dst.BasicLit{
					Kind: token.STRING,
				},
			},
			&dst.ImportSpec{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"`,
				},
			},
			&dst.ImportSpec{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"`,
				},
			},
		},
	}

	tests := []dst.Decl{}

	for _, response := range model.Responses {
		for _, test := range response.Tests {
			tests = append(tests, buildDecodeTestFunc(*response, test))
		}
	}

	for _, test := range model.ListenerEvent.Tests {
		tests = append(tests, buildDecodeTestFunc(model.ListenerEvent, test))
	}

	file := &dst.File{
		Name:  dst.NewIdent("decode"),
		Decls: append([]dst.Decl{imports}, tests...),
	}

	return file
}

func buildDecodeTestFunc(response lib.Response, test lib.ResponseTest) *dst.FuncDecl {
	name := fmt.Sprintf("Test%vResponse", codegen.TitleCase(test.Name))

	f := &dst.FuncDecl{
		Name: dst.NewIdent(name),
		Type: &dst.FuncType{
			Params: &dst.FieldList{
				List: []*dst.Field{
					{
						Names: []*dst.Ident{dst.NewIdent("t")},
						Type: &dst.StarExpr{
							X: &dst.SelectorExpr{
								X:   &dst.Ident{Name: "testing"},
								Sel: &dst.Ident{Name: "T"},
							},
						},
					},
				},
			},
		},
		Body: buildDecodeTestImpl(response, test),
	}

	f.Decs.After = dst.EmptyLine

	return f
}

func buildDecodeTestImpl(response lib.Response, test lib.ResponseTest) *dst.BlockStmt {
	packet := make([]dst.Expr, 64)
	for i, b := range test.Response {
		xx := &dst.BasicLit{
			Kind:  token.INT,
			Value: fmt.Sprintf("0x%02x", b),
		}

		if i%16 == 0 {
			xx.Decs.Before = dst.NewLine
		}

		if i == 63 {
			xx.Decs.After = dst.NewLine
		}

		packet[i] = xx
	}

	return &dst.BlockStmt{
		List: []dst.Stmt{
			// packet := []byte{...}
			&dst.AssignStmt{
				Lhs: []dst.Expr{
					dst.NewIdent("packet"),
				},
				Tok: token.DEFINE,
				Rhs: []dst.Expr{
					&dst.CompositeLit{
						Type: &dst.ArrayType{
							Len: nil,
							Elt: dst.NewIdent("byte"),
						},
						Elts: packet,
					},
				},

				Decs: dst.AssignStmtDecorations{
					NodeDecs: dst.NodeDecs{
						After: dst.EmptyLine,
					},
				},
			},

			// expected := ...
			&dst.AssignStmt{
				Lhs: []dst.Expr{
					&dst.Ident{Name: "expected"},
				},
				Tok: token.DEFINE,
				Rhs: []dst.Expr{
					buildDecodeTestExpected(response, test.Expected),
				},
			},

			buildDecodeTestExec(response),
			buildDecodeTestValidate(response),
		},
	}
}

func buildDecodeTestExpected(response lib.Response, values []lib.Value) dst.Expr {
	name := strings.TrimSuffix(codegen.TitleCase(response.Name), "Response")
	fields := []dst.Expr{}

	for _, field := range response.Fields {
		name := codegen.TitleCase(field.Name)

		for _, v := range values {
			if v.Name == field.Name {
				value := makeValue(field, v)

				f := dst.KeyValueExpr{
					Key:   &dst.Ident{Name: name},
					Value: value,
				}

				f.Decs.Before = dst.NewLine
				f.Decs.After = dst.NewLine

				fields = append(fields, &f)
			}
		}
	}

	composite := &dst.CompositeLit{
		Type: &dst.SelectorExpr{
			X:   &dst.Ident{Name: "responses"},
			Sel: &dst.Ident{Name: name},
		},
		Elts: fields,

		Decs: dst.CompositeLitDecorations{
			NodeDecs: dst.NodeDecs{
				After: dst.EmptyLine,
			},
		},
	}

	return composite
}

func buildDecodeTestExec(response lib.Response) dst.Stmt {
	name := codegen.TitleCase(response.Name)

	// response, err := XXXResponse(packet)
	assign := dst.AssignStmt{
		Lhs: []dst.Expr{
			&dst.Ident{Name: "response"},
			&dst.Ident{Name: "err"},
		},
		Tok: token.DEFINE,
		Rhs: []dst.Expr{
			&dst.CallExpr{
				Fun: &dst.Ident{Name: name},
				Args: []dst.Expr{
					&dst.Ident{Name: "packet"},
				},
			},
		},

		Decs: dst.AssignStmtDecorations{
			NodeDecs: dst.NodeDecs{
				After: dst.EmptyLine,
			},
		},
	}

	return &assign
}

func buildDecodeTestValidate(response lib.Response) dst.Stmt {
	return &dst.IfStmt{
		Cond: &dst.BinaryExpr{
			X:  &dst.Ident{Name: "err"},
			Op: token.NEQ,
			Y:  &dst.Ident{Name: "nil"},
		},
		Body: &dst.BlockStmt{
			List: []dst.Stmt{
				&dst.ExprStmt{
					X: &dst.CallExpr{
						Fun: &dst.SelectorExpr{
							X:   &dst.Ident{Name: "t"},
							Sel: &dst.Ident{Name: "Fatalf"},
						},
						Args: []dst.Expr{
							&dst.BasicLit{Kind: token.STRING, Value: `"%v"`},
							&dst.Ident{Name: "err"},
						},
					},
				},
			},
		},
		Else: &dst.IfStmt{
			Cond: &dst.UnaryExpr{
				Op: token.NOT,
				X: &dst.CallExpr{
					Fun: &dst.SelectorExpr{
						X:   &dst.Ident{Name: "reflect"},
						Sel: &dst.Ident{Name: "DeepEqual"},
					},
					Args: []dst.Expr{
						&dst.Ident{Name: "response"},
						&dst.Ident{Name: "expected"},
					},
				},
			},
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ExprStmt{
						X: &dst.CallExpr{
							Fun: &dst.SelectorExpr{
								X:   &dst.Ident{Name: "t"},
								Sel: &dst.Ident{Name: "Errorf"},
							},
							Args: []dst.Expr{
								&dst.BasicLit{
									Kind:  token.STRING,
									Value: `"incorrectly decoded response:\n   expected: %#v\n   got:      %#v"`,
								},
								&dst.Ident{Name: "expected"},
								&dst.Ident{Name: "response"},
							},
						},
					},
				},
			},
		},
	}
}
