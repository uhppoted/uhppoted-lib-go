package main

import (
	"flag"
	"fmt"

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
			fmt.Printf("   %v\n", v)
		}
	}

	return nil
}

func getController(u lib.Uhppoted, args []string) error {
	var controller uint

	flagset := flag.NewFlagSet("get-controller", flag.ExitOnError)

	flagset.UintVar(&controller, "controller", 0, "controller serial number")
	if err := flagset.Parse(args); err != nil {
		return err
	} else {
		if v, err := lib.GetController(u, uint32(controller), options.timeout); err != nil {
			return err
		} else {
			fmt.Printf("get-controller\n")
			fmt.Printf("   %v\n", v)
		}

		return nil
	}
}
