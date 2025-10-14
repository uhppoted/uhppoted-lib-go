package uhppoted

import (
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/entities"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"
)

// re-exported response types
type (
	GetControllerResponse            = responses.GetController
	SetIPv4Response                  = responses.SetIPv4
	GetTimeResponse                  = responses.GetTime
	SetTimeResponse                  = responses.SetTime
	GetListenerResponse              = responses.GetListener
	SetListenerResponse              = responses.SetListener
	GetListenerAddrPortResponse      = responses.GetListenerAddrPort
	SetListenerAddrPortResponse      = responses.SetListenerAddrPort
	GetDoorResponse                  = responses.GetDoor
	SetDoorResponse                  = responses.SetDoor
	SetDoorPasscodesResponse         = responses.SetDoorPasscodes
	OpenDoorResponse                 = responses.OpenDoor
	GetStatusResponse                = responses.GetStatus
	GetCardsResponse                 = responses.GetCards
	GetCardResponse                  = responses.GetCard
	GetCardAtIndexResponse           = responses.GetCardAtIndex
	PutCardResponse                  = responses.PutCard
	DeleteCardResponse               = responses.DeleteCard
	DeleteAllCardsResponse           = responses.DeleteAllCards
	GetEventResponse                 = responses.GetEvent
	GetEventIndexResponse            = responses.GetEventIndex
	SetEventIndexResponse            = responses.SetEventIndex
	RecordSpecialEventsResponse      = responses.RecordSpecialEvents
	GetTimeProfileResponse           = responses.GetTimeProfile
	SetTimeProfileResponse           = responses.SetTimeProfile
	ClearTimeProfilesResponse        = responses.ClearTimeProfiles
	AddTaskResponse                  = responses.AddTask
	RefreshTaskListResponse          = responses.RefreshTaskList
	ClearTaskListResponse            = responses.ClearTaskList
	SetPCControlResponse             = responses.SetPCControl
	SetInterlockResponse             = responses.SetInterlock
	ActivateKeypadsResponse          = responses.ActivateKeypads
	GetAntiPassbackResponse          = responses.GetAntiPassback
	SetAntiPassbackResponse          = responses.SetAntiPassback
	RestoreDefaultParametersResponse = responses.RestoreDefaultParameters
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
	EventType   = entities.EventType
	Direction   = entities.Direction
	Reason      = entities.Reason
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

// re-exported event directions
var (
	In  = entities.In
	Out = entities.Out
)

// re-exported event reasons
var (
	SwipeOk                        = entities.SwipeOk
	SwipeDeniedPCControl           = entities.SwipeDeniedPCControl
	SwipeDeniedNoAccess            = entities.SwipeDeniedNoAccess
	SwipeDeniedPassword            = entities.SwipeDeniedPassword
	SwipeDeniedAntiPassback        = entities.SwipeDeniedAntiPassback
	SwipeDeniedMoreCards           = entities.SwipeDeniedMoreCards
	SwipeDeniedFirstCardOpen       = entities.SwipeDeniedFirstCardOpen
	SwipeDeniedDoorNormallyClosed  = entities.SwipeDeniedDoorNormallyClosed
	SwipeDeniedInterLock           = entities.SwipeDeniedInterLock
	SwipeDeniedTimeProfile         = entities.SwipeDeniedTimeProfile
	SwipeDeniedTimezone            = entities.SwipeDeniedTimezone
	SwipeDenied                    = entities.SwipeDenied
	PushbuttonOk                   = entities.PushbuttonOk
	DoorOpen                       = entities.DoorOpen
	DoorClosed                     = entities.DoorClosed
	DoorOpenPasscode               = entities.DoorOpenPasscode
	ControllerPowerOn              = entities.ControllerPowerOn
	ControllerReset                = entities.ControllerReset
	PushbuttonDeniedDisabledByTask = entities.PushbuttonDeniedDisabledByTask
	PushButtonDeniedForcedLock     = entities.PushButtonDeniedForcedLock
	PushButtonDeniedOffline        = entities.PushButtonDeniedOffline
	PushButtonDeniedInterLock      = entities.PushButtonDeniedInterLock
	Threat                         = entities.Threat
	OpenTooLong                    = entities.OpenTooLong
	ForcedOpen                     = entities.ForcedOpen
	Fire                           = entities.Fire
	ForcedClose                    = entities.ForcedClose
	TamperDetect                   = entities.TamperDetect
	Zone24x7                       = entities.Zone24x7
	EmergencyCall                  = entities.EmergencyCall
	RemoteOpenDoor                 = entities.RemoteOpenDoor
	RemoteOpenDoorUSB              = entities.RemoteOpenDoorUSB
)
