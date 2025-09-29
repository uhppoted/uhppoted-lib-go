package main

import (
	"encoding/json"
	"flag"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

var interlocks = map[string]lib.Interlock{
	"disabled": lib.NoInterlock,
	"1&2":      lib.Interlock12,
	"3&4":      lib.Interlock34,
	"1&2,3&4":  lib.Interlock12_34,
	"1&2&3":    lib.Interlock123,
	"1&2&3&4":  lib.Interlock1234,
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
