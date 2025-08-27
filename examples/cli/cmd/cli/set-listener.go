package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/netip"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func setListener(u lib.Uhppoted, args []string) error {
	var listener string
	var interval uint

	flagset := flag.NewFlagSet("set-listener", flag.ExitOnError)

	flagset.StringVar(&listener, "listener", "", "event listener IPv4 address:port")
	flagset.UintVar(&interval, "interval", 0, "(optional) auto-send interval (seconds, [0..255]")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if addrport, err := netip.ParseAddrPort(listener); err != nil {
		return err
	} else if interval > 255 {
		return fmt.Errorf("invalid auto-send interval (%v)", interval)
	} else {
		f := func(c uint32) (any, error) {
			return lib.SetListenerAddrPort(u, c, addrport, uint8(interval), options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.SetListenerAddrPort(u, c, addrport, uint8(interval), options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("set-listener\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
