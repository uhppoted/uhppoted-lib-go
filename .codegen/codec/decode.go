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

func decode() {
	outfile := filepath.Join("codec", "decode", "generated.go")
	decl := buildDecode()

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

func buildDecode() *dst.File {
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
						Value: `"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"`,
					},
				},
			},
		},
	}

	for _, response := range model.Responses {
		if f := buildDecodeFunc(*response); f != nil {
			impl = append(impl, f)
		}
	}

	if f := buildDecodeFunc(model.ListenerEvent); f != nil {
		impl = append(impl, f)
	}

	return &dst.File{
		Name:  dst.NewIdent("decode"),
		Decls: impl,
	}
}

func buildDecodeFunc(r lib.Response) *dst.FuncDecl {
	name := fmt.Sprintf("%v", codegen.TitleCase(r.Name))
	returnType := strings.TrimSuffix(name, "Response")

	params := dst.FieldList{
		List: []*dst.Field{
			{
				Names: []*dst.Ident{dst.NewIdent("packet")},
				Type: &dst.ArrayType{
					Elt: dst.NewIdent("byte"),
				},
			},
		},
	}

	results := dst.FieldList{
		List: []*dst.Field{
			{
				Type: dst.NewIdent(fmt.Sprintf("responses.%v", returnType)),
			},
			{
				Type: dst.NewIdent("error"),
			},
		},
	}

	response := []dst.Expr{}
	for _, f := range r.Fields {
		response = append(response, unpack(f))
	}

	// ... body
	body := dst.BlockStmt{
		List: []dst.Stmt{
			ifPacketLengthNot64(returnType),
			ifPacketSOMNotValid(returnType),
			ifPacketFunctionCodeNotValid(r, returnType),

			// return responses.<T>{
			// 	...
			// }
			&dst.ReturnStmt{
				Results: []dst.Expr{
					&dst.CompositeLit{
						Type: &dst.SelectorExpr{
							X:   &dst.Ident{Name: "responses"},
							Sel: &dst.Ident{Name: strings.TrimSuffix(name, "Response")},
						},
						Elts: response,
					},
					&dst.Ident{Name: "nil"},
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
	f.Decs.Start.Append(fmt.Sprintf("// Decodes a %v from a 64 byte response packet.", name))

	return &f
}

//	if len(packet) != 64 {
//	    return responses.<R>{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
//	}
func ifPacketLengthNot64(returnType string) *dst.IfStmt {
	iff := dst.IfStmt{
		Cond: &dst.BinaryExpr{
			X: &dst.CallExpr{
				Fun:  &dst.Ident{Name: "len"},
				Args: []dst.Expr{&dst.Ident{Name: "packet"}},
			},
			Op: token.NEQ,
			Y:  &dst.BasicLit{Kind: token.INT, Value: "64"},
		},
		Body: &dst.BlockStmt{
			List: []dst.Stmt{
				&dst.ReturnStmt{
					Results: []dst.Expr{
						&dst.CompositeLit{
							Type: &dst.SelectorExpr{
								X:   &dst.Ident{Name: "responses"},
								Sel: &dst.Ident{Name: returnType},
							},
							Elts: nil,
						},
						&dst.CallExpr{
							Fun: &dst.Ident{Name: "fmt.Errorf"},
							Args: []dst.Expr{
								&dst.BasicLit{Kind: token.STRING, Value: `"invalid reply packet length (%v)"`},
								&dst.CallExpr{
									Fun:  &dst.Ident{Name: "len"},
									Args: []dst.Expr{&dst.Ident{Name: "packet"}},
								},
							},
						},
					},
				},
			},
		},
	}

	iff.Decs.After = dst.EmptyLine

	return &iff
}

// Ref. v6.62 firmware
//
//	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
//	   return responses.<R>{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
//	}
func ifPacketSOMNotValid(returnType string) *dst.IfStmt {
	iff := dst.IfStmt{
		Cond: &dst.BinaryExpr{
			Op: token.LAND, // &&
			X: &dst.BinaryExpr{
				Op: token.NEQ,
				X: &dst.IndexExpr{
					X:     &dst.Ident{Name: "packet"},
					Index: &dst.BasicLit{Kind: token.INT, Value: "0"},
				},
				Y: &dst.Ident{Name: "SOM"},
			},
			Y: &dst.BinaryExpr{
				Op: token.LOR, // ||
				X: &dst.BinaryExpr{
					Op: token.NEQ,
					X: &dst.IndexExpr{
						X:     &dst.Ident{Name: "packet"},
						Index: &dst.BasicLit{Kind: token.INT, Value: "0"},
					},
					Y: &dst.Ident{Name: "SOM_v6_62"},
				},
				Y: &dst.BinaryExpr{
					Op: token.NEQ,
					X: &dst.IndexExpr{
						X:     &dst.Ident{Name: "packet"},
						Index: &dst.BasicLit{Kind: token.INT, Value: "1"},
					},
					Y: &dst.BasicLit{Kind: token.INT, Value: "0x20"},
				},
			},
		},
		Body: &dst.BlockStmt{
			List: []dst.Stmt{
				&dst.ReturnStmt{
					Results: []dst.Expr{
						&dst.CompositeLit{
							Type: &dst.SelectorExpr{
								X:   &dst.Ident{Name: "responses"},
								Sel: &dst.Ident{Name: returnType},
							},
							Elts: nil,
						},
						&dst.CallExpr{
							Fun: &dst.Ident{Name: "fmt.Errorf"},
							Args: []dst.Expr{
								&dst.BasicLit{Kind: token.STRING, Value: `"invalid reply start of message byte (%02x)"`},
								&dst.IndexExpr{
									X:     &dst.Ident{Name: "packet"},
									Index: &dst.BasicLit{Kind: token.INT, Value: "0"},
								},
							},
						},
					},
				},
			},
		},
	}

	iff.Decs.After = dst.EmptyLine

	return &iff
}

//	if packet[1] != <code> {
//	    return responses.<R>{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
//	}
func ifPacketFunctionCodeNotValid(r lib.Response, returnType string) *dst.IfStmt {
	iff := dst.IfStmt{
		Cond: &dst.BinaryExpr{
			Op: token.NEQ,
			X: &dst.IndexExpr{
				X:     &dst.Ident{Name: "packet"},
				Index: &dst.BasicLit{Kind: token.INT, Value: "1"},
			},
			Y: &dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf("0x%02x", r.MsgType)},
		},
		Body: &dst.BlockStmt{
			List: []dst.Stmt{
				&dst.ReturnStmt{
					Results: []dst.Expr{
						&dst.CompositeLit{
							Type: &dst.SelectorExpr{
								X:   &dst.Ident{Name: "responses"},
								Sel: &dst.Ident{Name: returnType},
							},
							Elts: nil,
						},
						&dst.CallExpr{
							Fun: &dst.Ident{Name: "fmt.Errorf"},
							Args: []dst.Expr{
								&dst.BasicLit{Kind: token.STRING, Value: `"invalid reply function code (%02x)"`},
								&dst.IndexExpr{
									X:     &dst.Ident{Name: "packet"},
									Index: &dst.BasicLit{Kind: token.INT, Value: "1"},
								},
							},
						},
					},
				},
			},
		},
	}

	iff.Decs.After = dst.EmptyLine

	return &iff
}

func unpack(field lib.Field) dst.Expr {
	types := map[string]string{
		"bool":              "unpackBool",
		"uint8":             "unpackUint8",
		"uint16":            "unpackUint16",
		"uint32":            "unpackUint32",
		"datetime":          "unpackDateTime",
		"optional datetime": "unpackOptionalDateTime",
		"date":              "unpackDate",
		"shortdate":         "unpackShortDate",
		"optional date":     "unpackOptionalDate",
		"time":              "unpackTime",
		"HHmm":              "unpackHHmm",
		"IPv4":              "unpackIPv4",
		"address:port":      "unpackAddrPort",
		"MAC":               "unpackMAC",
		"pin":               "unpackPIN",
		"mode":              "unpackMode",
		"event-type":        "unpackEventType",
		"direction":         "unpackDirection",
		"reason":            "unpackReason",
		"version":           "unpackVersion",
	}

	if f, ok := types[field.Type]; !ok {
		panic(fmt.Sprintf("unknown response field type (%v)", field.Type))
	} else {
		kv := dst.KeyValueExpr{
			Key: &dst.Ident{
				Name: codegen.TitleCase(field.Name),
			},
			Value: &dst.CallExpr{
				Fun: &dst.Ident{Name: f},
				Args: []dst.Expr{
					&dst.Ident{Name: "packet"},
					&dst.BasicLit{
						Kind:  token.INT,
						Value: fmt.Sprintf("%v", field.Offset),
					}},
			},
		}

		kv.Decs.Before = dst.NewLine
		kv.Decs.After = dst.NewLine

		return &kv
	}
}
