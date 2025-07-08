package uhppoted

import (
	"fmt"
	"net"
	"net/netip"
	"os"
	"slices"
	"testing"
	"time"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

var bind = netip.MustParseAddrPort("0.0.0.0:0")
var broadcast = netip.MustParseAddrPort("255.255.255.255:50001")
var listen = netip.MustParseAddrPort("0.0.0.0:60001")

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
					for _, m := range messages {
						if slices.Compare(m.request, buffer[0:64]) == 0 {
							for _, packet := range m.response {
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

func TestGetAllControllers(t *testing.T) {
	u := lib.NewUhppoted(bind, broadcast, listen, true)

	controllers, err := u.GetAllControllers(1000 * time.Millisecond)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !slices.Equal(controllers, expected.getAllControllers) {
		t.Error("incorrect response")
	}
}
