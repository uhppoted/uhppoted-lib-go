package codec

import (
	"fmt"

	decoder "github.com/uhppoted/uhppoted-lib-go/src/uhppoted/codec/decode"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/entities"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"
)

type key[R any] struct {
	opcode byte
}

var decoders = map[any]func([]byte) (any, error){
	key[responses.GetListenerAddrPort]{0x92}: func(b []byte) (any, error) {
		return decoder.GetListenerAddrPortResponse(b)
	},

	key[responses.SetListenerAddrPort]{0x90}: func(b []byte) (any, error) {
		return decoder.SetListenerAddrPortResponse(b)
	},

	key[entities.Card]{0x5a}: func(b []byte) (any, error) {
		return decodeCardRecord(b)
	},

	key[entities.Card]{0x5c}: func(b []byte) (any, error) {
		return decodeCardRecordAtIndex(b)
	},

	key[entities.Status]{0x20}: func(b []byte) (any, error) {
		return decodeStatusRecord(b)
	},

	key[entities.Event]{0xb0}: func(b []byte) (any, error) {
		return decodeEventRecord(b)
	},

	key[entities.TimeProfile]{0x98}: func(b []byte) (any, error) {
		return decodeTimeProfileRecord(b)
	},
}

func Decode[R any](packet []byte) (R, error) {
	var zero R

	if len(packet) != 64 {
		return zero, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return zero, fmt.Errorf("invalid reply SOM byte (0x%02x)", packet[0])
	}

	k := key[R]{packet[1]}
	fn := decode
	if f, ok := decoders[k]; ok {
		fn = f
	}

	if v, err := fn(packet); err != nil {
		return zero, err
	} else if response, ok := v.(R); !ok {
		return zero, err
	} else {
		return response, nil
	}
}

func decodeCardRecord(packet []byte) (entities.Card, error) {
	if response, err := decoder.GetCardResponse(packet); err != nil {
		return entities.Card{}, err
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

func decodeCardRecordAtIndex(packet []byte) (entities.Card, error) {
	if response, err := decoder.GetCardAtIndexResponse(packet); err != nil {
		return entities.Card{}, err
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

func decodeStatusRecord(packet []byte) (entities.Status, error) {
	if response, err := decoder.GetStatusResponse(packet); err != nil {
		return entities.Status{}, err
	} else {
		datetime := entities.NewDateTime(
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
				Direction:     entities.Direction(response.EventDirection),
				Card:          response.EventCard,
				Timestamp:     response.EventTimestamp,
				Reason:        response.EventReason,
			},
		}, nil
	}
}

func decodeEventRecord(packet []byte) (entities.Event, error) {
	if response, err := decoder.GetEventResponse(packet); err != nil {
		return entities.Event{}, err
	} else {
		return entities.Event{
			Index:         response.Index,
			Event:         entities.EventType(response.EventType),
			AccessGranted: response.AccessGranted,
			Door:          response.Door,
			Direction:     entities.Direction(response.Direction),
			Card:          response.Card,
			Timestamp:     response.Timestamp,
			Reason:        response.Reason,
		}, nil
	}
}

func decodeTimeProfileRecord(packet []byte) (entities.TimeProfile, error) {
	if response, err := decoder.GetTimeProfileResponse(packet); err != nil {
		return entities.TimeProfile{}, err
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
