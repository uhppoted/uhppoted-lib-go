package main

import (
	"encoding/json"
	"fmt"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
)

func findControllers(u uhppoted.Uhppoted, args []string) error {
	if controllers, err := uhppoted.FindControllers(u, options.timeout); err != nil {
		return err
	} else {
		fmt.Printf("find-controllers\n")
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
