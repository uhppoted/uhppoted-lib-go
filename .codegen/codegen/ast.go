package codegen

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"go/ast"
	"go/printer"
	"go/token"
)

type AST struct {
	file *ast.File
}

func NewAST(pkg string, imports [][]string, types []*ast.GenDecl, functions []*ast.FuncDecl) AST {
	// ... imports
	imported := []*ast.ImportSpec{}

	for i, u := range imports {
		if i > 0 {
			imported = append(imported, &ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind: token.STRING,
				},
			})
		}

		for _, v := range u {
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

	// ... types
	for _, t := range types {
		decls = append(decls, t)
	}

	// ... functions
	for _, f := range functions {
		decls = append(decls, f)
	}

	// ... 'k, all done
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

		b := bytes.Buffer{}
		fileset := token.NewFileSet()

		printer.Fprint(&b, fileset, a.file)

		// ... remove 'lines intentionally left blank'
		lines := strings.Split(b.String(), "\n")
		for i, line := range lines {
			if strings.HasPrefix(line, "// -- line intentionally left blank --") {
				lines[i] = ""
			}
		}

		cleaned := strings.Join(lines, "\n")

		if _, err = f.WriteString(cleaned); err != nil {
			return err
		}

		return nil
	}
}
