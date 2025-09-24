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
		cfg := &printer.Config{
			Mode:     printer.UseSpaces,
			Tabwidth: 4,
		}

		// printer.Fprint(&b, fileset, a.file)
		cfg.Fprint(&b, fileset, a.file)

		// ... remove 'lines intentionally left blank'
		lines := strings.Split(b.String(), "\n")
		out := []string{}
		for _, line := range lines {
			if strings.HasPrefix(line, "// -- line intentionally left blank --") {
				out = append(out, "")
				continue
			}

			// ... reformat response struct (ewwwwww :-()
			re := regexp.MustCompile(`^(\s*return\s+responses\.(?:.*?)Response\s*\{)([^}]*)(\}.*)`)
			if match := re.FindStringSubmatch(line); len(match) == 4 {
				out = append(out, match[1])

				fields := regexp.MustCompile(`(.*?):\s*(.*?)(\(.*?\))(?:,\s*)?`).FindAllStringSubmatch(match[2], -1)
				for _, f := range fields {
					out = append(out, fmt.Sprintf("        %v: %v%v,", f[1], f[2], f[3]))
				}

				out = append(out, "    "+match[3])
				continue
			}

			// ... nothing to do
			out = append(out, line)
		}

		cleaned := strings.Join(out, "\n")

		if _, err = f.WriteString("// generated code - ** DO NOT EDIT **\n\n"); err != nil {
			return err
		}

		if _, err = f.WriteString(cleaned); err != nil {
			return err
		}

		return nil
	}
}
