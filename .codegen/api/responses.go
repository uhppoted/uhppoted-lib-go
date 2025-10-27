package api

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"go/printer"
	"go/token"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"

	"github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
	"codegen/model"
)

func Responses() {
	outfile := filepath.Join("responses", "generated.go")
	decl := buildResponses()

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

	// ... write to file
	if f, err := os.Create(outfile); err != nil {
		log.Fatalf("error creating file %s (%v)", outfile, err)
	} else {
		defer f.Close()

		writeln(f, "// generated code - ** DO NOT EDIT **")
		writeln(f, "")
		writeln(f, buf.String())
	}
}

func buildResponses() *dst.File {
	impl := []dst.Decl{
		&dst.GenDecl{
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
						Kind: token.STRING,
					},
				},
				&dst.ImportSpec{
					Path: &dst.BasicLit{
						Kind:  token.STRING,
						Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"`,
					},
				},
			},
		},
	}

	for _, response := range model.Responses {
		if t := typedef(*response); t != nil {
			impl = append(impl, t)
		}
	}

	if t := typedef(model.ListenerEvent); t != nil {
		impl = append(impl, t)
	}

	return &dst.File{
		Name: dst.NewIdent("responses"),

		// Imports: []*dst.ImportSpec{
		// 	{
		// 		Path: &dst.BasicLit{
		// 			Kind:  token.STRING,
		// 			Value: `"net/netip"`,
		// 		},
		// 	},
		// 	{
		// 		Path: &dst.BasicLit{
		// 			Kind:  token.STRING,
		// 			Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"`,
		// 		},
		// 	},
		// },

		Decls: impl,
	}
}

func typedef(r types.Response) *dst.GenDecl {
	name := strings.TrimSuffix(codegen.TitleCase(r.Name), "Response")
	// description := godoc(r)
	fields := []*dst.Field{}

	for _, f := range r.Fields {
		ident := dst.NewIdent(codegen.TitleCase(f.Name))
		ftype := gotype(f)
		tag := fmt.Sprintf("`json:%v`", gotag(f))

		field := dst.Field{
			Names: []*dst.Ident{ident},
			Type:  dst.NewIdent(ftype),
			Tag:   &dst.BasicLit{Kind: token.STRING, Value: tag},
		}

		fields = append(fields, &field)
	}

	decl := dst.GenDecl{
		Tok: token.TYPE,
		Specs: []dst.Spec{
			&dst.TypeSpec{
				Name: dst.NewIdent(name),
				Type: &dst.StructType{
					Fields: &dst.FieldList{
						List: fields,
					},
				},
			},
		},
	}

	// godoc
	for _, line := range r.Description {
		decl.Decs.Start.Append(fmt.Sprintf("// %v", line))
	}

	decl.Decs.Before = dst.EmptyLine

	return &decl
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
		return "types.DateTime"

	case "optional datetime":
		return "types.DateTime"

	case "date":
		return "types.Date"

	case "shortdate":
		return "types.Date"

	case "optional date":
		return "types.Date"

	case "time":
		return "types.Time"

	case "HHmm":
		return "types.HHmm"

	case "IPv4":
		return "netip.Addr"

	case "address:port":
		return "netip.AddrPort"

	case "MAC":
		return "string"

	case "pin":
		return "uint32"

	case "mode":
		return "types.DoorMode"

	case "event-type":
		return "types.EventType"

	case "direction":
		return "types.Direction"

	case "reason":
		return "types.Reason"

	case "version":
		return "string"

	default:
		panic(fmt.Sprintf("unknown response field type (%v)", field.Type))
	}
}

func gotag(field types.Field) string {
	if field.Tag != "" {
		return fmt.Sprintf(`"%v"`, field.Tag)
	}

	return `"-"`
}
