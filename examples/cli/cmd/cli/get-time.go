package main

import (
	"encoding/json"
	"flag"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func getTime(u lib.Uhppoted, args []string) error {
	flagset := flag.NewFlagSet("get-time", flag.ExitOnError)

	f := func(c uint32) (any, error) {
		return lib.GetTime(u, c, options.timeout)
	}

	g := func(c lib.Controller) (any, error) {
		return lib.GetTime(u, c, options.timeout)
	}

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if v, err := get(controller, flagset, f, g); err != nil {
		return err
	} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
		return err
	} else {
		fmt.Printf("get-time\n")
		fmt.Printf("   %v\n", string(bytes))
		fmt.Println()

		return nil
	}
}
