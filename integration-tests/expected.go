package uhppoted

import (
	"net/netip"
	"time"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

var Expected = struct {
	GetAllControllers []lib.GetControllerResponse
	GetController     lib.GetControllerResponse
}{
	GetAllControllers: []lib.GetControllerResponse{
		lib.GetControllerResponse{
			Controller: 201020304,
			IpAddress:  netip.MustParseAddr("192.168.1.101"),
			SubnetMask: netip.MustParseAddr("255.255.255.0"),
			Gateway:    netip.MustParseAddr("192.168.1.1"),
			MACAddress: "52:fd:fc:07:21:82",
			Version:    "v6.62",
			Date:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
		},
		lib.GetControllerResponse{
			Controller: 303986753,
			IpAddress:  netip.MustParseAddr("192.168.1.100"),
			SubnetMask: netip.MustParseAddr("255.255.255.0"),
			Gateway:    netip.MustParseAddr("192.168.1.1"),
			MACAddress: "52:fd:fc:07:21:82",
			Version:    "v8.92",
			Date:       time.Date(2019, 8, 15, 0, 0, 0, 0, time.Local),
		},
		lib.GetControllerResponse{
			Controller: 405419896,
			IpAddress:  netip.MustParseAddr("192.168.1.100"),
			SubnetMask: netip.MustParseAddr("255.255.255.0"),
			Gateway:    netip.MustParseAddr("192.168.1.1"),
			MACAddress: "00:12:23:34:45:56",
			Version:    "v8.92",
			Date:       time.Date(2018, 11, 5, 0, 0, 0, 0, time.Local),
		},
	},

	GetController: lib.GetControllerResponse{
		Controller: 405419896,
		IpAddress:  netip.MustParseAddr("192.168.1.100"),
		SubnetMask: netip.MustParseAddr("255.255.255.0"),
		Gateway:    netip.MustParseAddr("192.168.1.1"),
		MACAddress: "00:12:23:34:45:56",
		Version:    "v8.92",
		Date:       time.Date(2018, 11, 5, 0, 0, 0, 0, time.Local),
	},
}
