package main

import (
	"encoding/json"
	"flag"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func setDoor(u lib.Uhppoted, args []string) error {
	var door uint
	var mode uint
	var delay uint

	flagset := flag.NewFlagSet("set-door", flag.ExitOnError)

	flagset.UintVar(&door, "door", 0, "door ID [1..4]")
	flagset.UintVar(&mode, "mode", 3, "control mode (1: normally open, 2: normally closed, 3:controlled)")
	flagset.UintVar(&delay, "delay", 5, "unlock delay (seconds)")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if door < 1 || door > 4 {
		return fmt.Errorf("invalid door (%v)", door)
	} else if mode != 1 && mode != 2 && mode != 3 {
		return fmt.Errorf("invalid control mode (%v)", mode)
	} else if delay > 255 {
		return fmt.Errorf("invalid unlock delay (%v)", delay)
	} else {
		f := func(c uint32) (any, error) {
			return lib.SetDoor(u, c, uint8(door), lib.DoorMode(mode), uint8(delay), options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.SetDoor(u, c, uint8(door), lib.DoorMode(mode), uint8(delay), options.timeout)
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
