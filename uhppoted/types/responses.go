package types

import (
	"time"
)

// // GetTimeResponse is a container struct for the response returned by a controller
// // when retrieving the system date/time.
// type GetTimeResponse struct {
// 	Controller uint32    `json:"controller"`
// 	DateTime   time.Time `json:"datetime"`
// }

// // SetTimeResponse is a container struct for the response returned by a controller
// // when setting the system date/time.
// type SetTimeResponse struct {
// 	Controller uint32    `json:"controller"`
// 	DateTime   time.Time `json:"datetime"`
// }

// // GetDoorResponse is a container struct for the response returned by a controller
// // when retrieving the control mode and unlock delay for a controller door.
// type GetDoorResponse struct {
// 	Controller uint32 `json:"controller"`
// 	Door       uint8  `json:"door"`
// 	Mode       uint8  `json:"mode"` // 1:normally open, 2:normally closed. 3:controlled
// 	Delay      uint8  `json:"delay"`
// }

// // SetDoorResponse is a container struct for the response returned by a controller
// // after setting the control mode and unlock delay for a controller door.
// type SetDoorResponse struct {
// 	Controller uint32 `json:"controller"`
// 	Door       uint8  `json:"door"`
// 	Mode       uint8  `json:"mode"` // 1:normally open, 2:normally closed. 3:controlled
// 	Delay      uint8  `json:"delay"`
// }

// // SetDoorPasscodesResponse is a container struct for the response returned by a controller
// // after setting the passcodes for a controller door.
// type SetDoorPasscodesResponse struct {
// 	Controller uint32 `json:"controller"`
// 	Ok         bool   `json:"ok"`
// }

// // OpenDoorResponse is a container struct for the response returned by a controller
// // after remotely opening a controller door.
// type OpenDoorResponse struct {
// 	Controller uint32 `json:"controller"`
// 	Ok         bool   `json:"ok"`
// }

// GetStatusResponse is a container struct for the response returned from a controller
// when retrieving the current runtime status.
type GetStatusResponse struct {
	Controller         uint32    `json:"controller"`
	SystemDate         time.Time `json:"system-date"`
	SystemTime         time.Time `json:"system-time"`
	Door1Open          bool      `json:"door-1-open"`
	Door2Open          bool      `json:"door-2-open"`
	Door3Open          bool      `json:"door-3-open"`
	Door4Open          bool      `json:"door-4-open"`
	Door1Button        bool      `json:"door-1-button"`
	Door2Button        bool      `json:"door-2-button"`
	Door3Button        bool      `json:"door-3-button"`
	Door4Button        bool      `json:"door-4-button"`
	Relays             uint8     `json:"relays"`
	Inputs             uint8     `json:"alarm-inputs"`
	SystemError        uint8     `json:"system-error"`
	SpecialInfo        uint8     `json:"special-info"`
	EventIndex         uint32    `json:"event-index"`
	EventType          uint8     `json:"event-type"`
	EventAccessGranted bool      `json:"event-granted"`
	EventDoor          uint8     `json:"event-door"`
	EventDirection     uint8     `json:"event-direction"`
	EventCard          uint32    `json:"event-card"`
	EventTimestamp     time.Time `json:"event-timestamp"`
	EventReason        uint8     `json:"event-reason"`
	SequenceNo         uint32    `json:"sequence-no"`
}

// GetCardsResponse is a container struct for the response returned from a controller
// when retrieving the number of cards stored on the controller.
type GetCardsResponse struct {
	Controller uint32 `json:"controller"`
	Cards      uint32 `json:"cards"`
}

// GetCardResponse is a container struct for the response returned from a controller
// when retrieving information for a card stored on the controller.
type GetCardResponse struct {
	Controller uint32    `json:"controller"`
	Card       uint32    `json:"card"`
	StartDate  time.Time `json:"start-date"`
	EndDate    time.Time `json:"end-date"`
	Door1      uint8     `json:"door-1"`
	Door2      uint8     `json:"door-2"`
	Door3      uint8     `json:"door-3"`
	Door4      uint8     `json:"door-4"`
	PIN        uint32    `json:"PIN"`
}

// GetCardAtIndexResponse is a container struct for the response returned from a controller
// when retrieving information for a card stored on the controller.
type GetCardAtIndexResponse struct {
	Controller uint32    `json:"controller"`
	Card       uint32    `json:"card"`
	StartDate  time.Time `json:"start-date"`
	EndDate    time.Time `json:"end-date"`
	Door1      uint8     `json:"door-1"`
	Door2      uint8     `json:"door-2"`
	Door3      uint8     `json:"door-3"`
	Door4      uint8     `json:"door-4"`
	PIN        uint32    `json:"PIN"`
}

// PutCardResponse is a container struct for the response returned by a controller
// after adding or updating a controller card record.
type PutCardResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// DeleteCardResponse is a container struct for the response returned by a controller
// after deleting a controller card record.
type DeleteCardResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// DeleteAllCardsResponse is a container struct for the response returned by a controller
// after deleting all controller card records.
type DeleteAllCardsResponse struct {
	Controller uint32 `json:"controller"`
	Ok         bool   `json:"ok"`
}

// GetEventResponse is a container struct for the response returned from a controller
// when retrieving the record for an event.
type GetEventResponse struct {
	Controller    uint32    `json:"controller"`
	Index         uint32    `json:"index"`
	Timestamp     time.Time `json:"timestamp"`
	EventType     uint8     `json:"event-type"`
	AccessGranted bool      `json:"access-granted"`
	Door          uint8     `json:"door"`
	Direction     uint8     `json:"direction"`
	Card          uint32    `json:"card"`
	Reason        uint8     `json:"reason"`
}

// // GetEventIndexResponse is a container struct for the response returned from a controller
// // when retrieving the downloaded event index.
// type GetEventIndexResponse struct {
// 	Controller uint32 `json:"controller"`
// 	Index      uint32 `json:"index"`
// }

// // SetEventIndexResponse is a container struct for the response returned from a controller
// // when setting the downloaded event index.
// type SetEventIndexResponse struct {
// 	Controller uint32 `json:"controller"`
// 	Ok         bool   `json:"ok"`
// }
