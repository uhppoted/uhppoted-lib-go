package main

import (
	"encoding/json"
	"flag"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func getAntiPassback(u lib.Uhppoted, args []string) error {
	flagset := flag.NewFlagSet("get-antipassback", flag.ExitOnError)

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else {
		f := func(c uint32) (any, error) {
			return lib.GetAntiPassback(u, c, options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.GetAntiPassback(u, c, options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("get-antipassback\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
