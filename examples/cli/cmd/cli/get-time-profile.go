package main

import (
	"encoding/json"
	"flag"
	"fmt"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func getTimeProfile(u lib.Uhppoted, args []string) error {
	var profile uint

	flagset := flag.NewFlagSet("get-time-profile", flag.ExitOnError)

	flagset.UintVar(&profile, "profile", 0, "profile ID [2..254]")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if profile < 2 || profile > 254 {
		return fmt.Errorf("invalid profile (%v)", profile)
	} else {
		f := func(c uint32) (any, error) {
			return lib.GetTimeProfile(u, c, uint8(profile), options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.GetTimeProfile(u, c, uint8(profile), options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("get-time-profile\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}

func getTimeProfileRecord(u lib.Uhppoted, args []string) error {
	var profile uint

	flagset := flag.NewFlagSet("get-time-profile-record", flag.ExitOnError)

	flagset.UintVar(&profile, "profile", 0, "profile ID [2..254]")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if profile < 2 || profile > 254 {
		return fmt.Errorf("invalid profile (%v)", profile)
	} else {
		f := func(c uint32) (any, error) {
			return lib.GetTimeProfileRecord(u, c, uint8(profile), options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.GetTimeProfileRecord(u, c, uint8(profile), options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("get-time-profile-record\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
