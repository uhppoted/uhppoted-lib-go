package uhppoted

import (
	"net/netip"
)

type Uhppoted struct {
	BindAddr      netip.AddrPort
	BroadcastAddr netip.AddrPort
	ListenAddr    netip.AddrPort
	Debug         bool
}
