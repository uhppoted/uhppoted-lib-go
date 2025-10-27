package codec

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"go/printer"
	"go/token"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"

	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
	"codegen/model"
)

func decoderTest() {
	outfile := filepath.Join("codec", "generated_test.go")
	decl := buildDecoderTest()

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

func buildDecoderTest() *dst.File {
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
					Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"`,
				},
			},
			&dst.ImportSpec{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"`,
				},
			},
		},
	}

	tests := []dst.Decl{}
	excluded := []*lib.Response{}
	responses := model.Responses
	for _, response := range responses {
		if slices.Contains(excluded, response) {
			log.Printf("skipping %v (excluded)", response.Name)
			continue
		}

		for _, test := range response.Tests {
			f := buildDecoderTestFunc(*response, test)

			tests = append(tests, f)
		}
	}

	file := &dst.File{
		Name: dst.NewIdent("codec"),

		Imports: []*dst.ImportSpec{
			{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"net/netip"`,
				},
			},
			{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"reflect"`,
				},
			},
			{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"testing"`,
				},
			},
			{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"`,
				},
			},
		},

		Decls: append([]dst.Decl{imports}, tests...),
	}

	return file
}

func buildDecoderTestFunc(response lib.Response, test lib.ResponseTest) *dst.FuncDecl {
	name := fmt.Sprintf("TestDecode%vResponse", codegen.TitleCase(test.Name))

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
		Body: buildDecoderTestImpl(response, test),
	}

	f.Decs.After = dst.EmptyLine

	return f
}

func buildDecoderTestImpl(response lib.Response, test lib.ResponseTest) *dst.BlockStmt {
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
			},

			// blank line
			&dst.ExprStmt{
				X: &dst.BasicLit{
					Kind: token.STRING,
				},
			},

			// expected := ...
			&dst.AssignStmt{
				Lhs: []dst.Expr{
					&dst.Ident{Name: "expected"},
				},
				Tok: token.DEFINE,
				Rhs: []dst.Expr{
					buildExpected(response, test.Expected),
				},
			},

			// blank line
			&dst.ExprStmt{
				X: &dst.BasicLit{
					Kind: token.STRING,
				},
			},

			// (exec)
			buildExec(response),
		},
	}
}

func buildExpected(response lib.Response, values []lib.Value) dst.Expr {
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
	}

	return composite
}

func buildExec(response lib.Response) dst.Stmt {
	name := codegen.TitleCase(response.Name)

	return &dst.IfStmt{
		Init: &dst.AssignStmt{
			Lhs: []dst.Expr{
				&dst.Ident{Name: "response"},
				&dst.Ident{Name: "err"},
			},
			Tok: token.DEFINE,
			Rhs: []dst.Expr{
				&dst.CallExpr{
					Fun: &dst.IndexExpr{
						X: &dst.Ident{Name: "Decode"},
						Index: &dst.SelectorExpr{
							X:   &dst.Ident{Name: "responses"},
							Sel: &dst.Ident{Name: strings.TrimSuffix(name, "Response")},
						},
					},
					Args: []dst.Expr{&dst.Ident{Name: "packet"}},
				},
			},
		},
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

func makeValue(field lib.Field, value lib.Value) dst.Expr {
	switch field.Type {
	case "uint8":
		return &dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%v", value.Value)}

	case "uint16":
		return &dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%v", value.Value)}

	case "uint32":
		return &dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%v", value.Value)}

	case "bool":
		return &dst.Ident{Name: fmt.Sprintf("%v", value.Value)}

	case "IPv4":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "netip"},
				Sel: &dst.Ident{Name: "MustParseAddr"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "address:port":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "netip"},
				Sel: &dst.Ident{Name: "MustParseAddrPort"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "date":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseDate"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "shortdate":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseDate"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "optional date":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseDate"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "time":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseTime"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "HHmm":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseHHmm"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "datetime":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseDateTime"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "optional datetime":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseDateTime"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "MAC":
		return &dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)}

	case "version":
		return &dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)}

	case "pin":
		return &dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%v", value.Value)}

	case "mode":
		return &dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%v", value.Value)}

	case "event-type":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "EventType"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf(`%v`, value.Value)},
			},
		}

	case "direction":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "Direction"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf(`%v`, value.Value)},
			},
		}

	case "reason":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "Reason"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf(`%v`, value.Value)},
			},
		}

	default:
		panic(fmt.Sprintf("%v", field.Type))
	}
}
