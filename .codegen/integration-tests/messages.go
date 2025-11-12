package integration_tests

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"go/printer"
	"go/token"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	// "codegen/codegen"
	"codegen/model"
)

func messagesAST() {
	outfile := filepath.Join(".", "_messages.go")
	decl := buildMessages()

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

func buildMessages() *dst.File {
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
		},
	}

	file := &dst.File{
		Name: dst.NewIdent("uhppoted"),
		Decls: []dst.Decl{
			imports,
			buildMessageStruct(),
			buildMessagesList(),
		},
	}

	return file
}

//	type message = struct {
//		Request  []byte
//		Response [][]byte
//	}
func buildMessageStruct() dst.Decl {
	return &dst.GenDecl{
		Tok: token.TYPE,
		Specs: []dst.Spec{
			&dst.TypeSpec{
				Name:   dst.NewIdent("message"),
				Assign: true,
				Type: &dst.StructType{
					Fields: &dst.FieldList{
						List: []*dst.Field{
							{
								Names: []*dst.Ident{dst.NewIdent("Request")},
								Type: &dst.ArrayType{
									Elt: dst.NewIdent("byte"), // []byte
								},
							},
							{
								Names: []*dst.Ident{dst.NewIdent("Response")},
								Type: &dst.ArrayType{
									Elt: &dst.ArrayType{
										Elt: dst.NewIdent("byte"), // [][]byte
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// var Messages = []message{...}
func buildMessagesList() dst.Decl {
	messages := []dst.Expr{}

	for _, f := range model.API {
		for i, test := range f.Tests {
			comment := ""
			if i == 0 {
				comment = fmt.Sprintf("// %v", f.Name)
			}

			messages = append(messages, buildMessage(test, comment))
		}
	}

	return &dst.GenDecl{
		Tok: token.VAR,
		Specs: []dst.Spec{
			&dst.ValueSpec{
				Names: []*dst.Ident{
					dst.NewIdent("Messages"),
				},
				Values: []dst.Expr{
					&dst.CompositeLit{
						Type: &dst.ArrayType{
							Elt: dst.NewIdent("message"),
						},
						Elts: messages,
					},
				},
			},
		},
	}
}

//	{
//		Request: []byte{
//			0x17, 0x94, ...
//		},
//		Response: [][]byte{
//			[]byte{
//				0x17, 0x94, ...
//			},
//		},
//	},
func buildMessage(test lib.FuncTest, comment string) dst.Expr {
	request := make([]dst.Expr, 64)
	for i, b := range test.Request {
		xx := &dst.BasicLit{
			Kind:  token.INT,
			Value: fmt.Sprintf("0x%02x", b),
		}

		if i%16 == 0 {
			xx.Decs.Before = dst.NewLine
		}

		if i == 63 {
			xx.Decs.After = dst.NewLine
		}

		request[i] = xx
	}

	responses := []dst.Expr{}
	for _, reply := range test.Replies {
		response := make([]dst.Expr, 64)
		for i, b := range reply.Message {
			xx := &dst.BasicLit{
				Kind:  token.INT,
				Value: fmt.Sprintf("0x%02x", b),
			}

			if i%16 == 0 {
				xx.Decs.Before = dst.NewLine
			}

			if i == 63 {
				xx.Decs.After = dst.NewLine
			}

			response[i] = xx
		}

		responses = append(responses, &dst.CompositeLit{
			Type: &dst.ArrayType{
				Elt: dst.NewIdent("byte"),
			},
			Elts: response,
			Decs: dst.CompositeLitDecorations{
				NodeDecs: dst.NodeDecs{
					Before: dst.NewLine,
					After:  dst.NewLine,
				},
			},
		})
	}

	return &dst.CompositeLit{
		Elts: []dst.Expr{
			&dst.KeyValueExpr{
				Key: dst.NewIdent("Request"),
				Value: &dst.CompositeLit{
					Type: &dst.ArrayType{
						Elt: dst.NewIdent("byte"),
					},
					Elts: request,
				},
				Decs: dst.KeyValueExprDecorations{
					NodeDecs: dst.NodeDecs{
						Before: dst.NewLine,
					},
				},
			},

			// Response: [][]byte{ ... }
			&dst.KeyValueExpr{
				Key: dst.NewIdent("Response"),
				Value: &dst.CompositeLit{
					Type: &dst.ArrayType{
						Elt: &dst.ArrayType{
							Elt: dst.NewIdent("byte"),
						},
					},
					Elts: responses,
				},

				Decs: dst.KeyValueExprDecorations{
					NodeDecs: dst.NodeDecs{
						Before: dst.NewLine,
						After:  dst.NewLine,
					},
				},
			},
		},

		Decs: dst.CompositeLitDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.NewLine,
				After:  dst.EmptyLine,
				Start:  []string{comment},
			},
		},
	}
}
