package io

import (
	"time"

	"codeberg.org/uhppoted/uhppoted-lib-go/src/uhppoted"
	"codeberg.org/uhppoted/uhppoted-lib-go/src/uhppoted/codec"
	"codeberg.org/uhppoted/uhppoted-lib-go/src/uhppoted/codec/encode"
	"codeberg.org/uhppoted/uhppoted-lib-go/src/uhppoted/responses"
	"codeberg.org/uhppoted/uhppoted-lib-go/src/uhppoted/types"
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
		return encode.SetFirstCardRequest(
			id,
			door,
			convert[types.HHmm](startTime),
			convert[types.HHmm](endTime),
			activeMode,
			inactiveMode,
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
