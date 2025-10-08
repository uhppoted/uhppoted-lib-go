package main

import (
	"encoding/json"
	"flag"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
)

func setPCControl(u lib.Uhppoted, args []string) error {
	var enabled bool

	flagset := flag.NewFlagSet("set-pc-control", flag.ExitOnError)

	flagset.BoolVar(&enabled, "enabled", true, "enables access control by remote host")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else {
		f := func(c uint32) (any, error) {
			return lib.SetPCControl(u, c, enabled, options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.SetPCControl(u, c, enabled, options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("set-pc-control\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
