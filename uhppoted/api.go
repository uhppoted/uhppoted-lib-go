package uhppoted

import (
	"fmt"
	"net/netip"
	"os"
	"time"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/decode"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/encode"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/entities"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/responses"
)

// FindControllers retrieves a list of all UHPPOTE controllers accessible on the local LAN.
//
// It broadcasts a UDP `get controller` request to the local network and returns a list of
// decoded responses from controllers that reply within the timeout. Responses that cannot
// be decoded are silently ignored.
func FindControllers(u Uhppoted, timeout time.Duration) ([]responses.GetControllerResponse, error) {
	if request, err := encode.GetControllerRequest(0); err != nil {
		return nil, err
	} else if replies, err := u.udp.broadcast(request, timeout); err != nil {
		return nil, err
	} else {
		responses := []responses.GetControllerResponse{}

		for _, reply := range replies {
			if response, err := decode.GetControllerResponse(reply); err == nil {
				responses = append(responses, response)
			}
		}

		return responses, nil
	}
}

// Retrieves the access controller event listener IPv4 address:port and auto-send interval.
func GetListenerAddrPort[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.GetListenerAddrPortResponse, error) {
	var zero responses.GetListenerAddrPortResponse

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
func SetListenerAddrPort[T TController](u Uhppoted, controller T, address netip.AddrPort, interval uint8, timeout time.Duration) (responses.SetListenerAddrPortResponse, error) {
	var zero responses.SetListenerAddrPortResponse

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

// Sets the access controller system date and time.
func SetTime[T TController, DT TDateTime](u Uhppoted, controller T, datetime DT, timeout time.Duration) (responses.SetTimeResponse, error) {
	f := func(id uint32) ([]byte, error) {
		switch dt := any(datetime).(type) {
		case entities.DateTime:
			return encode.SetTimeRequest(id, dt)
		case time.Time:
			return encode.SetTimeRequest(id, entities.DateTimeFromTime(dt))
		default:
			return nil, fmt.Errorf("unsupported datetime type %T", datetime)
		}
	}

	return exec[T, responses.SetTimeResponse](u, controller, f, timeout)
}

// Listens for access controller events sent to the listen address:port and routes received events
// to the events channel. Terminates on any signal sent to the interrupt channel.
func Listen(u Uhppoted, events chan ListenerEvent, errors chan error, interrupt chan os.Signal) error {
	ch := make(chan []uint8)

	go u.udp.listen(ch)

loop:
	for {
		select {
		case msg := <-ch:
			if evt, err := decode.ListenerEvent(msg); err != nil {
				errors <- err
			} else {
				events <- evt
			}

		case <-interrupt:
			break loop
		}
	}

	return nil
}

// // Sets the access controller system date and time.
// func SetTimeX[T TController, D TDateTime](u Uhppoted, controller T, datetime D, timeout time.Duration) (responses.SetTimeResponse, error) {
// 	fmt.Printf(">>>>>>>>>>>>>>>>>>>>> %T %v\n", datetime, datetime)
// 	return responses.SetTimeResponse{}, nil
// 	// f := func(id uint32) ([]byte, error) {
// 	// 	return encode.SetTimeRequest(id, datetime)
// 	// }
// 	//
// 	// return exec[T, responses.SetTimeResponse](u, controller, f, timeout)
// }

//go:generate ../.codegen/bin/codegen API
//go:generate ../.codegen/bin/codegen responses
//go:generate ../.codegen/bin/codegen README
