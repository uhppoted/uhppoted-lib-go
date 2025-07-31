package main

import (
	"flag"
	"strconv"
	"strings"
	"time"
)

type controller struct {
	controller uint
	dest       string
	tcp        bool
}

func parse(flagset *flag.FlagSet, args []string) (controller, error) {
	v := controller{}

	flagset.UintVar(&v.controller, "controller", 0, "controller serial number")
	flagset.StringVar(&v.dest, "dest", "", "controller IPv4 address (optional)")
	flagset.BoolVar(&v.tcp, "tcp", false, "use TCP/IP transport (optional)")

	if err := flagset.Parse(args); err != nil {
		return v, err
	} else {
		return v, nil
	}
}

func parseDateTime(v string) (time.Time, error) {
	if v == "" {
		return time.Now(), nil
	} else if datetime, err := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local); err != nil {
		return time.Now(), err
	} else {
		return datetime, nil
	}
}

func parsePasscodes(v string) []uint32 {
	passcodes := []uint32{}
	tokens := strings.Split(v, ",")

	for _, token := range tokens {
		if u32, err := strconv.ParseUint(token, 10, 32); err == nil && u32 < 1000000 {
			passcodes = append(passcodes, uint32(u32))
		}
	}

	return passcodes
}
