package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/netip"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func setIPv4(u lib.Uhppoted, args []string) error {
	var address string
	var netmask string
	var gateway string

	flagset := flag.NewFlagSet("set-IPv4", flag.ExitOnError)

	flagset.StringVar(&address, "address", "", "controller IPv4 address")
	flagset.StringVar(&netmask, "netmask", "", "controller IPv4 subnet mask")
	flagset.StringVar(&gateway, "gateway", "", "controller IPv4 gateway address")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if addr, err := netip.ParseAddr(address); err != nil {
		return err
	} else if mask, err := netip.ParseAddr(netmask); err != nil {
		return err
	} else if gw, err := netip.ParseAddr(gateway); err != nil {
		return err
	} else {
		f := func(c uint32) (any, error) {
			return lib.SetIPv4(u, c, addr, mask, gw, options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.SetIPv4(u, c, addr, mask, gw, options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("set-IP\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
