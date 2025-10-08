package main

import (
	"encoding/json"
	"flag"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
)

func setEventIndex(u lib.Uhppoted, args []string) error {
	var index uint

	flagset := flag.NewFlagSet("set-event-index", flag.ExitOnError)

	flagset.UintVar(&index, "index", 0, "downloaded event index")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else {
		f := func(c uint32) (any, error) {
			return lib.SetEventIndex(u, c, uint32(index), options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.SetEventIndex(u, c, uint32(index), options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("set-event-index\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
