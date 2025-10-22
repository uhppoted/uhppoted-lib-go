package api

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"go/printer"
	"go/token"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"

	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
	"codegen/model"
)

func API() {
	outfile := filepath.Join("generated.go")
	decl := buildAPI()

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

	// ... write to file
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

func buildAPI() *dst.File {
	impl := []dst.Decl{
		&dst.GenDecl{
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
						Value: `"time"`,
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
						Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/codec/encode"`,
					},
				},
				&dst.ImportSpec{
					Path: &dst.BasicLit{
						Kind:  token.STRING,
						Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/entities"`,
					},
				},
				&dst.ImportSpec{
					Path: &dst.BasicLit{
						Kind:  token.STRING,
						Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"`,
					},
				},
			},
		},
	}

	for _, api := range model.API[1:] {
		if f := buildFunction(*api); f != nil {
			impl = append(impl, f)
		}
	}

	return &dst.File{
		Name: dst.NewIdent("uhppoted"),

		Imports: []*dst.ImportSpec{
			{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"net/netip"`,
				},
			},
			{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"time"`,
				},
			},
			{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/codec/encode"`,
				},
			},
			{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/entities"`,
				},
			},
			{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"`,
				},
			},
		},

		Decls: impl,
	}

	// imports := [][]string{
	// 	[]string{
	// 		"net/netip",
	// 		"time",
	// 	},
	// 	[]string{
	// 		"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/codec/encode",
	// 		"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/entities",
	// 		"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses",
	// 	},
	// }
	//
	// types := []*dst.GenDecl{}
	// functions := []*dst.FuncDecl{}
	// excluded := []*lib.Function{}
	//
	// for _, f := range model.API[1:] {
	// 	if slices.Contains(excluded, f) {
	// 		log.Printf("skipping %v (excluded)", f.Name)
	// 		continue
	// 	}
	//
	// 	functions = append(functions, function(*f))
	// }
	//
	// AST := codegen.NewAST("uhppoted", imports, types, functions)
	//
	// if err := AST.Generate(outfile); err != nil {
	// 	log.Fatalf("error generating %v (%v)", outfile, err)
	// } else {
	// 	log.Printf("... generated %s", filepath.Base(outfile))
	// }
}

func buildFunction(f lib.Function) *dst.FuncDecl {
	name := codegen.TitleCase(f.Name)
	response := fmt.Sprintf("responses.%v", strings.TrimSuffix(codegen.TitleCase(f.Response.Name), "Response"))

	// ... function type
	ftype := []*dst.Field{}
	set := map[string]bool{}

	for _, arg := range f.Args {
		switch arg.Type {
		case "controller":
			if defined := set[arg.Type]; !defined {
				ftype = append(ftype, &dst.Field{
					Names: []*dst.Ident{dst.NewIdent("T")},
					Type:  dst.NewIdent("TController"),
				})

				set[arg.Type] = true
			}

		case "datetime":
			if defined := set[arg.Type]; !defined {
				ftype = append(ftype, &dst.Field{
					Names: []*dst.Ident{dst.NewIdent("DT")},
					Type:  dst.NewIdent("TDateTime"),
				})
				set[arg.Type] = true
			}

		case "date":
			if defined := set[arg.Type]; !defined {
				ftype = append(ftype, &dst.Field{
					Names: []*dst.Ident{dst.NewIdent("D")},
					Type:  dst.NewIdent("TDate"),
				})
				set[arg.Type] = true
			}

		case "HHmm":
			if defined := set[arg.Type]; !defined {
				ftype = append(ftype, &dst.Field{
					Names: []*dst.Ident{dst.NewIdent("H")},
					Type:  dst.NewIdent("THHmm"),
				})
				set[arg.Type] = true
			}
		}
	}

	// ... args
	args := []*dst.Field{}
	args = append(args, &dst.Field{
		Names: []*dst.Ident{
			{Name: "u"},
		},
		Type: &dst.Ident{Name: "Uhppoted"},
	})

	for _, arg := range f.Args {
		name := regexp.MustCompile(`[ \-]+`).ReplaceAllString(arg.Name, "")
		t := arg.Type

		switch arg.Type {
		case "controller":
			t = "T"

		case "IPv4":
			t = "netip.Addr"

		case "address:port":
			t = "netip.AddrPort"

		case "date":
			t = "D"

		case "datetime":
			t = "DT"

		case "HHmm":
			t = "H"

		case "pin":
			t = "uint32"

		case "mode":
			t = "DoorMode"

		case "task":
			t = "TaskType"

		case "interlock":
			t = "Interlock"
		}

		args = append(args, &dst.Field{
			Names: []*dst.Ident{
				{Name: name},
			},
			Type: &dst.Ident{Name: t},
		})
	}

	args = append(args, &dst.Field{
		Names: []*dst.Ident{
			{Name: "timeout"},
		},
		Type: &dst.Ident{Name: "time.Duration"},
	})

	// ... compose func
	decl := dst.FuncDecl{
		Name: dst.NewIdent(name),
		Type: &dst.FuncType{
			TypeParams: &dst.FieldList{
				List: ftype,
			},
			Params: &dst.FieldList{
				List: args,
			},
			Results: &dst.FieldList{
				List: []*dst.Field{
					{
						Type: dst.NewIdent(response),
					},
					{
						Type: dst.NewIdent("error"),
					},
				},
			},
		},
		Body: impl(f),
	}

	// godoc
	for _, line := range f.Description {
		decl.Decs.Start.Append(fmt.Sprintf("// %v", line))
	}

	decl.Decs.Before = dst.EmptyLine

	return &decl
}

