package entities

import (
	"encoding/json"
	"fmt"
	"time"
)

type DateTime struct {
	year   uint16
	month  time.Month
	day    uint8
	hour   uint8
	minute uint8
	second uint8
}

func NewDateTime(year uint16, month time.Month, day uint8, hour, minute, second uint8) DateTime {
	yyyy := year
	if yyyy < 1 {
		yyyy = 1
	}

	MM := month
	if MM < time.January || MM > time.December {
		MM = time.January
	}

	dd := day
	if dd < 1 {
		dd = 1
	}

	ss := second
	if ss > 59 {
		ss = 59
	}

	mm := minute
	if mm > 59 {
		mm = 59
	}

	hh := hour
	if hh >= 24 {
		hh = 24
		mm = 0
		ss = 0
	}

	return DateTime{
		year:   yyyy,
		month:  MM,
		day:    dd,
		hour:   hh,
		minute: mm,
		second: ss,
	}
}

// MustParseDateTime invokes ParseDateTime and panics on error.
//
// It is intended for use in tests with hard-coded strings.
func MustParseDateTime(s string) DateTime {
	if datetime, err := ParseDateTime(s); err != nil {
		panic(err)
	} else {
		return datetime
	}
}

// Parses a date/time string in "yyyy-mm-dd HH:mm:ss" format, returning a zero value DateTime{} and an
// error if the string is blank or not a valid date.
func ParseDateTime(s string) (DateTime, error) {
	if s == "" {
		return DateTime{}, fmt.Errorf("blank date string")
	} else if datetime, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local); err != nil {
		return DateTime{}, err
	} else {
		year, month, day := datetime.Date()
		hour, minute, second := datetime.Hour(), datetime.Minute(), datetime.Second()

		return NewDateTime(uint16(year), month, uint8(day), uint8(hour), uint8(minute), uint8(second)), nil
	}
}

func (dt DateTime) Year() uint16 {
	if dt.year < 1 {
		return 1
	} else {
		return dt.year
	}
}

// Returns the 'month' field value.
//
// The returned value is constrained to the interval [January..December].
func (dt DateTime) Month() time.Month {
	if dt.month < time.January {
		return time.January
	} else if dt.month > time.December {
		return time.December
	} else {
		return dt.month
	}
}

// Returns the 'day' field value.
//
// The returned value is constrained to the interval [1..].
func (dt DateTime) Day() uint8 {
	if dt.day < 1 {
		return 1
	} else {
		return dt.day
	}
}

// Returns the 'hour' field value.
//
// The returned value is constrained to the interval [0..24].
func (dt DateTime) Hour() uint8 {
	if dt.hour > 24 {
		return 24
	} else {
		return dt.hour
	}
}

// Returns the 'minute' field value.
//
// The returned value is constrained to the interval [0..59].
func (dt DateTime) Minute() uint8 {
	if dt.minute > 59 {
		return 59
	} else {
		return dt.minute
	}
}

// Returns the 'second' field value.
//
// The returned value is constrained to the interval [0..59].
func (dt DateTime) Second() uint8 {
	if dt.second > 59 {
		return 59
	} else {
		return dt.second
	}
}

func (dt DateTime) IsZero() bool {
	return dt.year <= 1 && dt.month <= 1 && dt.day <= 1 && dt.hour == 0 && dt.minute == 0 && dt.second == 0
}

func (dt DateTime) String() string {
	return fmt.Sprintf("%04v-%02v-%02v %02v:%02v:%02v",
		dt.Year(), uint8(dt.Month()), dt.Day(),
		dt.Hour(), dt.Minute(), dt.Second())
}

func (dt DateTime) MarshalJSON() ([]byte, error) {
	if dt.IsZero() {
		return json.Marshal("")
	} else {
		return json.Marshal(fmt.Sprintf("%v", dt))
	}
}

func (dt *DateTime) UnmarshalJSON(bytes []byte) error {
	var s string

	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	} else if s == "" {
		*dt = DateTime{
			year:   1,
			month:  time.January,
			day:    1,
			hour:   0,
			minute: 0,
			second: 0,
		}
	} else if v, err := ParseDateTime(s); err != nil {
		return err
	} else {
		*dt = v
	}

	return nil
}
