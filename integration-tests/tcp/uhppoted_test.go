package uhppoted

import (
	"errors"
	"fmt"
	"net"
	"net/netip"
	"os"
	"reflect"
	"slices"
	"testing"
	"time"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/entities"
	test "integration-tests"
)

var bind = netip.MustParseAddrPort("0.0.0.0:0")
var broadcast = netip.MustParseAddrPort("255.255.255.255:50001")
var listen = netip.MustParseAddrPort("0.0.0.0:60001")
var u = lib.NewUhppoted(bind, broadcast, listen, false)

const timeout = 1000 * time.Millisecond

func TestMain(m *testing.M) {
	if socket, err := setup(); err != nil {
		fmt.Printf("*** ERROR %v\n", err)
		os.Exit(1)
	} else {
		code := m.Run()
		teardown(socket)
		os.Exit(code)
	}
}

func setup() (*net.TCPListener, error) {
	bind := netip.MustParseAddrPort("0.0.0.0:50003")

	if socket, err := net.ListenTCP("tcp", net.TCPAddrFromAddrPort(bind)); err != nil {
		return nil, err
	} else {
		go func() {
			for {
				if client, err := socket.AcceptTCP(); err == nil {
					go func() {
						defer client.Close()

						buffer := make([]byte, 1024)
						if N, err := client.Read(buffer); err != nil {
							return
						} else if N == 64 {
							for _, m := range test.Messages {
								if slices.Compare(m.Request, buffer[0:64]) == 0 {
									for _, packet := range m.Response {
										client.Write(packet)
									}
								}
							}
						}
					}()
				}
			}
		}()

		return socket, nil
	}
}

func TestInvalidResponse(t *testing.T) {
	controller := lib.Controller{
		ID:       201020304,
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	_, err := lib.GetController(u, controller, timeout)

	if err == nil || !errors.Is(err, lib.ErrInvalidResponse) {
		t.Errorf("expected %v error, got:%v", lib.ErrInvalidResponse, err)
	}
}

func TestGetCardRecord(t *testing.T) {
	controller := lib.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	card := uint32(10058400)

	expected := lib.Card{
		Card:      10058400,
		StartDate: lib.MustParseDate("2025-01-01"),
		EndDate:   lib.MustParseDate("2025-12-31"),
		Permissions: map[uint8]uint8{
			1: 1,
			2: 0,
			3: 17,
			4: 1,
		},
		PIN: 7531,
	}

	record, err := lib.GetCardRecord(u, controller, card, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(record, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, record)
	}
}

func TestGetStatusRecord(t *testing.T) {
	controller := lib.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	expected := lib.Status{
		System: struct {
			Time  lib.DateTime `json:"datetime"`
			Error uint8        `json:"error"`
			Info  uint8        `json:"info"`
		}{
			Time:  lib.MustParseDateTime("2022-08-23 09:49:39"),
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

		Event: lib.Event{
			Index:         78,
			Event:         lib.EventDoor,
			AccessGranted: true,
			Door:          3,
			Direction:     1,
			Card:          8165537,
			Timestamp:     lib.MustParseDateTime("2022-08-23 09:47:06"),
			Reason:        44,
		},
	}

	record, err := lib.GetStatusRecord(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(record, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, record)
	}
}

func TestGetEventRecord(t *testing.T) {
	controller := lib.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	index := uint32(13579)

	expected := lib.Event{
		Index:         13579,
		Timestamp:     lib.MustParseDateTime("2025-11-17 12:34:56"),
		Event:         2,
		AccessGranted: true,
		Door:          4,
		Direction:     2,
		Card:          10058400,
		Reason:        21,
	}

	record, err := lib.GetEventRecord(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(record, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, record)
	}
}

func TestGetTimeProfileRecord(t *testing.T) {
	controller := lib.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	profile := uint8(37)

	expected := lib.TimeProfile{
		Profile:   37,
		StartDate: lib.MustParseDate("2025-11-26"),
		EndDate:   lib.MustParseDate("2025-12-29"),
		Weekdays: lib.Weekdays{
			Monday:    true,
			Tuesday:   true,
			Wednesday: false,
			Thursday:  true,
			Friday:    false,
			Saturday:  true,
			Sunday:    true,
		},
		Segments: []lib.TimeSegment{
			{
				Start: lib.MustParseHHmm("08:30"),
				End:   lib.MustParseHHmm("09:45"),
			},
			{
				Start: lib.MustParseHHmm("11:35"),
				End:   lib.MustParseHHmm("13:15"),
			},
			{
				Start: lib.MustParseHHmm("14:01"),
				End:   lib.MustParseHHmm("17:59"),
			},
		},
		LinkedProfile: 19,
	}

	record, err := lib.GetTimeProfileRecord(u, controller, profile, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(record, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, record)
	}
}

func TestSetTimeProfileRecord(t *testing.T) {
	controller := lib.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	record := lib.TimeProfile{
		Profile:   37,
		StartDate: lib.MustParseDate("2025-11-26"),
		EndDate:   lib.MustParseDate("2025-12-29"),
		Weekdays: lib.Weekdays{
			Monday:    true,
			Tuesday:   true,
			Wednesday: false,
			Thursday:  true,
			Friday:    false,
			Saturday:  true,
			Sunday:    true,
		},
		Segments: []lib.TimeSegment{
			{
				Start: lib.MustParseHHmm("8:30"),
				End:   lib.MustParseHHmm("9:45"),
			},
			{
				Start: lib.MustParseHHmm("11:35"),
				End:   lib.MustParseHHmm("13:15"),
			},
			{
				Start: lib.MustParseHHmm("14:01"),
				End:   lib.MustParseHHmm("17:59"),
			},
		},
		LinkedProfile: 19,
	}

	ok, err := lib.SetTimeProfileRecord(u, controller, record, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !ok {
		t.Errorf("incorrect response\n   expected:%v\n   got:     %v", true, ok)
	}
}

func TestAddTaskRecord(t *testing.T) {
	controller := lib.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	task := entities.Task{
		Task:      entities.LockDoor,
		StartDate: lib.MustParseDate("2025-01-01"),
		EndDate:   lib.MustParseDate("2025-12-31"),
		StartTime: lib.MustParseHHmm("08:45"),
		Weekdays: lib.Weekdays{
			Monday:    true,
			Tuesday:   true,
			Wednesday: false,
			Thursday:  true,
			Friday:    false,
			Saturday:  true,
			Sunday:    true,
		},
		Door:      3,
		MoreCards: 7,
	}

	ok, err := lib.AddTaskRecord(u, controller, task, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !ok {
		t.Errorf("incorrect response\n   expected:%v\n   got:     %v", true, ok)
	}
}

func teardown(socket *net.TCPListener) {
	if socket != nil {
		socket.Close()
	}
}
