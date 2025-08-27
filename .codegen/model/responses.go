package model

import (
	lib "github.com/uhppoted/uhppoted-codegen/model"

	"github.com/uhppoted/uhppoted-codegen/model/types"
)

var Responses = []*types.Response{
	&GetControllerResponse,
	&SetIPv4Response,
	&GetStatusResponse,
	&GetTimeResponse,
	&SetTimeResponse,
	&GetListenerResponse,
	&SetListenerResponse,
	&GetListenerAddrPortResponse,
	&SetListenerAddrPortResponse,
	&GetDoorResponse,
	&SetDoorResponse,
	&SetDoorPasscodesResponse,
	&OpenDoorResponse,
	&GetCardsResponse,
	&GetCardResponse,
	&GetCardAtIndexResponse,
	&PutCardResponse,
	&DeleteCardResponse,
	&DeleteAllCardsResponse,
	&GetEventResponse,
	&GetEventIndexResponse,
	&SetEventIndexResponse,
	&RecordSpecialEventsResponse,
	&GetTimeProfileResponse,
	&SetTimeProfileResponse,
	&ClearTimeProfilesResponse,
	&AddTaskResponse,
}

var GetControllerResponse = lib.GetControllerResponse
var SetIPv4Response = lib.SetIPv4Response
var GetTimeResponse = lib.GetTimeResponse
var SetTimeResponse = lib.SetTimeResponse
var GetListenerResponse = lib.GetListenerResponse
var SetListenerResponse = lib.SetListenerResponse
var GetListenerAddrPortResponse = lib.GetListenerAddrPortResponse
var SetListenerAddrPortResponse = lib.SetListenerAddrPortResponse
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
var SetEventIndexResponse = lib.SetEventIndexResponse
var RecordSpecialEventsResponse = lib.RecordSpecialEventsResponse
var GetTimeProfileResponse = lib.GetTimeProfileResponse
var SetTimeProfileResponse = lib.SetTimeProfileResponse
var ClearTimeProfilesResponse = lib.ClearTimeProfilesResponse
var AddTaskResponse = lib.AddTaskResponse
