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

func Structs() {
	file := filepath.Join("types", "generated.go")

	imports := []string{
		"net/netip",
		"time",
	}

	structs := []types.Response{
		model.GetControllerResponse,
		model.SetIPv4Response,
		model.GetTimeResponse,
		model.SetTimeResponse,
		model.GetListenerResponse,
		model.SetListenerResponse,
		model.GetListenerAddrPortResponse,
		model.SetListenerAddrPortResponse,
		model.GetDoorResponse,
		model.SetDoorResponse,
		model.SetDoorPasscodesResponse,
		model.OpenDoorResponse,
		// GetStatus,
		model.GetCardsResponse,
		// GetCard,
		// GetCardAtIndex,
		// PutCard,
		// DeleteCard,
		// DeleteAllCards,
		model.GetEventResponse,
		model.GetEventIndexResponse,
		model.SetEventIndexResponse,
		model.RecordSpecialEventsResponse,
		model.GetTimeProfileResponse,
		model.SetTimeProfileResponse,
		model.ClearTimeProfilesResponse,
		model.AddTaskResponse,
	}

	types := []*ast.GenDecl{}
	functions := []*ast.FuncDecl{}

	for _, f := range structs {
		types = append(types, typedef(f))
	}

	AST := codegen.NewAST("types", imports, types, functions)

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

	case "date":
		return "time.Time"

	case "optional date":
		return "time.Time"

	case "HHmm":
		return "time.Time"

	case "IPv4":
		return "netip.Addr"

	case "address:port":
		return "netip.AddrPort"

	case "MAC":
		return "string"

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
