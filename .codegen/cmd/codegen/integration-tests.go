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

//go:embed templates/integration-tests/default.template
var defaultTemplate string

//go:embed templates/integration-tests/udp.template
var udpTemplate string

//go:embed templates/integration-tests/tcp.template
var tcpTemplate string

func integrationTests() {
	messages()
	expected()
	broadcast()
	udp()
	tcp()
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
	const output = "default/api_test.go"

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	tmpl := template.Must(template.New("encode").Funcs(functions).Parse(defaultTemplate))
	if err := tmpl.Execute(f, model.API); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("... generated %s", filepath.Base(output))
}

func udp() {
	const output = "udp/api_test.go"

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	tmpl := template.Must(template.New("encode").Funcs(functions).Parse(udpTemplate))
	if err := tmpl.Execute(f, model.API); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("... generated %s", filepath.Base(output))
}

func tcp() {
	const output = "tcp/api_test.go"

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	tmpl := template.Must(template.New("encode").Funcs(functions).Parse(tcpTemplate))
	if err := tmpl.Execute(f, model.API); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("... generated %s", filepath.Base(output))
}
