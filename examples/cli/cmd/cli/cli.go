package main

import (
	"flag"
	"fmt"
	"net/netip"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

var commands = map[string]func(u lib.Uhppoted, args []string) error{
	"get-all-controllers": GetAllControllers,
	"get-controller":      getController,
}

func GetAllControllers(u lib.Uhppoted, args []string) error {
	if controllers, err := lib.GetAllControllers(u, options.timeout); err != nil {
		return err
	} else {
		fmt.Printf("get-all-controllers\n")
		for _, v := range controllers {
			fmt.Printf("   %v\n", v)
		}
	}

	return nil
}

func getController(u lib.Uhppoted, args []string) error {
	var controller uint
	var dest string

	flagset := flag.NewFlagSet("get-controller", flag.ExitOnError)

	flagset.UintVar(&controller, "controller", 0, "controller serial number")
	flagset.StringVar(&dest, "dest", "", "controller IPv4 address (optional)")

	if err := flagset.Parse(args); err != nil {
		return err
	}

	if dest == "" {
		if v, err := lib.GetController(u, uint32(controller), options.timeout); err != nil {
			return err
		} else {
			fmt.Printf("get-controller\n")
			fmt.Printf("   %v\n", v)

			return nil
		}
	} else if addrport, err := netip.ParseAddrPort(dest); err == nil {
		c := lib.Controller{uint32(controller), addrport, "udp"}

		if v, err := lib.GetController(u, c, options.timeout); err != nil {
			return err
		} else {
			fmt.Printf("get-controller\n")
			fmt.Printf("   %v\n", v)

			return nil
		}
	} else if addr, err := netip.ParseAddr(dest); err == nil {
		c := lib.Controller{uint32(controller), netip.AddrPortFrom(addr, 60000), "udp"}

		if v, err := lib.GetController(u, c, options.timeout); err != nil {
			return err
		} else {
			fmt.Printf("get-controller\n")
			fmt.Printf("   %v\n", v)

			return nil
		}
	} else {
		return fmt.Errorf("invalid address (%v)", dest)
	}
}
