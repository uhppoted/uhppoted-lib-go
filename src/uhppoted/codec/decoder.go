package codec

import (
	"fmt"

	decoder "github.com/uhppoted/uhppoted-lib-go/src/uhppoted/codec/decode"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"
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

	key[types.Card]{0x5a}: func(b []byte) (any, error) {
		return decodeCardRecord(b)
	},

	key[types.Card]{0x5c}: func(b []byte) (any, error) {
		return decodeCardRecordAtIndex(b)
	},

	key[types.Status]{0x20}: func(b []byte) (any, error) {
		return decodeStatusRecord(b)
	},

	key[types.Event]{0xb0}: func(b []byte) (any, error) {
		return decodeEventRecord(b)
	},

	key[types.TimeProfile]{0x98}: func(b []byte) (any, error) {
		return decodeTimeProfileRecord(b)
	},

	key[responses.ListenerEvent]{0x20}: func(b []byte) (any, error) {
		return decoder.ListenerEvent(b)
	},
}

// Decode is a convenience 'factory' decoder to decode a received packet into the appropriate
// response struct or entity.
func Decode[R any](packet []byte) (R, error) {
	var zero R

	if len(packet) != 64 {
		return zero, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
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

func decodeCardRecord(packet []byte) (types.Card, error) {
	if response, err := decoder.GetCardResponse(packet); err != nil {
		return types.Card{}, err
	} else {
		return types.Card{
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

func decodeCardRecordAtIndex(packet []byte) (types.Card, error) {
	if response, err := decoder.GetCardAtIndexResponse(packet); err != nil {
		return types.Card{}, err
	} else {
		return types.Card{
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

func decodeStatusRecord(packet []byte) (types.Status, error) {
	if response, err := decoder.GetStatusResponse(packet); err != nil {
		return types.Status{}, err
	} else {
		datetime := types.NewDateTime(
			response.SystemDate.Year(),
			response.SystemDate.Month(),
			response.SystemDate.Day(),
			response.SystemTime.Hour(),
			response.SystemTime.Minute(),
			response.SystemTime.Second())

		return types.Status{
			System: struct {
				Time  types.DateTime `json:"datetime"`
				Error uint8          `json:"error"`
				Info  uint8          `json:"info"`
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

			Event: types.Event{
				Index:         response.EventIndex,
				Event:         types.EventType(response.EventType),
				AccessGranted: response.EventAccessGranted,
				Door:          response.EventDoor,
				Direction:     types.Direction(response.EventDirection),
				Card:          response.EventCard,
				Timestamp:     response.EventTimestamp,
				Reason:        response.EventReason,
			},
		}, nil
	}
}

func decodeEventRecord(packet []byte) (types.Event, error) {
	if response, err := decoder.GetEventResponse(packet); err != nil {
		return types.Event{}, err
	} else {
		return types.Event{
			Index:         response.Index,
			Event:         types.EventType(response.EventType),
			AccessGranted: response.AccessGranted,
			Door:          response.Door,
			Direction:     types.Direction(response.Direction),
			Card:          response.Card,
			Timestamp:     response.Timestamp,
			Reason:        response.Reason,
		}, nil
	}
}

func decodeTimeProfileRecord(packet []byte) (types.TimeProfile, error) {
	if response, err := decoder.GetTimeProfileResponse(packet); err != nil {
		return types.TimeProfile{}, err
	} else {
		return types.TimeProfile{
			Profile:   response.Profile,
			StartDate: response.StartDate,
			EndDate:   response.EndDate,
			Weekdays: types.Weekdays{
				Monday:    response.Monday,
				Tuesday:   response.Tuesday,
				Wednesday: response.Wednesday,
				Thursday:  response.Thursday,
				Friday:    response.Friday,
				Saturday:  response.Saturday,
				Sunday:    response.Sunday,
			},
			Segments: []types.TimeSegment{
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
