package codec

import (
	"fmt"
	"log"
	"os"
	"slices"

	"go/ast"
	"go/printer"
	"go/token"

	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
	"codegen/model"
)

func decoder() {
	const output = "decoder.go"

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	decl := buildDecoder()

	// ... 'generated code' warning
	writeln(f, "// generated code - ** DO NOT EDIT **")
	writeln(f, "")

	printer.Fprint(f, token.NewFileSet(), decl)

	f.Close()
}

func buildDecoder() *ast.File {
	impl := []ast.Decl{
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
					Path: &ast.BasicLit{
						Kind: token.STRING,
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
	}

	impl = append(impl, buildDecoderMap()...)
	impl = append(impl, buildDecoderFunc())
	impl = append(impl, buildDecoderFactoryFunc())

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

		Decls: impl,
	}
}

func buildDecoderMap() []ast.Decl {
	key := ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent("key"),
				TypeParams: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{ast.NewIdent("R")},
							Type:  ast.NewIdent("any"),
						},
					},
				},
				Type: &ast.StructType{
					Fields: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{ast.NewIdent("opcode")},
								Type:  ast.NewIdent("byte"),
							},
						},
					},
				},
			},
		},
	}

	decoders := ast.GenDecl{
		Tok: token.VAR,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Names: []*ast.Ident{ast.NewIdent("decoders")},
				Values: []ast.Expr{
					&ast.CompositeLit{
						Type: &ast.MapType{
							Key: ast.NewIdent("any"),
							Value: &ast.FuncType{
								Params: &ast.FieldList{
									List: []*ast.Field{
										{
											Type: &ast.ArrayType{
												Elt: ast.NewIdent("byte"),
											},
										},
									},
								},
								Results: &ast.FieldList{
									List: []*ast.Field{
										{Type: ast.NewIdent("any")},
										{Type: ast.NewIdent("error")},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return []ast.Decl{&key, &decoders}
}

func buildDecoderFunc() *ast.FuncDecl {
	return &ast.FuncDecl{
		Name: ast.NewIdent("Decode"),
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
		Body: buildDecoderImpl(),
		Doc:  &ast.CommentGroup{},
	}
}

func buildDecoderImpl() *ast.BlockStmt {
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

			// blank line
			&ast.ExprStmt{
				X: &ast.BasicLit{
					Kind: token.STRING,
				},
			},

			// k := key[R]{packet[1]}
			&ast.AssignStmt{
				Tok: token.DEFINE, // :=

				Lhs: []ast.Expr{
					ast.NewIdent("k"),
				},

				Rhs: []ast.Expr{
					&ast.CompositeLit{
						Type: &ast.IndexExpr{
							X:     ast.NewIdent("key"),
							Index: ast.NewIdent("R"),
						},
						Elts: []ast.Expr{
							&ast.IndexExpr{
								// packet[1]
								X:     ast.NewIdent("packet"),
								Index: &ast.BasicLit{Kind: token.INT, Value: "1"},
							},
						},
					},
				},
			},

			// fn := decode
			&ast.AssignStmt{
				Tok: token.DEFINE,
				Lhs: []ast.Expr{
					ast.NewIdent("fn"),
				},
				Rhs: []ast.Expr{
					ast.NewIdent("decode"),
				},
			},

			// if f, ok := decoders[k]; ok {
			//     fn = f
			// }
			&ast.IfStmt{
				Init: &ast.AssignStmt{
					Tok: token.DEFINE,
					Lhs: []ast.Expr{
						ast.NewIdent("f"),
						ast.NewIdent("ok"),
					},
					Rhs: []ast.Expr{
						&ast.IndexExpr{
							X:     ast.NewIdent("decoders"),
							Index: ast.NewIdent("k"),
						},
					},
				},
				Cond: ast.NewIdent("ok"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Tok: token.ASSIGN,
							Lhs: []ast.Expr{ast.NewIdent("fn")},
							Rhs: []ast.Expr{ast.NewIdent("f")},
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

			// if v, err := fn(packet); err != nil {
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
							Fun:  ast.NewIdent("fn"),
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

func buildDecoderFactoryFunc() *ast.FuncDecl {
	return &ast.FuncDecl{
		Name: ast.NewIdent("decode"),
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
		Body: buildDecoderFactoryBody(),
		Doc:  &ast.CommentGroup{},
	}
}

func buildDecoderFactoryBody() *ast.BlockStmt {
	// switch packet[1] { ... }
	_switch := &ast.BlockStmt{
		List: []ast.Stmt{},
	}

	// ... message types
	excluded := []*lib.Response{
		&model.GetListenerAddrPortResponse,
		&model.SetListenerAddrPortResponse,
	}

	for _, response := range model.Responses {
		if slices.Contains(excluded, response) {
			log.Printf("skipping %v (excluded)", response.Name)
			continue
		}

		name := fmt.Sprintf("%v", codegen.TitleCase(response.Message.Name))

		clause := ast.CaseClause{
			List: []ast.Expr{
				&ast.BasicLit{
					Kind:  token.INT,
					Value: fmt.Sprintf("0x%02x", response.Message.MsgType),
				},
			},
			Body: []ast.Stmt{
				// return decode.<XXX>(packet)
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("decoder"),
								Sel: ast.NewIdent(name),
							},
							Args: []ast.Expr{ast.NewIdent("packet")},
						},
					},
				},
				// blank line
				&ast.ExprStmt{
					X: &ast.BasicLit{
						Kind: token.STRING,
					},
				},
			},
		}

		_switch.List = append(_switch.List, &clause)
	}

	// ... default
	_switch.List = append(_switch.List, &ast.CaseClause{
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
	})

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

			// blank line
			&ast.ExprStmt{
				X: &ast.BasicLit{
					Kind: token.STRING,
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

			// blank line
			&ast.ExprStmt{
				X: &ast.BasicLit{
					Kind: token.STRING,
				},
			},

			// switch packet[1] { ... }
			&ast.SwitchStmt{
				Tag: &ast.IndexExpr{
					X:     ast.NewIdent("packet"),
					Index: &ast.BasicLit{Kind: token.INT, Value: "1"},
				},
				Body: _switch,
			},
		},
	}
}
