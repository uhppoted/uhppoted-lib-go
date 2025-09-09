package entities

import (
	"encoding/json"
	"fmt"
	"time"
)

type Date struct {
	year  uint16
	month time.Month
	day   uint8
}

func NewDate(year uint16, month time.Month, day uint8) Date {
	yyyy := year
	if yyyy < 1 {
		yyyy = 1
	}

	mm := month
	if mm < time.January || mm > time.December {
		mm = time.January
	}

	dd := day
	if dd < 1 {
		dd = 1
	}

	return Date{
		year:  yyyy,
		month: mm,
		day:   dd,
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

// Parses a date string in "yyyy-mm-dd" format, returning a zero value Date{} and an
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

func (d Date) Year() uint16 {
	if d.year < 1 {
		return 1
	} else {
		return d.year
	}
}

func (d Date) Month() time.Month {
	if d.month < time.January {
		return time.January
	} else if d.month > time.December {
		return time.December
	} else {
		return d.month
	}
}

func (d Date) Day() uint8 {
	if d.day < 1 {
		return 1
	} else {
		return d.day
	}
}

func (d Date) IsZero() bool {
	return d.year <= 1 && d.month <= 1 && d.day <= 1
}

func (d Date) String() string {
	return fmt.Sprintf("%04v-%02v-%02v", d.Year(), uint8(d.Month()), d.Day())
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return json.Marshal("")
	} else {
		return json.Marshal(fmt.Sprintf("%v", d))
	}
}

func (d *Date) UnmarshalJSON(bytes []byte) error {
	var s string

	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	} else if s == "" {
		*d = Date{
			year:  1,
			month: time.January,
			day:   1,
		}
	} else if v, err := ParseDate(s); err != nil {
		return err
	} else {
		*d = v
	}

	return nil
}
