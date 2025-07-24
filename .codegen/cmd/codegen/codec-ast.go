package main

import (
	_ "embed"
	"log"
	"os"

	"go/ast"
	"go/printer"
	"go/token"
)

func decoderAST() {
	const output = "decodex.go"

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	decl := buildAST()

	printer.Fprint(f, token.NewFileSet(), decl)

	f.Close()
}

func buildAST() *ast.File {
	return &ast.File{
		Name: ast.NewIdent("codec"),

		Imports: []*ast.ImportSpec{
			{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"fmt"`,
				},
			},
			{
				Name: ast.NewIdent("decoder"),
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/decode"`,
				},
			},
		},

		Decls: []ast.Decl{
			&ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: `"fmt"`,
						},
					},
					&ast.ImportSpec{
						Name: ast.NewIdent("decoder"),
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: `"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/decode"`,
						},
					},
				},
			},
			buildDecodeFunc(),
			buildDecodeFactoryFunc(),
		},
	}
}

func buildDecodeFunc() *ast.FuncDecl {
	return &ast.FuncDecl{
		Name: ast.NewIdent("DecodeX"),
		Type: &ast.FuncType{
			TypeParams: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{ast.NewIdent("R")},
						Type:  ast.NewIdent("any"),
					},
				},
			},
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{ast.NewIdent("packet")},
						Type: &ast.ArrayType{
							Elt: ast.NewIdent("byte"),
						},
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent("R"),
					},
					{
						Type: ast.NewIdent("error"),
					},
				},
			},
		},
		Body: buildDecodeImpl(),
	}
}

func buildDecodeImpl() *ast.BlockStmt {
	return &ast.BlockStmt{
		List: []ast.Stmt{
			// var zero R
			&ast.DeclStmt{
				Decl: &ast.GenDecl{
					Tok: token.VAR,
					Specs: []ast.Spec{
						&ast.ValueSpec{
							Names: []*ast.Ident{ast.NewIdent("zero")},
							Type:  ast.NewIdent("R"),
						},
					},
				},
			},

			// if v, err := decode(packet); err != nil {
			//     return zero, fmt.Errorf("invalid packet")
			// } else if response, ok := v.(R); !ok {
			//     return zero, fmt.Errorf("invalid packet")
			// } else {
			//     return response, nil
			// }
			&ast.IfStmt{
				Init: &ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("v"),
						ast.NewIdent("err"),
					},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun:  ast.NewIdent("decodeX"),
							Args: []ast.Expr{ast.NewIdent("packet")},
						},
					},
				},
				Cond: &ast.BinaryExpr{
					X:  ast.NewIdent("err"),
					Op: token.NEQ,
					Y:  ast.NewIdent("nil"),
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								ast.NewIdent("zero"),
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("fmt"),
										Sel: ast.NewIdent("Errorf"),
									},
									Args: []ast.Expr{
										&ast.BasicLit{
											Kind:  token.STRING,
											Value: `"invalid packet"`,
										},
									},
								},
							},
						},
					},
				},
				Else: &ast.IfStmt{
					Init: &ast.AssignStmt{
						Lhs: []ast.Expr{
							ast.NewIdent("response"),
							ast.NewIdent("ok"),
						},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.TypeAssertExpr{
								X:    ast.NewIdent("v"),
								Type: ast.NewIdent("R"),
							},
						},
					},
					Cond: &ast.UnaryExpr{
						Op: token.NOT,
						X:  ast.NewIdent("ok"),
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ReturnStmt{
								Results: []ast.Expr{
									ast.NewIdent("zero"),
									&ast.CallExpr{
										Fun: &ast.SelectorExpr{
											X:   ast.NewIdent("fmt"),
											Sel: ast.NewIdent("Errorf"),
										},
										Args: []ast.Expr{
											&ast.BasicLit{
												Kind:  token.STRING,
												Value: `"invalid packet"`,
											},
										},
									},
								},
							},
						},
					},
					Else: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ReturnStmt{
								Results: []ast.Expr{
									ast.NewIdent("response"),
									ast.NewIdent("nil"),
								},
							},
						},
					},
				},
			},
		},
	}
}

