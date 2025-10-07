// generated code - ** DO NOT EDIT **

package responses

import (
	"net/netip"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/entities"
)

// Container struct for the response returned from an access controller when retrieving the
// network configuration, firmware version and firmware release date.
type GetControllerResponse struct {
	Controller uint32        `json:"controller"`
	IpAddress  netip.Addr    `json:"ip-address"`
	SubnetMask netip.Addr    `json:"netmask"`
	Gateway    netip.Addr    `json:"gateway"`
	MACAddress string        `json:"MAC"`
	Version    string        `json:"version"`
	Date       entities.Date `json:"date"`
}

// SetIPv4Response is a synthesized response provided to simplify code generation. The controller
// does not return a response to a 'set-IPv4' request.
type SetIPv4Response struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned from a controller when
// retrieving the current runtime status.
type GetStatusResponse struct {
	Controller         uint32            `json:"controller"`
	SystemDate         entities.Date     `json:"system-date"`
	SystemTime         entities.Time     `json:"system-time"`
	Door1Open          bool              `json:"door-1-open"`
	Door2Open          bool              `json:"door-2-open"`
	Door3Open          bool              `json:"door-3-open"`
	Door4Open          bool              `json:"door-4-open"`
	Door1Button        bool              `json:"door-1-button"`
	Door2Button        bool              `json:"door-2-button"`
	Door3Button        bool              `json:"door-3-button"`
	Door4Button        bool              `json:"door-4-button"`
	Relays             uint8             `json:"relays"`
	Inputs             uint8             `json:"alarm-inputs"`
	SystemError        uint8             `json:"system-error"`
	SpecialInfo        uint8             `json:"special-info"`
	EventIndex         uint32            `json:"event-index"`
	EventType          uint8             `json:"event-type"`
	EventAccessGranted bool              `json:"event-granted"`
	EventDoor          uint8             `json:"event-door"`
	EventDirection     uint8             `json:"event-direction"`
	EventCard          uint32            `json:"event-card"`
	EventTimestamp     entities.DateTime `json:"event-timestamp"`
	EventReason        uint8             `json:"event-reason"`
	SequenceNo         uint32            `json:"sequence-no"`
}

// Container struct for the response returned by a controller when retrieving the system date/time.
type GetTimeResponse struct {
	Controller uint32            `json:"controller"`
	DateTime   entities.DateTime `json:"date-time"`
}

// Container struct for the response returned by a controller after setting the system date/time.
type SetTimeResponse struct {
	Controller uint32            `json:"controller"`
	DateTime   entities.DateTime `json:"date-time"`
}

// Container struct for the response returned by a controller when retrieving
// the configured event listener IPv4 address and port.
type GetListenerResponse struct {
	Controller uint32     `json:"controller"`
	Address    netip.Addr `json:"address"`
	Port       uint16     `json:"port"`
	Interval   uint8      `json:"interval"`
}

// Container struct for the response returned by a controller when setting
// the event listener IPv4 address and port.
type SetListenerResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned by a controller when retrieving
// the configured event listener IPv4 address and port.
type GetListenerAddrPortResponse struct {
	Controller uint32         `json:"controller"`
	Listener   netip.AddrPort `json:"listener"`
	Interval   uint8          `json:"interval"`
}

// Container struct for the response returned by a controller when setting
// the event listener IPv4 address and port.
type SetListenerAddrPortResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned by a controller to a request
// for door configuration information.
type GetDoorResponse struct {
	Controller uint32            `json:"controller"`
	Door       uint8             `json:"door"`
	Mode       entities.DoorMode `json:"mode"`
	Delay      uint8             `json:"delay"`
}

// Container struct for the response returned by a controller after updating
// a door configuration.
type SetDoorResponse struct {
	Controller uint32 `json:"controller"`
	Door       uint8  `json:"door"`
	Mode       uint8  `json:"mode"`
	Delay      uint8  `json:"delay"`
}

// Container struct for the response returned by a controller after setting the
// supervisor passcodes for a door.
type SetDoorPasscodesResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned by a controller to a request
// to unlock a door.
type OpenDoorResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned from a controller when retrieving the number of
// cards stored on the controller.
type GetCardsResponse struct {
	Controller uint32 `json:"controller"`
	Cards      uint32 `json:"cards"`
}

// Container struct for the response returned from a controller when retrieving
// a card record stored on the controller.
type GetCardResponse struct {
	Controller uint32        `json:"controller"`
	Card       uint32        `json:"card"`
	StartDate  entities.Date `json:"start-date"`
	EndDate    entities.Date `json:"end-date"`
	Door1      uint8         `json:"door-1"`
	Door2      uint8         `json:"door-2"`
	Door3      uint8         `json:"door-3"`
	Door4      uint8         `json:"door-4"`
	PIN        uint32        `json:"PIN"`
}

// Container struct for the response returned from a controller when retrieving
// the card record stored at an index on the controller.
type GetCardAtIndexResponse struct {
	Controller uint32        `json:"controller"`
	Card       uint32        `json:"card"`
	StartDate  entities.Date `json:"start-date"`
	EndDate    entities.Date `json:"end-date"`
	Door1      uint8         `json:"door-1"`
	Door2      uint8         `json:"door-2"`
	Door3      uint8         `json:"door-3"`
	Door4      uint8         `json:"door-4"`
	PIN        uint32        `json:"PIN"`
}

