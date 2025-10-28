package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"
)

var antipassbacks = map[string]types.AntiPassback{
	"disabled":    types.NoAntiPassback,
	"(1:2);(3:4)": types.Readers12_34,
	"(1,3):(2,4)": types.Readers13_24,
	"1:(2,3)":     types.Readers1_23,
	"1:(2,3,4)":   types.Readers1_234,
}

func setAntiPassback(u uhppoted.Uhppoted, args []string) error {
	var antipassback string

	flagset := flag.NewFlagSet("set-antipassback", flag.ExitOnError)

	flagset.StringVar(&antipassback, "anti-passback", "", "sets the controller anti-passback mode")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if v, ok := antipassbacks[antipassback]; !ok {
		return fmt.Errorf("invalid anti-passback mode (%v)", antipassback)
	} else {
		f := func(c uint32) (any, error) {
			return uhppoted.SetAntiPassback(u, c, v, options.timeout)
		}

		g := func(c uhppoted.Controller) (any, error) {
			return uhppoted.SetAntiPassback(u, c, v, options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("set-antipassback\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
