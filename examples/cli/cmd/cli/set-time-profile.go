package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func setTimeProfile(u lib.Uhppoted, args []string) error {
	var profile uint
	var startDate, _ = time.Parse("2006-01-02", "2025-01-01")
	var endDate, _ = time.Parse("2006-01-02", "2025-12-31")
	var monday = true
	var tuesday = true
	var wednesday = false
	var thursday = false
	var friday = true
	var saturday = false
	var sunday = true
	var segment1start, _ = time.Parse("15:04", "08:30")
	var segment1end, _ = time.Parse("15:04", "11:30")
	var segment2start, _ = time.Parse("15:04", "12:15")
	var segment2end, _ = time.Parse("15:04", "16:30")
	var segment3start, _ = time.Parse("15:04", "17:45")
	var segment3end, _ = time.Parse("15:04", "22:00")
	var linked = uint8(33)

	flagset := flag.NewFlagSet("set-time-profile", flag.ExitOnError)

	flagset.UintVar(&profile, "profile", 0, "profile ID [2..254]")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if profile < 2 || profile > 254 {
		return fmt.Errorf("invalid profile (%v)", profile)
	} else {
		d := lib.NewDate(2025, 3, 2)
		e := lib.NewDate(2025, 11, 29)

		f := func(c uint32) (any, error) {
			return lib.SetTimeProfile(u,
				c,
				uint8(profile),
				startDate,
				endDate,
				monday, tuesday, wednesday, thursday, friday, saturday, sunday,
				segment1start, segment1end,
				segment2start, segment2end,
				segment3start, segment3end,
				linked,
				options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.SetTimeProfile(u,
				c,
				uint8(profile),
				d,
				e,
				monday, tuesday, wednesday, thursday, friday, saturday, sunday,
				segment1start, segment1end,
				segment2start, segment2end,
				segment3start, segment3end,
				linked,
				options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("set-time-profile\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}

func setTimeProfileRecord(u lib.Uhppoted, args []string) error {
	var profile uint
	var startDate, _ = lib.ParseDate("2025-01-01")
	var endDate, _ = lib.ParseDate("2025-12-31")
	var monday = true
	var tuesday = true
	var wednesday = false
	var thursday = false
	var friday = true
	var saturday = false
	var sunday = true
	var segment1start, _ = lib.ParseHHmm("08:30")
	var segment1end, _ = lib.ParseHHmm("11:30")
	var segment2start, _ = lib.ParseHHmm("12:15")
	var segment2end, _ = lib.ParseHHmm("16:30")
	var segment3start, _ = lib.ParseHHmm("17:45")
	var segment3end, _ = lib.ParseHHmm("22:00")
	var linked = uint8(33)

	flagset := flag.NewFlagSet("set-time-profile", flag.ExitOnError)

	flagset.UintVar(&profile, "profile", 0, "profile ID [2..254]")

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else if profile < 2 || profile > 254 {
		return fmt.Errorf("invalid profile (%v)", profile)
	} else {
		record := lib.TimeProfile{
			Profile:   uint8(profile),
			StartDate: startDate,
			EndDate:   endDate,
			Weekdays: lib.Weekdays{
				Monday:    monday,
				Tuesday:   tuesday,
				Wednesday: wednesday,
				Thursday:  thursday,
				Friday:    friday,
				Saturday:  saturday,
				Sunday:    sunday,
			},
			Segments: []lib.TimeSegment{
				{segment1start, segment1end},
				{segment2start, segment2end},
				{segment3start, segment3end},
			},
			LinkedProfile: linked,
		}

		f := func(c uint32) (any, error) {
			return lib.SetTimeProfileRecord(u, c, record, options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.SetTimeProfileRecord(u, c, record, options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("set-time-profile-record\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
