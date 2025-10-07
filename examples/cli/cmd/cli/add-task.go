package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func addTask(u lib.Uhppoted, args []string) error {
	var task = lib.LockDoor
	var startDate, _ = time.Parse("2006-01-02", "2025-01-01")
	var endDate, _ = time.Parse("2006-01-02", "2025-12-31")
	var monday = true
	var tuesday = true
	var wednesday = false
	var thursday = false
	var friday = true
	var saturday = false
	var sunday = true
	var startTime, _ = time.Parse("15:04", "08:30")
	var door = uint8(3)
	var moreCards = uint8(5)

	flagset := flag.NewFlagSet("add-task", flag.ExitOnError)

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else {
		d := lib.NewDate(2025, 3, 2)
		e := lib.NewDate(2025, 11, 29)

		f := func(c uint32) (any, error) {
			return lib.AddTask(u,
				c,
				task,
				startDate, endDate,
				monday, tuesday, wednesday, thursday, friday, saturday, sunday,
				startTime,
				door,
				moreCards,
				options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.AddTask(u,
				c,
				task,
				d, e,
				monday, tuesday, wednesday, thursday, friday, saturday, sunday,
				startTime,
				door,
				moreCards,
				options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("add-task\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}

func addTaskRecord(u lib.Uhppoted, args []string) error {
	var task = lib.LockDoor
	var startDate, _ = lib.ParseDate("2025-01-01")
	var endDate, _ = lib.ParseDate("2025-12-31")
	var startTime, _ = lib.ParseHHmm("08:30")
	var monday = true
	var tuesday = true
	var wednesday = false
	var thursday = false
	var friday = true
	var saturday = false
	var sunday = true
	var door = uint8(3)
	var moreCards = uint8(5)

	flagset := flag.NewFlagSet("add-task-record", flag.ExitOnError)

	if controller, err := parse(flagset, args); err != nil {
		return err
	} else {
		record := lib.Task{
			Task:      task,
			Door:      door,
			StartDate: startDate,
			EndDate:   endDate,
			StartTime: startTime,
			Weekdays: lib.Weekdays{
				Monday:    monday,
				Tuesday:   tuesday,
				Wednesday: wednesday,
				Thursday:  thursday,
				Friday:    friday,
				Saturday:  saturday,
				Sunday:    sunday,
			},
			MoreCards: moreCards,
		}

		f := func(c uint32) (any, error) {
			return lib.AddTaskRecord(u, c, record, options.timeout)
		}

		g := func(c lib.Controller) (any, error) {
			return lib.AddTaskRecord(u, c, record, options.timeout)
		}

		if v, err := exec(controller, flagset, f, g); err != nil {
			return err
		} else if bytes, err := json.MarshalIndent(v, "   ", "   "); err != nil {
			return err
		} else {
			fmt.Printf("add-task-record\n")
			fmt.Printf("   %v\n", string(bytes))
			fmt.Println()

			return nil
		}
	}
}
