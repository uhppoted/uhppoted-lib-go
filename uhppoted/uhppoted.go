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
	"net"
	"net/netip"
	"time"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/decode"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/encode"
)

type GetControllerResponse = codec.GetControllerResponse

type Uhppoted interface {
	GetAllControllers(timeout time.Duration) ([]GetControllerResponse, error)
}

type uhppoted struct {
	bindAddr      netip.AddrPort
	broadcastAddr netip.AddrPort
	listenAddr    netip.AddrPort
	debug         bool

	udp udp
}

// NewUhppoted creates a new instance of the uhppoted service, configured with the supplied
// local bind address, broadcast address, and listen address. The debug flag enables or
// disables logging of the network packets to the console.
//
// The bind, broadcast, and listen parameters are expected to be valid netip.AddPort
// addresses.
func NewUhppoted(bind, broadcast, listen netip.AddrPort, debug bool) uhppoted {
	return uhppoted{
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
//   - An error if the request could not be encoded or broadcast.
//
// Note: Responses that cannot be decoded are silently ignored.

func (u uhppoted) GetAllControllers(timeout time.Duration) ([]GetControllerResponse, error) {
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
