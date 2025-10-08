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

func decode() {
	file := filepath.Join("decode", "generated.go")

	imports := [][]string{
		[]string{
			"fmt",
		},
		[]string{
			"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses",
		},
	}

	responses := []*lib.Response{}
	responses = append(responses, model.Responses...)
	responses = append(responses, &model.ListenerEvent)

	types := []*ast.GenDecl{}
	functions := []*ast.FuncDecl{}

	for _, response := range responses {
		if f := buildDecode(*response); f != nil {
			functions = append(functions, f)
		}
	}

	AST := codegen.NewAST("decode", imports, types, functions)

	if err := AST.Generate(file); err != nil {
		log.Fatalf("error generating %v (%v)", file, err)
	} else {
		log.Printf("... generated %s", filepath.Base(file))
	}

}

func buildDecode(r lib.Response) *ast.FuncDecl {
	name := fmt.Sprintf("%v", codegen.TitleCase(r.Name))

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

	response := []ast.Expr{}
	for _, f := range r.Fields {
		response = append(response, unpack(f))
	}

	body := ast.BlockStmt{
		List: []ast.Stmt{
			// if len(packet) != 64 {
			//     return responses.<R>{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
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
								&ast.CompositeLit{
									Type: &ast.SelectorExpr{
										X:   &ast.Ident{Name: "responses"},
										Sel: &ast.Ident{Name: name},
									},
									Elts: nil,
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
			//    return responses.<R>{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
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
								&ast.CompositeLit{
									Type: &ast.SelectorExpr{
										X:   &ast.Ident{Name: "responses"},
										Sel: &ast.Ident{Name: name},
									},
									Elts: nil,
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

			// if packet[1] != <code> {
			//     return responses.<R>{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
			// }
			&ast.IfStmt{
				Cond: &ast.BinaryExpr{
					Op: token.NEQ,
					X: &ast.IndexExpr{
						X:     &ast.Ident{Name: "packet"},
						Index: &ast.BasicLit{Kind: token.INT, Value: "1"},
					},
					Y: &ast.BasicLit{Kind: token.INT, Value: fmt.Sprintf("0x%02x", r.MsgType)},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.CompositeLit{
									Type: &ast.SelectorExpr{
										X:   &ast.Ident{Name: "responses"},
										Sel: &ast.Ident{Name: name},
									},
									Elts: nil,
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
						Elts: response,
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

func unpack(field lib.Field) ast.Expr {
	types := map[string]string{
		"bool":              "unpackBool",
		"uint8":             "unpackUint8",
		"uint16":            "unpackUint16",
		"uint32":            "unpackUint32",
		"datetime":          "unpackDateTime",
		"optional datetime": "unpackOptionalDateTime",
		"date":              "unpackDate",
		"shortdate":         "unpackShortDate",
		"optional date":     "unpackOptionalDate",
		"time":              "unpackTime",
		"HHmm":              "unpackHHmm",
		"IPv4":              "unpackIPv4",
		"address:port":      "unpackAddrPort",
		"MAC":               "unpackMAC",
		"pin":               "unpackPIN",
		"mode":              "unpackMode",
		"event-type":        "unpackEventType",
		"version":           "unpackVersion",
	}

	if f, ok := types[field.Type]; !ok {
		panic(fmt.Sprintf("unknown response field type (%v)", field.Type))
	} else {
		return &ast.KeyValueExpr{
			Key: &ast.Ident{
				Name: codegen.TitleCase(field.Name),
			},
			Value: &ast.CallExpr{
				Fun: &ast.Ident{Name: f},
				Args: []ast.Expr{
					&ast.Ident{Name: "packet"},
					&ast.BasicLit{
						Kind:  token.INT,
						Value: fmt.Sprintf("%v", field.Offset),
					}},
			},
		}
	}
}
