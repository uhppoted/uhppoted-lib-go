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

// Retrieves the access controller event listener IPv4 address:port and auto-send interval.
func GetListenerAddrPort[T TController](u Uhppoted, controller T, timeout time.Duration) (GetListenerAddrPortResponse, error) {
	var zero GetListenerAddrPortResponse

	if c, err := resolve(controller); err != nil {
		return zero, err
	} else if request, err := encode.GetListenerAddrPortRequest(c.ID); err != nil {
		return zero, err
	} else if reply, err := send(u, c, request, timeout); err != nil {
		return zero, err
	} else {
		return decode.GetListenerAddrPortResponse(reply)
	}
}

// Sets the access controller event listener IPv4 address:port and auto-send interval.
func SetListenerAddrPort[T TController](u Uhppoted, controller T, address netip.AddrPort, interval uint8, timeout time.Duration) (SetListenerAddrPortResponse, error) {
	var zero SetListenerAddrPortResponse

	if c, err := resolve(controller); err != nil {
		return zero, err
	} else if request, err := encode.SetListenerAddrPortRequest(c.ID, address, interval); err != nil {
		return zero, err
	} else if reply, err := send(u, c, request, timeout); err != nil {
		return zero, err
	} else {
		return decode.SetListenerAddrPortResponse(reply)
	}
}

//go:generate ../.codegen/bin/codegen API
//go:generate ../.codegen/bin/codegen README
