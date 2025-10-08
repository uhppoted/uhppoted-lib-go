package main

import (
	"encoding/json"
	"flag"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
)

func getEventIndex(u lib.Uhppoted, args []string) error {
	flagset := flag.NewFlagSet("get-event-index", flag.ExitOnError)

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else {
		f := func(c uint32) (any, error) {
			return lib.GetEventIndex(u, c, options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.GetEventIndex(u, c, options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("get-event-index\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
