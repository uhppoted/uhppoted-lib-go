package main

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"codegen/model"
)

//go:embed templates/encode.template
var encodeTemplate string

//go:embed templates/encode_test.template
var encodeTestTemplate string

//go:embed templates/decode_test.template
var decodeTestTemplate string

func encode() {
	const output = "encode/requests.go"

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	tmpl := template.Must(template.New("encode").Funcs(functions).Parse(encodeTemplate))
	if err := tmpl.Execute(f, model.Requests); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("... generated %s", filepath.Base(output))
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

	log.Printf("... generated %s", filepath.Base(output))
}

func decodeTest() {
	const output = "decode/decode_test.go"

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	tmpl := template.Must(template.New("decode_test").Funcs(functions).Parse(decodeTestTemplate))
	if err := tmpl.Execute(f, model.Responses); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("... generated %s", filepath.Base(output))
}
