// Package codec defines the data structures and helper types for encoding and
// decoding messages exchanged with UHPPOTE controllers.
package codec

import (
	"net/netip"
	"time"
)

// GetControllerResponse represents the response returned from a controller
// when querying its network configuration, firmware version and firmeare release
// date.
type GetControllerResponse struct {
	Controller uint32     `json:"controller"`
	IpAddress  netip.Addr `json:"ip-address"`
	SubnetMask netip.Addr `json:"subnet-mask"`
	Gateway    netip.Addr `json:"gateway"`
	MACAddress string     `json:"MAC-address"`
	Version    string     `json:"version"`
	Date       time.Time  `json:"date"`
}

// SetIPv4Response is a synthesized response - the controller does not return a respons
// to a 'set-IPv4' request.
type SetIPv4Response struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}
