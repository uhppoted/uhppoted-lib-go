package main

import (
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func GetAllControllers(u lib.Uhppoted, args []string) error {
	if controllers, err := u.GetAllControllers(options.timeout); err != nil {
		return err
	} else {
		fmt.Printf("get-all-controllers\n")
		for _, v := range controllers {
			fmt.Printf("   %v\n", v)
		}
	}

	return nil
}
