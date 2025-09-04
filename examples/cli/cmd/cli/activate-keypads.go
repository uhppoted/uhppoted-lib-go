package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"strings"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func activateKeypads(u lib.Uhppoted, args []string) error {
	var keypads string

	flagset := flag.NewFlagSet("activate-keypads", flag.ExitOnError)

	flagset.StringVar(&keypads, "keypads", "", "activates reader keypads (e.g. 1,2,4)")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if v, err := parseKeypads(keypads); err != nil {
		return err
	} else {
		f := func(c uint32) (any, error) {
			return lib.ActivateKeypads(u, c, v[0], v[1], v[2], v[3], options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.ActivateKeypads(u, c, v[0], v[1], v[2], v[3], options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("activate-keypads\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}

func parseKeypads(s string) ([]bool, error) {
	keypads := map[uint8]bool{
		1: false,
		2: false,
		3: false,
		4: false,
	}

	tokens := strings.Split(s, ",")
	for _, token := range tokens {
		if keypad, err := strconv.ParseUint(token, 10, 8); err == nil {
			keypads[uint8(keypad)] = true
		}
	}

	return []bool{keypads[1], keypads[2], keypads[3], keypads[4]}, nil
}
