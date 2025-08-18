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
	file := filepath.Join("types", "_structs.go")

	imports := []string{
		"net/netip",
		"time",

		"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/encode",
	}

	types := []*ast.GenDecl{}
	functions := []*ast.FuncDecl{}

	for _, f := range model.Responses[:1] {
		types = append(types, typedef(f))
	}

	AST := codegen.NewAST("uhppoted", imports, types, functions)

	if err := AST.Generate(file); err != nil {
		log.Fatalf("error generating %v (%v)", file, err)
	} else {
		log.Printf("... generated %s", filepath.Base(file))
	}
}

func typedef(r types.Response) *ast.GenDecl {
	name := codegen.TitleCase(r.Name)
	fields := []*ast.Field{}

	for _, f := range r.Fields {
		ident := ast.NewIdent(codegen.TitleCase(f.Name))
		tag := fmt.Sprintf(`json:"%v"`, f.Tag)

		field := ast.Field{
			Names: []*ast.Ident{ident},
			Type:  ast.NewIdent("uint32"),
			Tag:   &ast.BasicLit{Kind: token.STRING, Value: tag},
		}

		fields = append(fields, &field)
	}

	decl := ast.GenDecl{
		Tok: token.TYPE,
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

// &ast.GenDecl{
//     Tok: token.TYPE,
//     Specs: []ast.Spec{
//         &ast.TypeSpec{
//             Name: ast.NewIdent("GetControllerResponse"),
//             Type: &ast.StructType{
//                 Fields: &ast.FieldList{
//                     List: []*ast.Field{
//                         {
//                             Names: []*ast.Ident{ast.NewIdent("Controller")},
//                             Type:  ast.NewIdent("uint32"),
//                             Tag:   &ast.BasicLit{Kind: token.STRING, Value: "`json:\"controller\"`"},
//                         },
//                         {
//                             Names: []*ast.Ident{ast.NewIdent("IpAddress")},
//                             Type: &ast.SelectorExpr{
//                                 X:   ast.NewIdent("netip"),
//                                 Sel: ast.NewIdent("Addr"),
//                             },
//                             Tag: &ast.BasicLit{Kind: token.STRING, Value: "`json:\"ip-address\"`"},
//                         },
//                         {
//                             Names: []*ast.Ident{ast.NewIdent("SubnetMask")},
//                             Type: &ast.SelectorExpr{
//                                 X:   ast.NewIdent("netip"),
//                                 Sel: ast.NewIdent("Addr"),
//                             },
//                             Tag: &ast.BasicLit{Kind: token.STRING, Value: "`json:\"subnet-mask\"`"},
//                         },
//                         {
//                             Names: []*ast.Ident{ast.NewIdent("Gateway")},
//                             Type: &ast.SelectorExpr{
//                                 X:   ast.NewIdent("netip"),
//                                 Sel: ast.NewIdent("Addr"),
//                             },
//                             Tag: &ast.BasicLit{Kind: token.STRING, Value: "`json:\"gateway\"`"},
//                         },
//                         {
//                             Names: []*ast.Ident{ast.NewIdent("MACAddress")},
//                             Type:  ast.NewIdent("string"),
//                             Tag:   &ast.BasicLit{Kind: token.STRING, Value: "`json:\"MAC-address\"`"},
//                         },
//                         {
//                             Names: []*ast.Ident{ast.NewIdent("Version")},
//                             Type:  ast.NewIdent("string"),
//                             Tag:   &ast.BasicLit{Kind: token.STRING, Value: "`json:\"version\"`"},
//                         },
//                         {
//                             Names: []*ast.Ident{ast.NewIdent("Date")},
//                             Type: &ast.SelectorExpr{
//                                 X:   ast.NewIdent("time"),
//                                 Sel: ast.NewIdent("Time"),
//                             },
//                             Tag: &ast.BasicLit{Kind: token.STRING, Value: "`json:\"date\"`"},
//                         },
//                     },
//                 },
//             },
//         },
//     },
// }
