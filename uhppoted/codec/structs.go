// Package codec defines the data structures and helper types for encoding and
// decoding messages exchanged with UHPPOTE controllers.
package codec

import (
	"net/netip"
	"time"
)

// GetControllerResponse is a container struct for the response returned from a controller
// when retrieving the network configuration, firmware version and firmware release date.
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

// GetStatusResponse is a container struct for the response returned from a controller
// when retrieving the current runtime status.
type GetStatusResponse struct {
	Controller         uint32    `json:"controller"`
	SystemDate         time.Time `json:"system-date"`
	SystemTime         time.Time `json:"system-time"`
	Door1Open          bool      `json:"door-1-open"`
	Door2Open          bool      `json:"door-2-open"`
	Door3Open          bool      `json:"door-3-open"`
	Door4Open          bool      `json:"door-4-open"`
	Door1Button        bool      `json:"door-1-button"`
	Door2Button        bool      `json:"door-2-button"`
	Door3Button        bool      `json:"door-3-button"`
	Door4Button        bool      `json:"door-4-button"`
	Relays             uint8     `json:"relays"`
	Inputs             uint8     `json:"alarm-inputs"`
	SystemError        uint8     `json:"system-error"`
	SpecialInfo        uint8     `json:"special-info"`
	EventIndex         uint32    `json:"event-index"`
	EventType          uint8     `json:"event-type"`
	EventAccessGranted bool      `json:"event-granted"`
	EventDoor          uint8     `json:"event-door"`
	EventDirection     uint8     `json:"event-direction"`
	EventCard          uint32    `json:"event-card"`
	EventTimestamp     time.Time `json:"event-timestamp"`
	EventReason        uint8     `json:"event-reason"`
	SequenceNo         uint32    `json:"sequence-no"`
}

// GetTimeResponse is a container struct for the response returned by a controller
// when retrieving the current date/time.
type GetTimeResponse struct {
	Controller uint32    `json:"controller"`
	DateTime   time.Time `json:"datetime"`
}
