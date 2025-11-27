package integration_tests

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"codegen/codegen"
	"codegen/model"
)

//go:embed templates/default.template
var defaultTemplate string

//go:embed templates/udp.template
var udpTemplate string

//go:embed templates/tcp.template
var tcpTemplate string

var functions = codegen.Functions

func IntegrationTests() {
	messages()
	expected()

	broadcast()
	broadcastAST()

	udp()
	udpAST()

	tcp()
}

func broadcast() {
	const output = "default/_api_test.go"

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
	const output = "udp/_api_test.go"

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	tmpl := template.Must(template.New("encode").Funcs(functions).Parse(udpTemplate))
	if err := tmpl.Execute(f, model.UDP); err != nil {
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
	if err := tmpl.Execute(f, model.TCP); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("... generated %s", filepath.Base(output))
}

func writeln(f *os.File, s string) {
	if _, err := f.WriteString(s + "\n"); err != nil {
		panic(fmt.Errorf("error writing to %v (%v)", f.Name(), err))
	}
}
