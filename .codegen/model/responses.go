package model

import (
	lib "github.com/uhppoted/uhppoted-codegen/model"
	// libx "github.com/uhppoted/uhppoted-codegen/model/types"

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
	GetCardAtIndexResponse,
	PutCardResponse,
	DeleteCardResponse,
	DeleteAllCardsResponse,
	GetEventResponse,
}

var GetControllerResponse = types.Response(lib.GetControllerResponse)
var SetIPv4Response = types.Response(lib.SetIPv4Response)
var GetTimeResponse = types.Response(lib.GetTimeResponse)
var SetTimeResponse = types.Response(lib.SetTimeResponse)
var GetListenerResponse = types.Response(lib.GetListenerAddrPortResponse)
var SetListenerResponse = types.Response(lib.SetListenerResponse)
var GetDoorResponse = types.Response(lib.GetDoorResponse)
var SetDoorResponse = types.Response(lib.SetDoorResponse)
var SetDoorPasscodesResponse = types.Response(lib.SetDoorPasscodesResponse)
var OpenDoorResponse = types.Response(lib.OpenDoorResponse)
var GetStatusResponse = types.Response(lib.GetStatusResponse)
var GetCardsResponse = types.Response(lib.GetCardsResponse)
var GetCardResponse = types.Response(lib.GetCardResponse)
var GetCardAtIndexResponse = types.Response(lib.GetCardAtIndexResponse)
var PutCardResponse = types.Response(lib.PutCardResponse)
var DeleteCardResponse = types.Response(lib.DeleteCardResponse)
var DeleteAllCardsResponse = types.Response(lib.DeleteAllCardsResponse)
var GetEventResponse = types.Response(lib.GetEventResponse)

// var SetIPv4Response = types.Response{
// 	Message: libx.Message{
// 		Name:    "set IPv4 response",
// 		MsgType: 0x96,
// 		Fields: []libx.Field{
// 			{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
// 			{"ok", "bool", 8, "controller IPv4 network configured"},
// 		},
// 	},
// 	Tests: []libx.ResponseTest{
// 		{
// 			Name: "set-IPv4",
// 			Response: []byte{
// 				0x17, 0x96, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
// 				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
// 				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
// 				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
// 			},
// 			Expected: []libx.Value{
// 				{
// 					Name:  "controller",
// 					Type:  "uint32",
// 					Value: 405419896,
// 				},
// 				{
// 					Name:  "ok",
// 					Type:  "bool",
// 					Value: true,
// 				},
// 			},
// 		},
// 	},
// }

// var SetDoorPasscodesResponse = types.Response{
// 	Message: libx.Message{
// 		Name:    "set door passcodes response",
// 		MsgType: 0x8c,
// 		Fields: []libx.Field{
// 			{"controller", "uint32", 4, "controller serial number, e.g. 405419896"},
// 			{"ok", "bool", 8, "set-door-passcodes succeeded/failed"},
// 		},
// 	},
// 	Tests: []libx.ResponseTest{
// 		{
// 			Name: "set-door-passcodes",
// 			Response: []byte{
// 				0x17, 0x8c, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
// 				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
// 				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
// 				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
// 			},
// 			Expected: []libx.Value{
// 				{
// 					Name:  "controller",
// 					Type:  "uint32",
// 					Value: 405419896,
// 				},
// 				{
// 					Name:  "ok",
// 					Type:  "bool",
// 					Value: true,
// 				},
// 			},
// 		},
// 	},
// }
