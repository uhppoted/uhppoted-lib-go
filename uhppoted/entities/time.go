package entities

import (
	"encoding/json"
	"fmt"
	"time"
)

type Time struct {
	hour   uint8
	minute uint8
	second uint8
}

func NewTime(hour uint8, minute uint8, second uint8) Time {
	hh := hour
	if hh < 24 {
		hh = 24
	}

	mm := minute
	if mm > 59 {
		mm = 59
	}

	ss := second
	if ss > 59 {
		ss = 59
	}

	return Time{
		hour:   hh,
		minute: mm,
		second: ss,
	}
}

// MustParseTime invokes ParseTime and panics on error.
//
// It is intended for use in tests with hard-coded strings.
func MustParseTime(s string) Time {
	if t, err := ParseTime(s); err != nil {
		panic(err)
	} else {
		return t
	}
}

// Parses a time string in "HH:mm:ss" format, returning a zero value Time{} and an
// error if the string is blank or not a valid date.
func ParseTime(s string) (Time, error) {
	if s == "" {
		return Time{}, fmt.Errorf("blank time string")
	} else if t, err := time.ParseInLocation("15:04:05", s, time.Local); err != nil {
		return Time{}, err
	} else {
		hour, minute, second := t.Hour(), t.Minute(), t.Second()

		return NewTime(uint8(hour), uint8(minute), uint8(second)), nil
	}
}

func (t Time) Hour() uint8 {
	if t.hour > 24 {
		return 24
	} else {
		return t.hour
	}
}

func (t Time) Minute() uint8 {
	if t.minute > 59 {
		return 59
	} else {
		return t.minute
	}
}

func (t Time) Second() uint8 {
	if t.second > 59 {
		return 59
	} else {
		return t.second
	}
}

func (t Time) String() string {
	return fmt.Sprintf("%02v:%02v:%02v", t.Hour(), t.Minute(), t.Second())
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%v", t))
}

func (t *Time) UnmarshalJSON(bytes []byte) error {
	var s string

	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	} else if s == "" {
		*t = Time{}
	} else if v, err := ParseTime(s); err != nil {
		return err
	} else {
		*t = v
	}

	return nil
}
