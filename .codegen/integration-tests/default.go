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

func broadcastAST() {
	outfile := filepath.Join(".", "default", "_api_test.go")
	decl := buildBroadcast()

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

func buildBroadcast() *dst.File {
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
					Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted"`,
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
					Kind: token.STRING,
				},
			},

			&dst.ImportSpec{
				Name: dst.NewIdent("tests"),
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/integration-tests"`,
				},
			},
		},
	}

	tests := []dst.Decl{}

	for _, fn := range model.API {
		for _, test := range fn.Tests {
			if test := buildBroadcastTestFunc(*fn, test); test != nil {
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

func buildBroadcastTestFunc(fn lib.Function, test lib.FuncTest) *dst.FuncDecl {
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
		Body: buildBroadcastTestImpl(fn, test),

		Decs: dst.FuncDeclDecorations{
			NodeDecs: dst.NodeDecs{
				After: dst.EmptyLine,
			},
		},
	}

	return f
}

func buildBroadcastTestImpl(fn lib.Function, test lib.FuncTest) *dst.BlockStmt {
	if fn.Name == "find-controllers" {
		return &dst.BlockStmt{
			List: []dst.Stmt{
				buildBroadcastTestExpected(fn, test),
				buildBroadcastTestValidate(fn, test),
			},
		}
	} else {
		return &dst.BlockStmt{
			List: []dst.Stmt{
				buildBroadcastTestController(fn, test),
				buildBroadcastTestExpected(fn, test),
				buildBroadcastTestValidate(fn, test),
			},
		}

	}
}

func buildBroadcastTestController(fn lib.Function, test lib.FuncTest) dst.Stmt {
	controller := dst.AssignStmt{
		Lhs: []dst.Expr{dst.NewIdent("controller")},
		Tok: token.DEFINE,
		Rhs: []dst.Expr{
			&dst.CallExpr{
				Fun: &dst.Ident{Name: "uint32"},
				Args: []dst.Expr{
					&dst.BasicLit{Kind: token.INT, Value: "405419896"}},
			},
		},

		Decs: dst.AssignStmtDecorations{
			NodeDecs: dst.NodeDecs{After: dst.EmptyLine},
		},
	}

	return &controller
}

func buildBroadcastTestExpected(fn lib.Function, test lib.FuncTest) dst.Stmt {
	name := codegen.TitleCase(test.Name)

	return &dst.AssignStmt{
		Lhs: []dst.Expr{dst.NewIdent("expected")},
		Tok: token.DEFINE,
		Rhs: []dst.Expr{
			&dst.SelectorExpr{
				X: &dst.SelectorExpr{
					X:   &dst.Ident{Name: "test"},
					Sel: &dst.Ident{Name: "Expected"},
				},

				Sel: &dst.Ident{Name: name},
			},
		},

		Decs: dst.AssignStmtDecorations{
			NodeDecs: dst.NodeDecs{After: dst.EmptyLine},
		},
	}
}

func buildBroadcastTestValidate(fn lib.Function, test lib.FuncTest) dst.Stmt {
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
									Value: `"incorrect response:\n   expected: %#v\n   got:      %#v"`,
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
