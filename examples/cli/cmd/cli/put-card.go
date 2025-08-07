package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func putCard(u lib.Uhppoted, args []string) error {
	var card uint
	var start string
	var end string
	var permissions string
	var PIN uint

	flagset := flag.NewFlagSet("put-card", flag.ExitOnError)

	flagset.UintVar(&card, "card", 0, "card number")
	flagset.StringVar(&start, "start-date", "", "card 'valid from' date e.g. 2025-01-01")
	flagset.StringVar(&end, "end-date", "", "card 'valid until' date e.g. 2025-12-31")
	flagset.StringVar(&permissions, "doors", "", "card access permissions e.g. 1,3:17,4 allows access to doors 1 and 4, door 3 access is managed by time profile 17")
	flagset.UintVar(&PIN, "PIN", 0, "(optional) PIN code [0..99999]")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if card == 0 || card > 25565535 {
		return fmt.Errorf("invalid card (%v)", card)
	} else if startdate, err := time.ParseInLocation("2006-01-02", start, time.Local); err != nil {
		return fmt.Errorf("invalid start date (%v)", start)
	} else if enddate, err := time.ParseInLocation("2006-01-02", end, time.Local); err != nil {
		return fmt.Errorf("invalid end date (%v)", start)
	} else if doors, err := parseCardPermissions(permissions); err != nil {
		return fmt.Errorf("invalid access permissions (%v)", permissions)
	} else if PIN > 999999 {
		return fmt.Errorf("invalid PIN (%v)", PIN)
	} else {
		f := func(c uint32) (any, error) {
			return lib.PutCard(u, c, uint32(card), startdate, enddate, doors[0], doors[1], doors[2], doors[3], uint32(PIN), options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.PutCard(u, c, uint32(card), startdate, enddate, doors[0], doors[1], doors[2], doors[3], uint32(PIN), options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("put-card\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}

func parseCardPermissions(v string) ([]uint8, error) {
	doors := []uint8{0, 0, 0, 0}

	re := regexp.MustCompile("([1234])(:([0-9]+))?")
	tokens := strings.Split(v, ",")
	for _, token := range tokens {
		match := re.FindStringSubmatch(token)

		if len(match) > 0 {
			if door, err := strconv.ParseUint(match[1], 10, 8); err != nil {
				return doors, err
			} else if len(match) > 3 && match[3] != "" {
				if profile, err := strconv.ParseUint(match[3], 10, 8); err != nil {
					return doors, err
				} else {
					doors[door-1] = uint8(profile)
				}
			} else {
				doors[door-1] = 1
			}
		}
	}

	return doors, nil
}
