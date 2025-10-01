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

	test "github.com/uhppoted/uhppoted-lib-go/integration-tests"
	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/entities"
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

	record, err := lib.GetCardRecord(u, controller, card, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(record, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, record)
	}
}

func TestGetEventRecord(t *testing.T) {
	controller := uint32(405419896)
	index := uint32(13579)

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

	record, err := lib.GetEventRecord(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(record, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, record)
	}
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

	events := make(chan lib.ListenerEvent)
	errors := make(chan error)
	interrupt := make(chan os.Signal, 1)

	defer close(events)
	defer close(errors)
	defer close(interrupt)

	received := struct {
		events []lib.ListenerEvent
		errors []error
	}{}

	go func() {
		for evt := range events {
			received.events = append(received.events, evt)
		}
	}()

	go func() {
		for err := range errors {
			received.errors = append(received.errors, err)
		}
	}()

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

	if err := lib.Listen(u, events, errors, interrupt); err != nil {
		t.Fatalf("%v", err)
	}

	if !slices.EqualFunc(received.events, expected.events, func(p lib.ListenerEvent, q lib.ListenerEvent) bool {
		return reflect.DeepEqual(p, q)
	}) {
		t.Errorf("event listen error\n   expected: %v\n   got:      %v", expected.events, received.events)
	}

	if !slices.EqualFunc(received.errors, expected.errors, func(p error, q error) bool {
		return reflect.DeepEqual(p, q)
	}) {
		t.Errorf("event listen error\n   expected: %v\n   got:      %v", expected.errors, received.errors)
	}
}

func teardown(socket *net.UDPConn) {
	if socket != nil {
		socket.Close()
	}
}
