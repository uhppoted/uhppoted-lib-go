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

// re-exported response types
type (
	GetControllerResponse            = responses.GetControllerResponse
	SetIPv4Response                  = responses.SetIPv4Response
	GetTimeResponse                  = responses.GetTimeResponse
	SetTimeResponse                  = responses.SetTimeResponse
	GetListenerResponse              = responses.GetListenerResponse
	SetListenerResponse              = responses.SetListenerResponse
	GetListenerAddrPortResponse      = responses.GetListenerAddrPortResponse
	SetListenerAddrPortResponse      = responses.SetListenerAddrPortResponse
	GetDoorResponse                  = responses.GetDoorResponse
	SetDoorResponse                  = responses.SetDoorResponse
	SetDoorPasscodesResponse         = responses.SetDoorPasscodesResponse
	OpenDoorResponse                 = responses.OpenDoorResponse
	GetStatusResponse                = responses.GetStatusResponse
	GetCardsResponse                 = responses.GetCardsResponse
	GetCardResponse                  = responses.GetCardResponse
	GetCardAtIndexResponse           = responses.GetCardAtIndexResponse
	PutCardResponse                  = responses.PutCardResponse
	DeleteCardResponse               = responses.DeleteCardResponse
	DeleteAllCardsResponse           = responses.DeleteAllCardsResponse
	GetEventResponse                 = responses.GetEventResponse
	GetEventIndexResponse            = responses.GetEventIndexResponse
	SetEventIndexResponse            = responses.SetEventIndexResponse
	RecordSpecialEventsResponse      = responses.RecordSpecialEventsResponse
	GetTimeProfileResponse           = responses.GetTimeProfileResponse
	SetTimeProfileResponse           = responses.SetTimeProfileResponse
	ClearTimeProfilesResponse        = responses.ClearTimeProfilesResponse
	AddTaskResponse                  = responses.AddTaskResponse
	RefreshTaskListResponse          = responses.RefreshTaskListResponse
	ClearTaskListResponse            = responses.ClearTaskListResponse
	SetPCControlResponse             = responses.SetPCControlResponse
	SetInterlockResponse             = responses.SetInterlockResponse
	ActivateKeypadsResponse          = responses.ActivateKeypadsResponse
	GetAntiPassbackResponse          = responses.GetAntiPassbackResponse
	SetAntiPassbackResponse          = responses.SetAntiPassbackResponse
	RestoreDefaultParametersResponse = responses.RestoreDefaultParametersResponse
	ListenerEvent                    = responses.ListenerEvent
)

// re-exported entities
type (
	Card        = entities.Card
	TimeProfile = entities.TimeProfile
	Event       = entities.Event
	DoorMode    = entities.DoorMode
	Task        = entities.Task
	Interlock   = entities.Interlock
	Weekdays    = entities.Weekdays
	TimeSegment = entities.TimeSegment
)

// re-exported door modes
var (
	NormallyOpen   = entities.NormallyOpen
	NormallyClosed = entities.NormallyClosed
	Controlled     = entities.Controlled
)

// re-exported tasks
var (
	ControlDoor          = entities.ControlDoor
	UnlockDoor           = entities.UnlockDoor
	LockDoor             = entities.LockDoor
	DisableTimeProfiles  = entities.DisableTimeProfiles
	EnableTimeProfiles   = entities.EnableTimeProfiles
	EnableCardNoPassword = entities.EnableCardNoPassword
	EnableCardInPassword = entities.EnableCardInPassword
	EnableCardPassword   = entities.EnableCardPassword
	EnableMoreCards      = entities.EnableMoreCards
	DisableMoreCards     = entities.DisableMoreCards
	TriggerOnce          = entities.TriggerOnce
	DisablePushbutton    = entities.DisablePushbutton
	EnablePushbutton     = entities.EnablePushbutton
)

// re-exported interlocks
var (
	NoInterlock    = entities.NoInterlock
	Interlock12    = entities.Interlock12
	Interlock34    = entities.Interlock34
	Interlock12_34 = entities.Interlock12_34
	Interlock123   = entities.Interlock123
	Interlock1234  = entities.Interlock1234
)

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
		return zero, ErrInvalidResponse
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
