package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/netip"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func setIPv4(u lib.Uhppoted, args []string) error {
	var controller uint
	var dest string
	var tcp bool

	var address string
	var netmask string
	var gateway string

	flagset := flag.NewFlagSet("get-controller", flag.ExitOnError)

	flagset.UintVar(&controller, "controller", 0, "controller serial number")
	flagset.StringVar(&dest, "dest", "", "controller IPv4 address (optional)")
	flagset.BoolVar(&tcp, "tcp", false, "use TCP/IP transport (optional)")
	flagset.StringVar(&address, "address", "", "controller IPv4 address")
	flagset.StringVar(&netmask, "netmask", "", "controller IPv4 subnet mask")
	flagset.StringVar(&gateway, "gateway", "", "controller IPv4 gateway address")

	f := func(c uint32, address netip.Addr, netmask netip.Addr, gateway netip.Addr) error {
		if v, err := lib.SetIPv4(u, c, address, netmask, gateway, options.timeout); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("set-ip\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}

	g := func(c lib.Controller, address netip.Addr, netmask netip.Addr, gateway netip.Addr) error {
		if v, err := lib.SetIPv4(u, c, address, netmask, gateway, options.timeout); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("get-controller\n")
			fmt.Printf("   %v\n", string(bytes))

			return nil
		}
	}

	if err := flagset.Parse(args); err != nil {
		return err
	} else if c, err := resolve(controller, dest, tcp); err != nil {
		return err
	} else if addr, err := netip.ParseAddr(address); err != nil {
		return err
	} else if mask, err := netip.ParseAddr(netmask); err != nil {
		return err
	} else if gw, err := netip.ParseAddr(gateway); err != nil {
		return err
	} else if c == nil {
		return f(uint32(controller), addr, mask, gw)
	} else {
		return g(*c, addr, mask, gw)
	}
}
