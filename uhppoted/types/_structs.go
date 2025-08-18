package uhppoted

import (
	"net/netip"
	"time"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/encode"
)

type GetControllerResponse struct {
	Controller	uint32	json:"controller"
	IpAddress	uint32	json:"ip-address"
	SubnetMask	uint32	json:"netmask"
	Gateway		uint32	json:"gateway"
	MACAddress	uint32	json:"MAC"
	Version		uint32	json:"version"
	Date		uint32	json:"version"
}
