package main

import (
	"encoding/json"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func setTime(u lib.Uhppoted, args []string) error {
	var datetime string

	flagset := NewFlagSet("set-time")

	flagset.StringVar(&datetime, "datetime", "", "(optional) date/time - defaults to current time")

	if now, err := parseDateTime(datetime); err != nil {
		return err
	} else {
		f := func(c uint32) (any, error) {
			return lib.SetTime(u, c, now, options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.SetTime(u, c, now, options.timeout)
		}

		if err := flagset.Parse(args); err != nil {
			return err
		} else if v, err := get(args, f, g); err != nil {
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