// Container struct for the response returned by a controller after adding or updating a
// controller card record.
type PutCardResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned by a controller after deleting a single
// stored card record.
type DeleteCardResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned by a controller after deleting all stored card records.
type DeleteAllCardsResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned from a controller for an event record request.
//
// Event types:
// - 0:   unknown
// - 1:   card
// - 2:   door
// - 3:   alarm
// - 4:   system
// - 255: overwritten
//
// Direction:
// - 1: in
// - 2: out
//
// Reasons:
// 0:      unknown
// 1:      card ok
// 5:      card denied (PC control)
// 6:      card denied (no access)
// 7:      card denied (invalid password)
// 8:      card denied (anti-passback)
// 9:      card denied (more cards)
// 10:     card denied (first card open)
// 11:     card denied (door normally closed)
// 12:     card denied (door interlock)
// 13:     card denied (limited times)
// 15:     card denied (invalid timezone)
// 18:     card denied
// 20:     door pushbutton
// 23:     door open
// 24:     door closed
// 25:     door supervisor password open
// 28:     controller power on
// 29:     controller reset
// 30:     pushbutton denied (disabled by task)
// 31:     pushbutton denied (forced lock)
// 32:     pushbutton denied (not online)
// 33:     pushbutton denied (door interlock
// 34:     alarm (threat)
// 37:     alarm (open too long)
// 38:     alarm (forced open)
// 39:     alarm (fire)
// 40:     alarm (forced close)
// 41:     alarm (tamper detect)
// 42:     alarm (24x7 zone)
// 43:     alarm (emergency call)
// 44:     remote open door
// 45:     remote open door (USB reader)
type GetEventResponse struct {
	Controller    uint32            `json:"controller"`
	Index         uint32            `json:"index"`
	EventType     uint8             `json:"event-type"`
	AccessGranted bool              `json:"granted"`
	Door          uint8             `json:"door"`
	Direction     uint8             `json:"direction"`
	Card          uint32            `json:"card"`
	Timestamp     entities.DateTime `json:"timestamp"`
	Reason        uint8             `json:"reason"`
}

type GetEventIndexResponse struct {
	Controller uint32 `json:"controller"`
	Index      uint32 `json:"index"`
}

// Container struct for the response returned from an access controller when setting the
// downloaded event index.
type SetEventIndexResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned from an access controller when enabling or
// disabling events for door opened, door closed and button pressed.
type RecordSpecialEventsResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned from an access controller when retrieving
// an access time profile.
type GetTimeProfileResponse struct {
	Controller    uint32        `json:"controller"`
	Profile       uint8         `json:"profile"`
	StartDate     entities.Date `json:"start-date"`
	EndDate       entities.Date `json:"end-date"`
	Monday        bool          `json:"monday"`
	Tuesday       bool          `json:"tuesday"`
	Wednesday     bool          `json:"wednesday"`
	Thursday      bool          `json:"thursday"`
	Friday        bool          `json:"friday"`
	Saturday      bool          `json:"saturday"`
	Sunday        bool          `json:"sunday"`
	Segment1Start entities.HHmm `json:"segment1-start"`
	Segment1End   entities.HHmm `json:"segment1-end"`
	Segment2Start entities.HHmm `json:"segment2-start"`
	Segment2End   entities.HHmm `json:"segment2-end"`
	Segment3Start entities.HHmm `json:"segment3-start"`
	Segment3End   entities.HHmm `json:"segment3-end"`
	LinkedProfile uint8         `json:"linked-profile"`
}

// Container struct for the response returned from an access controller when adding/updating
// an access time profile.
type SetTimeProfileResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned from an access controller when clearing all
// stored time profiles.
type ClearTimeProfilesResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned from an access controller when creating
// a scheduled task.
type AddTaskResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned from an access controller when updating
// the task list scheduler.
type RefreshTaskListResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned from an access controller when removing
// all scheduled tasks.
type ClearTaskListResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned from an access controller when enabling or
// disabling remote access control.
type SetPCControlResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned from an access controller after setting
// the door interlock mode.
type SetInterlockResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned from an access controller when enabling or
// disabling door reader keypads.
type ActivateKeypadsResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned from to a request for the current
// anti-passback mode:
// - 0: disabled
// - 1: readers 1:2; 3:4 (independently)
// - 2: readers (1,3):(2,4)
// - 3: readers 1:(2,3)
// - 4: readers 1:(2,3,4)
type GetAntiPassbackResponse struct {
	Controller   uint32 `json:"controller"`
	Antipassback uint8  `json:"antipassback"`
}

// Container struct for the response returned by a controller when setting
// the anti-passback mode.
type SetAntiPassbackResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for the response returned by a controller after restoring the manufacturer
// default settings.
type RestoreDefaultParametersResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// Container struct for an event message sent by a controller.
type ListenerEvent struct {
	Controller         uint32            `json:"controller"`
	SystemDate         entities.Date     `json:"system-date"`
	SystemTime         entities.Time     `json:"system-time"`
	Door1Open          bool              `json:"door-1-open"`
	Door2Open          bool              `json:"door-2-open"`
	Door3Open          bool              `json:"door-3-open"`
	Door4Open          bool              `json:"door-4-open"`
	Door1Button        bool              `json:"door-1-button"`
	Door2Button        bool              `json:"door-2-button"`
	Door3Button        bool              `json:"door-3-button"`
	Door4Button        bool              `json:"door-4-button"`
	Relays             uint8             `json:"relays"`
	Inputs             uint8             `json:"alarm-inputs"`
	SystemError        uint8             `json:"system-error"`
	SpecialInfo        uint8             `json:"special-info"`
	EventIndex         uint32            `json:"event-index"`
	EventType          uint8             `json:"event-type"`
	EventAccessGranted bool              `json:"event-granted"`
	EventDoor          uint8             `json:"event-door"`
	EventDirection     uint8             `json:"event-direction"`
	EventCard          uint32            `json:"event-card"`
	EventTimestamp     entities.DateTime `json:"event-timestamp"`
	EventReason        uint8             `json:"event-reason"`
	SequenceNo         uint32            `json:"sequence-no"`
}
