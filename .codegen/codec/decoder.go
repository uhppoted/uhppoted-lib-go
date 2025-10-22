package codec

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"

	"bytes"
	"go/printer"
	"go/token"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"

	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
	"codegen/model"
)

func decoder() {
	outfile := filepath.Join("codec", "generated.go")
	decl := buildDecoder()

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
		log.Fatalf("error creating  file %s (%v)", outfile, err)
	} else {
		defer f.Close()

		writeln(f, "// generated code - ** DO NOT EDIT **")
		writeln(f, "")
		writeln(f, buf.String())

		f.Close()
	}
}

func buildDecoder() *dst.File {
	impl := []dst.Decl{
		&dst.GenDecl{
			Tok: token.IMPORT,
			Specs: []dst.Spec{
				&dst.ImportSpec{
					Path: &dst.BasicLit{
						Kind:  token.STRING,
						Value: `"fmt"`,
					},
				},
				&dst.ImportSpec{
					Path: &dst.BasicLit{
						Kind: token.STRING,
					},
				},
				&dst.ImportSpec{
					Name: dst.NewIdent("decoder"),
					Path: &dst.BasicLit{
						Kind:  token.STRING,
						Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/codec/decode"`,
					},
				},
			},
		},
	}

	impl = append(impl, buildDecoderFactoryFunc())

	return &dst.File{
		Name: dst.NewIdent("codec"),

		Imports: []*dst.ImportSpec{
			{
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"fmt"`,
				},
			},
			{
				Name: dst.NewIdent("decoder"),
				Path: &dst.BasicLit{
					Kind:  token.STRING,
					Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/codec/decode"`,
				},
			},
		},

		Decls: impl,
	}
}

func buildDecoderFactoryFunc() *dst.FuncDecl {
	return &dst.FuncDecl{
		Name: dst.NewIdent("decode"),
		Type: &dst.FuncType{
			Params: &dst.FieldList{
				List: []*dst.Field{
					{
						Names: []*dst.Ident{dst.NewIdent("packet")},
						Type: &dst.ArrayType{
							Elt: dst.NewIdent("byte"),
						},
					},
				},
			},
			Results: &dst.FieldList{
				List: []*dst.Field{
					{
						Type: dst.NewIdent("any"),
					},
					{
						Type: dst.NewIdent("error"),
					},
				},
			},
		},
		Body: buildDecoderFactoryBody(),
	}
}

func buildDecoderFactoryBody() *dst.BlockStmt {
	// switch packet[1] { ... }
	_switch := &dst.BlockStmt{
		List: []dst.Stmt{},
	}

	// ... message types
	excluded := []*lib.Response{
		&model.GetListenerAddrPortResponse,
		&model.SetListenerAddrPortResponse,
	}

	for _, response := range model.Responses {
		if slices.Contains(excluded, response) {
			log.Printf("skipping %v (excluded)", response.Name)
			continue
		}

		name := fmt.Sprintf("%v", codegen.TitleCase(response.Message.Name))

		clause := dst.CaseClause{
			List: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.INT,
					Value: fmt.Sprintf("0x%02x", response.Message.MsgType),
				},
			},
			Body: []dst.Stmt{
				// return decode.<XXX>(packet)
				&dst.ReturnStmt{
					Results: []dst.Expr{
						&dst.CallExpr{
							Fun: &dst.SelectorExpr{
								X:   dst.NewIdent("decoder"),
								Sel: dst.NewIdent(name),
							},
							Args: []dst.Expr{dst.NewIdent("packet")},
						},
					},
				},
				// blank line
				&dst.ExprStmt{
					X: &dst.BasicLit{
						Kind: token.STRING,
					},
				},
			},
		}

		_switch.List = append(_switch.List, &clause)
	}

	// ... default
	_switch.List = append(_switch.List, &dst.CaseClause{
		List: nil,
		Body: []dst.Stmt{
			&dst.ReturnStmt{
				Results: []dst.Expr{
					dst.NewIdent("nil"),
					&dst.CallExpr{
						Fun: &dst.SelectorExpr{
							X:   dst.NewIdent("fmt"),
							Sel: dst.NewIdent("Errorf"),
						},
						Args: []dst.Expr{
							&dst.BasicLit{
								Kind:  token.STRING,
								Value: `"unknown message type (0x%02x)"`,
							},
							&dst.IndexExpr{
								X:     dst.NewIdent("packet"),
								Index: &dst.BasicLit{Kind: token.INT, Value: "1"},
							},
						},
					},
				},
			},
		},
	})

	return &dst.BlockStmt{
		List: []dst.Stmt{
			// switch packet[1] { ... }
			&dst.SwitchStmt{
				Tag: &dst.IndexExpr{
					X:     dst.NewIdent("packet"),
					Index: &dst.BasicLit{Kind: token.INT, Value: "1"},
				},
				Body: _switch,
			},
		},
	}
}
