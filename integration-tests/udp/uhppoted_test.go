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

	test "github.com/uhppoted/uhppoted-lib-go/integration-tests"
	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
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

func setup() (*net.UDPConn, error) {
	bind := netip.MustParseAddrPort("0.0.0.0:50002")

	if socket, err := net.ListenUDP("udp", net.UDPAddrFromAddrPort(bind)); err != nil {
		return nil, err
	} else {
		go func() {
			for {
				buffer := make([]byte, 1024)
				if N, addr, err := socket.ReadFromUDPAddrPort(buffer); err != nil {
					return
				} else {
					if N == 64 {
						for _, m := range test.Messages {
							if slices.Compare(m.Request, buffer[0:64]) == 0 {
								for _, packet := range m.Response {
									socket.WriteMsgUDPAddrPort(packet, nil, addr)
								}
							}
						}
					}
				}
			}
		}()

		return socket, nil
	}
}

func TestInvalidResponse(t *testing.T) {
	controller := lib.Controller{
		ID:       201020304,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	_, err := lib.GetController(u, controller, timeout)

	if err == nil || !errors.Is(err, lib.ErrInvalidResponse) {
		t.Errorf("expected %v error, got:%v", lib.ErrInvalidResponse, err)
	}
}

func TestGetCardRecord(t *testing.T) {
	controller := lib.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
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

func TestGetEventRecord(t *testing.T) {
	controller := lib.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
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
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
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
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
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
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	task := lib.Task{
		Task:      lib.LockDoor,
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

func teardown(socket *net.UDPConn) {
	if socket != nil {
		socket.Close()
	}
}
