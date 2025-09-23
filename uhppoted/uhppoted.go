package uhppoted

import (
	"errors"
	"fmt"
	"net"
	"net/netip"
	"reflect"
	"time"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/entities"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/responses"
)

type Uhppoted struct {
	bindAddr      netip.AddrPort
	broadcastAddr netip.AddrPort
	listenAddr    netip.AddrPort
	debug         bool

	udp udp
	tcp tcp
}

type TController interface {
	~uint32 | Controller
}

type TDateTime interface {
	time.Time | entities.DateTime
}

type TDate interface {
	time.Time | entities.Date
}

type TTime interface {
	time.Time | entities.Time
}

type THHmm interface {
	time.Time | entities.HHmm
}

type Controller struct {
	ID       uint32
	Address  netip.AddrPort
	Protocol string
}

type GetControllerResponse = responses.GetControllerResponse
type SetIPv4Response = responses.SetIPv4Response
type GetTimeResponse = responses.GetTimeResponse
type SetTimeResponse = responses.SetTimeResponse
type GetListenerResponse = responses.GetListenerResponse
type SetListenerResponse = responses.SetListenerResponse
type GetListenerAddrPortResponse = responses.GetListenerAddrPortResponse
type SetListenerAddrPortResponse = responses.SetListenerAddrPortResponse
type GetDoorResponse = responses.GetDoorResponse
type SetDoorResponse = responses.SetDoorResponse
type SetDoorPasscodesResponse = responses.SetDoorPasscodesResponse
type OpenDoorResponse = responses.OpenDoorResponse
type GetStatusResponse = responses.GetStatusResponse
type GetCardsResponse = responses.GetCardsResponse
type GetCardResponse = responses.GetCardResponse
type GetCardAtIndexResponse = responses.GetCardAtIndexResponse
type PutCardResponse = responses.PutCardResponse
type DeleteCardResponse = responses.DeleteCardResponse
type DeleteAllCardsResponse = responses.DeleteAllCardsResponse
type GetEventResponse = responses.GetEventResponse
type GetEventIndexResponse = responses.GetEventIndexResponse
type SetEventIndexResponse = responses.SetEventIndexResponse
type RecordSpecialEventsResponse = responses.RecordSpecialEventsResponse
type GetTimeProfileResponse = responses.GetTimeProfileResponse
type SetTimeProfileResponse = responses.SetTimeProfileResponse
type ClearTimeProfilesResponse = responses.ClearTimeProfilesResponse
type AddTaskResponse = responses.AddTaskResponse
type RefreshTaskListResponse = responses.RefreshTaskListResponse
type ClearTaskListResponse = responses.ClearTaskListResponse
type SetPCControlResponse = responses.SetPCControlResponse
type SetInterlockResponse = responses.SetInterlockResponse
type ActivateKeypadsResponse = responses.ActivateKeypadsResponse
type GetAntiPassbackResponse = responses.GetAntiPassbackResponse
type SetAntiPassbackResponse = responses.SetAntiPassbackResponse
type RestoreDefaultParametersResponse = responses.RestoreDefaultParametersResponse
type ListenerEvent = responses.ListenerEvent

// Error constants
var ErrInvalidResponse = errors.New("invalid response")

// NewUhppoted creates a new instance of the uhppoted service, configured with the supplied
// local bind address, broadcast address, and listen address. The debug flag enables or
// disables logging of the network packets to the console.
//
// The bind, broadcast, and listen parameters are expected to be valid netip.AddPort
// addresses.
func NewUhppoted(bind, broadcast, listen netip.AddrPort, debug bool) Uhppoted {
	return Uhppoted{
		bindAddr:      bind,
		broadcastAddr: broadcast,
		listenAddr:    listen,
		debug:         debug,

		udp: udp{
			bindAddr:      net.UDPAddrFromAddrPort(bind),
			broadcastAddr: net.UDPAddrFromAddrPort(broadcast),
			listenAddr:    net.UDPAddrFromAddrPort(listen),
			debug:         debug,
		},

		tcp: tcp{
			bindAddr: net.TCPAddrFromAddrPort(bind),
			debug:    debug,
		},
	}
}

func exec[T TController, R any](u Uhppoted, controller T, encode func(id uint32) ([]byte, error), timeout time.Duration) (R, error) {
	var zero R

	if c, err := resolve(controller); err != nil {
		return zero, err
	} else if request, err := encode(c.ID); err != nil {
		return zero, err
	} else if reply, err := send(u, c, request, timeout); err != nil {
		return zero, err
	} else if response, err := codec.Decode[R](reply); err != nil {
		return response, err
	} else if !valid(response, c.ID) {
		return response, ErrInvalidResponse
	} else {
		return response, nil
	}
}

func send(u Uhppoted, controller Controller, request []byte, timeout time.Duration) ([]byte, error) {
	zero := netip.AddrPort{}

	if controller.Address != zero && !controller.Address.IsValid() {
		return nil, fmt.Errorf("invalid address (%v)", controller.Address)
	} else if controller.Address != zero && controller.Protocol == "tcp" {
		return u.tcp.sendTo(request, controller.Address, timeout)
	} else if controller.Address != zero {
		return u.udp.sendTo(request, controller.Address, timeout)
	} else {
		return u.udp.broadcastTo(request, timeout)
	}
}

func resolve[T TController](controller T) (Controller, error) {
	switch v := any(controller).(type) {
	case uint32:
		var err error
		if v == 0 {
			err = fmt.Errorf("invalid controller ID (%v)", v)
		}

		return Controller{
			ID: v,
		}, err

	case Controller:
		var err error
		if v.ID == 0 {
			err = fmt.Errorf("invalid controller ID (%v)", v)
		}

		return v, err
	}

	return Controller{}, fmt.Errorf("unsupported type (%T)", controller)
}

func valid[R any](response R, controller uint32) bool {
	r := reflect.ValueOf(response)
	if r.Kind() == reflect.Ptr && !r.IsNil() {
		r = r.Elem()
	}

	if r.Kind() == reflect.Struct {
		f := r.FieldByName("Controller")
		if f.IsValid() && f.Kind() == reflect.Uint32 {
			return uint32(f.Uint()) == controller
		}
	}

	return true
}
