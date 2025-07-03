// Package uhppoted provides a Go wrapper for the UHPPOTE TCP/IP access controller API.
//
// It enables interaction with UHPPOTE controllers on a local LAN using either UDP
// broadcasts, UDP connected sockets or TCP/IP. The package supports querying controller
// information, managing access cards, and controlling doors.
//
// Typical usage:
//
//     client := uhppoted.Uhppoted{
//         BindAddr:      netip.MustParseAddrPort("0.0.0.0:0"),
//         BroadcastAddr: netip.MustParseAddrPort("255.255.255.255:60000"),
//         ListenAddr:    netip.MustParseAddrPort("0.0.0.0:60001"),
//         Debug:         false,
//     }
//
//     if controllers, err := client.GetAllControllers(2 * time.Second); err != nil {
//         log.Fatal(err)
//     } else {
//         for _, c := range controllers {
//             fmt.Printf("Controller: %v\n", c.SerialNumber)
//         }
//     }
package uhppoted

import (
	"net/netip"
	"time"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/decode"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/encode"
)

type GetControllerResponse = codec.GetControllerResponse

type Uhppoted struct {
	BindAddr      netip.AddrPort
	BroadcastAddr netip.AddrPort
	ListenAddr    netip.AddrPort
	Debug         bool

	udp udp
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

// GetAllControllers retrieves a list of all UHPPOTE controllers accessible on the local LAN.
//
// It broadcasts a `get controller` request to the local network and returns a list of
// responses from devices that reply within the specified timeout.
//
// Parameters:
//   - timeout: Optional operation timeout. If zero, a default of 2.5 seconds should be used.
//
// Returns:
//   - A slice of GetControllerResponse objects (as `any`), one for each responding controller.
//   - An error if the request could not be encoded or broadcast.
//
// Note: responses that cannot be decoded are silently ignored.
func (u Uhppoted) GetAllControllers(timeout time.Duration) ([]GetControllerResponse, error) {
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
