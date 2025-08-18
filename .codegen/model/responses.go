package model

import (
	lib "github.com/uhppoted/uhppoted-codegen/model"

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
	GetEventIndexResponse,
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
var GetEventIndexResponse = types.Response(lib.GetEventIndexResponse)
