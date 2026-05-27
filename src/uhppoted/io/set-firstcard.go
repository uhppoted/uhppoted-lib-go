package io

import (
	"time"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/codec"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/codec/encode"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"
)

// Sets the first card configuration for a access controller managed door.
func SetFirstCard[T uhppoted.TController](
	u uhppoted.Uhppoted,
	controller T,
	door uint8,
	startTime, endTime types.HHmm,
	activeMode, inactiveMode types.DoorMode,
	monday, tuesday, wednesday, thursday, friday, saturday, sunday bool,
	timeout time.Duration) (responses.SetFirstCard, error) {

	f := func(id uint32) ([]byte, error) {
		active := uint8(0)
		inactive := uint8(0)

		switch activeMode {
		case types.Controlled:
			active = uint8(0)
		case types.NormallyOpen:
			active = uint8(1)
		case types.NormallyClosed:
			active = uint8(2)
		}

		switch inactiveMode {
		case types.Controlled:
			inactive = uint8(0)
		case types.NormallyOpen:
			inactive = uint8(1)
		case types.NormallyClosed:
			inactive = uint8(2)
		case types.FirstCardOnly:
			inactive = uint8(3)
		}

		return encode.SetFirstCardRequest(
			id,
			door,
			convert[types.HHmm](startTime),
			convert[types.HHmm](endTime),
			active,
			inactive,
			monday,
			tuesday,
			wednesday,
			thursday,
			friday,
			saturday,
			sunday)
	}

	if c, err := resolve(controller); err != nil {
		return responses.SetFirstCard{}, err
	} else if request, err := f(c.ID); err != nil {
		return responses.SetFirstCard{}, err
	} else if reply, err := send(u, c, request, timeout); err != nil {
		return responses.SetFirstCard{}, err
	} else if response, err := codec.Decode[responses.SetFirstCard](reply); err != nil {
		return responses.SetFirstCard{}, err
	} else {
		return response, nil
	}
}
