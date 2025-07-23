package main

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"codegen/model"
)

//go:embed templates/integration-tests/messages.template
var messagesTemplate string

//go:embed templates/integration-tests/expected.template
var expectedTemplate string

//go:embed templates/integration-tests/broadcast.template
var broadcastTemplate string

func integrationTests() {
	messages()
	expected()
	broadcast()
}

func messages() {
	const output = "messages.go"

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	tmpl := template.Must(template.New("encode").Funcs(functions).Parse(messagesTemplate))
	if err := tmpl.Execute(f, model.API); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("... generated %s", filepath.Base(output))
}

func expected() {
	const output = "expected.go"

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	tmpl := template.Must(template.New("encode").Funcs(functions).Parse(expectedTemplate))
	if err := tmpl.Execute(f, model.API); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("... generated %s", filepath.Base(output))
}

func broadcast() {
	const output = "udp-broadcast/api_test.go"

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	tmpl := template.Must(template.New("encode").Funcs(functions).Parse(broadcastTemplate))
	if err := tmpl.Execute(f, model.API); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("... generated %s", filepath.Base(output))
}
