package main

import (
	"encoding/json"
	"flag"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func deleteCard(u lib.Uhppoted, args []string) error {
	var card uint

	flagset := flag.NewFlagSet("get-card", flag.ExitOnError)

	flagset.UintVar(&card, "card", 0, "card number")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else {
		facilityCode := card / 100000
		cardNumber := card % 100000
		if facilityCode > 255 || cardNumber > 65535 || (facilityCode == 0 && cardNumber == 0) {
			return fmt.Errorf("invalid card (%v)", card)
		} else {
			f := func(c uint32) (any, error) {
				return lib.DeleteCard(u, c, uint32(card), options.timeout)
			}

			g := func(c lib.Controller) (any, error) {
				return lib.DeleteCard(u, c, uint32(card), options.timeout)
			}

			if v, err := exec(controller, flagset, f, g); err != nil {
				return err
			} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
				return err
			} else {
				fmt.Printf("delete-card\n")
				fmt.Printf("   %v\n", string(bytes))
				fmt.Println()

				return nil
			}
		}
	}
}
