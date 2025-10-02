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
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/entities"
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

	expected := lib.Card{
		Card:      10058400,
		StartDate: entities.MustParseDate("2025-01-01"),
		EndDate:   entities.MustParseDate("2025-12-31"),
		Permissions: map[uint8]uint8{
			1: 1,
			2: 0,
			3: 17,
			4: 1,
		},
		PIN: 7531,
	}

	record, err := lib.GetCardRecord(u, controller, 10058400, timeout)

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

	expected := lib.Event{
		Index:         13579,
		Timestamp:     entities.MustParseDateTime("2025-11-17 12:34:56"),
		Event:         2,
		AccessGranted: true,
		Door:          4,
		Direction:     2,
		Card:          10058400,
		Reason:        21,
	}

	record, err := lib.GetEventRecord(u, controller, 13579, timeout)

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
		StartDate: entities.MustParseDate("2025-11-26"),
		EndDate:   entities.MustParseDate("2025-12-29"),
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
				Start: entities.MustParseHHmm("8:30"),
				End:   entities.MustParseHHmm("9:45"),
			},
			{
				Start: entities.MustParseHHmm("11:35"),
				End:   entities.MustParseHHmm("13:15"),
			},
			{
				Start: entities.MustParseHHmm("14:01"),
				End:   entities.MustParseHHmm("17:59"),
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

func teardown(socket *net.TCPListener) {
	if socket != nil {
		socket.Close()
	}
}
