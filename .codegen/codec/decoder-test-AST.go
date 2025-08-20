package codec

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"go/ast"
	"go/printer"
	"go/token"

	libx "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
	"codegen/model"
)

func decoderTest() {
	const output = "_decoder_test.go"

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	decl := buildDecoderTest()

	printer.Fprint(f, token.NewFileSet(), decl)

	f.Close()
}

func buildDecoderTest() *ast.File {
	return &ast.File{
		Name: ast.NewIdent("codec"),

		Imports: []*ast.ImportSpec{
			{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"testing"`,
				},
			},
			// {
			//     Name: ast.NewIdent("decoder"),
			//     Path: &ast.BasicLit{
			//         Kind:  token.STRING,
			//         Value: `"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/decode"`,
			//     },
			// },
		},

		Decls: []ast.Decl{
			&ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: `"testing"`,
						},
					},
					//     &ast.ImportSpec{
					//         Path: &ast.BasicLit{
					//             Kind: token.STRING,
					//         },
					//     },
					//     &ast.ImportSpec{
					//         Name: ast.NewIdent("decoder"),
					//         Path: &ast.BasicLit{
					//             Kind:  token.STRING,
					//             Value: `"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/decode"`,
					//         },
					//     },
				},
			},
			buildDecoderTestFunc(model.GetControllerResponse.Tests[0]),
		},
	}
}

func buildDecoderTestFunc(test libx.ResponseTest) *ast.FuncDecl {
	name := fmt.Sprintf("TestDecode%vResponse", codegen.TitleCase(test.Name))
	return &ast.FuncDecl{
		Name: ast.NewIdent(name),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{ast.NewIdent("t")},
						Type: &ast.ArrayType{
							Elt: ast.NewIdent("*testing.T"),
						},
					},
				},
			},
		},
		Body: buildDecoderTestImpl(test),
		Doc:  &ast.CommentGroup{},
	}
}

func buildDecoderTestImpl(test libx.ResponseTest) *ast.BlockStmt {
	packet := make([]ast.Expr, 64)
	for i, b := range test.Response {
		packet[i] = &ast.BasicLit{
			Kind:  token.INT,
			Value: fmt.Sprintf("0x%02x", b),
		}
	}

	return &ast.BlockStmt{
		List: []ast.Stmt{
			// packet := []byte{...}
			&ast.AssignStmt{
				Lhs: []ast.Expr{
					ast.NewIdent("packet"),
				},
				Tok: token.DEFINE,
				Rhs: []ast.Expr{
					&ast.CompositeLit{
						Type: &ast.ArrayType{
							Len: nil,
							Elt: ast.NewIdent("byte"),
						},
						Elts: packet,
					},
				},
			},

			// // blank line
			// &ast.ExprStmt{
			//     X: &ast.BasicLit{
			//         Kind: token.STRING,
			//     },
			// },

			// // if v, err := decode(packet); err != nil {
			// //     return zero, fmt.Errorf("invalid packet")
			// // } else if response, ok := v.(R); !ok {
			// //     return zero, fmt.Errorf("invalid packet")
			// // } else {
			// //     return response, nil
			// // }
			// &ast.IfStmt{
			//     Init: &ast.AssignStmt{
			//         Lhs: []ast.Expr{
			//             ast.NewIdent("v"),
			//             ast.NewIdent("err"),
			//         },
			//         Tok: token.DEFINE,
			//         Rhs: []ast.Expr{
			//             &ast.CallExpr{
			//                 Fun:  ast.NewIdent("decode"),
			//                 Args: []ast.Expr{ast.NewIdent("packet")},
			//             },
			//         },
			//     },
			//     Cond: &ast.BinaryExpr{
			//         X:  ast.NewIdent("err"),
			//         Op: token.NEQ,
			//         Y:  ast.NewIdent("nil"),
			//     },
			//     Body: &ast.BlockStmt{
			//         List: []ast.Stmt{
			//             &ast.ReturnStmt{
			//                 Results: []ast.Expr{
			//                     ast.NewIdent("zero"),
			//                     &ast.CallExpr{
			//                         Fun: &ast.SelectorExpr{
			//                             X:   ast.NewIdent("fmt"),
			//                             Sel: ast.NewIdent("Errorf"),
			//                         },
			//                         Args: []ast.Expr{
			//                             &ast.BasicLit{
			//                                 Kind:  token.STRING,
			//                                 Value: `"invalid packet"`,
			//                             },
			//                         },
			//                     },
			//                 },
			//             },
			//         },
			//     },
			//     Else: &ast.IfStmt{
			//         Init: &ast.AssignStmt{
			//             Lhs: []ast.Expr{
			//                 ast.NewIdent("response"),
			//                 ast.NewIdent("ok"),
			//             },
			//             Tok: token.DEFINE,
			//             Rhs: []ast.Expr{
			//                 &ast.TypeAssertExpr{
			//                     X:    ast.NewIdent("v"),
			//                     Type: ast.NewIdent("R"),
			//                 },
			//             },
			//         },
			//         Cond: &ast.UnaryExpr{
			//             Op: token.NOT,
			//             X:  ast.NewIdent("ok"),
			//         },
			//         Body: &ast.BlockStmt{
			//             List: []ast.Stmt{
			//                 &ast.ReturnStmt{
			//                     Results: []ast.Expr{
			//                         ast.NewIdent("zero"),
			//                         &ast.CallExpr{
			//                             Fun: &ast.SelectorExpr{
			//                                 X:   ast.NewIdent("fmt"),
			//                                 Sel: ast.NewIdent("Errorf"),
			//                             },
			//                             Args: []ast.Expr{
			//                                 &ast.BasicLit{
			//                                     Kind:  token.STRING,
			//                                     Value: `"invalid packet"`,
			//                                 },
			//                             },
			//                         },
			//                     },
			//                 },
			//             },
			//         },
			//         Else: &ast.BlockStmt{
			//             List: []ast.Stmt{
			//                 &ast.ReturnStmt{
			//                     Results: []ast.Expr{
			//                         ast.NewIdent("response"),
			//                         ast.NewIdent("nil"),
			//                     },
			//                 },
			//             },
			//         },
			//     },
			// },
		},
	}
}
