package codec

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"go/token"

	"github.com/dave/dst"

	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
	"codegen/model"
)

//go:embed templates/encode_test.template
var encodeTestTemplate string

var functions = codegen.Functions

func Codec() {
	encode()
	encodeTest()
	encodeTestAST()

	decode()
	decodeTest()

	decoder()
	decoderTest()
}

func encodeTest() {
	output := filepath.Join("codec", "encode", "encode_test.go")

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	tmpl := template.Must(template.New("encode_test").Funcs(functions).Parse(encodeTestTemplate))
	if err := tmpl.Execute(f, model.Requests); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("... generated %s", output)
}

func writeln(f *os.File, s string) {
	if _, err := f.WriteString(s + "\n"); err != nil {
		panic(fmt.Errorf("error writing to %v (%v)", f.Name(), err))
	}
}

func makeValue(field lib.Field, value lib.Value) dst.Expr {
	switch field.Type {
	case "uint8":
		return &dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%v", value.Value)}

	case "uint16":
		return &dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%v", value.Value)}

	case "uint32":
		return &dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%v", value.Value)}

	case "bool":
		return &dst.Ident{Name: fmt.Sprintf("%v", value.Value)}

	case "IPv4":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "netip"},
				Sel: &dst.Ident{Name: "MustParseAddr"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "address:port":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "netip"},
				Sel: &dst.Ident{Name: "MustParseAddrPort"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "date":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseDate"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "shortdate":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseDate"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "optional date":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseDate"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "time":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseTime"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "HHmm":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseHHmm"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "datetime":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseDateTime"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "optional datetime":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "MustParseDateTime"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)},
			},
		}

	case "MAC":
		return &dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)}

	case "version":
		return &dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%v"`, value.Value)}

	case "pin":
		return &dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%v", value.Value)}

	case "mode":
		return &dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%v", value.Value)}

	case "anti-passback":
		return &dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%v", value.Value)}

	case "event-type":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "EventType"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf(`%v`, value.Value)},
			},
		}

	case "direction":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "Direction"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf(`%v`, value.Value)},
			},
		}

	case "reason":
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "types"},
				Sel: &dst.Ident{Name: "Reason"},
			},
			Args: []dst.Expr{
				&dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf(`%v`, value.Value)},
			},
		}

	default:
		panic(fmt.Sprintf("%v", field.Type))
	}
}
