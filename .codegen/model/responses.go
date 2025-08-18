package model

import (
	lib "github.com/uhppoted/uhppoted-codegen/model"

	"github.com/uhppoted/uhppoted-codegen/model/types"
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

var GetControllerResponse = lib.GetControllerResponse
var SetIPv4Response = lib.SetIPv4Response
var GetTimeResponse = lib.GetTimeResponse
var SetTimeResponse = lib.SetTimeResponse
var GetListenerResponse = lib.GetListenerAddrPortResponse
var SetListenerResponse = lib.SetListenerResponse
var GetDoorResponse = lib.GetDoorResponse
var SetDoorResponse = lib.SetDoorResponse
var SetDoorPasscodesResponse = lib.SetDoorPasscodesResponse
var OpenDoorResponse = lib.OpenDoorResponse
var GetStatusResponse = lib.GetStatusResponse
var GetCardsResponse = lib.GetCardsResponse
var GetCardResponse = lib.GetCardResponse
var GetCardAtIndexResponse = lib.GetCardAtIndexResponse
var PutCardResponse = lib.PutCardResponse
var DeleteCardResponse = lib.DeleteCardResponse
var DeleteAllCardsResponse = lib.DeleteAllCardsResponse
var GetEventResponse = lib.GetEventResponse
var GetEventIndexResponse = lib.GetEventIndexResponse
