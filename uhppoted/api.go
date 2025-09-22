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

// Creates or updates a card record stored on an access controller.
func PutCard[T TController, D TDate](u Uhppoted, controller T, card uint32, startdate D, enddate D, door1 uint8, door2 uint8, door3 uint8, door4 uint8, PIN uint32, timeout time.Duration) (responses.PutCardResponse, error) {
	f := func(id uint32) ([]byte, error) {
		if start, err := convert(startdate); err != nil {
			return nil, err
		} else if end, err := convert(enddate); err != nil {
			return nil, err
		} else {
			return encode.PutCardRequest(id, card, start, end, door1, door2, door3, door4, PIN)
		}
	}

	return exec[T, responses.PutCardResponse](u, controller, f, timeout)
}

// Adds or updates an access time profile stored on a controller.
func SetTimeProfile[T TController, D TDate](u Uhppoted, controller T, profile uint8, startdate D, enddate D, monday bool, tuesday bool, wednesday bool, thursday bool, friday bool, saturday bool, sunday bool, segment1start time.Time, segment1end time.Time, segment2start time.Time, segment2end time.Time, segment3start time.Time, segment3end time.Time, linkedprofileid uint8, timeout time.Duration) (responses.SetTimeProfileResponse, error) {
	f := func(id uint32) ([]byte, error) {
		if start, err := convert(startdate); err != nil {
			return nil, err
		} else if end, err := convert(enddate); err != nil {
			return nil, err
		} else {
			return encode.SetTimeProfileRequest(id, profile, start, end, monday, tuesday, wednesday, thursday, friday, saturday, sunday, segment1start, segment1end, segment2start, segment2end, segment3start, segment3end, linkedprofileid)
		}
	}

	return exec[T, responses.SetTimeProfileResponse](u, controller, f, timeout)
}

// Creates a scheduled task.
//
// Task types
// 0:  control door
// 1:  unlock door
// 2:  lock door
// 3:  disable time profiles
// 4:  enable time profiles
// 5:  enable card, no password
// 6:  enable card+IN password
// 7:  enable card+password
// 8:  enable more cards
// 9:  disable more cards
// 10: trigger once
// 11: disable pushbutton
// 12: enable pushbutton
func AddTask[T TController, D TDate](u Uhppoted, controller T, task uint8, startdate D, enddate D, monday bool, tuesday bool, wednesday bool, thursday bool, friday bool, saturday bool, sunday bool, starttime time.Time, door uint8, morecards uint8, timeout time.Duration) (responses.AddTaskResponse, error) {
	f := func(id uint32) ([]byte, error) {
		if start, err := convert(startdate); err != nil {
			return nil, err
		} else if end, err := convert(enddate); err != nil {
			return nil, err
		} else {
			return encode.AddTaskRequest(id, task, start, end, monday, tuesday, wednesday, thursday, friday, saturday, sunday, starttime, door, morecards)
		}
	}

	return exec[T, responses.AddTaskResponse](u, controller, f, timeout)
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

func convert[D any](date D) (entities.Date, error) {
	switch d := any(date).(type) {
	case entities.Date:
		return d, nil
	case time.Time:
		return entities.DateFromTime(d), nil
	}

	return entities.Date{}, fmt.Errorf("unsupported date type %T", date)
}

//go:generate ../.codegen/bin/codegen API
//go:generate ../.codegen/bin/codegen responses
//go:generate ../.codegen/bin/codegen README