func impl(f lib.Function) *dst.BlockStmt {
	request := codegen.TitleCase(f.Request.Name)
	response := fmt.Sprintf("responses.%v", strings.TrimSuffix(codegen.TitleCase(f.Response.Name), "Response"))

	args := []dst.Expr{
		&dst.Ident{Name: "id"},
	}

loop:
	for _, arg := range f.Request.Fields[1:] {
		name := regexp.MustCompile(`\s+`).ReplaceAllString(arg.Name, "")

		switch arg.Type {
		case "magic":
			continue loop

		case "datetime":
			args = append(args, &dst.CallExpr{
				Fun: &dst.IndexExpr{
					X:     &dst.Ident{Name: "convert"},
					Index: &dst.Ident{Name: "entities.DateTime"},
				},
				Args: []dst.Expr{
					&dst.Ident{Name: name},
				},
			})

		case "date":
			args = append(args, &dst.CallExpr{
				Fun: &dst.IndexExpr{
					X:     &dst.Ident{Name: "convert"},
					Index: &dst.Ident{Name: "entities.Date"},
				},
				Args: []dst.Expr{
					&dst.Ident{Name: name},
				},
			})

		case "HHmm":
			args = append(args, &dst.CallExpr{
				Fun: &dst.IndexExpr{
					X:     &dst.Ident{Name: "convert"},
					Index: &dst.Ident{Name: "entities.HHmm"},
				},
				Args: []dst.Expr{
					&dst.Ident{Name: name},
				},
			})
		default:
			args = append(args, &dst.Ident{Name: name})
		}
	}

	return &dst.BlockStmt{
		List: []dst.Stmt{
			//	f := func(id uint32) ([]byte, error) {
			//       return d.XXX(id,...)
			//  }
			&dst.AssignStmt{
				Lhs: []dst.Expr{
					&dst.Ident{Name: "f"},
				},
				Tok: token.DEFINE,
				Rhs: []dst.Expr{
					&dst.FuncLit{
						Type: &dst.FuncType{
							Params: &dst.FieldList{
								List: []*dst.Field{
									&dst.Field{
										Names: []*dst.Ident{
											{Name: "id"},
										},
										Type: &dst.Ident{Name: "uint32"},
									},
								},
							},
							Results: &dst.FieldList{
								List: []*dst.Field{
									{Type: &dst.ArrayType{
										Elt: &dst.Ident{Name: "byte"},
									}},
									{Type: &dst.Ident{Name: "error"}},
								},
							},
						},
						Body: &dst.BlockStmt{
							List: []dst.Stmt{
								&dst.ReturnStmt{
									Results: []dst.Expr{
										&dst.CallExpr{
											Fun: &dst.SelectorExpr{
												X:   &dst.Ident{Name: "encode"},
												Sel: &dst.Ident{Name: request},
											},
											Args: args,
										},
									},
								},
							},
						},
					},
				},
			},

			blankline(),

			// return exec[T, R](u, controller, f, timeout)
			&dst.ReturnStmt{
				Results: []dst.Expr{
					&dst.CallExpr{
						Fun: &dst.IndexListExpr{
							X: &dst.Ident{Name: "exec"},
							Indices: []dst.Expr{
								&dst.Ident{Name: "T"},
								&dst.Ident{Name: response},
							},
						},
						Args: []dst.Expr{
							&dst.Ident{Name: "u"},
							&dst.Ident{Name: "controller"},
							&dst.Ident{Name: "f"},
							&dst.Ident{Name: "timeout"},
						},
					},
				},
			},
		},
	}
}

func blankline() dst.Stmt {
	return &dst.ExprStmt{
		X: &dst.BasicLit{
			Kind: token.STRING,
		},
	}
}
