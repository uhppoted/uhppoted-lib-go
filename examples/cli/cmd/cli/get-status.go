package main

import (
	"encoding/json"
	"flag"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func getStatus(u lib.Uhppoted, args []string) error {
	var controller uint
	var dest string
	var tcp bool

	flagset := flag.NewFlagSet("get-status", flag.ExitOnError)

	flagset.UintVar(&controller, "controller", 0, "controller serial number")
	flagset.StringVar(&dest, "dest", "", "controller IPv4 address (optional)")
	flagset.BoolVar(&tcp, "tcp", false, "use TCP/IP transport (optional)")

	f := func(c uint32) (any, error) {
		return lib.GetStatus(u, c, options.timeout)
	}

	g := func(c lib.Controller) (any, error) {
		return lib.GetStatus(u, c, options.timeout)
	}

	if err := flagset.Parse(args); err != nil {
		return err
	} else if v, err := exec(controller, dest, tcp, f, g); err != nil {
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
