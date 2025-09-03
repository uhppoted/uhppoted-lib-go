package main

import (
	"encoding/json"
	"flag"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

var interlocks = map[string]uint8{
	"disabled": 0,
	"1&2":      1,
	"3&4":      2,
	"1&2,3&4":  3,
	"1&2&3":    4,
	"1&2&3&4":  8,
}

func setInterlock(u lib.Uhppoted, args []string) error {
	var interlock string

	flagset := flag.NewFlagSet("set-interlock", flag.ExitOnError)

	flagset.StringVar(&interlock, "interlock", "", "sets the door interlock mode")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if v, ok := interlocks[interlock]; !ok {
		return fmt.Errorf("invalid interlock mode (%v)", interlock)
	} else {
		f := func(c uint32) (any, error) {
			return lib.SetInterlock(u, c, v, options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.SetInterlock(u, c, v, options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("set-interlock\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
