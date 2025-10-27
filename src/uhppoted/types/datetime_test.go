package types

import (
	"fmt"
	"testing"
	"time"
)

func TestNewDateTime(t *testing.T) {
	tests := []struct {
		year     uint16
		month    time.Month
		day      uint8
		hour     uint8
		minute   uint8
		second   uint8
		expected DateTime
	}{
		{2025, time.June, 23, 12, 34, 56, DateTime{2025, 6, 23, 12, 34, 56}},
	}

	for _, test := range tests {
		d := NewDateTime(test.year, uint8(test.month), test.day, test.hour, test.minute, test.second)

		if d != test.expected {
			t.Errorf("incorrectly constructed date (%v,%v,%v,%v,%v,%v) - expected:%v, got:%v",
				test.year, test.month, test.day,
				test.hour, test.minute, test.second,
				test.expected, d)
		}
	}

}

func TestDateTimeYear(t *testing.T) {
	tests := []struct {
		datetime DateTime
		expected uint16
	}{
		{DateTime{2025, 6, 23, 12, 34, 56}, 2025},
	}

	for _, test := range tests {
		year := test.datetime.Year()

		if year != test.expected {
			t.Errorf("%v: incorrect datetime 'year' - expected:%v, got:%v", test.datetime, test.expected, year)
		}
	}
}

func TestDateTimeMonth(t *testing.T) {
	tests := []struct {
		datetime DateTime
		expected time.Month
	}{
		{DateTime{2025, 6, 23, 12, 34, 56}, time.June},
	}

	for _, test := range tests {
		month := test.datetime.Month()

		if month != uint8(test.expected) {
			t.Errorf("%v: incorrect datetime 'month' - expected:%v, got:%v", test.datetime, test.expected, month)
		}
	}
}

func TestDateTimeDay(t *testing.T) {
	tests := []struct {
		datetime DateTime
		expected uint8
	}{
		{DateTime{2025, 6, 23, 12, 34, 56}, 23},
	}

	for _, test := range tests {
		day := test.datetime.Day()

		if day != test.expected {
			t.Errorf("%v: incorrect datetime 'day' - expected:%v, got:%v", test.datetime, test.expected, day)
		}
	}
}

func TestDateTimeStringer(t *testing.T) {
	tests := []struct {
		datetime DateTime
		expected string
	}{
		{DateTime{}, "0001-01-01 00:00:00"},
		{DateTime{2025, 6, 23, 12, 34, 56}, "2025-06-23 12:34:56"},
	}

	for _, test := range tests {
		s := fmt.Sprintf("%v", test.datetime)

		if s != test.expected {
			t.Errorf("incorrectly formatted datetime (%v,%v,%v,%v,%v,%v) - expected:%v, got:%v",
				test.datetime.year, test.datetime.month, test.datetime.day,
				test.datetime.hour, test.datetime.minute, test.datetime.second,
				test.expected, s)
		}
	}
}
