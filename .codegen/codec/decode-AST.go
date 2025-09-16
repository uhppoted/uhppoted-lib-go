package codec

import (
	"fmt"
	"log"
	"path/filepath"

	"go/ast"
	"go/token"

	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
	"codegen/model"
)

func decodeAST() {
	file := filepath.Join("decode", "_decode.go")

	imports := [][]string{
		[]string{
			"fmt",
		},
		[]string{
			"github.com/uhppoted/uhppoted-lib-go/uhppoted/responses",
		},
	}

	responses := []lib.Response{
		model.GetControllerResponse,
	}

	types := []*ast.GenDecl{}
	functions := []*ast.FuncDecl{}

	for _, response := range responses {
		if f := buildDecode(response); f != nil {
			functions = append(functions, f)
		}
	}

	AST := codegen.NewAST("responses", imports, types, functions)

	if err := AST.Generate(file); err != nil {
		log.Fatalf("error generating %v (%v)", file, err)
	} else {
		log.Printf("... generated %s", filepath.Base(file))
	}

}

func buildDecode(response lib.Response) *ast.FuncDecl {
	name := fmt.Sprintf("%v", codegen.TitleCase(response.Name))

	params := ast.FieldList{
		List: []*ast.Field{
			{
				Names: []*ast.Ident{ast.NewIdent("packet")},
				Type: &ast.ArrayType{
					Elt: ast.NewIdent("byte"),
				},
			},
		},
	}

	results := ast.FieldList{
		List: []*ast.Field{
			{
				Type: ast.NewIdent(fmt.Sprintf("responses.%v", name)),
			},
			{
				Type: ast.NewIdent("error"),
			},
		},
	}

	body := ast.BlockStmt{
		List: []ast.Stmt{
			// if len(packet) != 64 {
			//     return responses.GetControllerResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
			// }
			&ast.IfStmt{
				Cond: &ast.BinaryExpr{
					X: &ast.CallExpr{
						Fun:  &ast.Ident{Name: "len"},
						Args: []ast.Expr{&ast.Ident{Name: "packet"}},
					},
					Op: token.NEQ,
					Y:  &ast.BasicLit{Kind: token.INT, Value: "64"},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   &ast.Ident{Name: "responses"},
										Sel: &ast.Ident{Name: "GetControllerResponse"},
									},
									Args: nil,
								},
								&ast.CallExpr{
									Fun: &ast.Ident{Name: "fmt.Errorf"},
									Args: []ast.Expr{
										&ast.BasicLit{Kind: token.STRING, Value: `"invalid reply packet length (%v)"`},
										&ast.CallExpr{
											Fun:  &ast.Ident{Name: "len"},
											Args: []ast.Expr{&ast.Ident{Name: "packet"}},
										},
									},
								},
							},
						},
					},
				},
			},

			// blank line
			&ast.ExprStmt{
				X: &ast.BasicLit{
					Kind: token.STRING,
				},
			},

			// Ref. v6.62 firmware
			// if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
			//    return responses.GetControllerResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
			// }
			&ast.IfStmt{
				Cond: &ast.BinaryExpr{
					Op: token.LAND, // &&
					X: &ast.BinaryExpr{
						Op: token.NEQ,
						X: &ast.IndexExpr{
							X:     &ast.Ident{Name: "packet"},
							Index: &ast.BasicLit{Kind: token.INT, Value: "0"},
						},
						Y: &ast.Ident{Name: "SOM"},
					},
					Y: &ast.BinaryExpr{
						Op: token.LOR, // ||
						X: &ast.BinaryExpr{
							Op: token.NEQ,
							X: &ast.IndexExpr{
								X:     &ast.Ident{Name: "packet"},
								Index: &ast.BasicLit{Kind: token.INT, Value: "0"},
							},
							Y: &ast.Ident{Name: "SOM_v6_62"},
						},
						Y: &ast.BinaryExpr{
							Op: token.NEQ,
							X: &ast.IndexExpr{
								X:     &ast.Ident{Name: "packet"},
								Index: &ast.BasicLit{Kind: token.INT, Value: "1"},
							},
							Y: &ast.BasicLit{Kind: token.INT, Value: "0x20"},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   &ast.Ident{Name: "responses"},
										Sel: &ast.Ident{Name: "GetControllerResponse"},
									},
									Args: nil,
								},
								&ast.CallExpr{
									Fun: &ast.Ident{Name: "fmt.Errorf"},
									Args: []ast.Expr{
										&ast.BasicLit{Kind: token.STRING, Value: `"invalid reply start of message byte (%02x)"`},
										&ast.IndexExpr{
											X:     &ast.Ident{Name: "packet"},
											Index: &ast.BasicLit{Kind: token.INT, Value: "0"},
										},
									},
								},
							},
						},
					},
				},
			},

			// blank line
			&ast.ExprStmt{
				X: &ast.BasicLit{
					Kind: token.STRING,
				},
			},

			// if packet[1] != 148 {
			//     return responses.GetControllerResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
			// }
			&ast.IfStmt{
				Cond: &ast.BinaryExpr{
					Op: token.NEQ,
					X: &ast.IndexExpr{
						X:     &ast.Ident{Name: "packet"},
						Index: &ast.BasicLit{Kind: token.INT, Value: "1"},
					},
					Y: &ast.BasicLit{Kind: token.INT, Value: fmt.Sprintf("0x%02x", response.MsgType)},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   &ast.Ident{Name: "responses"},
										Sel: &ast.Ident{Name: "GetControllerResponse"},
									},
									Args: nil,
								},
								&ast.CallExpr{
									Fun: &ast.Ident{Name: "fmt.Errorf"},
									Args: []ast.Expr{
										&ast.BasicLit{Kind: token.STRING, Value: `"invalid reply function code (%02x)"`},
										&ast.IndexExpr{
											X:     &ast.Ident{Name: "packet"},
											Index: &ast.BasicLit{Kind: token.INT, Value: "1"},
										},
									},
								},
							},
						},
					},
				},
			},

			// blank line
			&ast.ExprStmt{
				X: &ast.BasicLit{
					Kind: token.STRING,
				},
			},

			// return responses.<T>{
			// 	...
			// }
			&ast.ReturnStmt{
				Results: []ast.Expr{
					&ast.CompositeLit{
						Type: &ast.SelectorExpr{
							X:   &ast.Ident{Name: "responses"},
							Sel: &ast.Ident{Name: name},
						},
						Elts: []ast.Expr{
							&ast.KeyValueExpr{
								Key: &ast.Ident{Name: "Controller"},
								Value: &ast.CallExpr{
									Fun:  &ast.Ident{Name: "unpackUint32"},
									Args: []ast.Expr{&ast.Ident{Name: "packet"}, &ast.BasicLit{Kind: token.INT, Value: "4"}},
								},
							},
						},
					},
					&ast.Ident{Name: "nil"},
				},
			},
		},
	}

	doc := ast.CommentGroup{}

	return &ast.FuncDecl{
		Name: ast.NewIdent(name),
		Type: &ast.FuncType{
			Params:  &params,
			Results: &results,
		},
		Body: &body,
		Doc:  &doc,
	}
}
