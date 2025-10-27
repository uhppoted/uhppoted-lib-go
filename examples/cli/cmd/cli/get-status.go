package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
)

func getStatus(u uhppoted.Uhppoted, args []string) error {
	flagset := flag.NewFlagSet("get-status", flag.ExitOnError)

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else {
		f := func(c uint32) (any, error) {
			return uhppoted.GetStatus(u, c, options.timeout)
		}

		g := func(c uhppoted.Controller) (any, error) {
			return uhppoted.GetStatus(u, c, options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("get-status\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}

func getStatusRecord(u uhppoted.Uhppoted, args []string) error {
	flagset := flag.NewFlagSet("get-status-record", flag.ExitOnError)

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else {
		f := func(c uint32) (any, error) {
			return uhppoted.GetStatusRecord(u, c, options.timeout)
		}

		g := func(c uhppoted.Controller) (any, error) {
			return uhppoted.GetStatusRecord(u, c, options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("get-status-record\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
