package uhppoted

import (
	"time"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/entities"
)

// Constructs a DateTime entity.
func NewDateTime(year uint16, month uint8, day uint8, hour, minute, second uint8) entities.DateTime {
	return entities.NewDateTime(year, month, day, hour, minute, second)
}

// Converts a Go time.Time to a DateTime entity.
func DateTimeFromTime(t time.Time) entities.DateTime {
	return entities.DateTimeFromTime(t)
}

// Converts a string in the format "yyyy-mm-dd HH:mm:ss" to a DateTime entity.
func ParseDateTime(s string) (entities.DateTime, error) {
	return entities.ParseDateTime(s)
}

// Converts a string in the format "yyyy-mm-dd HH:mm:ss" to a DateTime entity. Panics if the
// string cannot be converted.
func MustParseDateTime(s string) entities.DateTime {
	return entities.MustParseDateTime(s)
}

func NewDate(year uint16, month uint8, day uint8) entities.Date {
	return entities.NewDate(year, month, day)
}

// Converts a Go time.Time to a Date entity.
func DateFromTime(t time.Time) entities.Date {
	return entities.DateFromTime(t)
}

// Converts a "yyyy-mm-dd" formatted string to a Date entity.
func ParseDate(s string) (entities.Date, error) {
	return entities.ParseDate(s)
}

// Converts a "yyyy-mm-dd" formatted string to a Date entity. Panics if the
// string cannot be converted.
func MustParseDate(s string) entities.Date {
	return entities.MustParseDate(s)
}

func NewTime(hour, minute, second uint8) entities.Time {
	return entities.NewTime(hour, minute, second)
}

// Converts an "HH:mm" formatted string to an HHmm entity.
func ParseHHmm(s string) (entities.HHmm, error) {
	return entities.ParseHHmm(s)
}

// Converts an "HH:mm" formatted string to an HHmm entity. Panics if the string
// cannot be converted.
func MustParseHHmm(s string) entities.HHmm {
	return entities.MustParseHHmm(s)
}
