package codec

import (
	"errors"
	"reflect"
	"testing"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"
)

func TestInvalidPacket(t *testing.T) {
	packet := []byte{}
	expected := errors.New("invalid reply packet length (0)")

	if _, err := Decode[responses.GetController](packet); err == nil {
		t.Errorf("expected error '...', got:      %#v", err)
	} else if err.Error() != expected.Error() {
		t.Errorf("incorrect error - expected:%#v, got:%#v", expected, err)
	}
}

func TestInvalidSOM(t *testing.T) {
	packet := []byte{
		0x16, 0x94, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0xc0, 0xa8, 0x01, 0x64, 0xff, 0xff, 0xff, 0x00,
		0xc0, 0xa8, 0x01, 0x01, 0x00, 0x12, 0x23, 0x34, 0x45, 0x56, 0x08, 0x92, 0x20, 0x18, 0x11, 0x05,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	expected := errors.New("invalid reply SOM byte (0x16)")

	if _, err := Decode[responses.GetController](packet); err == nil {
		t.Errorf("expected error '...', got:      %#v", err)
	} else if err.Error() != expected.Error() {
		t.Errorf("incorrect error - expected:%#v, got:%#v", expected, err)
	}
}

func TestInvalidMsgType(t *testing.T) {
	packet := []byte{
		0x17, 0x00, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0xc0, 0xa8, 0x01, 0x64, 0xff, 0xff, 0xff, 0x00,
		0xc0, 0xa8, 0x01, 0x01, 0x00, 0x12, 0x23, 0x34, 0x45, 0x56, 0x08, 0x92, 0x20, 0x18, 0x11, 0x05,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	expected := errors.New("unknown message type (0x00)")

	if _, err := Decode[responses.GetController](packet); err == nil {
		t.Errorf("expected error '...', got:      %#v", err)
	} else if err.Error() != expected.Error() {
		t.Errorf("incorrect error - expected:%#v, got:%#v", expected, err)
	}
}

func TestDecodeGetCardRecord(t *testing.T) {
	packet := []byte{
		0x17, 0x5a, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0xa0, 0x7a, 0x99, 0x00, 0x20, 0x24, 0x01, 0x01,
		0x20, 0x24, 0x12, 0x31, 0x01, 0x00, 0x11, 0x01, 0x3f, 0x42, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	expected := types.Card{
		Card:      10058400,
		StartDate: types.MustParseDate("2024-01-01"),
		EndDate:   types.MustParseDate("2024-12-31"),
		Permissions: map[uint8]uint8{
			1: 1,
			2: 0,
			3: 17,
			4: 1,
		},
		PIN: 999999,
	}

	if response, err := Decode[types.Card](packet); err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrectly decoded response:\n   expected: %#v\n   got:      %#v", expected, response)
	}
}

func TestDecodeGetCardRecordAtIndex(t *testing.T) {
	packet := []byte{
		0x17, 0x5c, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0xa0, 0x7a, 0x99, 0x00, 0x20, 0x24, 0x01, 0x01,
		0x20, 0x24, 0x12, 0x31, 0x01, 0x00, 0x11, 0x01, 0x3f, 0x42, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	expected := types.Card{
		Card:      10058400,
		StartDate: types.MustParseDate("2024-01-01"),
		EndDate:   types.MustParseDate("2024-12-31"),
		Permissions: map[uint8]uint8{
			1: 1,
			2: 0,
			3: 17,
			4: 1,
		},
		PIN: 999999,
	}

	if response, err := Decode[types.Card](packet); err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrectly decoded response:\n   expected: %#v\n   got:      %#v", expected, response)
	}
}

func TestDecodeGetStatusRecord(t *testing.T) {
	packet := []byte{
		0x17, 0x20, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x4e, 0x00, 0x00, 0x00, 0x02, 0x01, 0x03, 0x01,
		0xa1, 0x98, 0x7c, 0x00, 0x20, 0x22, 0x08, 0x23, 0x09, 0x47, 0x06, 0x2c, 0x00, 0x01, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x03, 0x09, 0x49, 0x39, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x27, 0x07, 0x09, 0x22, 0x08, 0x23, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	expected := types.Status{
		System: struct {
			Time  types.DateTime `json:"datetime"`
			Error uint8          `json:"error"`
			Info  uint8          `json:"info"`
		}{
			Time:  types.MustParseDateTime("2022-08-23 09:49:39"),
			Error: 3,
			Info:  39,
		},

		Doors: map[uint8]struct {
			Open     bool `json:"open"`
			Button   bool `json:"button"`
			Unlocked bool `json:"unlocked"`
		}{
			1: struct {
				Open     bool `json:"open"`
				Button   bool `json:"button"`
				Unlocked bool `json:"unlocked"`
			}{
				Open:     false,
				Button:   false,
				Unlocked: true,
			},
			2: struct {
				Open     bool `json:"open"`
				Button   bool `json:"button"`
				Unlocked bool `json:"unlocked"`
			}{
				Open:     true,
				Button:   false,
				Unlocked: true,
			},
			3: struct {
				Open     bool `json:"open"`
				Button   bool `json:"button"`
				Unlocked bool `json:"unlocked"`
			}{
				Open:     false,
				Button:   false,
				Unlocked: true,
			},
			4: struct {
				Open     bool `json:"open"`
				Button   bool `json:"button"`
				Unlocked bool `json:"unlocked"`
			}{
				Open:     false,
				Button:   true,
				Unlocked: false,
			},
		},

		Alarms: struct {
			Flags      uint8 `json:"flags"`
			Fire       bool  `json:"fire"`
			LockForced bool  `json:"lock-forced"`
		}{
			Flags:      0x09,
			Fire:       true,
			LockForced: false,
		},

		Event: types.Event{
			Index:         78,
			Event:         types.EventDoor,
			AccessGranted: true,
			Door:          3,
			Direction:     1,
			Card:          8165537,
			Timestamp:     types.MustParseDateTime("2022-08-23 09:47:06"),
			Reason:        44,
		},
	}

	if response, err := Decode[types.Status](packet); err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrectly decoded response:\n   expected: %#v\n   got:      %#v", expected, response)
	}
}

func TestDecodeGetEventRecord(t *testing.T) {
	packet := []byte{
		0x17, 0xb0, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x0b, 0x35, 0x00, 0x00, 0x02, 0x01, 0x04, 0x02,
		0xa0, 0x7a, 0x99, 0x00, 0x20, 0x25, 0x11, 0x17, 0x12, 0x34, 0x56, 0x15, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	expected := types.Event{
		Index:         13579,
		Event:         types.EventType(2),
		AccessGranted: true,
		Door:          4,
		Direction:     2,
		Card:          10058400,
		Timestamp:     types.MustParseDateTime("2025-11-17 12:34:56"),
		Reason:        21,
	}

	if response, err := Decode[types.Event](packet); err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrectly decoded response:\n   expected: %#v\n   got:      %#v", expected, response)
	}
}

func TestDecodeGetTimeProfileRecord(t *testing.T) {
	packet := []byte{
		0x17, 0x98, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x25, 0x20, 0x25, 0x11, 0x26, 0x20, 0x25, 0x12,
		0x29, 0x01, 0x01, 0x00, 0x01, 0x00, 0x01, 0x01, 0x08, 0x30, 0x09, 0x45, 0x11, 0x35, 0x13, 0x15,
		0x14, 0x01, 0x17, 0x59, 0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	expected := types.TimeProfile{
		Profile:   37,
		StartDate: types.MustParseDate("2025-11-26"),
		EndDate:   types.MustParseDate("2025-12-29"),
		Weekdays: types.Weekdays{
			Monday:    true,
			Tuesday:   true,
			Wednesday: false,
			Thursday:  true,
			Friday:    false,
			Saturday:  true,
			Sunday:    true,
		},
		Segments: []types.TimeSegment{
			{
				Start: types.MustParseHHmm("08:30"),
				End:   types.MustParseHHmm("09:45"),
			},
			{
				Start: types.MustParseHHmm("11:35"),
				End:   types.MustParseHHmm("13:15"),
			},
			{
				Start: types.MustParseHHmm("14:01"),
				End:   types.MustParseHHmm("17:59"),
			},
		},
		LinkedProfile: 19,
	}

	if response, err := Decode[types.TimeProfile](packet); err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrectly decoded response:\n   expected: %#v\n   got:      %#v", expected, response)
	}
}

func TestDecodeListenerEvent(t *testing.T) {
	packet := []byte{
		0x17, 0x20, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x4e, 0x00, 0x00, 0x00, 0x02, 0x01, 0x03, 0x01,
		0xa1, 0x98, 0x7c, 0x00, 0x20, 0x22, 0x08, 0x23, 0x09, 0x47, 0x06, 0x2c, 0x00, 0x01, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x03, 0x09, 0x49, 0x39, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x27, 0x07, 0x09, 0x22, 0x08, 0x23, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	expected := responses.ListenerEvent{
		Controller:         405419896,
		SystemDate:         types.MustParseDate("2022-08-23"),
		SystemTime:         types.MustParseTime("09:49:39"),
		Door1Open:          false,
		Door2Open:          true,
		Door3Open:          false,
		Door4Open:          false,
		Door1Button:        false,
		Door2Button:        false,
		Door3Button:        false,
		Door4Button:        true,
		Relays:             7,
		Inputs:             9,
		SystemError:        3,
		SpecialInfo:        39,
		EventIndex:         78,
		EventType:          types.EventType(2),
		EventAccessGranted: true,
		EventDoor:          3,
		EventDirection:     1,
		EventCard:          8165537,
		EventTimestamp:     types.MustParseDateTime("2022-08-23 09:47:06"),
		EventReason:        44,
		SequenceNo:         0,
	}

	if response, err := Decode[responses.ListenerEvent](packet); err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrectly decoded response:\n   expected: %#v\n   got:      %#v", expected, response)
	}
}

func TestDecodeListenerEventV6_62(t *testing.T) {
	packet := []byte{
		0x19, 0x20, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x4e, 0x00, 0x00, 0x00, 0x02, 0x01, 0x03, 0x01,
		0xa1, 0x98, 0x7c, 0x00, 0x20, 0x22, 0x08, 0x23, 0x09, 0x47, 0x06, 0x2c, 0x00, 0x01, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x03, 0x09, 0x49, 0x39, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x27, 0x07, 0x09, 0x22, 0x08, 0x23, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	expected := responses.ListenerEvent{
		Controller:         405419896,
		SystemDate:         types.MustParseDate("2022-08-23"),
		SystemTime:         types.MustParseTime("09:49:39"),
		Door1Open:          false,
		Door2Open:          true,
		Door3Open:          false,
		Door4Open:          false,
		Door1Button:        false,
		Door2Button:        false,
		Door3Button:        false,
		Door4Button:        true,
		Relays:             7,
		Inputs:             9,
		SystemError:        3,
		SpecialInfo:        39,
		EventIndex:         78,
		EventType:          types.EventType(2),
		EventAccessGranted: true,
		EventDoor:          3,
		EventDirection:     1,
		EventCard:          8165537,
		EventTimestamp:     types.MustParseDateTime("2022-08-23 09:47:06"),
		EventReason:        44,
		SequenceNo:         0,
	}

	if response, err := Decode[responses.ListenerEvent](packet); err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrectly decoded response:\n   expected: %#v\n   got:      %#v", expected, response)
	}
}
