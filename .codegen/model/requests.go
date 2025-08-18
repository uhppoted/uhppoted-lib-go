package model

import (
	lib "github.com/uhppoted/uhppoted-codegen/model"

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
	DeleteAllCardsRequest,
	GetEventRequest,
	GetEventIndexRequest,
}

var GetControllerRequest = types.Request(lib.GetControllerRequest)
var SetIPv4Request = types.Request(lib.SetIPv4Request)
var GetTimeRequest = types.Request(lib.GetTimeRequest)
var SetTimeRequest = types.Request(lib.SetTimeRequest)
var GetListenerRequest = types.Request(lib.GetListenerRequest)
var SetListenerRequest = types.Request(lib.SetListenerAddrPortRequest)
var GetDoorRequest = types.Request(lib.GetDoorRequest)
var SetDoorRequest = types.Request(lib.SetDoorRequest)
var SetDoorPasscodesRequest = types.Request(lib.SetDoorPasscodesRequest)
var OpenDoorRequest = types.Request(lib.OpenDoorRequest)
var GetStatusRequest = types.Request(lib.GetStatusRequest)
var GetCardsRequest = types.Request(lib.GetCardsRequest)
var GetCardRequest = types.Request(lib.GetCardRequest)
var GetCardAtIndexRequest = types.Request(lib.GetCardAtIndexRequest)
var PutCardRequest = types.Request(lib.PutCardRequest)
var DeleteCardRequest = types.Request(lib.DeleteCardRequest)
var DeleteAllCardsRequest = types.Request(lib.DeleteAllCardsRequest)
var GetEventRequest = types.Request(lib.GetEventRequest)
var GetEventIndexRequest = types.Request(lib.GetEventIndexRequest)
