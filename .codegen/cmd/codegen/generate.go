package main

import (
	_ "embed"
	"os"

	"codegen/api"
	"codegen/codegen"
	"codegen/readme"
)

var functions = codegen.Functions

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]

		switch args[0] {
		case "codec":
			codec()

		case "integration-tests":
			integrationTests()

		case "API":
			api.API()
			api.Structs()

		case "README":
			readme.README()
		}
	}
}
