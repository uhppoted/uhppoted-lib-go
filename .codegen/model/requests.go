package model

import (
	lib "github.com/uhppoted/uhppoted-codegen/model"
	libx "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/model/types"
)

var Requests = []types.Request{
	GetControllerRequest,
	SetIPv4Request,
	GetStatusRequest,
	GetTimeRequest,
	SetTimeRequest,
	GetListenerRequest,
	SetListenerRequest,
	GetDoorRequest,
	SetDoorRequest,
	SetDoorPasscodesRequest,
	OpenDoorRequest,
	GetCardsRequest,
	GetCardRequest,
	GetCardAtIndexRequest,
	PutCardRequest,
	DeleteCardRequest,
}

var GetControllerRequest = types.Request(lib.GetControllerRequest)
var SetIPv4Request = types.Request(lib.SetIPv4Request)
var GetTimeRequest = types.Request(lib.GetTimeRequest)
var SetTimeRequest = types.Request(lib.SetTimeRequest)

var GetListenerRequest = types.Request(lib.GetListenerRequest)
var SetListenerRequest = types.Request(lib.SetListenerAddrPortRequest)
var GetDoorRequest = types.Request(lib.GetDoorRequest)
var OpenDoorRequest = types.Request(lib.OpenDoorRequest)
var GetStatusRequest = types.Request(lib.GetStatusRequest)
var GetCardsRequest = types.Request(lib.GetCardsRequest)
var GetCardRequest = types.Request(lib.GetCardRequest)
var GetCardAtIndexRequest = types.Request(lib.GetCardAtIndexRequest)
var PutCardRequest = types.Request(lib.PutCardRequest)
var DeleteCardRequest = types.Request(lib.DeleteCardRequest)

var SetDoorRequest = types.Request{
	Message: lib.SetDoorRequest,
	Tests: []libx.RequestTest{
		{
			Name: "set-door",
			Args: []libx.TestArg{
				{Arg: libx.Arg{Name: "controller", Type: "uint32"}, Value: uint32(405419896)},
				{Arg: libx.Arg{Name: "door", Type: "uint8"}, Value: uint8(3)},
				{Arg: libx.Arg{Name: "mode", Type: "uint8"}, Value: uint8(2)},
				{Arg: libx.Arg{Name: "delay", Type: "uint8"}, Value: uint8(17)},
			},
			Expected: []byte{
				0x17, 0x80, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x03, 0x02, 0x11, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
	},
}

var SetDoorPasscodesRequest = types.Request{
	Message: lib.SetDoorPasscodesRequest,
	Tests: []libx.RequestTest{
		{
			Name: "set-door-passcodes",
			Args: []libx.TestArg{
				{Arg: libx.Arg{Name: "controller", Type: "uint32"}, Value: uint32(405419896)},
				{Arg: libx.Arg{Name: "door", Type: "uint8"}, Value: uint8(3)},
				{Arg: libx.Arg{Name: "passcode1", Type: "uint32"}, Value: uint32(12345)},
				{Arg: libx.Arg{Name: "passcode2", Type: "uint32"}, Value: uint32(54321)},
				{Arg: libx.Arg{Name: "passcode3", Type: "uint32"}, Value: uint32(0)},
				{Arg: libx.Arg{Name: "passcode4", Type: "uint32"}, Value: uint32(999999)},
			},
			Expected: []byte{
				0x17, 0x8c, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x03, 0x00, 0x00, 0x00, 0x39, 0x30, 0x00, 0x00,
				0x31, 0xd4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3f, 0x42, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
	},
}
