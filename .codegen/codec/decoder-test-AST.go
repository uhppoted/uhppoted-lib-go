package codec

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"

	"go/ast"
	"go/printer"
	"go/token"

	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
	"codegen/model"
)

func decoderTest() {
	const output = "generated_test.go"

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	decl := buildDecoderTest()
	b := bytes.Buffer{}

	printer.Fprint(&b, token.NewFileSet(), decl)

	// ... format
	lines := strings.Split(b.String(), "\n")

	// ... reformat message packets and response structs
	re := map[string]*regexp.Regexp{
		"packet":   regexp.MustCompile(`^(\s*packet\s+:=\s+\[\]byte\s*{)(.*?)}\s*$`),
		"expected": regexp.MustCompile(`^(\s*expected\s+:=\s+responses\.(?:.*?)Response\s*{)(.*?)(}$)`),
	}

	for _, line := range lines {
		// ... reformat message packet?
		if match := re["packet"].FindStringSubmatch(line); len(match) > 1 {
			hex := match[2]

			writeln(f, match[1])
			for len(hex) > 96 {
				writeln(f, "		"+hex[:96])
				hex = hex[96:]
			}
			writeln(f, "		"+hex+",")
			writeln(f, "	}")
			continue
		}

		// ... reformat response struct ?
		if match := re["expected"].FindStringSubmatch(line); match != nil {
			fields := strings.Split(match[2], ",")

			writeln(f, match[1])
			for _, field := range fields {
				writeln(f, "		"+strings.Trim(field, " ")+",")
			}
			writeln(f, "	"+strings.Trim(match[3], " "))
			continue
		}

		// ... same old, same old
		writeln(f, line)
	}
}

func buildDecoderTest() *ast.File {
	imports := &ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"net/netip"`,
				},
			},

			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"reflect"`,
				},
			},

			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"testing"`,
				},
			},
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind: token.STRING,
				},
			},
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/uhppoted/entities"`,
				},
			},
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/uhppoted/responses"`,
				},
			},
		},
	}

	tests := []ast.Decl{}

	excluded := []*lib.Response{
		&model.GetListenerAddrPortResponse,
		&model.SetListenerAddrPortResponse,
	}

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

	return &ast.File{
		Name: ast.NewIdent("codec"),

		Imports: []*ast.ImportSpec{
			{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"net/netip"`,
				},
			},
			{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"reflect"`,
				},
			},
			{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"testing"`,
				},
			},
			// {
			// 	Path: &ast.BasicLit{
			// 		Kind:  token.STRING,
			// 		Value: `"github.com/uhppoted/uhppoted-lib-go/uhppoted/entities"`,
			// 	},
			// },
			{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/uhppoted/responses"`,
				},
			},
		},

		Decls: append([]ast.Decl{imports}, tests...),
	}
}

func buildDecoderTestFunc(response lib.Response, test lib.ResponseTest) *ast.FuncDecl {
	name := fmt.Sprintf("TestDecode%vResponse", codegen.TitleCase(test.Name))
	return &ast.FuncDecl{
		Name: ast.NewIdent(name),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{ast.NewIdent("t")},
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X:   &ast.Ident{Name: "testing"},
								Sel: &ast.Ident{Name: "T"},
							},
						},
					},
				},
			},
		},
		Body: buildDecoderTestImpl(response, test),
		Doc:  &ast.CommentGroup{},
	}
}

func buildDecoderTestImpl(response lib.Response, test lib.ResponseTest) *ast.BlockStmt {
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

			// blank line
			&ast.ExprStmt{
				X: &ast.BasicLit{
					Kind: token.STRING,
				},
			},

			// (exec)
			buildExec(response),
		},
	}
}

func buildExpected(response lib.Response, values []lib.Value) ast.Expr {
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

func buildExec(response lib.Response) ast.Stmt {
	name := codegen.TitleCase(response.Name)

	return &ast.IfStmt{
		Init: &ast.AssignStmt{
			Lhs: []ast.Expr{
				&ast.Ident{Name: "response"},
				&ast.Ident{Name: "err"},
			},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.IndexExpr{
						X: &ast.Ident{Name: "Decode"},
						Index: &ast.SelectorExpr{
							X:   &ast.Ident{Name: "responses"},
							Sel: &ast.Ident{Name: name},
						},
					},
					Args: []ast.Expr{&ast.Ident{Name: "packet"}},
				},
			},
		},
		Cond: &ast.BinaryExpr{
			X:  &ast.Ident{Name: "err"},
			Op: token.NEQ,
			Y:  &ast.Ident{Name: "nil"},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   &ast.Ident{Name: "t"},
							Sel: &ast.Ident{Name: "Fatalf"},
						},
						Args: []ast.Expr{
							&ast.BasicLit{Kind: token.STRING, Value: `"%v"`},
							&ast.Ident{Name: "err"},
						},
					},
				},
			},
		},
		Else: &ast.IfStmt{
			Cond: &ast.UnaryExpr{
				Op: token.NOT,
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "reflect"},
						Sel: &ast.Ident{Name: "DeepEqual"},
					},
					Args: []ast.Expr{
						&ast.Ident{Name: "response"},
						&ast.Ident{Name: "expected"},
					},
				},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ExprStmt{
						X: &ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   &ast.Ident{Name: "t"},
								Sel: &ast.Ident{Name: "Errorf"},
							},
							Args: []ast.Expr{
								&ast.BasicLit{
									Kind:  token.STRING,
									Value: `"incorrectly decoded response:\n   expected: %#v\n   got:      %#v"`,
								},
								&ast.Ident{Name: "expected"},
								&ast.Ident{Name: "response"},
							},
						},
					},
				},
			},
		},
	}
}

func makeValue(field lib.Field, value lib.Value) ast.Expr {
	switch field.Type {
	case "uint8":
		return &ast.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%v", value.Value)}

	case "uint16":
		return &ast.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%v", value.Value)}

	case "uint32":
		return &ast.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%v", value.Value)}

	case "bool":
		return &ast.Ident{Name: fmt.Sprintf("%v", value.Value)}

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

	case "address:port":
		return &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "netip"},
				Sel: &ast.Ident{Name: "MustParseAddrPort"},
			},
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

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

	case "shortdate":
		return &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "entities"},
				Sel: &ast.Ident{Name: "MustParseDate"},
			},
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "optional date":
		return &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "entities"},
				Sel: &ast.Ident{Name: "MustParseDate"},
			},
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "time":
		return &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "entities"},
				Sel: &ast.Ident{Name: "MustParseTime"},
			},
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "HHmm":
		return &ast.CallExpr{
			Fun: &ast.Ident{Name: "string2HHmm"},
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "datetime":
		return &ast.CallExpr{
			Fun: &ast.Ident{Name: "string2datetime"},
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "optional datetime":
		return &ast.CallExpr{
			Fun: &ast.Ident{Name: "string2datetime"},
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "MAC":
		return &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)}

	case "version":
		return &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)}

	case "pin":
		return &ast.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%v", value.Value)}

	default:
		panic(fmt.Sprintf("%v", field.Type))
		// return &ast.BasicLit{Kind: token.STRING, Value: `"???"`}
	}
}

func writeln(f *os.File, s string) {
	if _, err := f.WriteString(s + "\n"); err != nil {
		panic(fmt.Errorf("error writing to %v (%v)", f.Name(), err))
	}
}
