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
			buildDecoderTestFunc(model.GetControllerResponse, model.GetControllerResponse.Tests[0]),
		},
	}
}

func buildDecoderTestFunc(response libx.Response, test libx.ResponseTest) *ast.FuncDecl {
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
		Body: buildDecoderTestImpl(response, test),
		Doc:  &ast.CommentGroup{},
	}
}

func buildDecoderTestImpl(response libx.Response, test libx.ResponseTest) *ast.BlockStmt {
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

			// blank line
			&ast.ExprStmt{
				X: &ast.BasicLit{
					Kind: token.STRING,
				},
			},

			// expected := ...
			&ast.AssignStmt{
				Lhs: []ast.Expr{
					&ast.Ident{Name: "expected"},
				},
				Tok: token.DEFINE,
				Rhs: []ast.Expr{
					buildExpected(response, test.Expected),
				},
			},
		},
	}
}

func buildExpected(response libx.Response, values []libx.Value) ast.Expr {
	name := codegen.TitleCase(response.Name)
	fields := []ast.Expr{}

	for _, field := range response.Fields {
		name := codegen.TitleCase(field.Name)

		for _, v := range values {
			if v.Name == field.Name {
				value := makeValue(field, v)

				f := ast.KeyValueExpr{
					Key:   &ast.Ident{Name: name},
					Value: value,
				}

				fields = append(fields, &f)
			}
		}
	}

	return &ast.CompositeLit{
		Type: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "responses"},
			Sel: &ast.Ident{Name: name},
		},
		Elts: fields,
	}
}

func makeValue(field libx.Field, value libx.Value) ast.Expr {
	switch field.Type {
	case "uint32":
		return &ast.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%v", value.Value)}

	case "IPv4":
		return &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "netip"},
				Sel: &ast.Ident{Name: "MustParseAddr"},
			},
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "MAC":
		return &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)}

	case "version":
		return &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)}

	case "date":
		return &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "entities"},
				Sel: &ast.Ident{Name: "MustParseDate"},
			},
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	default:
		return &ast.BasicLit{Kind: token.STRING, Value: `"???"`}
	}
}
