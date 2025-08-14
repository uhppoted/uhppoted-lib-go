package main

import (
	"encoding/json"
	"flag"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func getEvent(u lib.Uhppoted, args []string) error {
	var index uint

	flagset := flag.NewFlagSet("get-event", flag.ExitOnError)

	flagset.UintVar(&index, "index", 0, "event index")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else {
		f := func(c uint32) (any, error) {
			return lib.GetEvent(u, c, uint32(index), options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.GetEvent(u, c, uint32(index), options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if v.(lib.GetEventResponse).EventType == 0 {
			return fmt.Errorf("no event at index %v", index)
		} else if v.(lib.GetEventResponse).EventType == 0xff {
			return fmt.Errorf("event at index %v overwritten", index)
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("get-event\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
