package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"codegen/models"
)

//go:embed templates/encode.template
var testTemplate string

func main() {
	const output = "encode_test.go"

	f, err := os.Create(output)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", output, err)
	}
	defer f.Close()

	functions := template.FuncMap{
		"titleCase": titleCase,
		"hex":       hex,
		"args":      args,
	}

	tmpl := template.Must(template.New("encode_test").Funcs(functions).Parse(testTemplate))
	if err := tmpl.Execute(f, models.Requests); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("... generated %s", filepath.Base(output))
}

func titleCase(s string) string {
	parts := strings.Split(s, "-")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}

	return strings.Join(parts, "")
}

func hex(bytes []byte) string {
	lines := []string{}
	hex := "0x%02x, 0x%02x, 0x%02x, 0x%02x, 0x%02x, 0x%02x, 0x%02x, 0x%02x,"

	for i := 0; i < 4; i++ {
		offset := i * 16
		u := bytes[offset : offset+8]
		v := bytes[offset+8 : offset+16]

		p := fmt.Sprintf(hex, u[0], u[1], u[2], u[3], u[4], u[5], u[6], u[7])
		q := fmt.Sprintf(hex, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7])

		lines = append(lines, fmt.Sprintf("            %v %v", p, q))
	}

	return strings.Join(lines, "\n")
}

func args(args []any) string {
	var parts []string
	for _, a := range args {
		parts = append(parts, fmt.Sprintf("%v", a))
	}

	return strings.Join(parts, ", ")
}
