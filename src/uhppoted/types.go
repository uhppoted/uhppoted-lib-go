package uhppoted

import (
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/entities"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"
)

// re-exported response types
type (
	GetControllerResponse            = responses.GetControllerResponse
	SetIPv4Response                  = responses.SetIPv4Response
	GetTimeResponse                  = responses.GetTimeResponse
	SetTimeResponse                  = responses.SetTimeResponse
	GetListenerResponse              = responses.GetListenerResponse
	SetListenerResponse              = responses.SetListenerResponse
	GetListenerAddrPortResponse      = responses.GetListenerAddrPortResponse
	SetListenerAddrPortResponse      = responses.SetListenerAddrPortResponse
	GetDoorResponse                  = responses.GetDoorResponse
	SetDoorResponse                  = responses.SetDoorResponse
	SetDoorPasscodesResponse         = responses.SetDoorPasscodesResponse
	OpenDoorResponse                 = responses.OpenDoorResponse
	GetStatusResponse                = responses.GetStatusResponse
	GetCardsResponse                 = responses.GetCardsResponse
	GetCardResponse                  = responses.GetCardResponse
	GetCardAtIndexResponse           = responses.GetCardAtIndexResponse
	PutCardResponse                  = responses.PutCardResponse
	DeleteCardResponse               = responses.DeleteCardResponse
	DeleteAllCardsResponse           = responses.DeleteAllCardsResponse
	GetEventResponse                 = responses.GetEventResponse
	GetEventIndexResponse            = responses.GetEventIndexResponse
	SetEventIndexResponse            = responses.SetEventIndexResponse
	RecordSpecialEventsResponse      = responses.RecordSpecialEventsResponse
	GetTimeProfileResponse           = responses.GetTimeProfileResponse
	SetTimeProfileResponse           = responses.SetTimeProfileResponse
	ClearTimeProfilesResponse        = responses.ClearTimeProfilesResponse
	AddTaskResponse                  = responses.AddTaskResponse
	RefreshTaskListResponse          = responses.RefreshTaskListResponse
	ClearTaskListResponse            = responses.ClearTaskListResponse
	SetPCControlResponse             = responses.SetPCControlResponse
	SetInterlockResponse             = responses.SetInterlockResponse
	ActivateKeypadsResponse          = responses.ActivateKeypadsResponse
	GetAntiPassbackResponse          = responses.GetAntiPassbackResponse
	SetAntiPassbackResponse          = responses.SetAntiPassbackResponse
	RestoreDefaultParametersResponse = responses.RestoreDefaultParametersResponse
	ListenerEvent                    = responses.ListenerEvent
)

// re-exported entities
type (
	Card        = entities.Card
	TimeProfile = entities.TimeProfile
	Task        = entities.Task
	Event       = entities.Event
	Status      = entities.Status
	DateTime    = entities.DateTime
	Date        = entities.Date
	Time        = entities.Time
	HHmm        = entities.HHmm
	DoorMode    = entities.DoorMode
	TaskType    = entities.TaskType
	Interlock   = entities.Interlock
	Weekdays    = entities.Weekdays
	TimeSegment = entities.TimeSegment
)

// re-exported door modes
var (
	NormallyOpen   = entities.NormallyOpen
	NormallyClosed = entities.NormallyClosed
	Controlled     = entities.Controlled
)

// re-exported tasks
var (
	ControlDoor          = entities.ControlDoor
	UnlockDoor           = entities.UnlockDoor
	LockDoor             = entities.LockDoor
	DisableTimeProfiles  = entities.DisableTimeProfiles
	EnableTimeProfiles   = entities.EnableTimeProfiles
	EnableCardNoPassword = entities.EnableCardNoPassword
	EnableCardInPassword = entities.EnableCardInPassword
	EnableCardPassword   = entities.EnableCardPassword
	EnableMoreCards      = entities.EnableMoreCards
	DisableMoreCards     = entities.DisableMoreCards
	TriggerOnce          = entities.TriggerOnce
	DisablePushbutton    = entities.DisablePushbutton
	EnablePushbutton     = entities.EnablePushbutton
)

// re-exported interlocks
var (
	NoInterlock    = entities.NoInterlock
	Interlock12    = entities.Interlock12
	Interlock34    = entities.Interlock34
	Interlock12_34 = entities.Interlock12_34
	Interlock123   = entities.Interlock123
	Interlock1234  = entities.Interlock1234
)

// re-exported event types
var (
	EventUnknown     = entities.EventUnknown
	EventSwipe       = entities.EventSwipe
	EventDoor        = entities.EventDoor
	EventAlarm       = entities.EventAlarm
	EventOverwritten = entities.EventOverwritten
)
