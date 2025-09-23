package entities

import (
	"encoding/json"
	"fmt"
	"time"
)

type HHmm struct {
	hour   uint8
	minute uint8
}

func NewHHmm(hour uint8, minute uint8) HHmm {
	mm := minute
	if mm > 59 {
		mm = 59
	}

	hh := hour
	if hh >= 24 {
		hh = 24
		mm = 0
	}

	return HHmm{
		hour:   hh,
		minute: mm,
	}
}

// MustParseHHmm invokes ParseHHmm and panics on error.
//
// It is intended for use in tests with hard-coded strings.
func MustParseHHmm(s string) HHmm {
	if t, err := ParseHHmm(s); err != nil {
		panic(err)
	} else {
		return t
	}
}

// Parses a time string in "HH:mm" format, returning a zero value HHmm{} and an
// error if the string is blank or not a valid date.
func ParseHHmm(s string) (HHmm, error) {
	if s == "" {
		return HHmm{}, fmt.Errorf("blank HHmm string")
	} else if t, err := time.ParseInLocation("15:04", s, time.Local); err != nil {
		return HHmm{}, err
	} else {
		hour, minute := t.Hour(), t.Minute()

		return NewHHmm(uint8(hour), uint8(minute)), nil
	}
}

// Creates an HHmm from the time.Time hour and minute fields.
func HHmmFromTime(t time.Time) HHmm {
	hour := minmax(t.Hour(), 0, 255)
	minute := minmax(t.Minute(), 0, 255)

	return NewHHmm(uint8(hour), uint8(minute))
}

// Returns the 'hour' field value.
//
// The returned value is constrained to the interval [0..24]
func (t HHmm) Hour() uint8 {
	if t.hour > 24 {
		return 24
	} else {
		return t.hour
	}
}

// Returns the 'minute' field value.
//
// The returned value is constrained to the interval [0..59]
func (t HHmm) Minute() uint8 {
	if t.minute > 59 {
		return 59
	} else {
		return t.minute
	}
}

// Returns the time value formatted as HH:mm.
func (t HHmm) String() string {
	return fmt.Sprintf("%02v:%02v", t.Hour(), t.Minute())
}

// Marshals the time value as a JSON string in the format "HH:mm".
func (t HHmm) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%v", t))
}

// Unmarshals a time value from a JSON string formatted as "HH:mm".
func (t *HHmm) UnmarshalJSON(bytes []byte) error {
	var s string

	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	} else if s == "" {
		*t = HHmm{}
	} else if v, err := ParseHHmm(s); err != nil {
		return err
	} else {
		*t = v
	}

	return nil
}
