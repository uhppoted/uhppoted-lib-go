package integration_tests

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"go/printer"
	"go/token"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"

	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
	"codegen/model"
)

func udpAST() {
	outfile := filepath.Join(".", "udp", "api_test.go")
	decl := buildUDP()

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

func buildUDP() *dst.File {
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
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted"`,
				},

				Decs: dst.ImportSpecDecorations{
					NodeDecs: dst.NodeDecs{
						Before: dst.EmptyLine,
					},
				},
			},
			&dst.ImportSpec{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"`,
				},
			},

			&dst.ImportSpec{
				Name: dst.NewIdent("test"),
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"integration-tests"`,
				},

				Decs: dst.ImportSpecDecorations{
					NodeDecs: dst.NodeDecs{
						Before: dst.EmptyLine,
					},
				},
			},
		},
	}

	tests := []dst.Decl{}

	for _, fn := range model.API {
		if fn.Name == "find-controllers" {
			continue
		}

		for _, test := range fn.Tests {
			if test := buildUDPTestFunc(*fn, test); test != nil {
				tests = append(tests, test)
			}
		}
	}

	decls := []dst.Decl{}
	decls = append(decls, imports)
	decls = append(decls, tests...)

	file := &dst.File{
		Name:  dst.NewIdent("uhppoted"),
		Decls: decls,
	}

	return file
}

func buildUDPTestFunc(fn lib.Function, test lib.FuncTest) *dst.FuncDecl {
	name := fmt.Sprintf("Test%v", codegen.TitleCase(test.Name))

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
		Body: buildUDPTestImpl(fn, test),

		Decs: dst.FuncDeclDecorations{
			NodeDecs: dst.NodeDecs{
				After: dst.EmptyLine,
			},
		},
	}

	return f
}

func buildUDPTestImpl(fn lib.Function, test lib.FuncTest) *dst.BlockStmt {
	block := &dst.BlockStmt{
		List: []dst.Stmt{},
	}

	block.List = append(block.List, buildAPITestExpected(fn, test))
	block.List = append(block.List, buildUDPTestArgs(fn, test)...)
	block.List = append(block.List, buildAPITestExec(fn, test))
	block.List = append(block.List, buildAPITestValidate(fn, test))

	return block
}

func buildUDPTestArgs(fn lib.Function, test lib.FuncTest) []dst.Stmt {
	args := []dst.Stmt{}

	for _, arg := range test.Args {
		name := codegen.CamelCase(arg.Name)

		if name == "controller" {
			args = append(args, &dst.AssignStmt{
				Lhs: []dst.Expr{
					dst.NewIdent(name),
				},
				Tok: token.DEFINE,
				Rhs: []dst.Expr{
					buildUDPTestController(arg),
				},
			})
		} else {
			args = append(args, &dst.AssignStmt{
				Lhs: []dst.Expr{
					dst.NewIdent(name),
				},
				Tok: token.DEFINE,
				Rhs: []dst.Expr{
					buildAPITestArg(arg),
				},
			})
		}
	}

	return args
}

func buildUDPTestController(arg lib.Arg) dst.Expr {
	controller := dst.KeyValueExpr{
		Key: &dst.Ident{
			Name: "ID",
		},
		Value: &dst.BasicLit{
			Kind:  token.INT,
			Value: fmt.Sprintf(`%v`, arg.Value),
		},
		Decs: dst.KeyValueExprDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.NewLine,
				After:  dst.NewLine,
			},
		},
	}

	address := dst.KeyValueExpr{
		Key: &dst.Ident{
			Name: "Address",
		},
		Value: &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "netip"},
				Sel: &dst.Ident{Name: "MustParseAddrPort"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: `"127.0.0.1:50002"`,
				},
			},
		},
		Decs: dst.KeyValueExprDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.NewLine,
				After:  dst.NewLine,
			},
		},
	}

	protocol := dst.KeyValueExpr{
		Key: &dst.Ident{
			Name: "Protocol",
		},
		Value: &dst.BasicLit{
			Kind:  token.STRING,
			Value: `"udp"`,
		},
		Decs: dst.KeyValueExprDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.NewLine,
				After:  dst.NewLine,
			},
		},
	}

	return &dst.CompositeLit{
		Type: &dst.SelectorExpr{
			X:   &dst.Ident{Name: "uhppoted"},
			Sel: &dst.Ident{Name: "Controller"},
		},
		Elts: []dst.Expr{
			&controller,
			&address,
			&protocol,
		},
		Decs: dst.CompositeLitDecorations{
			NodeDecs: dst.NodeDecs{
				After: dst.EmptyLine,
			},
		},
	}
}
