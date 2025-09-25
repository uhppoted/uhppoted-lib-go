package codec

import (
	_ "embed"
	"log"
	"os"
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
	encode()
	encodeTest()

	decodeAST()
	decodeTest()

	decoder()
	decoderTest()
}

func encode() {
	const output = "encode/generated.go"

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
	const output = "encode/encode_test.go"

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

// func decode() {
// 	const output = "decode/generated.go"
//
// 	f, err := os.Create(output)
// 	if err != nil {
// 		log.Fatalf("Failed to create file %s: %v", output, err)
// 	}
// 	defer f.Close()
//
// 	responses := append([]*lib.Response(nil), model.Responses...)
// 	responses = append(responses, &model.ListenerEvent)
//
// 	tmpl := template.Must(template.New("decode").Funcs(functions).Parse(decodeTemplate))
// 	if err := tmpl.Execute(f, responses); err != nil {
// 		log.Fatalf("Failed to execute template: %v", err)
// 	}
//
// 	log.Printf("... generated %s", output)
// }

func decodeTest() {
	const output = "decode/decode_test.go"

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	responses := append([]*lib.Response(nil), model.Responses...)
	responses = append(responses, &model.ListenerEvent)

	tmpl := template.Must(template.New("decode_test").Funcs(functions).Parse(decodeTestTemplate))
	if err := tmpl.Execute(f, responses); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("... generated %s", output)
}
