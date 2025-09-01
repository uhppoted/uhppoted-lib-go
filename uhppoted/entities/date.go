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
