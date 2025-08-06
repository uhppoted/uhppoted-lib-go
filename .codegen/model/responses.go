package model

import (
	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/model/types"
)

var Responses = []types.Response{
	GetControllerResponse,
	SetIPv4Response,
	GetStatusResponse,
	GetTimeResponse,
	SetTimeResponse,
	GetListenerResponse,
	SetListenerResponse,
	GetDoorResponse,
	SetDoorResponse,
	SetDoorPasscodesResponse,
	OpenDoorResponse,
	GetCardsResponse,
	GetCardResponse,
}

// var GetCardResponse = types.Response(lib.GetCardResponse)

var GetControllerResponse = types.Response{
	Message: lib.Message{
		Name:    "get controller",
		MsgType: 0x94,
		Fields: []lib.Field{
			{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
			{"ip address", "IPv4", 8, "controller IPv4 address, e.g. 192.168.1.100"},
			{"subnet mask", "IPv4", 12, "controller IPv4 netmask, e.g. 255.255.255.0"},
			{"gateway", "IPv4", 16, "controller IPv4 gateway address, e.g. 192.168.1.1"},
			{"MAC address", "MAC", 20, "controller MAC address, e.g. 52:fd:fc:07:21:82"},
			{"version", "version", 26, "controller firmware version, e.g. v6.62"},
			{"date", "date", 28, "controller firmware release date, e.g. 2020-12-31"},
		},
	},
	Tests: []types.ResponseTest{
		{
			Name: "get-controller",
			Response: []byte{
				0x17, 0x94, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0xc0, 0xa8, 0x01, 0x64, 0xff, 0xff, 0xff, 0x00,
				0xc0, 0xa8, 0x01, 0x01, 0x00, 0x12, 0x23, 0x34, 0x45, 0x56, 0x08, 0x92, 0x20, 0x18, 0x11, 0x05,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Expected: []types.Value{
				{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				{
					Name:  "ip address",
					Type:  "IPv4",
					Value: "192.168.1.100",
				},
				{
					Name:  "subnet mask",
					Type:  "IPv4",
					Value: "255.255.255.0",
				},
				{
					Name:  "gateway",
					Type:  "IPv4",
					Value: "192.168.1.1",
				},
				{
					Name:  "MAC address",
					Type:  "MAC",
					Value: "00:12:23:34:45:56",
				},
				{
					Name:  "version",
					Type:  "version",
					Value: "v8.92",
				},
				{
					Name:  "date",
					Type:  "date",
					Value: "2018-11-05",
				},
			},
		},
	},
}

var SetIPv4Response = types.Response{
	Message: lib.Message{
		Name:    "set IPv4",
		MsgType: 0x96,
		Fields: []lib.Field{
			{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
			{"ok", "bool", 8, "controller IPv4 network configured"},
		},
	},
	Tests: []types.ResponseTest{
		{
			Name: "set-IPv4",
			Response: []byte{
				0x17, 0x96, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Expected: []types.Value{
				{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				{
					Name:  "ok",
					Type:  "bool",
					Value: true,
				},
			},
		},
	},
}

var GetStatusResponse = types.Response{
	Message: lib.Message{
		Name:    "get status",
		MsgType: 0x20,
		Fields: []lib.Field{
			{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
			{"system-date", "shortdate", 51, "current date, e.g. 2025-07-21"},
			{"system-time", "time", 37, "current time, e.g. 13:25:47"},
			{"door-1-open", "bool", 28, "true if door 1 is open"},
			{"door-2-open", "bool", 29, "true if door 2 is open"},
			{"door-3-open", "bool", 30, "true if door 3 is open"},
			{"door-4-open", "bool", 31, "true if door 4 is open"},
			{"door-1-button", "bool", 32, "true if door 1 button is pressed"},
			{"door-2-button", "bool", 33, "true if door 2 button is pressed"},
			{"door-3-button", "bool", 34, "true if door 3 button is pressed"},
			{"door-4-button", "bool", 35, "true if door 4 button is pressed"},
			{"relays", "uint8", 49, "bitset of door unlock relay states"},
			{"inputs", "uint8", 50, "bitset of alarm inputs"},
			{"system-error", "uint8", 36, "system error code"},
			{"special-info", "uint8", 48, "absolutely no idea"},
			{"event-index", "uint32", 8, "last event index"},
			{"event-type", "uint8", 12, "last event type"},
			{"event-access-granted", "bool", 13, "last event access granted"},
			{"event-door", "uint8", 14, "last event door"},
			{"event-direction", "uint8", 15, "last event door direction (0: in, 1: out)"},
			{"event-card", "uint32", 16, "last event card number"},
			{"event-timestamp", "datetime", 20, "last event timestamp"},
			{"event-reason", "uint8", 27, "last event reason"},
			{"sequence-no", "uint32", 40, "packet sequence number"},
		},
	},
	Tests: []types.ResponseTest{
		{
			Name: "get-status",
			Response: []byte{
				0x17, 0x20, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x4e, 0x00, 0x00, 0x00, 0x02, 0x01, 0x03, 0x01,
				0xa1, 0x98, 0x7c, 0x00, 0x20, 0x22, 0x08, 0x23, 0x09, 0x47, 0x06, 0x2c, 0x00, 0x01, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x01, 0x03, 0x09, 0x49, 0x39, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x27, 0x07, 0x09, 0x22, 0x08, 0x23, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Expected: []types.Value{
				{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				{
					Name:  "system-date",
					Type:  "date",
					Value: "2022-08-23",
				},
				{
					Name:  "system-time",
					Type:  "time",
					Value: "09:49:39",
				},
				{
					Name:  "door-1-open",
					Type:  "bool",
					Value: false,
				},
				{
					Name:  "door-2-open",
					Type:  "bool",
					Value: true,
				},
				{
					Name:  "door-3-open",
					Type:  "bool",
					Value: false,
				},
				{
					Name:  "door-4-open",
					Type:  "bool",
					Value: false,
				},
				{
					Name:  "door-1-button",
					Type:  "bool",
					Value: false,
				},
				{
					Name:  "door-2-button",
					Type:  "bool",
					Value: false,
				},
				{
					Name:  "door-3-button",
					Type:  "bool",
					Value: false,
				},
				{
					Name:  "door-4-button",
					Type:  "bool",
					Value: true,
				},
				{
					Name:  "relays",
					Type:  "uint8",
					Value: 0x07,
				},
				{
					Name:  "inputs",
					Type:  "uint8",
					Value: 0x09,
				},
				{
					Name:  "system-error",
					Type:  "uint8",
					Value: 3,
				},
				{
					Name:  "special-info",
					Type:  "uint8",
					Value: 39,
				},
				{
					Name:  "event-index",
					Type:  "uint32",
					Value: 78,
				},
				{
					Name:  "event-type",
					Type:  "uint8",
					Value: 2,
				},
				{
					Name:  "event-access-granted",
					Type:  "bool",
					Value: true,
				},
				{
					Name:  "event-door",
					Type:  "uint8",
					Value: 3,
				},
				{
					Name:  "event-direction",
					Type:  "uint8",
					Value: 1,
				},
				{
					Name:  "event-card",
					Type:  "uint32",
					Value: 8165537,
				},
				{
					Name:  "event-timestamp",
					Type:  "datetime",
					Value: "2022-08-23 09:47:06",
				},
				{
					Name:  "event-reason",
					Type:  "uint8",
					Value: 44,
				},
				{
					Name:  "sequence-no",
					Type:  "uint32",
					Value: 0,
				},
			},
		},
	},
}

var GetTimeResponse = types.Response{
	Message: lib.Message{
		Name:    "get time",
		MsgType: 0x32,
		Fields: []lib.Field{
			{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
			{"date-time", "datetime", 8, "controller system date/time"},
		},
	},
	Tests: []types.ResponseTest{
		{
			Name: "get-time",
			Response: []byte{
				0x17, 0x32, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x20, 0x24, 0x11, 0x01, 0x12, 0x34, 0x56, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Expected: []types.Value{
				{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				{
					Name:  "date-time",
					Type:  "datetime",
					Value: "2024-11-01 12:34:56",
				},
			},
		},
	},
}

var SetTimeResponse = types.Response{
	Message: lib.Message{
		Name:    "set time",
		MsgType: 0x30,
		Fields: []lib.Field{
			{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
			{"date-time", "datetime", 8, "controller system date/time"},
		},
	},
	Tests: []types.ResponseTest{
		{
			Name: "set-time",
			Response: []byte{
				0x17, 0x30, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x20, 0x24, 0x11, 0x01, 0x12, 0x34, 0x56, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Expected: []types.Value{
				{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				{
					Name:  "date-time",
					Type:  "datetime",
					Value: "2024-11-01 12:34:56",
				},
			},
		},
	},
}

var GetListenerResponse = types.Response{
	Message: lib.Message{
		Name:    "get listener",
		MsgType: 0x92,
		Fields: []lib.Field{
			{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
			{"address", "addrport", 8, "event listener IPv4 address:port"},
			{"interval", "uint8", 14, "auto-send interval (seconds)"},
		},
	},
	Tests: []types.ResponseTest{
		{
			Name: "get-listener",
			Response: []byte{
				0x17, 0x92, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0xc0, 0xa8, 0x01, 0x64, 0x61, 0xea, 0x11, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Expected: []types.Value{
				{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				{
					Name:  "address",
					Type:  "addrport",
					Value: "192.168.1.100:60001",
				},
				{
					Name:  "interval",
					Type:  "uint8",
					Value: "17",
				},
			},
		},
	},
}

var SetListenerResponse = types.Response{
	Message: lib.Message{
		Name:    "set listener",
		MsgType: 0x90,
		Fields: []lib.Field{
			{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
			{"ok", "bool", 8, "set-listener succeeded/failed"},
		},
	},
	Tests: []types.ResponseTest{
		{
			Name: "set-listener",
			Response: []byte{
				0x17, 0x90, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Expected: []types.Value{
				{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				{
					Name:  "ok",
					Type:  "bool",
					Value: true,
				},
			},
		},
	},
}

var GetDoorResponse = types.Response{
	Message: lib.Message{
		Name:    "get door",
		MsgType: 0x82,
		Fields: []lib.Field{
			{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
			{"door", "uint8", 8, "door ID ([1..4]"},
			{"mode", "uint8", 9, "control mode (1:normally open, 2:normally closed. 3:controlled)"},
			{"delay", "uint8", 10, "unlock delay (seconds)"},
		},
	},
	Tests: []types.ResponseTest{
		{
			Name: "get-door",
			Response: []byte{
				0x17, 0x82, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x03, 0x02, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Expected: []types.Value{
				{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				{
					Name:  "door",
					Type:  "uint8",
					Value: 3,
				},
				{
					Name:  "mode",
					Type:  "uint8",
					Value: 2,
				},
				{
					Name:  "delay",
					Type:  "uint8",
					Value: 7,
				},
			},
		},
	},
}

var SetDoorResponse = types.Response{
	Message: lib.Message{
		Name:    "set door",
		MsgType: 0x80,
		Fields: []lib.Field{
			{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
			{"door", "uint8", 8, "door ID ([1..4]"},
			{"mode", "uint8", 9, "control mode (1:normally open, 2:normally closed. 3:controlled)"},
			{"delay", "uint8", 10, "unlock delay (seconds)"},
		},
	},
	Tests: []types.ResponseTest{
		{
			Name: "set-door",
			Response: []byte{
				0x17, 0x80, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x03, 0x02, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Expected: []types.Value{
				{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				{
					Name:  "door",
					Type:  "uint8",
					Value: 3,
				},
				{
					Name:  "mode",
					Type:  "uint8",
					Value: 2,
				},
				{
					Name:  "delay",
					Type:  "uint8",
					Value: 7,
				},
			},
		},
	},
}

var SetDoorPasscodesResponse = types.Response{
	Message: lib.Message{
		Name:    "set door passcodes",
		MsgType: 0x8c,
		Fields: []lib.Field{
			{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
			{"ok", "bool", 8, "set-door-passcodes succeeded/failed"},
		},
	},
	Tests: []types.ResponseTest{
		{
			Name: "set-door-passcodes",
			Response: []byte{
				0x17, 0x8c, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Expected: []types.Value{
				{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				{
					Name:  "ok",
					Type:  "bool",
					Value: true,
				},
			},
		},
	},
}

var OpenDoorResponse = types.Response{
	Message: lib.Message{
		Name:    "open door",
		MsgType: 0x40,
		Fields: []lib.Field{
			{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
			{"ok", "bool", 8, "open-door succeeded/failed"},
		},
	},
	Tests: []types.ResponseTest{
		{
			Name: "open-door",
			Response: []byte{
				0x17, 0x40, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Expected: []types.Value{
				{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				{
					Name:  "ok",
					Type:  "bool",
					Value: true,
				},
			},
		},
	},
}

var GetCardsResponse = types.Response{
	Message: lib.Message{
		Name:    "get cards",
		MsgType: 0x58,
		Fields: []lib.Field{
			{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
			{"cards", "uint32", 8, "number of stored cards"},
		},
	},
	Tests: []types.ResponseTest{
		{
			Name: "get-cards",
			Response: []byte{
				0x17, 0x58, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x0b, 0x35, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Expected: []types.Value{
				{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				{
					Name:  "cards",
					Type:  "uint32",
					Value: 13579,
				},
			},
		},
	},
}

var GetCardResponse = types.Response{
	Message: lib.Message{
		Name:    "get card",
		MsgType: 0x5a,
		Fields: []lib.Field{
			{
				Name:        "controller",
				Type:        "uint32",
				Offset:      4,
				Description: "controller serial number",
			},
			{
				Name:   "card",
				Type:   "uint32",
				Offset: 8,
			},
			{
				Name:   "start date",
				Type:   "optional date",
				Offset: 12,
			},
			{
				Name:   "end date",
				Type:   "optional date",
				Offset: 16,
			},
			{
				Name:   "door 1",
				Type:   "uint8",
				Offset: 20,
			},
			{
				Name:   "door 2",
				Type:   "uint8",
				Offset: 21,
			},
			{
				Name:   "door 3",
				Type:   "uint8",
				Offset: 22,
			},
			{
				Name:   "door 4",
				Type:   "uint8",
				Offset: 23,
			},
			{
				Name:   "PIN",
				Type:   "pin",
				Offset: 24,
			},
		},
	},
	Tests: []types.ResponseTest{
		{
			Name: "get-card",
			Response: []byte{
				0x17, 0x5a, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0xa0, 0x7a, 0x99, 0x00, 0x20, 0x24, 0x01, 0x01,
				0x20, 0x24, 0x12, 0x31, 0x01, 0x00, 0x11, 0x01, 0x3f, 0x42, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Expected: []types.Value{
				{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				{
					Name:  "card",
					Type:  "uint32",
					Value: 10058400,
				},
				{
					Name:  "start date",
					Type:  "date",
					Value: "2024-01-01",
				},
				{
					Name:  "end date",
					Type:  "date",
					Value: "2024-12-31",
				},
				{
					Name:  "door 1",
					Type:  "uint8",
					Value: 1,
				},
				{
					Name:  "door 2",
					Type:  "uint8",
					Value: 0,
				},
				{
					Name:  "door 3",
					Type:  "uint8",
					Value: 17,
				},
				{
					Name:  "door 4",
					Type:  "uint8",
					Value: 1,
				},
				{
					Name:  "PIN",
					Type:  "uint32",
					Value: 999999,
				},
			},
		},
	},
}
