package readme

import (
	_ "embed"
	"log"
	"os"
	"text/template"

	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
	"codegen/model"
)

//go:embed templates/API.template
var apiTemplate string

func API() {
	const file = "../_API.md"

	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", file, err)
	}
	defer f.Close()

	var data = struct {
		API       []*lib.Function
		Responses []*lib.Response
	}{
		API:       model.API,
		Responses: model.Responses,
	}

	tmpl := template.Must(template.New("encode").Funcs(codegen.Functions).Parse(apiTemplate))
	if err := tmpl.Execute(f, data); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("... generated %s", file)

}
