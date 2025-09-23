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

type convertable interface {
	entities.DateTime | entities.Date | entities.HHmm
}

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

// // Sets the access controller system date and time.
// func SetTime[T TController, DT TDateTime](u Uhppoted, controller T, datetime DT, timeout time.Duration) (responses.SetTimeResponse, error) {
// 	f := func(id uint32) ([]byte, error) {
// 		return encode.SetTimeRequest(id, convert[entities.DateTime](datetime))
// 	}
//
// 	return exec[T, responses.SetTimeResponse](u, controller, f, timeout)
// }

// // Creates or updates a card record stored on an access controller.
// func PutCard[T TController, D TDate](u Uhppoted, controller T, card uint32, startdate D, enddate D, door1 uint8, door2 uint8, door3 uint8, door4 uint8, PIN uint32, timeout time.Duration) (responses.PutCardResponse, error) {
// 	f := func(id uint32) ([]byte, error) {
// 		return encode.PutCardRequest(
// 			id, card,
// 			convert[entities.Date](startdate),
// 			convert[entities.Date](enddate),
// 			door1, door2, door3, door4,
// 			PIN)
// 	}
//
// 	return exec[T, responses.PutCardResponse](u, controller, f, timeout)
// }

// // Adds or updates an access time profile stored on a controller.
// func SetTimeProfile[T TController, D TDate, H THHmm](u Uhppoted, controller T, profile uint8,
// 	startdate D, enddate D,
// 	monday, tuesday, wednesday, thursday, friday, saturday, sunday bool,
// 	segment1start, segment1end H,
// 	segment2start, segment2end H,
// 	segment3start, segment3end H,
// 	linkedprofileid uint8,
// 	timeout time.Duration) (responses.SetTimeProfileResponse, error) {
// 	f := func(id uint32) ([]byte, error) {
// 		return encode.SetTimeProfileRequest(id, profile,
// 			convert[entities.Date](startdate), convert[entities.Date](enddate),
// 			monday, tuesday, wednesday, thursday, friday, saturday, sunday,
// 			convert[entities.HHmm](segment1start), convert[entities.HHmm](segment1end),
// 			convert[entities.HHmm](segment2start), convert[entities.HHmm](segment2end),
// 			convert[entities.HHmm](segment3start), convert[entities.HHmm](segment3end),
// 			linkedprofileid)
// 	}
//
// 	return exec[T, responses.SetTimeProfileResponse](u, controller, f, timeout)
// }

// // Creates a scheduled task.
// //
// // Task types
// // 0:  control door
// // 1:  unlock door
// // 2:  lock door
// // 3:  disable time profiles
// // 4:  enable time profiles
// // 5:  enable card, no password
// // 6:  enable card+IN password
// // 7:  enable card+password
// // 8:  enable more cards
// // 9:  disable more cards
// // 10: trigger once
// // 11: disable pushbutton
// // 12: enable pushbutton
// func AddTask[T TController, D TDate, H THHmm](u Uhppoted, controller T, task uint8,
// 	startdate D, enddate D,
// 	monday, tuesday, wednesday, thursday, friday, saturday, sunday bool,
// 	startTime H, door uint8, morecards uint8,
// 	timeout time.Duration) (responses.AddTaskResponse, error) {
// 	f := func(id uint32) ([]byte, error) {
// 		return encode.AddTaskRequest(
// 			id, task,
// 			convert[entities.Date](startdate),
// 			convert[entities.Date](enddate),
// 			monday, tuesday, wednesday, thursday, friday, saturday, sunday,
// 			convert[entities.HHmm](startTime),
// 			door, morecards)
// 	}
//
// 	return exec[T, responses.AddTaskResponse](u, controller, f, timeout)
// }

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

func convert[R convertable](t any) R {
	var zero R

	switch v := any(t).(type) {
	case R:
		return v

	case time.Time:
		switch any(zero).(type) {
		case entities.DateTime:
			return any(entities.DateTimeFromTime(v)).(R)

		case entities.Date:
			return any(entities.DateFromTime(v)).(R)

		case entities.HHmm:
			return any(entities.HHmmFromTime(v)).(R)
		}
	}

	panic(fmt.Sprintf("unsupported conversion from %T to %T", t, zero))
}

//go:generate ../.codegen/bin/codegen API
//go:generate ../.codegen/bin/codegen responses
//go:generate ../.codegen/bin/codegen README
