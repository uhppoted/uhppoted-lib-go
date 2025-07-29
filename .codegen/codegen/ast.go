package codegen

import (
	"os"

	"go/ast"
	"go/printer"
	"go/token"
)

type AST struct {
	file *ast.File
}

func NewAST(pkg string) AST {
	return AST{
		file: &ast.File{
			Name:    ast.NewIdent(pkg),
			Imports: []*ast.ImportSpec{},
			Decls:   []ast.Decl{},
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
