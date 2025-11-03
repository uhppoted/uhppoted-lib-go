package codec

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
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
	params := encodeParams(r)

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

	// ... body
	body := []dst.Stmt{
		makePacket(),
		setSOM(),
		setMsgType(r),
	}

	for _, f := range r.Fields {
		body = append(body, pack(f))
	}

	// return packet,nil
	body = append(body, &dst.ReturnStmt{
		Results: []dst.Expr{
			&dst.Ident{
				Name: "packet",
			},
			&dst.Ident{
				Name: "nil",
			},
		},

		Decs: dst.ReturnStmtDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.EmptyLine,
			},
		},
	})

	// ... assemble func
	f := dst.FuncDecl{
		Name: dst.NewIdent(name),
		Type: &dst.FuncType{
			Params:  &params,
			Results: &results,
		},

		Body: &dst.BlockStmt{
			List: body,
		},

		Decs: dst.FuncDeclDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.EmptyLine,
				After:  dst.EmptyLine,
			},
		},
	}

	// godoc
	f.Decs.Start.Append(fmt.Sprintf("// Encodes a %v request to a 64 byte packet.", name))

	return &f
}

func encodeParams(r lib.Request) dst.FieldList {
	args := []*dst.Field{}

	for _, arg := range r.Fields {
		name := regexp.MustCompile(`[ \-]+`).ReplaceAllString(arg.Name, "")
		t := arg.Type

		switch arg.Type {
		case "controller":
			t = "T"

		case "IPv4":
			t = "netip.Addr"

		case "address:port":
			t = "netip.AddrPort"

		case "date":
			t = "D"

		case "datetime":
			t = "DT"

		case "HHmm":
			t = "H"

		case "pin":
			t = "uint32"

		case "mode":
			t = "types.DoorMode"

		case "task":
			t = "types.TaskType"

		case "interlock":
			t = "types.Interlock"

		case "anti-passback":
			t = "types.AntiPassback"

		case "magic":
			continue
		}

		args = append(args, &dst.Field{
			Names: []*dst.Ident{
				{Name: name},
			},
			Type: &dst.Ident{Name: t},
		})
	}

	return dst.FieldList{
		List: args,
	}
}

// packet := make([]byte, 64)
func makePacket() *dst.AssignStmt {
	stmt := dst.AssignStmt{
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

	stmt.Decs.Before = dst.NewLine
	stmt.Decs.After = dst.EmptyLine

	return &stmt
}

// packet[0] = SOM
func setSOM() *dst.AssignStmt {
	stmt := dst.AssignStmt{
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

	return &stmt
}

// packet[1] = <msgtype>
func setMsgType(r lib.Request) *dst.AssignStmt {
	stmt := dst.AssignStmt{
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

	stmt.Decs.After = dst.EmptyLine

	return &stmt
}

// packXXX()
func pack(field lib.Field) *dst.ExprStmt {
	f := func(fn string) *dst.ExprStmt {
		return &dst.ExprStmt{
			X: &dst.CallExpr{
				Fun: &dst.Ident{
					Name: fn,
				},
				Args: []dst.Expr{
					&dst.Ident{
						Name: field.Name,
					},
					&dst.Ident{
						Name: "packet",
					},
					&dst.BasicLit{
						Kind:  token.INT,
						Value: fmt.Sprintf("%v", field.Offset),
					},
				},
			},
		}
	}

	switch field.Type {
	case "bool":
		return f("packBool")

	case "uint8":
		return f("packUint8")

	case "uint16":
		return f("packUint16")

	case "uint32":
		return f("packUint32")

	case "IPv4":
		return f("packIPv4")

	case "address:port":
		return f("packAddrPort")

	case "datetime":
		return f("packDateTime")

	case "date":
		return f("packDate")

	case "HHmm":
		return f("packHHmm")

	case "pin":
		return f("packPIN")

	case "mode":
		return f("packMode")

	case "task":
		return f("packTaskType")

	case "interlock":
		return f("packInterlock")

	case "anti-passback":
		return f("packAntiPassback")

	// case "passcode":
	// 	return fmt.Sprintf("packPasscode(%v, packet, %v)", name, field.Offset)

	case "magic":
		return &dst.ExprStmt{
			X: &dst.CallExpr{
				Fun: &dst.Ident{
					Name: "packUint32",
				},
				Args: []dst.Expr{
					&dst.BasicLit{
						Kind:  token.INT,
						Value: "0x55aaaa55",
					},
					&dst.Ident{
						Name: "packet",
					},
					&dst.BasicLit{
						Kind:  token.INT,
						Value: fmt.Sprintf("%v", field.Offset),
					},
				},
			},
		}

	default:
		panic(fmt.Sprintf("*** unsupported field type (%v)", field.Type))
	}
}
