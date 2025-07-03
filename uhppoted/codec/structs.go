package codec

import (
	"net/netip"
	"time"
)

type GetControllerResponse struct {
	Controller uint32     `json:"controller"`
	IpAddress  netip.Addr `json:"ip-address"`
	SubnetMask netip.Addr `json:"subnet-mask"`
	Gateway    netip.Addr `json:"gateway"`
	MACAddress string     `json:"MAC-address"`
	Version    string     `json:"version"`
	Date       time.Time  `json:"date"`
}
