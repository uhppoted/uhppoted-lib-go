package uhppoted

import (
	"time"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/entities"
)

func NewDateTime(year uint16, month time.Month, day uint8, hour, minute, second uint8) entities.DateTime {
	return entities.NewDateTime(year, month, day, hour, minute, second)
}

func NewDate(year uint16, month time.Month, day uint8) entities.Date {
	return entities.NewDate(year, month, day)
}

func NewTime(hour, minute, second uint8) entities.Time {
	return entities.NewTime(hour, minute, second)
}
