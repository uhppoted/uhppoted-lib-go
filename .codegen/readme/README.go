package readme

import (
	_ "embed"
	"log"
	"os"
	"text/template"

	"codegen/codegen"
	"codegen/model"
	"codegen/model/types"
)

//go:embed templates/README.template
var readmeTemplate string

func README() {
	const file = "../README.md"

	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", file, err)
	}
	defer f.Close()

	var API = []types.Function{}
	API = append(API, model.API...)
	API = append(API, model.GetListenerAddrPort)
	API = append(API, model.SetListenerAddrPort)

	tmpl := template.Must(template.New("encode").Funcs(codegen.Functions).Parse(readmeTemplate))
	if err := tmpl.Execute(f, API); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("... generated %s", file)

}
