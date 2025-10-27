package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"
)

func setTime(u uhppoted.Uhppoted, args []string) error {
	var datetime string

	flagset := flag.NewFlagSet("set-time", flag.ExitOnError)

	flagset.StringVar(&datetime, "datetime", "", "(optional) date/time - defaults to current time")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if now, err := parseDateTime(datetime); err != nil {
		return err
	} else {
		// (for demo purposes only - there is actually a uhppoted.DateTimeFromTime helper function)
		dt := types.NewDateTime(uint16(now.Year()), uint8(now.Month()), uint8(now.Day()), uint8(now.Hour()), uint8(now.Minute()), uint8(now.Second()))

		f := func(c uint32) (any, error) {
			return uhppoted.SetTime(u, c, now, options.timeout)
		}

		g := func(c uhppoted.Controller) (any, error) {
			return uhppoted.SetTime(u, c, dt, options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("set-time\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
