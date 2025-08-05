package api

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"

	"go/ast"
	"go/token"

	"codegen/codegen"
	"codegen/model"
)

func API() {
	const file = "generated.go"

	imports := []string{
		"time",
		"",
		"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/encode",
	}

	f := []*ast.FuncDecl{
		function(model.GetController),
		// function(model.SetIPv4),
		function(model.GetTime),
		// function(model.SetTime),
		function(model.GetListener),
		// function(model.SetListener),
		function(model.GetDoor),
		function(model.SetDoor),
		function(model.SetDoorPasscodes),
		function(model.OpenDoor),
		function(model.GetStatus),
		function(model.GetCards),
	}

	AST := codegen.NewAST("uhppoted", imports, f)

	if err := AST.Generate(file); err != nil {
		log.Fatalf("error generating %v (%v)", file, err)
	} else {
		log.Printf("... generated %s", filepath.Base(file))
	}
}

func function(f model.Func) *ast.FuncDecl {
	name := codegen.TitleCase(f.Name)
	response := fmt.Sprintf("%vResponse", codegen.TitleCase(f.Response.Name))

	// ... args
	args := []*ast.Field{}
	args = append(args, &ast.Field{
		Names: []*ast.Ident{
			{Name: "u"},
		},
		Type: &ast.Ident{Name: "Uhppoted"},
	})

	args = append(args, &ast.Field{
		Names: []*ast.Ident{
			{Name: "controller"},
		},
		Type: &ast.Ident{Name: "T"},
	})

	for _, arg := range f.Request.Fields[1:] {
		name := regexp.MustCompile(`\s+`).ReplaceAllString(arg.Name, "")

		args = append(args, &ast.Field{
			Names: []*ast.Ident{
				{Name: name},
			},
			Type: &ast.Ident{Name: arg.Type},
		})
	}

	args = append(args, &ast.Field{
		Names: []*ast.Ident{
			{Name: "timeout"},
		},
		Type: &ast.Ident{Name: "time.Duration"},
	})

	// ... godoc
	godoc := ast.CommentGroup{
		List: []*ast.Comment{
			{Text: fmt.Sprintf("// -- line intentionally left blank --")},
			{Text: fmt.Sprintf("// %v", f.Description)},
		},
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
		Doc:  &godoc,
	}
}

func impl(f model.Func) *ast.BlockStmt {
	request := codegen.TitleCase(f.Request.Name)
	response := fmt.Sprintf("%vResponse", codegen.TitleCase(f.Response.Name))

	args := []ast.Expr{
		&ast.Ident{Name: "id"},
	}

	for _, arg := range f.Request.Fields[1:] {
		name := regexp.MustCompile(`\s+`).ReplaceAllString(arg.Name, "")

		args = append(args, &ast.Ident{Name: name})
	}

	return &ast.BlockStmt{
		List: []ast.Stmt{
			//	f := func(id uint32) ([]byte, error) {
			//       return encode.XXX(id,...)
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
