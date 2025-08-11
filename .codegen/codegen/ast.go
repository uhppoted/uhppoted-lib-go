package codegen

import (
	"bytes"
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

		// // ... format message packets
		// re := regexp.MustCompile(`^    packet := \[\]byte\{(?:0x[0-9a-fA-F]{2})(, ?:0x[0-9a-fA-F]{2}){63})\}`)
		// for _, line := range lines {
		// 	if match := re.FindStringSubmatch(line); len(match) > 0 {
		// 		fmt.Printf(">>>> %v\n", line)
		// 		fmt.Printf(">>>> %v\n", match)
		// 	}
		// }

		cleaned := strings.Join(lines, "\n")

		if _, err = f.WriteString(cleaned); err != nil {
			return err
		}

		return nil
	}
}

func TitleCase(s string) string {
	re := regexp.MustCompile(`[ \-:]+`)
	parts := re.Split(s, -1)
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}

	return strings.Join(parts, "")
}
