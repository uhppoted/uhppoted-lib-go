// generated code - ** DO NOT EDIT **

package uhppoted

import (
	"net/netip"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"
)

var Expected = struct {
	FindControllers          []responses.GetController
	GetController            responses.GetController
	SetIPv4                  responses.SetIPv4
	GetTime                  responses.GetTime
	SetTime                  responses.SetTime
	GetListener              responses.GetListener
	SetListener              responses.SetListener
	GetListenerAddrPort      responses.GetListenerAddrPort
	SetListenerAddrPort      responses.SetListenerAddrPort
	GetDoor                  responses.GetDoor
	SetDoor                  responses.SetDoor
	SetDoorPasscodes         responses.SetDoorPasscodes
	OpenDoor                 responses.OpenDoor
	GetStatus                responses.GetStatus
	GetStatusNoEvent         responses.GetStatus
	GetCards                 responses.GetCards
	GetCard                  responses.GetCard
	GetCardNotFound          responses.GetCard
	GetCardAtIndex           responses.GetCardAtIndex
	GetCardAtIndexNotFound   responses.GetCardAtIndex
	GetCardAtIndexDeleted    responses.GetCardAtIndex
	PutCard                  responses.PutCard
	DeleteCard               responses.DeleteCard
	DeleteAllCards           responses.DeleteAllCards
	GetEvent                 responses.GetEvent
	GetEventNotFound         responses.GetEvent
	GetEventOverwritten      responses.GetEvent
	GetEventIndex            responses.GetEventIndex
	SetEventIndex            responses.SetEventIndex
	RecordSpecialEvents      responses.RecordSpecialEvents
	GetTimeProfile           responses.GetTimeProfile
	SetTimeProfile           responses.SetTimeProfile
	ClearTimeProfiles        responses.ClearTimeProfiles
	AddTask                  responses.AddTask
	RefreshTaskList          responses.RefreshTaskList
	ClearTaskList            responses.ClearTaskList
	SetPcControl             responses.SetPCControl
	SetInterlock             responses.SetInterlock
	ActivateKeypads          responses.ActivateKeypads
	GetAntipassback          responses.GetAntiPassback
	SetAntipassback          responses.SetAntiPassback
	RestoreDefaultParameters responses.RestoreDefaultParameters
}{
	FindControllers: []responses.GetController{
		responses.GetController{
			Controller: 201020304,
			IpAddress:  netip.MustParseAddr("192.168.1.101"),
			SubnetMask: netip.MustParseAddr("255.255.255.0"),
			Gateway:    netip.MustParseAddr("192.168.1.1"),
			MACAddress: "52:fd:fc:07:21:82",
			Version:    "v6.62",
			Date:       types.MustParseDate("2020-01-01"),
		},
		responses.GetController{
			Controller: 303986753,
			IpAddress:  netip.MustParseAddr("192.168.1.100"),
			SubnetMask: netip.MustParseAddr("255.255.255.0"),
			Gateway:    netip.MustParseAddr("192.168.1.1"),
			MACAddress: "52:fd:fc:07:21:82",
			Version:    "v8.92",
			Date:       types.MustParseDate("2019-08-15"),
		},
		responses.GetController{
			Controller: 405419896,
			IpAddress:  netip.MustParseAddr("192.168.1.100"),
			SubnetMask: netip.MustParseAddr("255.255.255.0"),
			Gateway:    netip.MustParseAddr("192.168.1.1"),
			MACAddress: "00:12:23:34:45:56",
			Version:    "v8.92",
			Date:       types.MustParseDate("2018-11-05"),
		},
	},

	GetController: responses.GetController{
		Controller: 405419896,
		IpAddress:  netip.MustParseAddr("192.168.1.100"),
		SubnetMask: netip.MustParseAddr("255.255.255.0"),
		Gateway:    netip.MustParseAddr("192.168.1.1"),
		MACAddress: "00:12:23:34:45:56",
		Version:    "v8.92",
		Date:       types.MustParseDate("2018-11-05"),
	},

	SetIPv4: responses.SetIPv4{
		Controller: 405419896,
		Ok:         true,
	},

	GetTime: responses.GetTime{
		Controller: 405419896,
		DateTime:   types.MustParseDateTime("2025-11-01 12:34:56"),
	},

	SetTime: responses.SetTime{
		Controller: 405419896,
		DateTime:   types.MustParseDateTime("2025-11-01 12:34:56"),
	},

	GetListener: responses.GetListener{
		Controller: 405419896,
		Address:    netip.MustParseAddr("192.168.1.100"),
		Port:       60001,
		Interval:   13,
	},

	SetListener: responses.SetListener{
		Controller: 405419896,
		Ok:         true,
	},

	GetListenerAddrPort: responses.GetListenerAddrPort{
		Controller: 405419897,
		Listener:   netip.MustParseAddrPort("192.168.1.100:60001"),
		Interval:   13,
	},

	SetListenerAddrPort: responses.SetListenerAddrPort{
		Controller: 405419897,
		Ok:         true,
	},

	GetDoor: responses.GetDoor{
		Controller: 405419896,
		Door:       4,
		Mode:       types.DoorMode(3),
		Delay:      7,
	},

	SetDoor: responses.SetDoor{
		Controller: 405419896,
		Door:       4,
		Mode:       types.DoorMode(2),
		Delay:      17,
	},

	SetDoorPasscodes: responses.SetDoorPasscodes{
		Controller: 405419896,
		Ok:         true,
	},

	OpenDoor: responses.OpenDoor{
		Controller: 405419896,
		Ok:         true,
	},

	GetStatus: responses.GetStatus{
		Controller:         405419896,
		SystemDate:         types.MustParseDate("2022-08-23"),
		SystemTime:         types.MustParseTime("09:49:39"),
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
		EventType:          types.EventType(2),
		EventAccessGranted: true,
		EventDoor:          3,
		EventDirection:     types.Direction(1),
		EventCard:          8165537,
		EventTimestamp:     types.MustParseDateTime("2022-08-23 09:47:06"),
		EventReason:        types.Reason(44),
		SequenceNo:         0,
	},

	GetStatusNoEvent: responses.GetStatus{
		Controller:         405419897,
		SystemDate:         types.MustParseDate("2025-11-23"),
		SystemTime:         types.MustParseTime("14:37:53"),
		Door1Open:          true,
		Door2Open:          false,
		Door3Open:          true,
		Door4Open:          true,
		Door1Button:        true,
		Door2Button:        true,
		Door3Button:        false,
		Door4Button:        true,
		Relays:             7,
		Inputs:             9,
		SystemError:        27,
		SpecialInfo:        39,
		EventIndex:         0,
		EventType:          types.EventType(0),
		EventAccessGranted: false,
		EventDoor:          0,
		EventDirection:     types.Direction(0),
		EventCard:          0,
		EventTimestamp:     types.MustParseDateTime("0001-01-01 00:00:00"),
		EventReason:        types.Reason(0),
		SequenceNo:         21987,
	},

	GetCards: responses.GetCards{
		Controller: 405419896,
		Cards:      13579,
	},

	GetCard: responses.GetCard{
		Controller: 405419896,
		Card:       10058400,
		StartDate:  types.MustParseDate("2025-01-01"),
		EndDate:    types.MustParseDate("2025-12-31"),
		Door1:      1,
		Door2:      0,
		Door3:      17,
		Door4:      1,
		PIN:        7531,
	},

	GetCardNotFound: responses.GetCard{
		Controller: 405419896,
		Card:       0,
		StartDate:  types.MustParseDate("0001-01-01"),
		EndDate:    types.MustParseDate("0001-01-01"),
		Door1:      0,
		Door2:      0,
		Door3:      0,
		Door4:      0,
		PIN:        0,
	},

	GetCardAtIndex: responses.GetCardAtIndex{
		Controller: 405419896,
		Card:       10058400,
		StartDate:  types.MustParseDate("2025-01-01"),
		EndDate:    types.MustParseDate("2025-12-31"),
		Door1:      1,
		Door2:      0,
		Door3:      17,
		Door4:      1,
		PIN:        7531,
	},

	GetCardAtIndexNotFound: responses.GetCardAtIndex{
		Controller: 405419896,
		Card:       0,
		StartDate:  types.MustParseDate("0001-01-01"),
		EndDate:    types.MustParseDate("0001-01-01"),
		Door1:      0,
		Door2:      0,
		Door3:      0,
		Door4:      0,
		PIN:        0,
	},

	GetCardAtIndexDeleted: responses.GetCardAtIndex{
		Controller: 405419896,
		Card:       4294967295,
		StartDate:  types.MustParseDate("0001-01-01"),
		EndDate:    types.MustParseDate("0001-01-01"),
		Door1:      0,
		Door2:      0,
		Door3:      0,
		Door4:      0,
		PIN:        0,
	},

	PutCard: responses.PutCard{
		Controller: 405419896,
		Ok:         true,
	},

	DeleteCard: responses.DeleteCard{
		Controller: 405419896,
		Ok:         true,
	},

	DeleteAllCards: responses.DeleteAllCards{
		Controller: 405419896,
		Ok:         true,
	},

	GetEvent: responses.GetEvent{
		Controller:    405419896,
		Index:         13579,
		Timestamp:     types.MustParseDateTime("2025-11-17 12:34:56"),
		EventType:     types.EventType(2),
		AccessGranted: true,
		Door:          4,
		Direction:     types.Direction(2),
		Card:          10058400,
		Reason:        types.Reason(21),
	},

	GetEventNotFound: responses.GetEvent{
		Controller:    405419896,
		Index:         24680,
		Timestamp:     types.MustParseDateTime("0001-01-01 00:00:00"),
		EventType:     types.EventType(0),
		AccessGranted: false,
		Door:          0,
		Direction:     types.Direction(0),
		Card:          0,
		Reason:        types.Reason(0),
	},

	GetEventOverwritten: responses.GetEvent{
		Controller:    405419896,
		Index:         98765,
		Timestamp:     types.MustParseDateTime("0001-01-01 00:00:00"),
		EventType:     types.EventType(255),
		AccessGranted: false,
		Door:          0,
		Direction:     types.Direction(0),
		Card:          0,
		Reason:        types.Reason(0),
	},

	GetEventIndex: responses.GetEventIndex{
		Controller: 405419896,
		Index:      13579,
	},

	SetEventIndex: responses.SetEventIndex{
		Controller: 405419896,
		Ok:         true,
	},

	RecordSpecialEvents: responses.RecordSpecialEvents{
		Controller: 405419896,
		Ok:         true,
	},

	GetTimeProfile: responses.GetTimeProfile{
		Controller:    405419896,
		Profile:       37,
		StartDate:     types.MustParseDate("2025-11-26"),
		EndDate:       types.MustParseDate("2025-12-29"),
		Monday:        true,
		Tuesday:       true,
		Wednesday:     false,
		Thursday:      true,
		Friday:        false,
		Saturday:      true,
		Sunday:        true,
		Segment1Start: types.MustParseHHmm("08:30"),
		Segment1End:   types.MustParseHHmm("09:45"),
		Segment2Start: types.MustParseHHmm("11:35"),
		Segment2End:   types.MustParseHHmm("13:15"),
		Segment3Start: types.MustParseHHmm("14:01"),
		Segment3End:   types.MustParseHHmm("17:59"),
		LinkedProfile: 19,
	},

	SetTimeProfile: responses.SetTimeProfile{
		Controller: 405419896,
		Ok:         true,
	},

	ClearTimeProfiles: responses.ClearTimeProfiles{
		Controller: 405419896,
		Ok:         true,
	},

	AddTask: responses.AddTask{
		Controller: 405419896,
		Ok:         true,
	},

	RefreshTaskList: responses.RefreshTaskList{
		Controller: 405419896,
		Ok:         true,
	},

	ClearTaskList: responses.ClearTaskList{
		Controller: 405419896,
		Ok:         true,
	},

	SetPcControl: responses.SetPCControl{
		Controller: 405419896,
		Ok:         true,
	},

	SetInterlock: responses.SetInterlock{
		Controller: 405419896,
		Ok:         true,
	},

	ActivateKeypads: responses.ActivateKeypads{
		Controller: 405419896,
		Ok:         true,
	},

	GetAntipassback: responses.GetAntiPassback{
		Controller:   405419896,
		Antipassback: 2,
	},

	SetAntipassback: responses.SetAntiPassback{
		Controller: 405419896,
		Ok:         true,
	},

	RestoreDefaultParameters: responses.RestoreDefaultParameters{
		Controller: 405419896,
		Ok:         true,
	},
}
