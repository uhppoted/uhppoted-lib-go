package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"iter"
	"strings"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"
)

func setFirstCard(u uhppoted.Uhppoted, args []string) error {
	var door uint
	var firstcard string

	flagset := flag.NewFlagSet("set-firstcard", flag.ExitOnError)

	flagset.UintVar(&door, "door", 1, "door to configure with first-card access")
	flagset.StringVar(&firstcard, "firstcard", "", "first-card configuration for the door")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if fc, err := parseFirstCard(firstcard); err != nil {
		return fmt.Errorf("invalid firstcard '%v' (%v)", firstcard, err)
	} else {
		f := func(c uint32) (any, error) {
			return uhppoted.SetFirstCard(u, c, uint8(door), fc, options.timeout)
		}

		g := func(c uhppoted.Controller) (any, error) {
			return uhppoted.SetFirstCard(u, c, uint8(door), fc, options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("set-firstcard\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}

func parseFirstCard(v string) (types.FirstCard, error) {
	modes := map[string]types.DoorMode{
		"controlled":      types.Controlled,
		"normally open":   types.NormallyOpen,
		"normally closed": types.NormallyClosed,
		"firstcard only":  types.FirstCardOnly,
	}

	firstcard := types.FirstCard{}

	next, stop := iter.Pull(strings.SplitSeq(v, ","))
	defer stop()

	if token, ok := next(); !ok {
		return firstcard, fmt.Errorf("missing start time")
	} else if hhmm, err := types.ParseHHmm(strings.Trim(token, " ")); err != nil {
		return firstcard, fmt.Errorf("invalid start time (%v)", token)
	} else {
		firstcard.StartTime = hhmm
	}

	if token, ok := next(); !ok {
		return firstcard, fmt.Errorf("missing end time")
	} else if hhmm, err := types.ParseHHmm(strings.Trim(token, " ")); err != nil {
		return firstcard, fmt.Errorf("invalid end time (%v)", strings.Trim(token, " "))
	} else {
		firstcard.EndTime = hhmm
	}

	if token, ok := next(); !ok {
		return firstcard, fmt.Errorf("missing active mode")
	} else if mode, ok := modes[strings.Trim(token, " ")]; !ok {
		return firstcard, fmt.Errorf("invalid active mode (%v)", strings.Trim(token, " "))
	} else {
		firstcard.ActiveMode = mode
	}

	if token, ok := next(); !ok {
		return firstcard, fmt.Errorf("missing inactive mode")
	} else if mode, ok := modes[strings.Trim(token, " ")]; !ok {
		return firstcard, fmt.Errorf("invalid inactive mode (%v)", strings.Trim(token, " "))
	} else {
		firstcard.InactiveMode = mode
	}

	if token, ok := next(); ok {
		token = strings.ToLower(token)

		if strings.Contains(token, "mon") {
			firstcard.Weekdays.Monday = true
		}

		if strings.Contains(token, "tue") {
			firstcard.Weekdays.Tuesday = true
		}

		if strings.Contains(token, "wed") {
			firstcard.Weekdays.Wednesday = true
		}

		if strings.Contains(token, "thu") {
			firstcard.Weekdays.Thursday = true
		}

		if strings.Contains(token, "fri") {
			firstcard.Weekdays.Friday = true
		}

		if strings.Contains(token, "sat") {
			firstcard.Weekdays.Saturday = true
		}

		if strings.Contains(token, "sun") {
			firstcard.Weekdays.Sunday = true
		}
	}

	return firstcard, nil
}
