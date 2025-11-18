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

func expectedAST() {
	outfile := filepath.Join(".", "_expected.go")
	decl := buildExpected()

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

func buildExpected() *dst.File {
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
					Kind: token.STRING,
				},
			},
			&dst.ImportSpec{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"`,
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

	file := &dst.File{
		Name: dst.NewIdent("uhppoted"),
		Decls: []dst.Decl{
			imports,
			buildExpectedVar(),
		},
	}

	return file
}

// var Expected = struct{...}{...}
func buildExpectedVar() dst.Decl {
	// ... struct fields
	fields := []*dst.Field{}
	for _, fn := range model.API {
		for _, test := range fn.Tests {
			field := buildStructField(*fn, test)

			fields = append(fields, &field)
		}
	}

	// ... struct initialisation
	values := []dst.Expr{}
	for _, fn := range model.API[1:] {
		for _, test := range fn.Tests {
			value := buildStructValue(*fn, test)

			values = append(values, &value)
		}
	}

	return &dst.GenDecl{
		Tok: token.VAR,
		Specs: []dst.Spec{
			&dst.ValueSpec{
				Names: []*dst.Ident{
					dst.NewIdent("Expected"),
				},
				Values: []dst.Expr{
					&dst.CompositeLit{
						Type: &dst.StructType{
							Fields: &dst.FieldList{
								List: fields,
							},
						},
						Elts: values,
					},
				},
			},
		},
	}
}

func buildStructField(fn lib.Function, test lib.FuncTest) dst.Field {
	name := codegen.TitleCase(test.Name)
	response := codegen.TitleCase(fn.Response.Name)

	if name == "FindControllers" {
		return dst.Field{
			Names: []*dst.Ident{
				dst.NewIdent(name),
			},
			Type: &dst.ArrayType{
				Elt: &dst.SelectorExpr{
					X:   dst.NewIdent("responses"),
					Sel: dst.NewIdent("GetController"),
				},
			},
		}
	}

	return dst.Field{
		Names: []*dst.Ident{
			dst.NewIdent(name),
		},
		Type: &dst.SelectorExpr{
			X:   dst.NewIdent("responses"),
			Sel: dst.NewIdent(response),
		},
	}
}

func buildStructValue(fn lib.Function, test lib.FuncTest) dst.KeyValueExpr {
	name := codegen.TitleCase(test.Name)
	response := codegen.TitleCase(fn.Response.Name)
	fields := []dst.Expr{}

	if len(test.Replies) > 1 {
		panic("not expecting more than one response")
	}

	for _, reply := range test.Replies[:1] {
		for _, v := range reply.Response {
			ident := codegen.TitleCase(v.Name)

			e := dst.KeyValueExpr{
				Key:   dst.NewIdent(ident),
				Value: buildFieldValue(v),
				Decs: dst.KeyValueExprDecorations{
					NodeDecs: dst.NodeDecs{
						Before: dst.NewLine,
						After:  dst.NewLine,
					},
				},
			}

			fields = append(fields, &e)
		}
	}

	value := dst.CompositeLit{
		Type: &dst.SelectorExpr{
			X:   dst.NewIdent("responses"),
			Sel: dst.NewIdent(response),
		},

		Elts: fields,

		Decs: dst.CompositeLitDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.NewLine,
				After:  dst.EmptyLine,
			},
		},
	}

	return dst.KeyValueExpr{
		Key:   dst.NewIdent(name),
		Value: &value,
		Decs: dst.KeyValueExprDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.NewLine,
				After:  dst.EmptyLine,
			},
		},
	}
}

func buildFieldValue(value lib.Value) dst.Expr {
	switch value.Type {
	case "bool":
		return &dst.Ident{
			Name: fmt.Sprintf("%v", value.Value),
		}

	case "uint8":
		return &dst.BasicLit{
			Kind:  token.INT,
			Value: fmt.Sprintf("%v", value.Value),
		}

	case "uint16":
		return &dst.BasicLit{
			Kind:  token.INT,
			Value: fmt.Sprintf("%v", value.Value),
		}

	case "uint32":
		return &dst.BasicLit{
			Kind:  token.INT,
			Value: fmt.Sprintf("%v", value.Value),
		}

	case "IPv4":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "netip.MustParseAddr"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, value.Value),
				},
			},
		}

	case "address:port":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "netip.MustParseAddrPort"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, value.Value),
				},
			},
		}

	case "datetime":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseDateTime"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "optional datetime":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseDateTime"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "date":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseDate"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "shortdate":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseDate"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "optional date":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseDate"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "time":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseTime"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "HHmm":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseHHmm"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, value.Value),
				},
			},
		}

	case "MAC":
		return &dst.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf(`"%v"`, value.Value),
		}

	case "version":
		return &dst.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf(`"%v"`, value.Value),
		}

	case "pin":
		return &dst.BasicLit{
			Kind:  token.INT,
			Value: fmt.Sprintf("%v", value.Value),
		}

	case "mode":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "DoorMode"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf(`%v`, value.Value)},
			},
		}

	case "anti-passback":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "AntiPassback"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf(`%v`, value.Value)},
			},
		}

	case "event-type":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "EventType"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf(`%v`, value.Value)},
			},
		}

	case "direction":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "Direction"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf(`%v`, value.Value)},
			},
		}

	case "reason":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "Reason"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf(`%v`, value.Value)},
			},
		}

	default:
		panic(fmt.Sprintf("unknown response field type (%v)", value.Type))
	}
}
