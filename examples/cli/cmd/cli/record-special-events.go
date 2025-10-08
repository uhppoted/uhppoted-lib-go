package main

import (
	"encoding/json"
	"flag"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
)

func recordSpecialEvents(u lib.Uhppoted, args []string) error {
	var enabled bool

	flagset := flag.NewFlagSet("record-special-events", flag.ExitOnError)

	flagset.BoolVar(&enabled, "enabled", true, "enables door open/close and button pressed events")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else {
		f := func(c uint32) (any, error) {
			return lib.RecordSpecialEvents(u, c, enabled, options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.RecordSpecialEvents(u, c, enabled, options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("record-special-events\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
