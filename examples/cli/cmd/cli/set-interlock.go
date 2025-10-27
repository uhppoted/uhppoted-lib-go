package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"
)

var interlocks = map[string]types.Interlock{
	"disabled": types.NoInterlock,
	"1&2":      types.Interlock12,
	"3&4":      types.Interlock34,
	"1&2,3&4":  types.Interlock12_34,
	"1&2&3":    types.Interlock123,
	"1&2&3&4":  types.Interlock1234,
}

func setInterlock(u uhppoted.Uhppoted, args []string) error {
	var interlock string

	flagset := flag.NewFlagSet("set-interlock", flag.ExitOnError)

	flagset.StringVar(&interlock, "interlock", "", "sets the door interlock mode")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if v, ok := interlocks[interlock]; !ok {
		return fmt.Errorf("invalid interlock mode (%v)", interlock)
	} else {
		f := func(c uint32) (any, error) {
			return uhppoted.SetInterlock(u, c, v, options.timeout)
		}

		g := func(c uhppoted.Controller) (any, error) {
			return uhppoted.SetInterlock(u, c, v, options.timeout)
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
