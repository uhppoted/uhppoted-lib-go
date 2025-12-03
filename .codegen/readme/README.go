package readme

import (
	"bytes"
	_ "embed"
	"log"
	"os"
	"regexp"
	"text/template"

	"codegen/codegen"
	"codegen/model"
)

//go:embed templates/README.template
var readme string

func README() {
	var b bytes.Buffer

	// ... load README.md
	src, err := os.ReadFile("../../README.md")
	if err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	// ... generate API section
	var API = model.API

	tmpl := template.Must(template.New("encode").Funcs(codegen.Functions).Parse(readme))
	if err := tmpl.Execute(&b, API); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	// ... update API section of README

	re := regexp.MustCompile(`(?s)(.*?\n## API)(.*?)(\n## License.*)`)
	replace := "$1\n" + b.String() + "\n$3"
	updated := re.ReplaceAllString(string(src), replace)

	// ... (conditionally) write to README.md
	if string(src) != updated {
		const file = "../../README.md"

		f, err := os.Create(file)
		if err != nil {
			log.Fatalf("Failed to create file %s: %v", file, err)
		}
		defer f.Close()

		f.Write([]byte(updated))

		log.Printf("**** WARNING: README.md updated")
	}
}
