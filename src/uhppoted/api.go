package uhppoted

import (
	"fmt"
	"os"
	"time"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/codec/decode"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/codec/encode"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/entities"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"
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

// Retrieves a controller status record.
func GetStatusRecord[T TController](u Uhppoted, controller T, timeout time.Duration) (entities.Status, error) {
	var zero entities.Status

	if c, err := resolve(controller); err != nil {
		return zero, err
	} else if request, err := encode.GetStatusRequest(c.ID); err != nil {
		return zero, err
	} else if reply, err := send(u, c, request, timeout); err != nil {
		return zero, err
	} else if response, err := decode.GetStatusResponse(reply); err != nil {
		return zero, err
	} else if !valid(response, c.ID) {
		return zero, ErrInvalidResponse
	} else {
		datetime := NewDateTime(
			response.SystemDate.Year(),
			response.SystemDate.Month(),
			response.SystemDate.Day(),
			response.SystemTime.Hour(),
			response.SystemTime.Minute(),
			response.SystemTime.Second())

		return entities.Status{
			System: struct {
				Time  entities.DateTime `json:"datetime"`
				Error uint8             `json:"error"`
				Info  uint8             `json:"info"`
			}{
				Time:  datetime,
				Error: response.SystemError,
				Info:  response.SpecialInfo,
			},

			Doors: map[uint8]struct {
				Open     bool `json:"open"`
				Button   bool `json:"button"`
				Unlocked bool `json:"unlocked"`
			}{
				1: struct {
					Open     bool `json:"open"`
					Button   bool `json:"button"`
					Unlocked bool `json:"unlocked"`
				}{
					Open:     response.Door1Open,
					Button:   response.Door1Button,
					Unlocked: response.Relays&0x01 == 0x01,
				},
				2: struct {
					Open     bool `json:"open"`
					Button   bool `json:"button"`
					Unlocked bool `json:"unlocked"`
				}{
					Open:     response.Door2Open,
					Button:   response.Door2Button,
					Unlocked: response.Relays&0x02 == 0x02,
				},
				3: struct {
					Open     bool `json:"open"`
					Button   bool `json:"button"`
					Unlocked bool `json:"unlocked"`
				}{
					Open:     response.Door3Open,
					Button:   response.Door3Button,
					Unlocked: response.Relays&0x04 == 0x04,
				},
				4: struct {
					Open     bool `json:"open"`
					Button   bool `json:"button"`
					Unlocked bool `json:"unlocked"`
				}{
					Open:     response.Door4Open,
					Button:   response.Door4Button,
					Unlocked: response.Relays&0x08 == 0x08,
				},
			},

			Alarms: struct {
				Flags      uint8 `json:"flags"`
				Fire       bool  `json:"fire"`
				LockForced bool  `json:"lock-forced"`
			}{
				Flags:      response.Inputs,
				Fire:       response.Inputs&0x01 == 0x01,
				LockForced: response.Inputs&0x02 == 0x02,
			},

			Event: entities.Event{
				Index:         response.EventIndex,
				Event:         entities.EventType(response.EventType),
				AccessGranted: response.EventAccessGranted,
				Door:          response.EventDoor,
				Direction:     response.EventDirection,
				Card:          response.EventCard,
				Timestamp:     response.EventTimestamp,
				Reason:        response.EventReason,
			},
		}, nil
	}
}

// Retrieves the event record for the even at an index.
func GetEventRecord[T TController](u Uhppoted, controller T, index uint32, timeout time.Duration) (entities.Event, error) {
	var zero entities.Event

	if c, err := resolve(controller); err != nil {
		return zero, err
	} else if request, err := encode.GetEventRequest(c.ID, index); err != nil {
		return zero, err
	} else if reply, err := send(u, c, request, timeout); err != nil {
		return zero, err
	} else if response, err := decode.GetEventResponse(reply); err != nil {
		return zero, err
	} else if !valid(response, c.ID) {
		return zero, ErrInvalidResponse
	} else {
		return entities.Event{
			Index:         response.Index,
			Event:         entities.EventType(response.EventType),
			AccessGranted: response.AccessGranted,
			Door:          response.Door,
			Direction:     response.Direction,
			Card:          response.Card,
			Timestamp:     response.Timestamp,
			Reason:        response.Reason,
		}, nil
	}
}

// Retrieves the requested access time profile record from a controller.
func GetTimeProfileRecord[T TController](u Uhppoted, controller T, profile uint8, timeout time.Duration) (entities.TimeProfile, error) {
	var zero entities.TimeProfile

	if c, err := resolve(controller); err != nil {
		return zero, err
	} else if request, err := encode.GetTimeProfileRequest(c.ID, profile); err != nil {
		return zero, err
	} else if reply, err := send(u, c, request, timeout); err != nil {
		return zero, err
	} else if response, err := decode.GetTimeProfileResponse(reply); err != nil {
		return zero, err
	} else if !valid(response, c.ID) {
		return zero, ErrInvalidResponse
	} else {
		return entities.TimeProfile{
			Profile:   response.Profile,
			StartDate: response.StartDate,
			EndDate:   response.EndDate,
			Weekdays: entities.Weekdays{
				Monday:    response.Monday,
				Tuesday:   response.Tuesday,
				Wednesday: response.Wednesday,
				Thursday:  response.Thursday,
				Friday:    response.Friday,
				Saturday:  response.Saturday,
				Sunday:    response.Sunday,
			},
			Segments: []entities.TimeSegment{
				{
					Start: response.Segment1Start,
					End:   response.Segment1End,
				},
				{
					Start: response.Segment2Start,
					End:   response.Segment2End,
				},
				{
					Start: response.Segment3Start,
					End:   response.Segment3End,
				},
			},
			LinkedProfile: response.LinkedProfile,
		}, nil
	}
}

// Adds or updates an access time profile record stored on a controller.
func SetTimeProfileRecord[T TController](u Uhppoted, controller T, record entities.TimeProfile, timeout time.Duration) (bool, error) {
	f := func(id uint32) ([]byte, error) {
		segment1start := entities.HHmm{}
		segment1end := entities.HHmm{}
		segment2start := entities.HHmm{}
		segment2end := entities.HHmm{}
		segment3start := entities.HHmm{}
		segment3end := entities.HHmm{}

		if len(record.Segments) > 0 {
			segment1start = record.Segments[0].Start
			segment1end = record.Segments[0].End
		}

		if len(record.Segments) > 1 {
			segment2start = record.Segments[1].Start
			segment2end = record.Segments[1].End
		}

		if len(record.Segments) > 2 {
			segment3start = record.Segments[2].Start
			segment3end = record.Segments[2].End
		}

		return encode.SetTimeProfileRequest(id,
			record.Profile,
			record.StartDate,
			record.EndDate,
			record.Weekdays.Monday,
			record.Weekdays.Tuesday,
			record.Weekdays.Wednesday,
			record.Weekdays.Thursday,
			record.Weekdays.Friday,
			record.Weekdays.Saturday,
			record.Weekdays.Sunday,
			segment1start, segment1end,
			segment2start, segment2end,
			segment3start, segment3end,
			record.LinkedProfile)
	}

	if c, err := resolve(controller); err != nil {
		return false, err
	} else if request, err := f(c.ID); err != nil {
		return false, err
	} else if reply, err := send(u, c, request, timeout); err != nil {
		return false, err
	} else if response, err := decode.SetTimeProfileResponse(reply); err != nil {
		return false, err
	} else if !valid(response, c.ID) {
		return false, ErrInvalidResponse
	} else {
		return response.Ok, nil
	}
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
func AddTaskRecord[T TController](u Uhppoted, controller T, record Task, timeout time.Duration) (bool, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.AddTaskRequest(id,
			record.Task,
			record.StartDate,
			record.EndDate,
			record.Weekdays.Monday,
			record.Weekdays.Tuesday,
			record.Weekdays.Wednesday,
			record.Weekdays.Thursday,
			record.Weekdays.Friday,
			record.Weekdays.Saturday,
			record.Weekdays.Sunday,
			record.StartTime,
			record.Door,
			record.MoreCards)
	}

	if c, err := resolve(controller); err != nil {
		return false, err
	} else if request, err := f(c.ID); err != nil {
		return false, err
	} else if reply, err := send(u, c, request, timeout); err != nil {
		return false, err
	} else if response, err := decode.AddTaskResponse(reply); err != nil {
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

//go:generate ../../.codegen/bin/codegen API
//go:generate ../../.codegen/bin/codegen responses
//go:generate ../../.codegen/bin/codegen README
