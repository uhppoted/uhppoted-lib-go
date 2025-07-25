package main

import (
	"flag"
	"fmt"
	"net/netip"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

var commands = map[string]func(u lib.Uhppoted, args []string) error{
	"get-all-controllers": getAllControllers,
	"get-controller":      getController,
	"set-IPv4":            setIPv4,
	"get-status":          getStatus,
	"get-time":            getTime,
	"set-time":            setTime,
	"get-listener":        getListener,
}

func get(args []string, f func(c uint32) (any, error), g func(c lib.Controller) (any, error)) (any, error) {
	var controller uint
	var dest string
	var tcp bool

	flagset := flag.NewFlagSet("get-time", flag.ContinueOnError)

	flagset.UintVar(&controller, "controller", 0, "controller serial number")
	flagset.StringVar(&dest, "dest", "", "controller IPv4 address (optional)")
	flagset.BoolVar(&tcp, "tcp", false, "use TCP/IP transport (optional)")

	// FIXME remove and replace with halfway decent argparse
	flagset.String("datetime", "", "(optional) date/time - defaults to current time")

	if err := flagset.Parse(args); err != nil {
		return nil, err
	} else if c, err := resolve(controller, dest, tcp); err != nil {
		return nil, err
	} else if c == nil {
		return f(uint32(controller))
	} else {
		return g(*c)
	}
}

func resolve(controller uint, dest string, tcp bool) (*lib.Controller, error) {
	if dest == "" {
		return nil, nil
	}

	if addrport, err := netip.ParseAddrPort(dest); err == nil && tcp {
		return &lib.Controller{
			ID:       uint32(controller),
			Address:  addrport,
			Protocol: "tcp",
		}, nil
	}

	if addrport, err := netip.ParseAddrPort(dest); err == nil {
		return &lib.Controller{
			ID:       uint32(controller),
			Address:  addrport,
			Protocol: "udp",
		}, nil
	}

	if addr, err := netip.ParseAddr(dest); err == nil && tcp {
		return &lib.Controller{
			ID:       uint32(controller),
			Address:  netip.AddrPortFrom(addr, 60000),
			Protocol: "tcp",
		}, nil
	}

	if addr, err := netip.ParseAddr(dest); err == nil {
		return &lib.Controller{
			ID:       uint32(controller),
			Address:  netip.AddrPortFrom(addr, 60000),
			Protocol: "udp",
		}, nil
	}

	return nil, fmt.Errorf("invalid controller IPv4 address (%v)", dest)
}
