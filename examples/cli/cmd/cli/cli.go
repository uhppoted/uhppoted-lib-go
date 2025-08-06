package main

import (
	"flag"
	"fmt"
	"net/netip"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

var commands = map[string]func(u lib.Uhppoted, args []string) error{
	"find-controllers":   findControllers,
	"get-controller":     getController,
	"set-IPv4":           setIPv4,
	"get-time":           getTime,
	"set-time":           setTime,
	"get-listener":       getListener,
	"set-listener":       setListener,
	"get-door":           getDoor,
	"set-door":           setDoor,
	"set-door-passcodes": setDoorPasscodes,
	"open-door":          openDoor,
	"get-status":         getStatus,
	"get-cards":          getCards,
	"get-card":           getCard,
}

func exec(args controller, flagset *flag.FlagSet, f func(c uint32) (any, error), g func(c lib.Controller) (any, error)) (any, error) {
	if c, err := resolve(args.controller, args.dest, args.tcp); err != nil {
		return nil, err
	} else if c == nil {
		return f(uint32(args.controller))
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
