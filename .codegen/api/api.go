package api

import (
	"log"
	"path/filepath"

	"codegen/codegen"
)

func API() {
	const file = "_generated.go"

	AST := codegen.NewAST("uhppoted")

	if err := AST.Generate(file); err != nil {
		log.Fatalf("error generating %v (%v)", file, err)
	} else {
		log.Printf("... generated %s", filepath.Base(file))
	}
}
