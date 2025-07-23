package main

import (
	"encoding/json"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func getController(u lib.Uhppoted, args []string) error {
	f := func(c uint32) (any, error) {
		return lib.GetController(u, c, options.timeout)
	}

	g := func(c lib.Controller) (any, error) {
		return lib.GetController(u, c, options.timeout)
	}

	if v, err := get(args, f, g); err != nil {
		return err
	} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
		return err
	} else {
		fmt.Printf("get-status\n")
		fmt.Printf("   %v\n", string(bytes))
		fmt.Println()

		return nil
	}
}
