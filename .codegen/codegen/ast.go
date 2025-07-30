package codegen

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"go/ast"
	"go/printer"
	"go/token"
)

type AST struct {
	file *ast.File
}

func NewAST(pkg string, imports []string, functions []*ast.FuncDecl) AST {
	imported := []*ast.ImportSpec{}

	for _, v := range imports {
		if v == "" {
			imported = append(imported, &ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind: token.STRING,
				},
			})
		} else {
			imported = append(imported, &ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, v),
				},
			})
		}
	}

	decls := []ast.Decl{
		&ast.GenDecl{
			Tok:   token.IMPORT,
			Specs: []ast.Spec{},
		},
	}

	for _, v := range imported {
		if g, ok := decls[0].(*ast.GenDecl); ok {
			g.Specs = append(g.Specs, v)
		}
	}

	for _, f := range functions {
		decls = append(decls, f)
	}

	return AST{
		file: &ast.File{
			Name:    ast.NewIdent(pkg),
			Imports: imported,
			Decls:   decls,
		},
	}
}

func (a AST) Generate(file string) error {
	if f, err := os.Create(file); err != nil {
		return err
	} else {
		defer f.Close()

		fileset := token.NewFileSet()

		printer.Fprint(f, fileset, a.file)

		return nil
	}
}

func TitleCase(s string) string {
	re := regexp.MustCompile(`[ -]+`)
	parts := re.Split(s, -1)
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}

	return strings.Join(parts, "")
}
