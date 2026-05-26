package io

import (
	"fmt"
	"time"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"
)

// type Uhppoted uhppoted.Uhppoted
// type TController uhppoted.TController
// type Controller uhppoted.Controller

type convertable interface {
	types.DateTime | types.Date | types.HHmm
}

func send(u uhppoted.Uhppoted, controller uhppoted.Controller, request []byte, timeout time.Duration) ([]byte, error) {
	return u.Send(controller, request, timeout)
}

func resolve[T uhppoted.TController](controller T) (uhppoted.Controller, error) {
	switch v := any(controller).(type) {
	case uint32:
		var err error
		if v == 0 {
			err = fmt.Errorf("invalid controller ID (%v)", v)
		}

		return uhppoted.Controller{
			ID: v,
		}, err

	case uhppoted.Controller:
		var err error
		if v.ID == 0 {
			err = fmt.Errorf("invalid controller ID (%v)", v)
		}

		return v, err
	}

	return uhppoted.Controller{}, fmt.Errorf("unsupported type (%T)", controller)
}

func convert[R convertable](t any) R {
	var zero R

	switch v := any(t).(type) {
	case R:
		return v

	case time.Time:
		switch any(zero).(type) {
		case types.DateTime:
			return any(types.DateTimeFromTime(v)).(R)

		case types.Date:
			return any(types.DateFromTime(v)).(R)

		case types.HHmm:
			return any(types.HHmmFromTime(v)).(R)
		}
	}

	panic(fmt.Sprintf("unsupported conversion from %T to %T", t, zero))
}
