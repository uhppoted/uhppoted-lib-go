// generated code - ** DO NOT EDIT **

package uhppoted

import (
	"net/netip"
	"time"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/codec/encode"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"
)

// Retrieves the system information for an access controller.
func GetController[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.GetController, error) {
	f := func(id uint32) ([]byte, error) { return encode.GetControllerRequest(id) }

	return exec[T, responses.GetController](u, controller, f, timeout)
}

// Sets the controller IPv4 address, netmask and gateway address.
func SetIPv4[T TController](u Uhppoted, controller T, address netip.Addr, netmask netip.Addr, gateway netip.Addr, timeout time.Duration) (responses.SetIPv4, error) {
	f := func(id uint32) ([]byte, error) { return encode.SetIPv4Request(id, address, netmask, gateway) }

	return exec[T, responses.SetIPv4](u, controller, f, timeout)
}

// Retrieves the access controller system date and time.
func GetTime[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.GetTime, error) {
	f := func(id uint32) ([]byte, error) { return encode.GetTimeRequest(id) }

	return exec[T, responses.GetTime](u, controller, f, timeout)
}

// Sets the access controller system date and time.
func SetTime[T TController, DT TDateTime](u Uhppoted, controller T, datetime DT, timeout time.Duration) (responses.SetTime, error) {
	f := func(id uint32) ([]byte, error) { return encode.SetTimeRequest(id, convert[types.DateTime](datetime)) }

	return exec[T, responses.SetTime](u, controller, f, timeout)
}

// Retrieves the access controller event listener IPv4 address:port and auto-send interval.
func GetListener[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.GetListener, error) {
	f := func(id uint32) ([]byte, error) { return encode.GetListenerRequest(id) }

	return exec[T, responses.GetListener](u, controller, f, timeout)
}

// Sets the access controller event listener IPv4 address:port and auto-send interval.
func SetListener[T TController](u Uhppoted, controller T, address netip.Addr, port uint16, interval uint8, timeout time.Duration) (responses.SetListener, error) {
	f := func(id uint32) ([]byte, error) { return encode.SetListenerRequest(id, address, port, interval) }

	return exec[T, responses.SetListener](u, controller, f, timeout)
}

// Retrieves the access controller event listener IPv4 address:port and auto-send interval.
func GetListenerAddrPort[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.GetListenerAddrPort, error) {
	f := func(id uint32) ([]byte, error) { return encode.GetListenerAddrPortRequest(id) }

	return exec[T, responses.GetListenerAddrPort](u, controller, f, timeout)
}

// Sets the access controller event listener IPv4 address:port and auto-send interval.
func SetListenerAddrPort[T TController](u Uhppoted, controller T, listener netip.AddrPort, interval uint8, timeout time.Duration) (responses.SetListenerAddrPort, error) {
	f := func(id uint32) ([]byte, error) { return encode.SetListenerAddrPortRequest(id, listener, interval) }

	return exec[T, responses.SetListenerAddrPort](u, controller, f, timeout)
}

// Retrieves the control mode and unlock delay time for an access controller door.
func GetDoor[T TController](u Uhppoted, controller T, door uint8, timeout time.Duration) (responses.GetDoor, error) {
	f := func(id uint32) ([]byte, error) { return encode.GetDoorRequest(id, door) }

	return exec[T, responses.GetDoor](u, controller, f, timeout)
}

// Sets the control mode and unlock delay time for an access controller door.
func SetDoor[T TController](u Uhppoted, controller T, door uint8, mode types.DoorMode, delay uint8, timeout time.Duration) (responses.SetDoor, error) {
	f := func(id uint32) ([]byte, error) { return encode.SetDoorRequest(id, door, mode, delay) }

	return exec[T, responses.SetDoor](u, controller, f, timeout)
}

// Sets up to 4 passcodes for a controller door.
func SetDoorPasscodes[T TController](u Uhppoted, controller T, door uint8, passcode1 uint32, passcode2 uint32, passcode3 uint32, passcode4 uint32, timeout time.Duration) (responses.SetDoorPasscodes, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetDoorPasscodesRequest(id, door, passcode1, passcode2, passcode3, passcode4)
	}

	return exec[T, responses.SetDoorPasscodes](u, controller, f, timeout)
}

// Unlocks a door controlled by an access controller.
func OpenDoor[T TController](u Uhppoted, controller T, door uint8, timeout time.Duration) (responses.OpenDoor, error) {
	f := func(id uint32) ([]byte, error) { return encode.OpenDoorRequest(id, door) }

	return exec[T, responses.OpenDoor](u, controller, f, timeout)
}

// Retrieves the system status from an access controller.
func GetStatus[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.GetStatus, error) {
	f := func(id uint32) ([]byte, error) { return encode.GetStatusRequest(id) }

	return exec[T, responses.GetStatus](u, controller, f, timeout)
}

