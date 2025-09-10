package codec

import (
	"log"
	"path/filepath"

	"go/ast"
	"go/token"

	"codegen/codegen"
)

func decodeAST() {
	// const output = "decode/_decode.go"
	//
	// f, err := os.Create(output)
	// if err != nil {
	// 	log.Fatalf("Failed to create file %s: %v", output, err)
	// }
	// defer f.Close()
	//
	// decl := buildDecode()
	//
	// printer.Fprint(f, token.NewFileSet(), decl)
	//
	// f.Close()

	file := filepath.Join("decode", "_decode.go")

	imports := [][]string{
		[]string{
			"fmt",
		},
		[]string{
			"github.com/uhppoted/uhppoted-lib-go/uhppoted/responses",
		},
	}

	types := []*ast.GenDecl{}
	functions := []*ast.FuncDecl{}

	// for _, f := range model.Responses {
	// 	types = append(types, typedef(*f))
	// }

	AST := codegen.NewAST("responses", imports, types, functions)

	if err := AST.Generate(file); err != nil {
		log.Fatalf("error generating %v (%v)", file, err)
	} else {
		log.Printf("... generated %s", filepath.Base(file))
	}

}

func buildDecode() *ast.File {
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
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/uhppoted/responses"`,
				},
			},
		},

		Decls: []ast.Decl{
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
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: `"github.com/uhppoted/uhppoted-lib-go/uhppoted/responses"`,
						},
					},
				},
			},
			// buildDecoderFunc(),
			// buildDecoderFactoryFunc(),
		},
	}
}
