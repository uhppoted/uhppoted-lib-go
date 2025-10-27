package uhppoted

import (
	"errors"
	"fmt"
	"net"
	"net/netip"
	"os"
	"reflect"
	"slices"
	"syscall"
	"testing"
	"time"

	lib "github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"
	test "integration-tests"
)

var bind = netip.MustParseAddrPort("0.0.0.0:0")
var broadcast = netip.MustParseAddrPort("255.255.255.255:50001")
var listen = netip.MustParseAddrPort("0.0.0.0:60005")
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

func setup() (*net.UDPConn, error) {
	bind := netip.MustParseAddrPort("0.0.0.0:50001")

	if socket, err := net.ListenUDP("udp", net.UDPAddrFromAddrPort(bind)); err != nil {
		return nil, err
	} else {
		go func() {
			for {
				buffer := make([]byte, 1024)
				if N, addr, err := socket.ReadFromUDPAddrPort(buffer); err != nil {
					return
				} else if N == 64 {
					for _, m := range test.Messages {
						if slices.Compare(m.Request, buffer[0:64]) == 0 {
							for _, packet := range m.Response {
								socket.WriteMsgUDPAddrPort(packet, nil, addr)
							}
						}
					}
				}
			}
		}()

		return socket, nil
	}
}

func teardown(socket *net.UDPConn) {
	if socket != nil {
		socket.Close()
	}
}

func TestInvalidResponse(t *testing.T) {
	controller := uint32(201020304)

	_, err := lib.GetController(u, controller, timeout)

	if err == nil || !errors.Is(err, lib.ErrInvalidResponse) {
		t.Errorf("expected %v error, got:%v", lib.ErrInvalidResponse, err)
	}
}

func TestGetCardRecord(t *testing.T) {
	controller := uint32(405419896)
	card := uint32(10058400)

	expected := types.Card{
		Card:      10058400,
		StartDate: types.MustParseDate("2025-01-01"),
		EndDate:   types.MustParseDate("2025-12-31"),
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
	controller := uint32(405419896)

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

	record, err := lib.GetStatusRecord(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(record, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, record)
	}
}

func TestGetEventRecord(t *testing.T) {
	controller := uint32(405419896)
	index := uint32(13579)

	expected := types.Event{
		Index:         13579,
		Timestamp:     types.MustParseDateTime("2025-11-17 12:34:56"),
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
	controller := uint32(405419896)
	profile := uint8(37)

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

	record, err := lib.GetTimeProfileRecord(u, controller, profile, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(record, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, record)
	}
}

func TestSetTimeProfileRecord(t *testing.T) {
	controller := uint32(405419896)

	record := types.TimeProfile{
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
				Start: types.MustParseHHmm("8:30"),
				End:   types.MustParseHHmm("9:45"),
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

	ok, err := lib.SetTimeProfileRecord(u, controller, record, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !ok {
		t.Errorf("incorrect response\n   expected:%v\n   got:     %v", true, ok)
	}
}

func TestAddTaskRecord(t *testing.T) {
	controller := uint32(405419896)

	task := types.Task{
		Task:      types.LockDoor,
		StartDate: types.MustParseDate("2025-01-01"),
		EndDate:   types.MustParseDate("2025-12-31"),
		StartTime: types.MustParseHHmm("08:45"),
		Weekdays: types.Weekdays{
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

type listener struct {
	events []lib.ListenerEvent
	errors []error
}

func (l *listener) OnEvent(evt lib.ListenerEvent) {
	l.events = append(l.events, evt)
}

func (l *listener) OnError(err error) {
	l.errors = append(l.errors, err)
}

func TestListen(t *testing.T) {
	expected := struct {
		events []lib.ListenerEvent
		errors []error
	}{}

	for _, v := range test.Events {
		if v.Event != nil {
			expected.events = append(expected.events, *v.Event)
		}

		if v.Error != nil {
			expected.errors = append(expected.errors, v.Error)
		}
	}

	interrupt := make(chan os.Signal, 1)

	defer close(interrupt)

	l := listener{}

	go func() {
		for _, v := range test.Events {

			time.Sleep(100 * time.Millisecond)

			if socket, err := net.DialUDP("udp", net.UDPAddrFromAddrPort(bind), net.UDPAddrFromAddrPort(listen)); err != nil {
				t.Errorf("%v", err)
			} else {
				defer socket.Close()

				if _, err := socket.Write(v.Packet); err != nil {
					t.Errorf("%v", err)
				}
			}
		}

		time.Sleep(1 * time.Second)
		interrupt <- syscall.SIGINT
	}()

	if err := lib.Listen(u, &l, interrupt); err != nil {
		t.Fatalf("%v", err)
	}

	if !slices.EqualFunc(l.events, expected.events, func(p lib.ListenerEvent, q lib.ListenerEvent) bool {
		return reflect.DeepEqual(p, q)
	}) {
		t.Errorf("event listen error\n   expected: %v\n   got:      %v", expected.events, l.events)
	}

	if !slices.EqualFunc(l.errors, expected.errors, func(p error, q error) bool {
		return reflect.DeepEqual(p, q)
	}) {
		t.Errorf("event listen error\n   expected: %v\n   got:      %v", expected.errors, l.errors)
	}
}
