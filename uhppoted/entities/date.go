package entities

import (
	"fmt"
	"time"
)

type Date struct {
	year  uint16
	month time.Month
	day   uint8
}

func NewDate(year uint16, month time.Month, day uint8) Date {
	return Date{
		year:  year,
		month: month,
		day:   day,
	}
}

// MustParseDate invokes ParseDate and panics on error.
//
// It is intended for use in tests with hard-coded strings.
func MustParseDate(s string) Date {
	if date, err := ParseDate(s); err != nil {
		panic(err)
	} else {
		return date
	}
}

// Parses a date string, returning a zero value Date{} and an
// error if the string is blank or not a valid date.
func ParseDate(s string) (Date, error) {
	if s == "" {
		return Date{}, fmt.Errorf("blank date string")
	} else if date, err := time.ParseInLocation("2006-01-02", s, time.Local); err != nil {
		return Date{}, err
	} else {
		year, month, day := date.Date()

		return NewDate(uint16(year), month, uint8(day)), nil
	}
}

func (d Date) String() string {
	return fmt.Sprintf("%04v-%02v-%02v", d.year, uint8(d.month), d.day)
}

func (d Date) Year() uint16 {
	return d.year
}

func (d Date) Month() time.Month {
	return d.month
}

func (d Date) Day() uint8 {
	return d.day
}
