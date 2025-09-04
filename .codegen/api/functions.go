package api

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"slices"

	"go/ast"
	"go/token"

	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
	"codegen/model"
)

func API() {
	const file = "generated.go"

	imports := []string{
		"net/netip",
		"time",

		"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/encode",
		"github.com/uhppoted/uhppoted-lib-go/uhppoted/responses",
	}

	types := []*ast.GenDecl{}
	functions := []*ast.FuncDecl{}
	excluded := []*lib.Function{
		&model.GetListenerAddrPort,
		&model.SetListenerAddrPort,
	}

	for _, f := range model.API[1:] {
		if slices.Contains(excluded, f) {
			log.Printf("skipping %v (excluded)", f.Name)
			continue
		}

		functions = append(functions, function(*f))
	}

	AST := codegen.NewAST("uhppoted", imports, types, functions)

	if err := AST.Generate(file); err != nil {
		log.Fatalf("error generating %v (%v)", file, err)
	} else {
		log.Printf("... generated %s", filepath.Base(file))
	}
}

func function(f lib.Function) *ast.FuncDecl {
	name := codegen.TitleCase(f.Name)
	response := fmt.Sprintf("responses.%v", codegen.TitleCase(f.Response.Name))

	// ... args
	args := []*ast.Field{}
	args = append(args, &ast.Field{
		Names: []*ast.Ident{
			{Name: "u"},
		},
		Type: &ast.Ident{Name: "Uhppoted"},
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
			t = "time.Time"

		case "datetime":
			t = "time.Time"

		case "HHmm":
			t = "time.Time"

		case "pin":
			t = "uint32"
		}

		args = append(args, &ast.Field{
			Names: []*ast.Ident{
				{Name: name},
			},
			Type: &ast.Ident{Name: t},
		})
	}

	args = append(args, &ast.Field{
		Names: []*ast.Ident{
			{Name: "timeout"},
		},
		Type: &ast.Ident{Name: "time.Duration"},
	})

	// ... godoc
	// godoc := ast.CommentGroup{
	godoc := []*ast.Comment{
		{Text: fmt.Sprintf("// -- line intentionally left blank --")},
	}

	for _, line := range f.Description {
		text := fmt.Sprintf("// %v", line)
		comment := ast.Comment{
			Text: text,
		}

		godoc = append(godoc, &comment)
	}

	// ... compose func
	return &ast.FuncDecl{
		Name: ast.NewIdent(name),
		Type: &ast.FuncType{
			TypeParams: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{ast.NewIdent("T")},
						Type:  ast.NewIdent("TController"),
					},
				},
			},
			Params: &ast.FieldList{
				List: args,
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent(response),
					},
					{
						Type: ast.NewIdent("error"),
					},
				},
			},
		},
		Body: impl(f),
		Doc: &ast.CommentGroup{
			List: godoc,
		},
	}
}

func impl(f lib.Function) *ast.BlockStmt {
	request := codegen.TitleCase(f.Request.Name)
	response := fmt.Sprintf("responses.%v", codegen.TitleCase(f.Response.Name))

	args := []ast.Expr{
		&ast.Ident{Name: "id"},
	}

	for _, arg := range f.Request.Fields[1:] {
		name := regexp.MustCompile(`\s+`).ReplaceAllString(arg.Name, "")

		if arg.Type != "magic" {
			args = append(args, &ast.Ident{Name: name})
		}
	}

	return &ast.BlockStmt{
		List: []ast.Stmt{
			//	f := func(id uint32) ([]byte, error) {
			//       return d.XXX(id,...)
			//  }
			&ast.AssignStmt{
				Lhs: []ast.Expr{
					&ast.Ident{Name: "f"},
				},
				Tok: token.DEFINE,
				Rhs: []ast.Expr{
					&ast.FuncLit{
						Type: &ast.FuncType{
							Params: &ast.FieldList{
								List: []*ast.Field{
									&ast.Field{
										Names: []*ast.Ident{
											{Name: "id"},
										},
										Type: &ast.Ident{Name: "uint32"},
									},
								},
							},
							Results: &ast.FieldList{
								List: []*ast.Field{
									{Type: &ast.ArrayType{
										Elt: &ast.Ident{Name: "byte"},
									}},
									{Type: &ast.Ident{Name: "error"}},
								},
							},
						},
						Body: &ast.BlockStmt{
							List: []ast.Stmt{
								&ast.ReturnStmt{
									Results: []ast.Expr{
										&ast.CallExpr{
											Fun: &ast.SelectorExpr{
												X:   &ast.Ident{Name: "encode"},
												Sel: &ast.Ident{Name: request},
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
			&ast.ReturnStmt{
				Results: []ast.Expr{
					&ast.CallExpr{
						Fun: &ast.IndexListExpr{
							X: &ast.Ident{Name: "exec"},
							Indices: []ast.Expr{
								&ast.Ident{Name: "T"},
								&ast.Ident{Name: response},
							},
						},
						Args: []ast.Expr{
							&ast.Ident{Name: "u"},
							&ast.Ident{Name: "controller"},
							&ast.Ident{Name: "f"},
							&ast.Ident{Name: "timeout"},
						},
					},
				},
			},
		},
	}
}

func blankline() ast.Stmt {
	return &ast.ExprStmt{
		X: &ast.BasicLit{
			Kind: token.STRING,
		},
	}
}
