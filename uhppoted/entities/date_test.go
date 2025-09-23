package entities

import (
	"fmt"
	"testing"
	"time"
)

func TestNewDate(t *testing.T) {
	tests := []struct {
		year     uint16
		month    time.Month
		day      uint8
		expected Date
	}{
		{2025, time.June, 23, Date{2025, 6, 23}},
	}

	for _, test := range tests {
		d := NewDate(test.year, uint8(test.month), test.day)

		if d != test.expected {
			t.Errorf("incorrectly constructed date (%v,%v,%v) - expected:%v, got:%v", test.year, test.month, test.day, test.expected, d)
		}
	}

}

func TestDateYear(t *testing.T) {
	tests := []struct {
		date     Date
		expected uint16
	}{
		{Date{2025, 6, 23}, 2025},
	}

	for _, test := range tests {
		year := test.date.Year()

		if year != test.expected {
			t.Errorf("%v: incorrect date 'year' - expected:%v, got:%v", test.date, test.expected, year)
		}
	}
}

func TestDateMonth(t *testing.T) {
	tests := []struct {
		date     Date
		expected time.Month
	}{
		{Date{2025, 6, 23}, time.June},
	}

	for _, test := range tests {
		month := test.date.Month()

		if month != uint8(test.expected) {
			t.Errorf("%v: incorrect date 'month' - expected:%v, got:%v", test.date, test.expected, month)
		}
	}
}

func TestDateDay(t *testing.T) {
	tests := []struct {
		date     Date
		expected uint8
	}{
		{Date{2025, 6, 23}, 23},
	}

	for _, test := range tests {
		day := test.date.Day()

		if day != test.expected {
			t.Errorf("%v: incorrect date 'day' - expected:%v, got:%v", test.date, test.expected, day)
		}
	}
}

func TestDateStringer(t *testing.T) {
	tests := []struct {
		date     Date
		expected string
	}{
		{Date{}, "0001-01-01"},
		{Date{2025, 6, 23}, "2025-06-23"},
	}

	for _, test := range tests {
		s := fmt.Sprintf("%v", test.date)

		if s != test.expected {
			t.Errorf("incorrectly formatted date (%v,%v,%v) - expected:%v, got:%v", test.date.year, test.date.month, test.date.day, test.expected, s)
		}
	}
}
