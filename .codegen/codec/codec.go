package codec

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
	"codegen/model"
)

//go:embed templates/encode.template
var encodeTemplate string

//go:embed templates/encode_test.template
var encodeTestTemplate string

//go:embed templates/decode_test.template
var decodeTestTemplate string

var functions = codegen.Functions

func Codec() {
	encodeAST()
	encode()
	encodeTest()

	decode()
	decodeTest()

	decoder()
	decoderTest()
}

func encode() {
	output := filepath.Join("codec", "encode", "_generated.go")

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	tmpl := template.Must(template.New("encode").Funcs(functions).Parse(encodeTemplate))
	if err := tmpl.Execute(f, model.Requests); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("... generated %s", output)
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

func decodeTest() {
	output := filepath.Join("codec", "decode", "decode_test.go")

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	responses := []*lib.Response{}
	responses = append(responses, model.Responses...)
	responses = append(responses, &model.ListenerEvent)

	tmpl := template.Must(template.New("decode_test").Funcs(functions).Parse(decodeTestTemplate))
	if err := tmpl.Execute(f, responses); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("... generated %s", output)
}

func writeln(f *os.File, s string) {
	if _, err := f.WriteString(s + "\n"); err != nil {
		panic(fmt.Errorf("error writing to %v (%v)", f.Name(), err))
	}
}
