package codec

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"

	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"

	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
	"codegen/model"
)

func decoder() {
	const output = "generated.go"

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	decl := buildDecoder()

	// ... pretty print
	b := bytes.Buffer{}

	printer.Fprint(&b, token.NewFileSet(), decl)

	// ... 'generated code' warning
	writeln(f, "// generated code - ** DO NOT EDIT **")
	writeln(f, "")

	lines := strings.Split(b.String(), "\n")
	for _, line := range lines {
		// ... replace 'insert newline here'
		re := regexp.MustCompile(`"// -- insert newline here --"[,]?`)
		writeln(f, re.ReplaceAllString(line, "\n"))
	}

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
								Value: `"unknown message type (0x%02x)"`,
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
