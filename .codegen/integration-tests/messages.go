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

	"codegen/model"
)

func messages() {
	outfile := filepath.Join(".", "messages.go")
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
	file := &dst.File{
		Name: dst.NewIdent("uhppoted"),
		Decls: []dst.Decl{
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

	messages = append(messages, buildInvalidResponseMessage())

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
	responses := []dst.Expr{}
	for _, reply := range test.Replies {
		responses = append(responses, buildResponseMessage(reply.Message))
	}

	return &dst.CompositeLit{
		Elts: []dst.Expr{
			buildRequestMessage(test.Request),

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

// invalid response
func buildInvalidResponseMessage() dst.Expr {
	comment := "// invalid response"

	request := []byte{
		0x17, 0x94, 0x00, 0x00, 0x90, 0x53, 0xfb, 0x0b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	replies := [][]byte{
		[]byte{
			0x17, 0x94, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0xc0, 0xa8, 0x01, 0x64, 0xff, 0xff, 0xff, 0x00,
			0xc0, 0xa8, 0x01, 0x01, 0x00, 0x12, 0x23, 0x34, 0x45, 0x56, 0x08, 0x92, 0x20, 0x18, 0x11, 0x05,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
	}

	responses := []dst.Expr{}
	for _, reply := range replies {
		responses = append(responses, buildResponseMessage(reply))
	}

	return &dst.CompositeLit{
		Elts: []dst.Expr{
			buildRequestMessage(request),

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

// Request: []byte{ ... }
func buildRequestMessage(bytes []byte) dst.Expr {
	request := make([]dst.Expr, 64)

	for i, b := range bytes {
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

	return &dst.KeyValueExpr{
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
	}
}

// []byte{ ... }
func buildResponseMessage(reply []byte) dst.Expr {
	response := make([]dst.Expr, 64)
	for i, b := range reply {
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

	return &dst.CompositeLit{
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
	}
}
