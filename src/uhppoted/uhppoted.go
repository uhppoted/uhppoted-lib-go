package uhppoted

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"net/netip"
	"reflect"
	"time"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/codec"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"
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
	time.Time | types.DateTime
}

type TDate interface {
	time.Time | types.Date
}

type TTime interface {
	time.Time | types.Time
}

type THHmm interface {
	time.Time | types.HHmm
}

type Controller struct {
	ID       uint32
	Address  netip.AddrPort
	Protocol string
}

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

	if r.Kind() == reflect.Slice && r.Type().Elem().Kind() == reflect.Uint8 && r.Len() == 64 {
		packet := r.Bytes()
		id := binary.LittleEndian.Uint32(packet[4:8])

		return id == controller
	}

	return true
}
