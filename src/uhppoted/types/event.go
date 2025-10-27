package types

type Event struct {
	Index         uint32    `json:"index"`
	Event         EventType `json:"event"`
	AccessGranted bool      `json:"granted"`
	Door          uint8     `json:"door"`
	Direction     Direction `json:"direction"`
	Card          uint32    `json:"card"`
	Timestamp     DateTime  `json:"timestamp"`
	Reason        Reason    `json:"reason"`
}

type EventType uint8

const (
	EventUnknown     EventType = 0
	EventSwipe       EventType = 1
	EventDoor        EventType = 2
	EventAlarm       EventType = 3
	EventOverwritten EventType = 255
)

func (e EventType) String() string {
	switch e {
	case EventSwipe:
		return "card swipe"
	case EventDoor:
		return "door"
	case EventAlarm:
		return "alarm"
	case EventOverwritten:
		return "overwritten"
	default:
		return "unknown"
	}
}

type Direction uint8

const (
	In  Direction = 1
	Out Direction = 2
)

func (d Direction) String() string {
	switch d {
	case In:
		return "in"
	case Out:
		return "out"
	default:
		return "unknown"
	}
}

type Reason uint8

const (
	SwipeOk                        Reason = 1
	Reserved1                      Reason = 2
	Reserved2                      Reason = 3
	Reserved3                      Reason = 4
	SwipeDeniedPCControl           Reason = 5
	SwipeDeniedNoAccess            Reason = 6
	SwipeDeniedPassword            Reason = 7
	SwipeDeniedAntiPassback        Reason = 8
	SwipeDeniedMoreCards           Reason = 9
	SwipeDeniedFirstCardOpen       Reason = 10
	SwipeDeniedDoorNormallyClosed  Reason = 11
	SwipeDeniedInterLock           Reason = 12
	SwipeDeniedTimeProfile         Reason = 13
	Reserved4                      Reason = 14
	SwipeDeniedTimezone            Reason = 15
	Reserved5                      Reason = 16
	Reserved6                      Reason = 17
	SwipeDenied                    Reason = 18
	Reserved7                      Reason = 19
	PushbuttonOk                   Reason = 20
	Reserved8                      Reason = 21
	Reserved9                      Reason = 22
	DoorOpen                       Reason = 23
	DoorClosed                     Reason = 24
	DoorOpenPasscode               Reason = 25
	Reserved10                     Reason = 26
	Reserved11                     Reason = 27
	ControllerPowerOn              Reason = 28
	ControllerReset                Reason = 29
	PushbuttonDeniedDisabledByTask Reason = 30
	PushButtonDeniedForcedLock     Reason = 31
	PushButtonDeniedOffline        Reason = 32
	PushButtonDeniedInterLock      Reason = 33
	Threat                         Reason = 34
	Reserved12                     Reason = 35
	Reserved13                     Reason = 36
	OpenTooLong                    Reason = 37
	ForcedOpen                     Reason = 38
	Fire                           Reason = 39
	ForcedClose                    Reason = 40
	TamperDetect                   Reason = 41
	Zone24x7                       Reason = 42
	EmergencyCall                  Reason = 43
	RemoteOpenDoor                 Reason = 44
	RemoteOpenDoorUSB              Reason = 45
)

