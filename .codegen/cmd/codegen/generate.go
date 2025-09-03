package main

import (
	_ "embed"
	"os"

	"codegen/api"
	"codegen/codec"
	"codegen/integration-tests"
	"codegen/readme"
)

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]

		switch args[0] {
		case "codec":
			codec.Codec()

		case "integration-tests":
			integration_tests.IntegrationTests()

		case "API":
			api.API()

		case "responses":
			api.Responses()

		case "README":
			readme.README()
			readme.API()
		}
	}
}
