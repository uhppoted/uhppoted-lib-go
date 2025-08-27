package uhppoted

import (
	"fmt"
	"net"
	"net/netip"
	"os"
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
