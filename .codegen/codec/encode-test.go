package codec

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

	"codegen/codegen"
	"codegen/model"
	lib "github.com/uhppoted/uhppoted-codegen/model/types"
)

func encodeTest() {
	outfile := filepath.Join("codec", "encode", "generated_test.go")
	decl := buildEncodeTest()

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

func buildEncodeTest() *dst.File {
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
					Value: `"slices"`,
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
		},
	}

	tests := []dst.Decl{}

	for _, request := range model.Requests {
		for _, test := range request.Tests {
			tests = append(tests, buildEncodeTestFunc(request, test))
		}
	}

	file := &dst.File{
		Name:  dst.NewIdent("encode"),
		Decls: append([]dst.Decl{imports}, tests...),
	}

	return file
}

func buildEncodeTestFunc(request lib.Request, test lib.RequestTest) *dst.FuncDecl {
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
		Body: buildEncodeTestImpl(request, test),
	}

	f.Decs.After = dst.EmptyLine

	return f
}

func buildEncodeTestImpl(request lib.Request, test lib.RequestTest) *dst.BlockStmt {
	return &dst.BlockStmt{
		List: []dst.Stmt{
			buildEncodeTestExpected(test),
			buildEncodeTestExec(request, test),
			buildEncodeTestValidate(request),
		},
	}
}

// expected := []byte{...}
func buildEncodeTestExpected(test lib.RequestTest) dst.Stmt {
	packet := make([]dst.Expr, 64)
	for i, b := range test.Expected {
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

	assign := dst.AssignStmt{
		Lhs: []dst.Expr{
			dst.NewIdent("expected"),
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
	}

	return &assign
}

// packet, err := <T>Request(...)
func buildEncodeTestExec(request lib.Request, test lib.RequestTest) dst.Stmt {
	name := codegen.TitleCase(request.Name)

	args := []dst.Expr{}
	for _, arg := range test.Args {
		args = append(args, buildEncodeTestArg(arg))
	}

	return &dst.AssignStmt{
		Lhs: []dst.Expr{
			&dst.Ident{Name: "packet"},
			&dst.Ident{Name: "err"},
		},
		Tok: token.DEFINE,
		Rhs: []dst.Expr{
			&dst.CallExpr{
				Fun:  &dst.Ident{Name: name},
				Args: args,
			},
		},

		Decs: dst.AssignStmtDecorations{
			NodeDecs: dst.NodeDecs{
				After: dst.EmptyLine,
			},
		},
	}
}

func buildEncodeTestArg(arg lib.Arg) dst.Expr {
	switch arg.Type {
	case "bool":
		return &dst.BasicLit{
			Kind:  token.IDENT,
			Value: fmt.Sprintf(`%v`, arg.Value),
		}

	case "uint8":
		return &dst.BasicLit{
			Kind:  token.INT,
			Value: fmt.Sprintf(`%v`, arg.Value),
		}

	case "uint16":
		return &dst.BasicLit{
			Kind:  token.INT,
			Value: fmt.Sprintf(`%v`, arg.Value),
		}

	case "uint32":
		return &dst.BasicLit{
			Kind:  token.INT,
			Value: fmt.Sprintf(`%v`, arg.Value),
		}

	case "IPv4":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "netip.MustParseAddr"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, arg.Value),
				},
			},
		}

	case "address:port":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "netip.MustParseAddrPort"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, arg.Value),
				},
			},
		}

	case "datetime":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "types.MustParseDateTime"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, arg.Value),
				},
			},
		}

	case "date":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "types.MustParseDate"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, arg.Value),
				},
			},
		}

	case "HHmm":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "types.MustParseHHmm"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, arg.Value),
				},
			},
		}

	case "pin":
		return &dst.BasicLit{
			Kind:  token.INT,
			Value: fmt.Sprintf(`%v`, arg.Value),
		}

	case "mode":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "types.DoorMode"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.INT,
					Value: fmt.Sprintf(`%v`, arg.Value),
				},
			},
		}

	case "task":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "types.TaskType"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.INT,
					Value: fmt.Sprintf(`%v`, arg.Value),
				},
			},
		}

	case "interlock":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "types.Interlock"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.INT,
					Value: fmt.Sprintf(`%v`, arg.Value),
				},
			},
		}

	case "anti-passback":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "types.AntiPassback"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.INT,
					Value: fmt.Sprintf(`%v`, arg.Value),
				},
			},
		}

	default:
		panic(fmt.Sprintf("unknown arg type '%v'", arg.Type))
	}
}

//	if err != nil {
//		t.Fatalf("%v", err)
//	} else if !slices.Equal(packet, expected) {
//
//		t.Errorf("incorrectly encoded request:\n   expected: %v\n   got:      %v", expected, packet)
//	}
func buildEncodeTestValidate(request lib.Request) dst.Stmt {
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
						X:   &dst.Ident{Name: "slices"},
						Sel: &dst.Ident{Name: "Equal"},
					},
					Args: []dst.Expr{
						&dst.Ident{Name: "packet"},
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
									Value: `"incorrectly encoded request:\n   expected: %v\n   got:      %v"`,
								},
								&dst.Ident{Name: "expected"},
								&dst.Ident{Name: "packet"},
							},
						},
					},
				},
			},
		},
	}
}
