package main

import (
	"encoding/json"
	"flag"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/entities"
)

var interlocks = map[string]entities.Interlock{
	"disabled": entities.NoInterlock,
	"1&2":      entities.Interlock12,
	"3&4":      entities.Interlock34,
	"1&2,3&4":  entities.Interlock12_34,
	"1&2&3":    entities.Interlock123,
	"1&2&3&4":  entities.Interlock1234,
}

func setInterlock(u lib.Uhppoted, args []string) error {
	var interlock string

	flagset := flag.NewFlagSet("set-interlock", flag.ExitOnError)

	flagset.StringVar(&interlock, "interlock", "", "sets the door interlock mode")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if v, ok := interlocks[interlock]; !ok {
		return fmt.Errorf("invalid interlock mode (%v)", interlock)
	} else {
		f := func(c uint32) (any, error) {
			return lib.SetInterlock(u, c, v, options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.SetInterlock(u, c, v, options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("set-interlock\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
