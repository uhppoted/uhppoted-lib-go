package model

import (
	lib "github.com/uhppoted/uhppoted-codegen/model"
	"github.com/uhppoted/uhppoted-codegen/model/types"
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

var GetControllerRequest = lib.GetControllerRequest
var SetIPv4Request = lib.SetIPv4Request
var GetTimeRequest = lib.GetTimeRequest
var SetTimeRequest = lib.SetTimeRequest
var GetListenerRequest = lib.GetListenerRequest
var SetListenerRequest = lib.SetListenerAddrPortRequest
var GetDoorRequest = lib.GetDoorRequest
var SetDoorRequest = lib.SetDoorRequest
var SetDoorPasscodesRequest = lib.SetDoorPasscodesRequest
var OpenDoorRequest = lib.OpenDoorRequest
var GetStatusRequest = lib.GetStatusRequest
var GetCardsRequest = lib.GetCardsRequest
var GetCardRequest = lib.GetCardRequest
var GetCardAtIndexRequest = lib.GetCardAtIndexRequest
var PutCardRequest = lib.PutCardRequest
var DeleteCardRequest = lib.DeleteCardRequest
var DeleteAllCardsRequest = lib.DeleteAllCardsRequest
var GetEventRequest = lib.GetEventRequest
var GetEventIndexRequest = lib.GetEventIndexRequest
