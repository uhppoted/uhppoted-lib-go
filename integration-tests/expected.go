package uhppoted

import (
	"fmt"
	"net/netip"
	"time"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

var Expected = struct {
	FindControllers        []lib.GetControllerResponse
	GetController          lib.GetControllerResponse
	SetIPv4                lib.SetIPv4Response
	GetTime                lib.GetTimeResponse
	SetTime                lib.SetTimeResponse
	GetListener            lib.GetListenerResponse
	SetListener            lib.SetListenerResponse
	GetDoor                lib.GetDoorResponse
	SetDoor                lib.SetDoorResponse
	SetDoorPasscodes       lib.SetDoorPasscodesResponse
	OpenDoor               lib.OpenDoorResponse
	GetStatus              lib.GetStatusResponse
	GetCards               lib.GetCardsResponse
	GetCard                lib.GetCardResponse
	GetCardNotFound        lib.GetCardResponse
	GetCardAtIndex         lib.GetCardAtIndexResponse
	GetCardAtIndexNotFound lib.GetCardAtIndexResponse
	GetCardAtIndexDeleted  lib.GetCardAtIndexResponse
	PutCard                lib.PutCardResponse
	DeleteCard             lib.DeleteCardResponse
	DeleteAllCards         lib.DeleteAllCardsResponse
	GetEvent               lib.GetEventResponse
	GetEventNotFound       lib.GetEventResponse
	GetEventOverwritten    lib.GetEventResponse
	GetEventIndex          lib.GetEventIndexResponse
	SetEventIndex          lib.SetEventIndexResponse
}{
	FindControllers: []lib.GetControllerResponse{
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
		IpAddress:  IPv4("192.168.1.100"),
		SubnetMask: IPv4("255.255.255.0"),
		Gateway:    IPv4("192.168.1.1"),
		MACAddress: "00:12:23:34:45:56",
		Version:    "v8.92",
		Date:       string2date("2018-11-05"),
	},

	SetIPv4: lib.SetIPv4Response{
		Controller: 405419896,
		Ok:         true,
	},

	GetTime: lib.GetTimeResponse{
		Controller: 405419896,
		DateTime:   string2datetime("2025-11-01 12:34:56"),
	},

	SetTime: lib.SetTimeResponse{
		Controller: 405419896,
		DateTime:   string2datetime("2025-11-01 12:34:56"),
	},

	GetListener: lib.GetListenerResponse{
		Controller: 405419896,
		Listener:   addrport("192.168.1.100:60001"),
		Interval:   17,
	},

	SetListener: lib.SetListenerResponse{
		Controller: 405419896,
		Ok:         true,
	},

	GetDoor: lib.GetDoorResponse{
		Controller: 405419896,
		Door:       3,
		Mode:       2,
		Delay:      7,
	},

	SetDoor: lib.SetDoorResponse{
		Controller: 405419896,
		Door:       3,
		Mode:       2,
		Delay:      17,
	},

	SetDoorPasscodes: lib.SetDoorPasscodesResponse{
		Controller: 405419896,
		Ok:         true,
	},

	OpenDoor: lib.OpenDoorResponse{
		Controller: 405419896,
		Ok:         true,
	},

	GetStatus: lib.GetStatusResponse{
		Controller:         405419896,
		SystemDate:         string2date("2022-08-23"),
		SystemTime:         string2time("09:49:39"),
		Door1Open:          false,
		Door2Open:          true,
		Door3Open:          false,
		Door4Open:          false,
		Door1Button:        false,
		Door2Button:        false,
		Door3Button:        false,
		Door4Button:        true,
		Relays:             7,
		Inputs:             9,
		SystemError:        3,
		SpecialInfo:        39,
		EventIndex:         78,
		EventType:          2,
		EventAccessGranted: true,
		EventDoor:          3,
		EventDirection:     1,
		EventCard:          8165537,
		EventTimestamp:     string2datetime("2022-08-23 09:47:06"),
		EventReason:        44,
		SequenceNo:         0,
	},

	GetCards: lib.GetCardsResponse{
		Controller: 405419896,
		Cards:      13579,
	},

	GetCard: lib.GetCardResponse{
		Controller: 405419896,
		Card:       10058400,
		StartDate:  string2date("2025-01-01"),
		EndDate:    string2date("2025-12-31"),
		Door1:      1,
		Door2:      0,
		Door3:      17,
		Door4:      1,
		PIN:        7531,
	},

	GetCardNotFound: lib.GetCardResponse{
		Controller: 405419896,
		Card:       0,
		StartDate:  string2date("0001-01-01"),
		EndDate:    string2date("0001-01-01"),
		Door1:      0,
		Door2:      0,
		Door3:      0,
		Door4:      0,
		PIN:        0,
	},

	GetCardAtIndex: lib.GetCardAtIndexResponse{
		Controller: 405419896,
		Card:       10058400,
		StartDate:  string2date("2025-01-01"),
		EndDate:    string2date("2025-12-31"),
		Door1:      1,
		Door2:      0,
		Door3:      17,
		Door4:      1,
		PIN:        7531,
	},

	GetCardAtIndexNotFound: lib.GetCardAtIndexResponse{
		Controller: 405419896,
		Card:       0,
		StartDate:  string2date("0001-01-01"),
		EndDate:    string2date("0001-01-01"),
		Door1:      0,
		Door2:      0,
		Door3:      0,
		Door4:      0,
		PIN:        0,
	},

	GetCardAtIndexDeleted: lib.GetCardAtIndexResponse{
		Controller: 405419896,
		Card:       4294967295,
		StartDate:  string2date("0001-01-01"),
		EndDate:    string2date("0001-01-01"),
		Door1:      0,
		Door2:      0,
		Door3:      0,
		Door4:      0,
		PIN:        0,
	},

	PutCard: lib.PutCardResponse{
		Controller: 405419896,
		Ok:         true,
	},

	DeleteCard: lib.DeleteCardResponse{
		Controller: 405419896,
		Ok:         true,
	},

	DeleteAllCards: lib.DeleteAllCardsResponse{
		Controller: 405419896,
		Ok:         true,
	},

	GetEvent: lib.GetEventResponse{
		Controller:    405419896,
		Index:         13579,
		Timestamp:     string2datetime("2025-11-17 12:34:56"),
		EventType:     2,
		AccessGranted: true,
		Door:          4,
		Direction:     2,
		Card:          10058400,
		Reason:        21,
	},

	GetEventNotFound: lib.GetEventResponse{
		Controller:    405419896,
		Index:         24680,
		Timestamp:     string2datetime("0001-01-01 00:00:00"),
		EventType:     0,
		AccessGranted: false,
		Door:          0,
		Direction:     0,
		Card:          0,
		Reason:        0,
	},

	GetEventOverwritten: lib.GetEventResponse{
		Controller:    405419896,
		Index:         98765,
		Timestamp:     string2datetime("0001-01-01 00:00:00"),
		EventType:     255,
		AccessGranted: false,
		Door:          0,
		Direction:     0,
		Card:          0,
		Reason:        0,
	},

	GetEventIndex: lib.GetEventIndexResponse{
		Controller: 405419896,
		Index:      13579,
	},

	SetEventIndex: lib.SetEventIndexResponse{
		Controller: 405419896,
		Ok:         true,
	},
}

func IPv4(v string) netip.Addr {
	return netip.MustParseAddr(v)
}

func addrport(v string) netip.AddrPort {
	return netip.MustParseAddrPort(v)
}

func string2datetime(v string) time.Time {
	if d, err := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local); err != nil {
		panic(fmt.Sprintf("invalid datetime (%v)", v))
	} else {
		return d
	}
}

func string2date(v string) time.Time {
	if d, err := time.ParseInLocation("2006-01-02", v, time.Local); err != nil {
		panic(fmt.Sprintf("invalid date (%v)", v))
	} else {
		return d
	}
}

func string2time(v string) time.Time {
	if d, err := time.ParseInLocation("15:04:05", v, time.Local); err != nil {
		panic(fmt.Sprintf("invalid time (%v)", v))
	} else {
		return d
	}
}
