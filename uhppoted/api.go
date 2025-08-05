package uhppoted

import (
	"net/netip"
	"time"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/decode"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/encode"
)

// FindControllers retrieves a list of all UHPPOTE controllers accessible on the local LAN.
//
// It broadcasts a UDP `get controller` request to the local network and returns a list of
// decoded responses from controllers that reply within the timeout. Responses that cannot
// be decoded are silently ignored.
func FindControllers(u Uhppoted, timeout time.Duration) ([]GetControllerResponse, error) {
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

//go:generate ../.codegen/bin/codegen API

// SetIP sets the controller IPv4 address, netmask and gateway address.
//
// Parameters:
//   - controller: Either a uint32 controller serial number or a controller struct with the
//     controller serial number, IPv4 address and transport.
//   - address:    controller IPv4 address e.g. 192.168.1.100.
//   - netmask:    controller IPv4 subnet mask e.g. 255.255.255.0.
//   - gateway:    controller IPv4 gateway address e.g. 192.168.1.1.
//   - timeout: The maximum time to wait for a response.
//
// Returns:
//   - A SetIPResponse struct.
//   - An error if the request could not be executed.
func SetIPv4[T TController](u Uhppoted, controller T, address, netmask, gateway netip.Addr, timeout time.Duration) (SetIPv4Response, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetIPv4Request(id, address, netmask, gateway)
	}

	return exec[T, SetIPv4Response](u, controller, f, timeout)
}

// SetTime sets the access controller system date and time.
//
// Parameters:
//   - controller: Either a uint32 controller serial number or a controller struct with the
//     controller serial number, IPv4 address and transport.
//   - datetime:   Date/time to which to set the controller system date/time.
//   - timeout: The maximum time to wait for a response.
//
// Returns:
//   - A SetTimeResponse struct.
//   - An error if the request could not be executed.
func SetTime[T TController](u Uhppoted, controller T, datetime time.Time, timeout time.Duration) (SetTimeResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetTimeRequest(id, datetime)
	}

	return exec[T, SetTimeResponse](u, controller, f, timeout)
}

// SetListener sets the access controller event listener IPv4 address:port and auto-send
// interval.
//
// Parameters:
//   - controller: Either a uint32 controller serial number or a controller struct with the
//     controller serial number, IPv4 address and transport.
//   - listener: IPv4 address:port of host listening for controller events.
//   - interval: status auto-send interval (seconds) for events (0 disables auto-send).
//   - timeout: The maximum time to wait for a response.
//
// Returns:
//   - A SetListenerResponse struct.
//   - An error if the request could not be executed.
func SetListener[T TController](u Uhppoted, controller T, listener netip.AddrPort, interval uint8, timeout time.Duration) (SetListenerResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetListenerRequest(id, listener, interval)
	}

	return exec[T, SetListenerResponse](u, controller, f, timeout)
}