// Retrieves the number of cards stored on an access controller.
func GetCards[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.GetCards, error) {
	f := func(id uint32) ([]byte, error) { return encode.GetCardsRequest(id) }

	return exec[T, responses.GetCards](u, controller, f, timeout)
}

// Retrieves the card information for a given card number.
func GetCard[T TController](u Uhppoted, controller T, cardnumber uint32, timeout time.Duration) (responses.GetCard, error) {
	f := func(id uint32) ([]byte, error) { return encode.GetCardRequest(id, cardnumber) }

	return exec[T, responses.GetCard](u, controller, f, timeout)
}

// Retrieves the card information stored at a given index.
func GetCardAtIndex[T TController](u Uhppoted, controller T, index uint32, timeout time.Duration) (responses.GetCardAtIndex, error) {
	f := func(id uint32) ([]byte, error) { return encode.GetCardAtIndexRequest(id, index) }

	return exec[T, responses.GetCardAtIndex](u, controller, f, timeout)
}

// Creates or updates the card information stored on an access controller.
func PutCard[T TController, D TDate](u Uhppoted, controller T, card uint32, startdate D, enddate D, door1 uint8, door2 uint8, door3 uint8, door4 uint8, PIN uint32, timeout time.Duration) (responses.PutCard, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.PutCardRequest(id, card, convert[types.Date](startdate), convert[types.Date](enddate), door1, door2, door3, door4, PIN)
	}

	return exec[T, responses.PutCard](u, controller, f, timeout)
}

// Removes a card record stored on a controller.
func DeleteCard[T TController](u Uhppoted, controller T, cardnumber uint32, timeout time.Duration) (responses.DeleteCard, error) {
	f := func(id uint32) ([]byte, error) { return encode.DeleteCardRequest(id, cardnumber) }

	return exec[T, responses.DeleteCard](u, controller, f, timeout)
}

// Deletes all card records stored on an access controller.
func DeleteAllCards[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.DeleteAllCards, error) {
	f := func(id uint32) ([]byte, error) { return encode.DeleteCardsRequest(id) }

	return exec[T, responses.DeleteAllCards](u, controller, f, timeout)
}

// Retrieves an event record stored on an access controller.
func GetEvent[T TController](u Uhppoted, controller T, eventindex uint32, timeout time.Duration) (responses.GetEvent, error) {
	f := func(id uint32) ([]byte, error) { return encode.GetEventRequest(id, eventindex) }

	return exec[T, responses.GetEvent](u, controller, f, timeout)
}

// Retrieves the downloaded event index from an access controller.
func GetEventIndex[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.GetEventIndex, error) {
	f := func(id uint32) ([]byte, error) { return encode.GetEventIndexRequest(id) }

	return exec[T, responses.GetEventIndex](u, controller, f, timeout)
}

// Sets the downloaded event index on an access controller.
func SetEventIndex[T TController](u Uhppoted, controller T, eventindex uint32, timeout time.Duration) (responses.SetEventIndex, error) {
	f := func(id uint32) ([]byte, error) { return encode.SetEventIndexRequest(id, eventindex) }

	return exec[T, responses.SetEventIndex](u, controller, f, timeout)
}

// Enables/disables events for door opened, door closed and door button pressed.
func RecordSpecialEvents[T TController](u Uhppoted, controller T, enabled bool, timeout time.Duration) (responses.RecordSpecialEvents, error) {
	f := func(id uint32) ([]byte, error) { return encode.RecordSpecialEventsRequest(id, enabled) }

	return exec[T, responses.RecordSpecialEvents](u, controller, f, timeout)
}

// Retrieves the requested access time profile from a controller.
func GetTimeProfile[T TController](u Uhppoted, controller T, profile uint8, timeout time.Duration) (responses.GetTimeProfile, error) {
	f := func(id uint32) ([]byte, error) { return encode.GetTimeProfileRequest(id, profile) }

	return exec[T, responses.GetTimeProfile](u, controller, f, timeout)
}

// Adds or updates an access time profile stored on a controller.
func SetTimeProfile[T TController, D TDate, H THHmm](u Uhppoted, controller T, profile uint8, startdate D, enddate D, monday bool, tuesday bool, wednesday bool, thursday bool, friday bool, saturday bool, sunday bool, segment1start H, segment1end H, segment2start H, segment2end H, segment3start H, segment3end H, linkedprofileid uint8, timeout time.Duration) (responses.SetTimeProfile, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetTimeProfileRequest(id, profile, convert[types.Date](startdate), convert[types.Date](enddate), monday, tuesday, wednesday, thursday, friday, saturday, sunday, convert[types.HHmm](segment1start), convert[types.HHmm](segment1end), convert[types.HHmm](segment2start), convert[types.HHmm](segment2end), convert[types.HHmm](segment3start), convert[types.HHmm](segment3end), linkedprofileid)
	}

	return exec[T, responses.SetTimeProfile](u, controller, f, timeout)
}

