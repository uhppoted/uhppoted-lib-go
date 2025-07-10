// Package uhppoted provides a Go wrapper for the UHPPOTE TCP/IP access controller API.
//
// It enables interaction with UHPPOTE controllers on a local LAN using either UDP
// broadcasts, UDP connected sockets or TCP/IP. The package supports querying controller
// information, managing access cards, and controlling doors.
//
// Typical usage:
//
//	client := uhppoted.Uhppoted{
//	    BindAddr:      netip.MustParseAddrPort("0.0.0.0:0"),
//	    BroadcastAddr: netip.MustParseAddrPort("255.255.255.255:60000"),
//	    ListenAddr:    netip.MustParseAddrPort("0.0.0.0:60001"),
//	    Debug:         false,
//	}
//
//	if controllers, err := client.GetAllControllers(2 * time.Second); err != nil {
//	    log.Fatal(err)
//	} else {
//	    for _, c := range controllers {
//	        fmt.Printf("Controller: %v\n", c.SerialNumber)
//	    }
//	}
package uhppoted

import (
	"fmt"
	"net"
	"net/netip"
	"time"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/decode"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/encode"
)

type Uhppoted struct {
	bindAddr      netip.AddrPort
	broadcastAddr netip.AddrPort
	listenAddr    netip.AddrPort
	debug         bool

	udp udp
}

type TController interface {
	~uint32 | Controller
}

type Controller struct {
	ID       uint32
	Address  netip.AddrPort
	Protocol string
}

type GetControllerResponse = codec.GetControllerResponse

// NewUhppoted creates a new instance of the uhppoted service, configured with the supplied
// local bind address, broadcast address, and listen address. The debug flag enables or
// disables logging of the network packets to the console.
//
// The bind, broadcast, and listen parameters are expected to be valid netip.AddPort
// addresses.
func NewUhppoted(bind, broadcast, listen netip.AddrPort, debug bool) Uhppoted {
	return Uhppoted{
		bindAddr:      bind,
		broadcastAddr: broadcast,
		listenAddr:    listen,
		debug:         debug,

		udp: udp{
			bindAddr:      net.UDPAddrFromAddrPort(bind),
			broadcastAddr: net.UDPAddrFromAddrPort(broadcast),
			listenAddr:    net.UDPAddrFromAddrPort(listen),
			debug:         debug,
		},
	}
}

// GetAllControllers retrieves a list of all UHPPOTE controllers accessible on the local LAN.
//
// It broadcasts a `get controller` request to the local network and returns a list of
// decoded responses from controllers that reply within the specified timeout.
//
// Parameters:
//   - timeout: The maximum time to wait for responses.
//
// Returns:
//   - A slice of GetControllerResponse structs, one for each responding controller.
//   - An error if the request could not be encoded or sent.
//
// Note: Responses that cannot be decoded are silently ignored.

func GetAllControllers(u Uhppoted, timeout time.Duration) ([]GetControllerResponse, error) {
	if request, err := encode.GetControllerRequest(0); err != nil {
		return nil, err
	} else if replies, err := u.udp.broadcast(request, timeout); err != nil {
		return nil, err
	} else {
		responses := []GetControllerResponse{}

		for _, reply := range replies {
			if response, err := decode.GetControllerResponse(reply); err == nil {
				responses = append(responses, response)
			}
		}

		return responses, nil
	}
}

// GetController retrieves the controller system information for a single access controller.
//
// Parameters:
//   - controller: Either a uint32 controller serial number or a controller struct with the
//     controller serial number, IPv4 address and transport.
//   - timeout: The maximum time to wait for a response.
//
// Returns:
//   - A GetControllerResponse structs.
//   - An error if the request could not be executed.
func GetController[T TController](u Uhppoted, controller T, timeout time.Duration) (GetControllerResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetControllerRequest(id)
	}

	g := func(b []byte) (GetControllerResponse, error) {
		return decode.GetControllerResponse(b)
	}

	if reply, err := exec[T, GetControllerResponse](u, controller, f, g, timeout); err != nil {
		return GetControllerResponse{}, err
	} else {
		return reply, nil
	}
}

func exec[T TController, R any](u Uhppoted, controller T, f func(id uint32) ([]byte, error), g func(packet []byte) (R, error), timeout time.Duration) (R, error) {
	var zero R

	if c, err := resolve(controller); err != nil {
		return zero, err
	} else if request, err := f(c.ID); err != nil {
		return zero, err
	} else if reply, err := send(u, c, request, timeout); err != nil {
		return zero, err
	} else {
		return g(reply)
	}
}

func send(u Uhppoted, controller Controller, request []byte, timeout time.Duration) ([]byte, error) {
	zero := netip.AddrPort{}

	if controller.Address != zero && !controller.Address.IsValid() {
		return nil, fmt.Errorf("invalid address (%v)", controller.Address)
	} else if controller.Address != zero {
		return u.udp.sendTo(request, controller.Address, timeout)
	} else {
		return u.udp.broadcastTo(request, timeout)
	}
}

func resolve[T TController](controller T) (Controller, error) {
	switch v := any(controller).(type) {
	case uint32:
		return Controller{
			ID: v,
		}, nil

	case Controller:
		return v, nil
	}

	return Controller{}, fmt.Errorf("unsupported type (%T)", controller)
}
