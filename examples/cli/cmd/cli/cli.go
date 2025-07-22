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
}

func exec(controller uint, dest string, tcp bool, f func(c uint32) (any, error), g func(c lib.Controller) (any, error)) (any, error) {
	if c, err := resolve(controller, dest, tcp); err != nil {
		return nil, err
	} else if c == nil {
		return f(uint32(controller))
	} else {
		return g(*c)
	}
}

func exex(args []string, f func(c uint32) (any, error), g func(c lib.Controller) (any, error)) (any, error) {
	var controller uint
	var dest string
	var tcp bool

	flagset := flag.NewFlagSet("get-time", flag.ExitOnError)

	flagset.UintVar(&controller, "controller", 0, "controller serial number")
	flagset.StringVar(&dest, "dest", "", "controller IPv4 address (optional)")
	flagset.BoolVar(&tcp, "tcp", false, "use TCP/IP transport (optional)")

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
	} else if addrport, err := netip.ParseAddrPort(dest); err == nil && tcp {
		return &lib.Controller{uint32(controller), addrport, "tcp"}, nil
	} else if addrport, err := netip.ParseAddrPort(dest); err == nil {
		return &lib.Controller{uint32(controller), addrport, "udp"}, nil
	} else if addr, err := netip.ParseAddr(dest); err == nil && tcp {
		return &lib.Controller{uint32(controller), netip.AddrPortFrom(addr, 60000), "tcp"}, nil
	} else if addr, err := netip.ParseAddr(dest); err == nil {
		return &lib.Controller{uint32(controller), netip.AddrPortFrom(addr, 60000), "udp"}, nil
	} else {
		return nil, fmt.Errorf("invalid controller IPv4 address (%v)", dest)
	}
}
