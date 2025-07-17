package main

import (
	"encoding/json"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func getAllControllers(u lib.Uhppoted, args []string) error {
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
