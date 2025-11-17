package integration_tests

import (
	"bytes"
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

	for _, f := range model.API {
		for _, test := range f.Tests {
			field := buildStructField(*f, test)

			fields = append(fields, &field)
		}
	}

	// ... struct initialisation
	values := []dst.Expr{}

	value := dst.CompositeLit{
		Type: &dst.SelectorExpr{
			X:   dst.NewIdent("responses"),
			Sel: dst.NewIdent("GetController"),
		},

		Elts: []dst.Expr{
			&dst.KeyValueExpr{
				Key:   dst.NewIdent("Controller"),
				Value: &dst.BasicLit{Kind: token.INT, Value: "303986753"},
				Decs: dst.KeyValueExprDecorations{
					NodeDecs: dst.NodeDecs{
						Before: dst.NewLine,
						After:  dst.NewLine,
					},
				},
			},

			&dst.KeyValueExpr{
				Key: dst.NewIdent("IpAddress"),
				Value: &dst.CallExpr{
					Fun: &dst.SelectorExpr{
						X:   dst.NewIdent("netip"),
						Sel: dst.NewIdent("MustParseAddr"),
					},
					Args: []dst.Expr{
						&dst.BasicLit{Kind: token.STRING, Value: "\"192.168.1.100\""},
					},
				},
				Decs: dst.KeyValueExprDecorations{
					NodeDecs: dst.NodeDecs{
						Before: dst.NewLine,
						After:  dst.NewLine,
					},
				},
			},
		},

		Decs: dst.CompositeLitDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.NewLine,
				After:  dst.EmptyLine,
			},
		},
	}

	values = append(values, &value)

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
