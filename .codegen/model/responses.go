package model

type Response struct {
	Name    string
	MsgType byte
	Fields  []Field
	Tests   []Test
}

var Responses = []Response{
	GetControllerResponse,
	SetIPv4Response,
	GetStatusResponse,
	GetTimeResponse,
	SetTimeResponse,
}

var GetControllerResponse = Response{
	Name:    "get controller",
	MsgType: 0x94,
	Fields: []Field{
		{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
		{"ip address", "IPv4", 8, "controller IPv4 address, e.g. 192.168.1.100"},
		{"subnet mask", "IPv4", 12, "controller IPv4 netmask, e.g. 255.255.255.0"},
		{"gateway", "IPv4", 16, "controller IPv4 gateway address, e.g. 192.168.1.1"},
		{"MAC address", "MAC", 20, "controller MAC address, e.g. 52:fd:fc:07:21:82"},
		{"version", "version", 26, "controller firmware version, e.g. v6.62"},
		{"date", "date", 28, "controller firmware release date, e.g. 2020-12-31"},
	},
	Tests: []Test{
		{
			Name: "get-controller",
			Response: []byte{
				0x17, 0x94, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0xc0, 0xa8, 0x01, 0x64, 0xff, 0xff, 0xff, 0x00,
				0xc0, 0xa8, 0x01, 0x01, 0x00, 0x12, 0x23, 0x34, 0x45, 0x56, 0x08, 0x92, 0x20, 0x18, 0x11, 0x05,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Values: []Value{
				Value{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				Value{
					Name:  "ip address",
					Type:  "IPv4",
					Value: "192.168.1.100",
				},
				Value{
					Name:  "subnet mask",
					Type:  "IPv4",
					Value: "255.255.255.0",
				},
				Value{
					Name:  "gateway",
					Type:  "IPv4",
					Value: "192.168.1.1",
				},
				Value{
					Name:  "MAC address",
					Type:  "MAC",
					Value: "00:12:23:34:45:56",
				},
				Value{
					Name:  "version",
					Type:  "version",
					Value: "v8.92",
				},
				Value{
					Name:  "date",
					Type:  "date",
					Value: "2018-11-05",
				},
			},
		},
	},
}

var SetIPv4Response = Response{
	Name:    "set IPv4",
	MsgType: 0x96,
	Fields: []Field{
		{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
		{"ok", "bool", 8, "controller IPv4 network configured"},
	},
	Tests: []Test{
		{
			Name: "set-IPv4",
			Response: []byte{
				0x17, 0x96, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Values: []Value{
				Value{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				Value{
					Name:  "ok",
					Type:  "bool",
					Value: true,
				},
			},
		},
	},
}

var GetStatusResponse = Response{
	Name:    "get status",
	MsgType: 0x20,
	Fields: []Field{
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
	Tests: []Test{
		{
			Name: "get-status",
			Response: []byte{
				0x17, 0x20, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x4e, 0x00, 0x00, 0x00, 0x02, 0x01, 0x03, 0x01,
				0xa1, 0x98, 0x7c, 0x00, 0x20, 0x22, 0x08, 0x23, 0x09, 0x47, 0x06, 0x2c, 0x00, 0x01, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x01, 0x03, 0x09, 0x49, 0x39, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x27, 0x07, 0x09, 0x22, 0x08, 0x23, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Values: []Value{
				Value{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				Value{
					Name:  "system-date",
					Type:  "date",
					Value: "2022-08-23",
				},
				Value{
					Name:  "system-time",
					Type:  "time",
					Value: "09:49:39",
				},
				Value{
					Name:  "door-1-open",
					Type:  "bool",
					Value: false,
				},
				Value{
					Name:  "door-2-open",
					Type:  "bool",
					Value: true,
				},
				Value{
					Name:  "door-3-open",
					Type:  "bool",
					Value: false,
				},
				Value{
					Name:  "door-4-open",
					Type:  "bool",
					Value: false,
				},
				Value{
					Name:  "door-1-button",
					Type:  "bool",
					Value: false,
				},
				Value{
					Name:  "door-2-button",
					Type:  "bool",
					Value: false,
				},
				Value{
					Name:  "door-3-button",
					Type:  "bool",
					Value: false,
				},
				Value{
					Name:  "door-4-button",
					Type:  "bool",
					Value: true,
				},
				Value{
					Name:  "relays",
					Type:  "uint8",
					Value: 0x07,
				},
				Value{
					Name:  "inputs",
					Type:  "uint8",
					Value: 0x09,
				},
				Value{
					Name:  "system-error",
					Type:  "uint8",
					Value: 3,
				},
				Value{
					Name:  "special-info",
					Type:  "uint8",
					Value: 39,
				},
				Value{
					Name:  "event-index",
					Type:  "uint32",
					Value: 78,
				},
				Value{
					Name:  "event-type",
					Type:  "uint8",
					Value: 2,
				},
				Value{
					Name:  "event-access-granted",
					Type:  "bool",
					Value: true,
				},
				Value{
					Name:  "event-door",
					Type:  "uint8",
					Value: 3,
				},
				Value{
					Name:  "event-direction",
					Type:  "uint8",
					Value: 1,
				},
				Value{
					Name:  "event-card",
					Type:  "uint32",
					Value: 8165537,
				},
				Value{
					Name:  "event-timestamp",
					Type:  "datetime",
					Value: "2022-08-23 09:47:06",
				},
				Value{
					Name:  "event-reason",
					Type:  "uint8",
					Value: 44,
				},
				Value{
					Name:  "sequence-no",
					Type:  "uint32",
					Value: 0,
				},
			},
		},
	},
}

var GetTimeResponse = Response{
	Name:    "get time",
	MsgType: 0x32,
	Fields: []Field{
		{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
		{"date-time", "datetime", 8, "current date/time"},
	},
	Tests: []Test{
		{
			Name: "get-time",
			Response: []byte{
				0x17, 0x32, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x20, 0x24, 0x11, 0x01, 0x12, 0x34, 0x56, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Values: []Value{
				Value{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				Value{
					Name:  "date-time",
					Type:  "datetime",
					Value: "2024-11-01 12:34:56",
				},
			},
		},
	},
}

var SetTimeResponse = Response{
	Name:    "set time",
	MsgType: 0x30,
	Fields: []Field{
		{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
		{"date-time", "datetime", 8, "current date/time"},
	},
	Tests: []Test{
		{
			Name: "set-time",
			Response: []byte{
				0x17, 0x30, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x20, 0x24, 0x11, 0x01, 0x12, 0x34, 0x56, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			Values: []Value{
				Value{
					Name:  "controller",
					Type:  "uint32",
					Value: 405419896,
				},
				Value{
					Name:  "date-time",
					Type:  "datetime",
					Value: "2024-11-01 12:34:56",
				},
			},
		},
	},
}