func buildDecodeFactoryFunc() *ast.FuncDecl {
	return &ast.FuncDecl{
		Name: ast.NewIdent("decodeX"),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{ast.NewIdent("packet")},
						Type: &ast.ArrayType{
							Elt: ast.NewIdent("byte"),
						},
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent("any"),
					},
					{
						Type: ast.NewIdent("error"),
					},
				},
			},
		},
		Body: buildDecodeFactoryBody(),
	}
}

func buildDecodeFactoryBody() *ast.BlockStmt {
	return &ast.BlockStmt{
		List: []ast.Stmt{
			// if len(packet) != 64 {
			&ast.IfStmt{
				Cond: &ast.BinaryExpr{
					X: &ast.CallExpr{
						Fun:  ast.NewIdent("len"),
						Args: []ast.Expr{ast.NewIdent("packet")},
					},
					Op: token.NEQ,
					Y: &ast.BasicLit{
						Kind:  token.INT,
						Value: "64",
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								ast.NewIdent("nil"),
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("fmt"),
										Sel: ast.NewIdent("Errorf"),
									},
									Args: []ast.Expr{
										&ast.BasicLit{
											Kind:  token.STRING,
											Value: `"invalid reply packet length (%v)"`,
										},
										&ast.CallExpr{
											Fun:  ast.NewIdent("len"),
											Args: []ast.Expr{ast.NewIdent("packet")},
										},
									},
								},
							},
						},
					},
				},
			},

			// if packet[0] != SOM {
			&ast.IfStmt{
				Cond: &ast.BinaryExpr{
					X: &ast.IndexExpr{
						X:     ast.NewIdent("packet"),
						Index: &ast.BasicLit{Kind: token.INT, Value: "0"},
					},
					Op: token.NEQ,
					Y:  ast.NewIdent("SOM"),
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								ast.NewIdent("nil"),
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("fmt"),
										Sel: ast.NewIdent("Errorf"),
									},
									Args: []ast.Expr{
										&ast.BasicLit{
											Kind:  token.STRING,
											Value: `"invalid reply start of message byte (%02x)"`,
										},
										&ast.IndexExpr{
											X:     ast.NewIdent("packet"),
											Index: &ast.BasicLit{Kind: token.INT, Value: "0"},
										},
									},
								},
							},
						},
					},
				},
			},

			// switch packet[1] { ... }
			&ast.SwitchStmt{
				Tag: &ast.IndexExpr{
					X:     ast.NewIdent("packet"),
					Index: &ast.BasicLit{Kind: token.INT, Value: "1"},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						// case 0x94:
						&ast.CaseClause{
							List: []ast.Expr{
								&ast.BasicLit{
									Kind:  token.INT,
									Value: "0x94",
								},
							},
							Body: []ast.Stmt{
								&ast.ReturnStmt{
									Results: []ast.Expr{
										&ast.CallExpr{
											Fun: &ast.SelectorExpr{
												X:   ast.NewIdent("decoder"),
												Sel: ast.NewIdent("GetControllerResponse"),
											},
											Args: []ast.Expr{ast.NewIdent("packet")},
										},
									},
								},
							},
						},

						// default:
						&ast.CaseClause{
							List: nil,
							Body: []ast.Stmt{
								&ast.ReturnStmt{
									Results: []ast.Expr{
										ast.NewIdent("nil"),
										&ast.CallExpr{
											Fun: &ast.SelectorExpr{
												X:   ast.NewIdent("fmt"),
												Sel: ast.NewIdent("Errorf"),
											},
											Args: []ast.Expr{
												&ast.BasicLit{
													Kind:  token.STRING,
													Value: `"unknown message type (%02x)"`,
												},
												&ast.IndexExpr{
													X:     ast.NewIdent("packet"),
													Index: &ast.BasicLit{Kind: token.INT, Value: "1"},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
