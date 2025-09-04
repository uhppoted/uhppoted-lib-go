package uhppoted

import (
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/encode"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/responses"
	"net/netip"
	"time"
)

// Retrieves the system information for an access controller.
func GetController[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.GetControllerResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetControllerRequest(id)
	}

	return exec[T, responses.GetControllerResponse](u, controller, f, timeout)
}

// Sets the controller IPv4 address, netmask and gateway address.
func SetIPv4[T TController](u Uhppoted, controller T, address netip.Addr, netmask netip.Addr, gateway netip.Addr, timeout time.Duration) (responses.SetIPv4Response, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetIPv4Request(id, address, netmask, gateway)
	}

	return exec[T, responses.SetIPv4Response](u, controller, f, timeout)
}

// Retrieves the access controller system date and time.
func GetTime[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.GetTimeResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetTimeRequest(id)
	}

	return exec[T, responses.GetTimeResponse](u, controller, f, timeout)
}

// Sets the access controller system date and time.
func SetTime[T TController](u Uhppoted, controller T, datetime time.Time, timeout time.Duration) (responses.SetTimeResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetTimeRequest(id, datetime)
	}

	return exec[T, responses.SetTimeResponse](u, controller, f, timeout)
}

// Retrieves the access controller event listener IPv4 address:port and auto-send interval.
func GetListener[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.GetListenerResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetListenerRequest(id)
	}

	return exec[T, responses.GetListenerResponse](u, controller, f, timeout)
}

// Sets the access controller event listener IPv4 address:port and auto-send interval.
func SetListener[T TController](u Uhppoted, controller T, address netip.Addr, port uint16, interval uint8, timeout time.Duration) (responses.SetListenerResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetListenerRequest(id, address, port, interval)
	}

	return exec[T, responses.SetListenerResponse](u, controller, f, timeout)
}

// Retrieves the control mode and unlock delay time for an access controller door.
func GetDoor[T TController](u Uhppoted, controller T, door uint8, timeout time.Duration) (responses.GetDoorResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetDoorRequest(id, door)
	}

	return exec[T, responses.GetDoorResponse](u, controller, f, timeout)
}

// Sets the control mode and unlock delay time for an access controller door.
func SetDoor[T TController](u Uhppoted, controller T, door uint8, mode uint8, delay uint8, timeout time.Duration) (responses.SetDoorResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetDoorRequest(id, door, mode, delay)
	}

	return exec[T, responses.SetDoorResponse](u, controller, f, timeout)
}

// Sets up to 4 passcodes for a controller door.
func SetDoorPasscodes[T TController](u Uhppoted, controller T, door uint8, passcode1 uint32, passcode2 uint32, passcode3 uint32, passcode4 uint32, timeout time.Duration) (responses.SetDoorPasscodesResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetDoorPasscodesRequest(id, door, passcode1, passcode2, passcode3, passcode4)
	}

	return exec[T, responses.SetDoorPasscodesResponse](u, controller, f, timeout)
}

// Unlocks a door controlled by an access controller.
func OpenDoor[T TController](u Uhppoted, controller T, door uint8, timeout time.Duration) (responses.OpenDoorResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.OpenDoorRequest(id, door)
	}

	return exec[T, responses.OpenDoorResponse](u, controller, f, timeout)
}

// Retrieves the system status from an access controller.
func GetStatus[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.GetStatusResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetStatusRequest(id)
	}

	return exec[T, responses.GetStatusResponse](u, controller, f, timeout)
}

// Retrieves the number of cards stored on an access controller.
func GetCards[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.GetCardsResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetCardsRequest(id)
	}

	return exec[T, responses.GetCardsResponse](u, controller, f, timeout)
}

// Retrieves the card record for a given card number.
func GetCard[T TController](u Uhppoted, controller T, cardnumber uint32, timeout time.Duration) (responses.GetCardResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetCardRequest(id, cardnumber)
	}

	return exec[T, responses.GetCardResponse](u, controller, f, timeout)
}

// Retrieves the card record stored at a given index.
func GetCardAtIndex[T TController](u Uhppoted, controller T, index uint32, timeout time.Duration) (responses.GetCardAtIndexResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetCardAtIndexRequest(id, index)
	}

	return exec[T, responses.GetCardAtIndexResponse](u, controller, f, timeout)
}

// Creates or updates a card record stored on an access controller.
func PutCard[T TController](u Uhppoted, controller T, card uint32, startdate time.Time, enddate time.Time, door1 uint8, door2 uint8, door3 uint8, door4 uint8, PIN uint32, timeout time.Duration) (responses.PutCardResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.PutCardRequest(id, card, startdate, enddate, door1, door2, door3, door4, PIN)
	}

	return exec[T, responses.PutCardResponse](u, controller, f, timeout)
}

// Removes a card record stored on a controller.
func DeleteCard[T TController](u Uhppoted, controller T, cardnumber uint32, timeout time.Duration) (responses.DeleteCardResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.DeleteCardRequest(id, cardnumber)
	}

	return exec[T, responses.DeleteCardResponse](u, controller, f, timeout)
}

// Deletes all card records stored on an access controller.
func DeleteAllCards[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.DeleteAllCardsResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.DeleteCardsRequest(id)
	}

	return exec[T, responses.DeleteAllCardsResponse](u, controller, f, timeout)
}

