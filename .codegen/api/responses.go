package api

import (
	"fmt"
	"log"
	"path/filepath"

	"go/ast"
	"go/token"

	"github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
	"codegen/model"
)

func Responses() {
	file := filepath.Join("responses", "generated.go")

	imports := []string{
		"net/netip",
		"time",
		"github.com/uhppoted/uhppoted-lib-go/uhppoted/entities",
	}

	types := []*ast.GenDecl{}
	functions := []*ast.FuncDecl{}

	for _, f := range model.Responses {
		types = append(types, typedef(*f))
	}

	AST := codegen.NewAST("responses", imports, types, functions)

	if err := AST.Generate(file); err != nil {
		log.Fatalf("error generating %v (%v)", file, err)
	} else {
		log.Printf("... generated %s", filepath.Base(file))
	}
}

func typedef(r types.Response) *ast.GenDecl {
	name := codegen.TitleCase(r.Name)
	description := godoc(r)
	fields := []*ast.Field{}

	for _, f := range r.Fields {
		ident := ast.NewIdent(codegen.TitleCase(f.Name))
		ftype := gotype(f)
		tag := fmt.Sprintf("`json:%v`", gotag(f))

		field := ast.Field{
			Names: []*ast.Ident{ident},
			Type:  ast.NewIdent(ftype),
			Tag:   &ast.BasicLit{Kind: token.STRING, Value: tag},
		}

		fields = append(fields, &field)
	}

	decl := ast.GenDecl{
		Tok: token.TYPE,
		Doc: &ast.CommentGroup{
			List: description,
		},
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(name),
				Type: &ast.StructType{
					Fields: &ast.FieldList{
						List: fields,
					},
				},
			},
		},
	}

	return &decl
}

func godoc(r types.Response) []*ast.Comment {
	doc := []*ast.Comment{
		{Text: fmt.Sprintf("// -- line intentionally left blank --")},
	}

	for _, line := range r.Description {
		text := fmt.Sprintf("// %v", line)
		comment := ast.Comment{
			Text: text,
		}

		doc = append(doc, &comment)
	}

	return doc
}

func gotype(field types.Field) string {
	switch field.Type {
	case "bool":
		return "bool"

	case "uint8":
		return "uint8"

	case "uint16":
		return "uint16"

	case "uint32":
		return "uint32"

	case "datetime":
		return "time.Time"

	case "optional datetime":
		return "time.Time"

	// case "date":
	// 	return "time.Time"
	//
	// case "shortdate":
	// 	return "time.Time"
	//
	// case "optional date":
	// 	return "time.Time"
	//
	// case "time":
	// 	return "time.Time"

	case "date":
		return "entities.Date"

	case "shortdate":
		return "entities.Date"

	case "optional date":
		return "entities.Date"

	case "time":
		return "entities.Time"

	case "HHmm":
		return "time.Time"

	case "IPv4":
		return "netip.Addr"

	case "address:port":
		return "netip.AddrPort"

	case "MAC":
		return "string"

	case "pin":
		return "uint32"

	case "version":
		return "string"

	default:
		return "unknown"
	}
}

func gotag(field types.Field) string {
	if field.Tag != "" {
		return fmt.Sprintf(`"%v"`, field.Tag)
	}

	return `"-"`
}
