package main

import (
	"encoding/json"
	"flag"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func setDoorPasscodes(u lib.Uhppoted, args []string) error {
	var door uint
	var passcodes string

	flagset := flag.NewFlagSet("set-door-passcodes", flag.ExitOnError)

	flagset.UintVar(&door, "door", 0, "door ID [1..4]")
	flagset.StringVar(&passcodes, "passcodes", "", "comma separated list of up to four passcodes, e.g. 1234,54321")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else {
		passcode1 := uint32(0)
		passcode2 := uint32(0)
		passcode3 := uint32(0)
		passcode4 := uint32(0)

		codes := parsePasscodes(passcodes)
		if len(codes) > 0 {
			passcode1 = codes[0]
		}

		if len(codes) > 1 {
			passcode2 = codes[1]
		}

		if len(codes) > 2 {
			passcode3 = codes[2]
		}

		if len(codes) > 3 {
			passcode4 = codes[0]
		}

		f := func(c uint32) (any, error) {
			return lib.SetDoorPasscodes(u, c, uint8(door), passcode1, passcode2, passcode3, passcode4, options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.SetDoorPasscodes(u, c, uint8(door), passcode1, passcode2, passcode3, passcode4, options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("set-door\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
