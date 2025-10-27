package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
)

func clearTimeProfiles(u uhppoted.Uhppoted, args []string) error {
	flagset := flag.NewFlagSet("clear-time-profiles", flag.ExitOnError)

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else {
		f := func(c uint32) (any, error) {
			return uhppoted.ClearTimeProfiles(u, c, options.timeout)
		}

		g := func(c uhppoted.Controller) (any, error) {
			return uhppoted.ClearTimeProfiles(u, c, options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("clear-time-profiles\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