// Retrieves an event record stored on an access controller.
func GetEvent[T TController](u Uhppoted, controller T, eventindex uint32, timeout time.Duration) (responses.GetEventResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetEventRequest(id, eventindex)
	}

	return exec[T, responses.GetEventResponse](u, controller, f, timeout)
}

// Retrieves the downloaded event index from an access controller.
func GetEventIndex[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.GetEventIndexResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetEventIndexRequest(id)
	}

	return exec[T, responses.GetEventIndexResponse](u, controller, f, timeout)
}

// Sets the downloaded event index on an access controller.
func SetEventIndex[T TController](u Uhppoted, controller T, eventindex uint32, timeout time.Duration) (responses.SetEventIndexResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetEventIndexRequest(id, eventindex)
	}

	return exec[T, responses.SetEventIndexResponse](u, controller, f, timeout)
}

// Enables/disables events for door opened, door closed and door button pressed.
func RecordSpecialEvents[T TController](u Uhppoted, controller T, enabled bool, timeout time.Duration) (responses.RecordSpecialEventsResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.RecordSpecialEventsRequest(id, enabled)
	}

	return exec[T, responses.RecordSpecialEventsResponse](u, controller, f, timeout)
}

// Retrieves the requested access time profile from a controller.
func GetTimeProfile[T TController](u Uhppoted, controller T, profile uint8, timeout time.Duration) (responses.GetTimeProfileResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetTimeProfileRequest(id, profile)
	}

	return exec[T, responses.GetTimeProfileResponse](u, controller, f, timeout)
}

// Adds or updates an access time profile stored on a controller.
func SetTimeProfile[T TController](u Uhppoted, controller T, profile uint8, startdate time.Time, enddate time.Time, monday bool, tuesday bool, wednesday bool, thursday bool, friday bool, saturday bool, sunday bool, segment1start time.Time, segment1end time.Time, segment2start time.Time, segment2end time.Time, segment3start time.Time, segment3end time.Time, linkedprofileid uint8, timeout time.Duration) (responses.SetTimeProfileResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetTimeProfileRequest(id, profile, startdate, enddate, monday, tuesday, wednesday, thursday, friday, saturday, sunday, segment1start, segment1end, segment2start, segment2end, segment3start, segment3end, linkedprofileid)
	}

	return exec[T, responses.SetTimeProfileResponse](u, controller, f, timeout)
}

// Clears all access time profiles stored on a controller.
func ClearTimeProfiles[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.ClearTimeProfilesResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.ClearTimeProfilesRequest(id)
	}

	return exec[T, responses.ClearTimeProfilesResponse](u, controller, f, timeout)
}

// Creates a scheduled task.
//
// Task types
// 0:  control door
// 1:  unlock door
// 2:  lock door
// 3:  disable time profiles
// 4:  enable time profiles
// 5:  enable card, no password
// 6:  enable card+IN password
// 7:  enable card+password
// 8:  enable more cards
// 9:  disable more cards
// 10: trigger once
// 11: disable pushbutton
// 12: enable pushbutton
func AddTask[T TController](u Uhppoted, controller T, task uint8, startdate time.Time, enddate time.Time, monday bool, tuesday bool, wednesday bool, thursday bool, friday bool, saturday bool, sunday bool, starttime time.Time, door uint8, morecards uint8, timeout time.Duration) (responses.AddTaskResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.AddTaskRequest(id, task, startdate, enddate, monday, tuesday, wednesday, thursday, friday, saturday, sunday, starttime, door, morecards)
	}

	return exec[T, responses.AddTaskResponse](u, controller, f, timeout)
}

// Updates scheduler with newly created scheduled tasks.
func RefreshTaskList[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.RefreshTaskListResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.RefreshTaskListRequest(id)
	}

	return exec[T, responses.RefreshTaskListResponse](u, controller, f, timeout)
}

// Removes all scheduled tasks.
func ClearTaskList[T TController](u Uhppoted, controller T, timeout time.Duration) (responses.ClearTasklistResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.ClearTasklistRequest(id)
	}

	return exec[T, responses.ClearTasklistResponse](u, controller, f, timeout)
}

// Enables remote access control. Remote access control will remain in effect provided the controller
// receives a message from the host at least once every 30 seconds.
func SetPCControl[T TController](u Uhppoted, controller T, enabled bool, timeout time.Duration) (responses.SetPCControlResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetPCControlRequest(id, enabled)
	}

	return exec[T, responses.SetPCControlResponse](u, controller, f, timeout)
}

// Sets the door interlock mode on an access controller.
//
// The following interlock modes are supported:
// - 0: disabled
// - 1: doors 1&2
// - 2: doors 3&4
// - 3: doors 1&2, doors 3&4
// - 4: doors 1,2&3
// - 8: doors 1,2,3&4
func SetInterlock[T TController](u Uhppoted, controller T, interlock uint8, timeout time.Duration) (responses.SetInterlockResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetInterlockRequest(id, interlock)
	}

	return exec[T, responses.SetInterlockResponse](u, controller, f, timeout)
}

// Enables/disables door keypad readers.
func ActivateKeypads[T TController](u Uhppoted, controller T, reader1 bool, reader2 bool, reader3 bool, reader4 bool, timeout time.Duration) (responses.ActivateKeypadsResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.ActivateKeypadsRequest(id, reader1, reader2, reader3, reader4)
	}

	return exec[T, responses.ActivateKeypadsResponse](u, controller, f, timeout)
}
