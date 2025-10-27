package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
)

func openDoor(u uhppoted.Uhppoted, args []string) error {
	var door uint

	flagset := flag.NewFlagSet("open-door", flag.ExitOnError)

	flagset.UintVar(&door, "door", 0, "door ID [1..4]")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if door < 1 || door > 4 {
		return fmt.Errorf("invalid door (%v)", door)
	} else {
		f := func(c uint32) (any, error) {
			return uhppoted.OpenDoor(u, c, uint8(door), options.timeout)
		}

		g := func(c uhppoted.Controller) (any, error) {
			return uhppoted.OpenDoor(u, c, uint8(door), options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("open-door\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
