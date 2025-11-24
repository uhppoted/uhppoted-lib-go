package integration_tests

import (
	"bytes"
	"log"
	"os"
	"path/filepath"

	"go/printer"
	"go/token"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"

	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/model"
)

func broadcastAST() {
	outfile := filepath.Join(".", "default", "_api_test.go")
	decl := buildBroadcast()

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

	// ... header comment
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

func buildBroadcast() *dst.File {
	imports := &dst.GenDecl{
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
					Value: `"reflect"`,
				},
			},

			&dst.ImportSpec{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"testing"`,
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
					Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted"`,
				},
			},
			&dst.ImportSpec{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"`,
				},
			},

			&dst.ImportSpec{
				Path: &dst.BasicLit{
					Kind: token.STRING,
				},
			},

			&dst.ImportSpec{
				Name: dst.NewIdent("tests"),
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/integration-tests"`,
				},
			},
		},
	}

	tests := []dst.Decl{}

	for _, fn := range model.API {
		for _, test := range fn.Tests {
			if test := buildBroadcastTest(*fn, test); test != nil {
				tests = append(tests, test)
			}
		}
	}

	decls := []dst.Decl{}
	decls = append(decls, imports)
	decls = append(decls, tests...)

	file := &dst.File{
		Name:  dst.NewIdent("uhppoted"),
		Decls: decls,
	}

	return file
}

func buildBroadcastTest(fn lib.Function, test lib.FuncTest) dst.Decl {
	return nil
}
