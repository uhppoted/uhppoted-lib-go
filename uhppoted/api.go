package uhppoted

import (
	"fmt"
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

// Retrieves the card record for a given card number.
func GetCardRecord[T TController](u Uhppoted, controller T, cardnumber uint32, timeout time.Duration) (entities.Card, error) {
	var zero entities.Card

	if c, err := resolve(controller); err != nil {
		return zero, err
	} else if request, err := encode.GetCardRequest(c.ID, cardnumber); err != nil {
		return zero, err
	} else if reply, err := send(u, c, request, timeout); err != nil {
		return zero, err
	} else if response, err := decode.GetCardResponse(reply); err != nil {
		return zero, err
	} else if !valid(response, c.ID) {
		return zero, ErrInvalidResponse
	} else {
		return entities.Card{
			Card:      response.Card,
			StartDate: response.StartDate,
			EndDate:   response.EndDate,
			Permissions: map[uint8]uint8{
				1: response.Door1,
				2: response.Door2,
				3: response.Door3,
				4: response.Door4,
			},
			PIN: response.PIN,
		}, nil
	}
}

// Retrieves the card record stored at a given index.
func GetCardRecordAtIndex[T TController](u Uhppoted, controller T, index uint32, timeout time.Duration) (entities.Card, error) {
	var zero entities.Card

	if c, err := resolve(controller); err != nil {
		return zero, err
	} else if request, err := encode.GetCardAtIndexRequest(c.ID, index); err != nil {
		return zero, err
	} else if reply, err := send(u, c, request, timeout); err != nil {
		return zero, err
	} else if response, err := decode.GetCardAtIndexResponse(reply); err != nil {
		return zero, err
	} else if !valid(response, c.ID) {
		return zero, ErrInvalidResponse
	} else {
		return entities.Card{
			Card:      response.Card,
			StartDate: response.StartDate,
			EndDate:   response.EndDate,
			Permissions: map[uint8]uint8{
				1: response.Door1,
				2: response.Door2,
				3: response.Door3,
				4: response.Door4,
			},
			PIN: response.PIN,
		}, nil
	}
}

// Creates or updates a card record stored on an access controller.
func PutCardRecord[T TController](u Uhppoted, controller T, card entities.Card, timeout time.Duration) (bool, error) {
	permissions := map[uint8]uint8{}
	if card.Permissions != nil {
		permissions = card.Permissions
	}

	door1 := permissions[1]
	door2 := permissions[2]
	door3 := permissions[3]
	door4 := permissions[4]

	if c, err := resolve(controller); err != nil {
		return false, err
	} else if request, err := encode.PutCardRequest(c.ID, card.Card, card.StartDate, card.EndDate, door1, door2, door3, door4, card.PIN); err != nil {
		return false, err
	} else if reply, err := send(u, c, request, timeout); err != nil {
		return false, err
	} else if response, err := decode.PutCardResponse(reply); err != nil {
		return false, err
	} else if !valid(response, c.ID) {
		return false, ErrInvalidResponse
	} else {
		return response.Ok, nil
	}
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
