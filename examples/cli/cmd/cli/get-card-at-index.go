package main

import (
	"encoding/json"
	"flag"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
)

func getCardAtIndex(u lib.Uhppoted, args []string) error {
	var index uint

	flagset := flag.NewFlagSet("get-card-at-index", flag.ExitOnError)

	flagset.UintVar(&index, "index", 0, "index of card record")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else {
		f := func(c uint32) (any, error) {
			return lib.GetCardAtIndex(u, c, uint32(index), options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.GetCardAtIndex(u, c, uint32(index), options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if v.(lib.GetCardAtIndexResponse).Card == 0 {
			return fmt.Errorf("no card at index %v", index)
		} else if v.(lib.GetCardAtIndexResponse).Card == 0xffffff {
			return fmt.Errorf("card at index %v deleted", index)
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("get-card-at-index\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}

func getCardRecordAtIndex(u lib.Uhppoted, args []string) error {
	var index uint

	flagset := flag.NewFlagSet("get-card-at-index", flag.ExitOnError)

	flagset.UintVar(&index, "index", 0, "index of card record")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else {
		f := func(c uint32) (any, error) {
			return lib.GetCardRecordAtIndex(u, c, uint32(index), options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.GetCardRecordAtIndex(u, c, uint32(index), options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if v.(lib.Card).Card == 0 {
			return fmt.Errorf("no card at index %v", index)
		} else if v.(lib.Card).Card == 0xffffff {
			return fmt.Errorf("card at index %v deleted", index)
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("get-card-at-index\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
