// Package uhppoted provides a Go wrapper for the UHPPOTE TCP/IP access controller API.
//
// It enables interaction with UHPPOTE controllers on a local LAN using either UDP
// broadcasts, UDP connected sockets or TCP/IP. The package supports querying controller
// information, managing access cards, and controlling doors.
//
// Typical usage:
//
//	client := uhppoted.Uhppoted{
//	    BindAddr:      netip.MustParseAddrPort("0.0.0.0:0"),
//	    BroadcastAddr: netip.MustParseAddrPort("255.255.255.255:60000"),
//	    ListenAddr:    netip.MustParseAddrPort("0.0.0.0:60001"),
//	    Debug:         false,
//	}
//
//	if controllers, err := client.GetAllControllers(2 * time.Second); err != nil {
//	    log.Fatal(err)
//	} else {
//	    for _, c := range controllers {
//	        fmt.Printf("Controller: %v\n", c.SerialNumber)
//	    }
//	}
package uhppoted

import (
	"fmt"
	"net"
	"net/netip"
	"time"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/log"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/types"
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

type Controller struct {
	ID       uint32
	Address  netip.AddrPort
	Protocol string
}

type GetControllerResponse = types.GetControllerResponse
type SetIPv4Response = types.SetIPv4Response
type GetTimeResponse = types.GetTimeResponse
type SetTimeResponse = types.SetTimeResponse
type GetListenerAddrPortResponse = types.GetListenerAddrPortResponse
type SetListenerAddrPortResponse = types.SetListenerAddrPortResponse
type GetDoorResponse = types.GetDoorResponse
type SetDoorResponse = types.SetDoorResponse
type SetDoorPasscodesResponse = types.SetDoorPasscodesResponse
type OpenDoorResponse = types.OpenDoorResponse
type GetStatusResponse = types.GetStatusResponse
type GetCardsResponse = types.GetCardsResponse
type GetCardResponse = types.GetCardResponse
type GetCardAtIndexResponse = types.GetCardAtIndexResponse
type PutCardResponse = types.PutCardResponse
type DeleteCardResponse = types.DeleteCardResponse
type DeleteAllCardsResponse = types.DeleteAllCardsResponse
type GetEventResponse = types.GetEventResponse
type GetEventIndexResponse = types.GetEventIndexResponse
type SetEventIndexResponse = types.SetEventIndexResponse
type RecordSpecialEventsResponse = types.RecordSpecialEventsResponse

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
	} else {
		return codec.Decode[R](reply)
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
		return Controller{
			ID: v,
		}, nil

	case Controller:
		return v, nil
	}

	return Controller{}, fmt.Errorf("unsupported type (%T)", controller)
}

func debugf(tag string, format string, args ...any) {
	log.Debugf(tag, format, args...)
}

// func infof(tag string, format string, args ...any) {
// 	log.Infof(tag, format, args...)
// }

// func warnf(tag string, format string, args ...any) {
// 	log.Warnf(tag, format, args...)
// }

// func errorf(tag string, format string, args ...any) {
// 	log.Errorf(tag, format, args...)
// }
