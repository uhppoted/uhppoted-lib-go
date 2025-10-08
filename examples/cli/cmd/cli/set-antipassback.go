package main

import (
	"encoding/json"
	"flag"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
)

var antipassbacks = map[string]uint8{
	"disabled":    0,
	"(1:2);(3:4)": 1,
	"(1,3):(2,4)": 2,
	"1:(2,3)":     3,
	"1:(2,3,4)":   4,
}

func setAntiPassback(u lib.Uhppoted, args []string) error {
	var antipassback string

	flagset := flag.NewFlagSet("set-antipassback", flag.ExitOnError)

	flagset.StringVar(&antipassback, "anti-passback", "", "sets the controller anti-passback mode")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if v, ok := antipassbacks[antipassback]; !ok {
		return fmt.Errorf("invalid anti-passback mode (%v)", antipassback)
	} else {
		f := func(c uint32) (any, error) {
			return lib.SetAntiPassback(u, c, v, options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.SetAntiPassback(u, c, v, options.timeout)
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
