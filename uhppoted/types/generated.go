package types

import (
	"net/netip"
	"time"
)

// Container struct for the response returned from an access controller when retrieving the
// network configuration, firmware version and firmware release date.
type GetControllerResponse struct {
	Controller uint32     `json:"controller"`
	IpAddress  netip.Addr `json:"ip-address"`
	SubnetMask netip.Addr `json:"netmask"`
	Gateway    netip.Addr `json:"gateway"`
	MACAddress string     `json:"MAC"`
	Version    string     `json:"version"`
	Date       time.Time  `json:"date"`
}

// SetIPv4Response is a synthesized response provided to simplify code generation. The controller
// does not return a response to a 'set-IPv4' request.
type SetIPv4Response struct {
	Controller uint32 `json:"-"`
	Ok         bool   `json:"-"`
}
