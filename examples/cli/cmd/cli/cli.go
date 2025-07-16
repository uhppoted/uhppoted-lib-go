package main

import (
	"encoding/json"
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
			if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
				return err
			} else {
				fmt.Printf("   %v\n", string(bytes))
			}
		}

		fmt.Println()
	}

	return nil
}

func getController(u lib.Uhppoted, args []string) error {
	var controller uint
	var dest string
	var tcp bool

	flagset := flag.NewFlagSet("get-controller", flag.ExitOnError)

	flagset.UintVar(&controller, "controller", 0, "controller serial number")
	flagset.StringVar(&dest, "dest", "", "controller IPv4 address (optional)")
	flagset.BoolVar(&tcp, "tcp", false, "use TCP/IP transport (optional)")

	f := func(c lib.Controller) error {
		if v, err := lib.GetController(u, c, options.timeout); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("get-controller\n")
			fmt.Printf("   %v\n", string(bytes))

			return nil
		}
	}

	g := func(c uint32) error {
		if v, err := lib.GetController(u, c, options.timeout); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("get-controller\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}

	if err := flagset.Parse(args); err != nil {
		return err
	} else if c, err := resolve(controller, dest, tcp); err != nil {
		return err
	} else if c != nil {
		return f(*c)
	} else {
		return g(uint32(controller))
	}
}

// func exec[T lib.TController](controller uint, dest string, tcp bool, f func(c T) (any, error)) error {
// 	if c, err := resolve(controller, dest, tcp); err != nil {
// 		return err
// 	} else if c != nil {
// 		if v, err := f(*c); err != nil {
// 			return err
// 		} else {
// 			fmt.Printf("get-controller\n")
// 			fmt.Printf("   %v\n", v)
//
// 			return nil
// 		}
// 	} else {
// 		if v, err := f(uint32(controller)); err != nil {
// 			return err
// 		} else {
// 			fmt.Printf("get-controller\n")
// 			fmt.Printf("   %v\n", v)
//
// 			return nil
// 		}
// 	}
// }

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
