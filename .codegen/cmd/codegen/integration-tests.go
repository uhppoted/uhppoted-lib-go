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

func integrationTests() {
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