// Clears all access time profiles stored on a controller.
func ClearTimeProfiles[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.ClearTimeProfiles, error) {
	f := func(id uint32) ([]byte, error) { return encode.ClearTimeProfilesRequest(id) }

	return exec[T, responses.ClearTimeProfiles](u, controller, f, timeout)
}

// Creates a scheduled task.
//
// Task types:
//   - 0:  control door
//   - 1:  unlock door
//   - 2:  lock door
//   - 3:  disable time profiles
//   - 4:  enable time profiles
//   - 5:  enable card, no password
//   - 6:  enable card+IN password
//   - 7:  enable card+password
//   - 8:  enable more cards
//   - 9:  disable more cards
//   - 10: trigger once
//   - 11: disable pushbutton
//   - 12: enable pushbutton
func AddTask[T TController, D TDate, H THHmm](u Uhppoted, controller T, task types.TaskType, startdate D, enddate D, monday bool, tuesday bool, wednesday bool, thursday bool, friday bool, saturday bool, sunday bool, starttime H, door uint8, morecards uint8, timeout time.Duration) (responses.AddTask, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.AddTaskRequest(id, task, convert[types.Date](startdate), convert[types.Date](enddate), monday, tuesday, wednesday, thursday, friday, saturday, sunday, convert[types.HHmm](starttime), door, morecards)
	}

	return exec[T, responses.AddTask](u, controller, f, timeout)
}

// Updates scheduler with newly created scheduled tasks.
func RefreshTaskList[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.RefreshTaskList, error) {
	f := func(id uint32) ([]byte, error) { return encode.RefreshTaskListRequest(id) }

	return exec[T, responses.RefreshTaskList](u, controller, f, timeout)
}

// Removes all scheduled tasks.
func ClearTaskList[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.ClearTaskList, error) {
	f := func(id uint32) ([]byte, error) { return encode.ClearTasklistRequest(id) }

	return exec[T, responses.ClearTaskList](u, controller, f, timeout)
}

// Enables remote access control. Remote access control will remain in effect provided the controller
// receives a message from the host at least once every 30 seconds.
func SetPCControl[T TController](u Uhppoted, controller T, enabled bool, timeout time.Duration) (responses.SetPCControl, error) {
	f := func(id uint32) ([]byte, error) { return encode.SetPCControlRequest(id, enabled) }

	return exec[T, responses.SetPCControl](u, controller, f, timeout)
}

// Sets the door interlock mode on an access controller.
//
// The following interlock modes are supported:
//   - 0: disabled
//   - 1: doors 1&2
//   - 2: doors 3&4
//   - 3: doors 1&2, doors 3&4
//   - 4: doors 1,2&3
//   - 8: doors 1,2,3&4
func SetInterlock[T TController](u Uhppoted, controller T, interlock types.Interlock, timeout time.Duration) (responses.SetInterlock, error) {
	f := func(id uint32) ([]byte, error) { return encode.SetInterlockRequest(id, interlock) }

	return exec[T, responses.SetInterlock](u, controller, f, timeout)
}

// Enables/disables door keypad readers.
func ActivateKeypads[T TController](u Uhppoted, controller T, reader1 bool, reader2 bool, reader3 bool, reader4 bool, timeout time.Duration) (responses.ActivateKeypads, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.ActivateKeypadsRequest(id, reader1, reader2, reader3, reader4)
	}

	return exec[T, responses.ActivateKeypads](u, controller, f, timeout)
}

// Retrieves the anti-passback mode for an access controller. The anti-passback mode
// will be one of the following:
//   - 0: disabled
//   - 1: readers 1:2; 3:4 (independently)
//   - 2: readers (1,3):(2,4)
//   - 3: readers 1:(2,3)
//   - 4: readers 1:(2,3,4)
func GetAntiPassback[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.GetAntiPassback, error) {
	f := func(id uint32) ([]byte, error) { return encode.GetAntipassbackRequest(id) }

	return exec[T, responses.GetAntiPassback](u, controller, f, timeout)
}

// Sets the access controller anti-passback mode.
//
// The following modes are supported:
//   - 0: disabled
//   - 1: doors 1&2, doors 3&4
//   - 2: doors 1&3, doors 2&4
//   - 3: door 1 & doors 2,3
//   - 4: door 1 & doors 1,2,3
func SetAntiPassback[T TController](u Uhppoted, controller T, antipassback uint8, timeout time.Duration) (responses.SetAntiPassback, error) {
	f := func(id uint32) ([]byte, error) { return encode.SetAntipassbackRequest(id, antipassback) }

	return exec[T, responses.SetAntiPassback](u, controller, f, timeout)
}

// Restores the controller configuration to the manufacturer defaults.
func RestoreDefaultParameters[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.RestoreDefaultParameters, error) {
	f := func(id uint32) ([]byte, error) { return encode.RestoreDefaultParametersRequest(id) }

	return exec[T, responses.RestoreDefaultParameters](u, controller, f, timeout)
}
