package main

import (
	"flag"
	"time"
)

func NewFlagSet(name string) *flag.FlagSet {
	var controller uint
	var dest string
	var tcp bool

	flagset := flag.NewFlagSet(name, flag.ExitOnError)

	flagset.UintVar(&controller, "controller", 0, "controller serial number")
	flagset.StringVar(&dest, "dest", "", "controller IPv4 address (optional)")
	flagset.BoolVar(&tcp, "tcp", false, "use TCP/IP transport (optional)")

	return flagset
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