func (r Reason) String() string {
	switch r {
	case SwipeOk:
		return "swipe"
	case Reserved1:
		return "reserved" // swipe open
	case Reserved2:
		return "reserved" // swipe closed
	case Reserved3:
		return "reserved"
	case SwipeDeniedPCControl:
		return "swipe:denied (system)"
	case SwipeDeniedNoAccess:
		return "no access rights"
	case SwipeDeniedPassword:
		return "incorrect PIN"
	case SwipeDeniedAntiPassback:
		return "anti-passback"
	case SwipeDeniedMoreCards:
		return "more cards"
	case SwipeDeniedFirstCardOpen:
		return "first card open"
	case SwipeDeniedDoorNormallyClosed:
		return "door normally closed"
	case SwipeDeniedInterLock:
		return "interlock"
	case SwipeDeniedTimeProfile:
		return "not allowed in time period"
	case Reserved4:
		return "reserved"
	case SwipeDeniedTimezone:
		return "invalid timezone"
	case Reserved5:
		return "reserved"
	case Reserved6:
		return "reserved"
	case SwipeDenied:
		return "access denied"
	case Reserved7:
		return "reserved"
	case PushbuttonOk:
		return "pushbutton ok"
	case Reserved8:
		return "reserved"
	case Reserved9:
		return "reserved"
	case DoorOpen:
		return "door opened"
	case DoorClosed:
		return "door closed"
	case DoorOpenPasscode:
		return "door opened (supervisor password)"
	case Reserved10:
		return "reserved"
	case Reserved11:
		return "reserved"
	case ControllerPowerOn:
		return "controller power on"
	case ControllerReset:
		return "controller reset"
	case PushbuttonDeniedDisabledByTask:
		return "pushbutton denied (disabled)"
	case PushButtonDeniedForcedLock:
		return "pushbutton denied (lock forced)"
	case PushButtonDeniedOffline:
		return "pushbutton denied (offline)"
	case PushButtonDeniedInterLock:
		return "pushbutton denied (interlock)"
	case Threat:
		return "threat"
	case Reserved12:
		return "reserved"
	case Reserved13:
		return "reserved"
	case OpenTooLong:
		return "door open too long"
	case ForcedOpen:
		return "door forced open"
	case Fire:
		return "fire"
	case ForcedClose:
		return "door forced closed"
	case TamperDetect:
		return "tamper detect"
	case Zone24x7:
		return "24x7 zone"
	case EmergencyCall:
		return "emergency"
	case RemoteOpenDoor:
		return "remote open door"
	case RemoteOpenDoorUSB:
		return "remote open door (USB)"
	default:
		return "unknown"
	}
}

// Reason Codes
// 1    01  SwipePass   Swipe
// 2    02  (Reserved)
// 3    03  (Reserved)
// 4    04  (Reserved)
// 5    05  SwipeNOPass Denied Access: PC Control
// 6    06  SwipeNOPass Denied Access: No PRIVILEGE
// 7    07  SwipeNOPass Denied Access: Wrong PASSWORD
// 8    08  SwipeNOPass Denied Access: AntiBack
// 9    09  SwipeNOPass Denied Access: More Cards
// 10   0A  SwipeNOPass Denied Access: First Card Open
// 11   0B  SwipeNOPass Denied Access: Door Set NC
// 12   0C  SwipeNOPass Denied Access: InterLock
// 13   0D  SwipeNOPass Denied Access: Limited Times
// 14   0E  (Reserved)
// 15   0F  SwipeNOPass Denied Access: Invalid Timezone
// 16   10  (Reserved)
// 17   11  (Reserved)
// 18   12  SwipeNOPass Denied Access
// 19   13  (Reserved)
// 20   14  ValidEvent  Push Button
// 21   15  (Reserved)
// 22   16  (Reserved)
// 23   17  ValidEvent  Door Open
// 24   18  ValidEvent  Door Closed
// 25   19  ValidEvent  Super Password Open Door
// 26   1A  (Reserved)
// 27   1B  (Reserved)
// 28   1C  Warn    Controller Power On
// 29   1D  Warn    Controller Reset
// 30   1E  (Reserved | Pushbutton disabled by task)
// 31   1F  Warn    Push Button Invalid: Forced Lock
// 32   20  Warn    Push Button Invalid: Not On Line
// 33   21  Warn    Push Button Invalid: InterLock
// 34   22  Warn    Threat
// 35   23  (Reserved)
// 36   24  (Reserved)
// 37   25  Warn    Open too long
// 38   26  Warn    Forced Open
// 39   27  Warn    Fire
// 40   28  Warn    Forced Close
// 41   29  Warn    Guard Against Theft
// 42   2A  Warn    7*24Hour Zone
// 43   2B  Warn    Emergency Call
// 44   2C  RemoteOpen  Remote Open Door
// 45   2D  RemoteOpen  Remote Open Door By USB Reader
// )
