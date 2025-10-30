package codec

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

	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
	"codegen/model"
)

func encodeAST() {
	outfile := filepath.Join("codec", "encode", "_generated.go")
	decl := buildEncode()

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

		f.Close()
	}
}

func buildEncode() *dst.File {
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
					Path: &dst.BasicLit{
						Kind:  token.STRING,
						Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"`,
					},
				},
			},
		},
	}

	for _, request := range model.Requests {
		if f := buildEncodeFunc(request); f != nil {
			impl = append(impl, f)
		}
	}

	return &dst.File{
		Name:  dst.NewIdent("encode"),
		Decls: impl,
	}
}

func buildEncodeFunc(r lib.Request) *dst.FuncDecl {
	name := strings.TrimSuffix(fmt.Sprintf("%v", codegen.TitleCase(r.Name)), "Request")

	params := dst.FieldList{
		List: []*dst.Field{
			// {
			// 	Names: []*dst.Ident{dst.NewIdent("packet")},
			// 	Type: &dst.ArrayType{
			// 		Elt: dst.NewIdent("byte"),
			// 	},
			// },
		},
	}

	//	packet := make([]byte, 64)
	makePacket := dst.AssignStmt{
		Lhs: []dst.Expr{
			&dst.Ident{
				Name: "packet",
			},
		},
		Tok: token.DEFINE, // ":="
		Rhs: []dst.Expr{
			&dst.CallExpr{
				Fun: &dst.Ident{
					Name: "make",
				},
				Args: []dst.Expr{
					&dst.ArrayType{
						Len: nil,
						Elt: &dst.Ident{
							Name: "byte",
						},
					},
					&dst.BasicLit{
						Kind:  token.INT,
						Value: "64",
					},
				},
			},
		},
	}

	makePacket.Decs.Before = dst.NewLine
	makePacket.Decs.After = dst.EmptyLine

	// packet[0] = SOM
	SOM := dst.AssignStmt{
		Lhs: []dst.Expr{
			&dst.IndexExpr{
				X: &dst.Ident{
					Name: "packet",
				},
				Index: &dst.BasicLit{
					Kind:  token.INT,
					Value: "0",
				},
			},
		},
		Tok: token.ASSIGN,
		Rhs: []dst.Expr{
			&dst.Ident{
				Name: "SOM",
			},
		},
	}

	// packet[1] = <msgtype>
	msgType := dst.AssignStmt{
		Lhs: []dst.Expr{
			&dst.IndexExpr{
				X: &dst.Ident{
					Name: "packet",
				},
				Index: &dst.BasicLit{
					Kind:  token.INT,
					Value: "1",
				},
			},
		},
		Tok: token.ASSIGN,
		Rhs: []dst.Expr{
			&dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf("0x%02x", r.MsgType)},
		},
	}

	msgType.Decs.After = dst.EmptyLine

	results := dst.FieldList{
		List: []*dst.Field{
			{
				Type: &dst.ArrayType{
					Elt: dst.NewIdent("byte"),
				},
			},
			{
				Type: dst.NewIdent("error"),
			},
		},
	}

	// request := []dst.Expr{}
	// request = append(request, makePacket)
	// for _, f := range r.Fields {
	// 	request = append(request, pack(f))
	// }

	// ... body
	body := dst.BlockStmt{
		List: []dst.Stmt{
			&makePacket,
			&SOM,
			&msgType,

			// return packet,nil
			&dst.ReturnStmt{
				Results: []dst.Expr{
					&dst.Ident{
						Name: "packet",
					},
					&dst.Ident{
						Name: "nil",
					},
				},
			},
		},
	}

	f := dst.FuncDecl{
		Name: dst.NewIdent(name),
		Type: &dst.FuncType{
			Params:  &params,
			Results: &results,
		},
		Body: &body,
	}

	f.Decs.Before = dst.EmptyLine
	f.Decs.After = dst.EmptyLine

	// godoc
	f.Decs.Start.Append(fmt.Sprintf("// Encodes a %v request to a 64 byte packet.", name))

	return &f
}
