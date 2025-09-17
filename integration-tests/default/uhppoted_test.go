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

func string2datetime(v string) time.Time {
	if d, err := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local); err != nil {
		panic(fmt.Sprintf("invalid datetime (%v)", v))
	} else {
		return d
	}
}

func string2date(v string) time.Time {
	if d, err := time.ParseInLocation("2006-01-02", v, time.Local); err != nil {
		panic(fmt.Sprintf("invalid date (%v)", v))
	} else {
		return d
	}
}

func string2HHmm(v string) time.Time {
	if t, err := time.ParseInLocation("15:04", v, time.Local); err != nil {
		panic(fmt.Sprintf("invalid time (%v)", t))
	} else {
		return t
	}
}
