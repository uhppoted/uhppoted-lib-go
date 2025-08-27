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
	GetListenerAddrPortRequest,
	SetListenerAddrPortRequest,
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
	SetEventIndexRequest,
	RecordSpecialEventsRequest,
	GetTimeProfileRequest,
	SetTimeProfileRequest,
	ClearTimeProfilesRequest,
	AddTaskRequest,
}

var GetControllerRequest = lib.GetControllerRequest
var SetIPv4Request = lib.SetIPv4Request
var GetTimeRequest = lib.GetTimeRequest
var SetTimeRequest = lib.SetTimeRequest
var GetListenerRequest = lib.GetListenerRequest
var SetListenerRequest = lib.SetListenerRequest
var GetListenerAddrPortRequest = lib.GetListenerAddrPortRequest
var SetListenerAddrPortRequest = lib.SetListenerAddrPortRequest
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
var SetEventIndexRequest = lib.SetEventIndexRequest
var RecordSpecialEventsRequest = lib.RecordSpecialEventsRequest
var GetTimeProfileRequest = lib.GetTimeProfileRequest
var SetTimeProfileRequest = lib.SetTimeProfileRequest
var ClearTimeProfilesRequest = lib.ClearTimeProfilesRequest
var AddTaskRequest = lib.AddTaskRequest
