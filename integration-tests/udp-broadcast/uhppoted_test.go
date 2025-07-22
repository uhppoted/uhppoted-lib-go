package uhppoted

import (
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

var u = lib.NewUhppoted(bind, broadcast, listen, true)
var controller = uint32(405419896)

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

func TestGetAllControllers(t *testing.T) {
	controllers, err := lib.GetAllControllers(u, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !slices.Equal(controllers, test.Expected.GetAllControllers) {
		t.Error("incorrect response")
	}
}

func TestGetController(t *testing.T) {
	c, err := lib.GetController(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(c, test.Expected.GetController) {
		t.Error("incorrect response")
	}
}

func TestSetIPv4(t *testing.T) {
	address := netip.MustParseAddr("192.168.1.125")
	netmask := netip.MustParseAddr("25.255.255.0")
	gateway := netip.MustParseAddr("192.168.1.1")

	c, err := lib.SetIPv4(u, controller, address, netmask, gateway, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(c, test.Expected.SetIPv4) {
		t.Error("incorrect response")
	}
}

func TestGetStatus(t *testing.T) {
	c, err := lib.GetStatus(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(c, test.Expected.GetStatus) {
		t.Error("incorrect response")
	}
}

func TestGetTime(t *testing.T) {
	c, err := lib.GetTime(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(c, test.Expected.GetTime) {
		t.Error("incorrect response")
	}
}
