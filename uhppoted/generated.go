package uhppoted

import (
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/encode"
	"net/netip"
	"time"
)

// GetController retrieves the system information from an access controller.
func GetController[T TController](u Uhppoted, controller T, timeout time.Duration) (GetControllerResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetControllerRequest(id)
	}

	return exec[T, GetControllerResponse](u, controller, f, timeout)
}

// SetIPv4 sets the controller IPv4 address, netmask and gateway address.
func SetIPv4[T TController](u Uhppoted, controller T, address netip.Addr, netmask netip.Addr, gateway netip.Addr, timeout time.Duration) (SetIPv4Response, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetIPv4Request(id, address, netmask, gateway)
	}

	return exec[T, SetIPv4Response](u, controller, f, timeout)
}

// GetTime retrieves the access controller system date and time.
func GetTime[T TController](u Uhppoted, controller T, timeout time.Duration) (GetTimeResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetTimeRequest(id)
	}

	return exec[T, GetTimeResponse](u, controller, f, timeout)
}

// SetTime sets the access controller system date and time.
func SetTime[T TController](u Uhppoted, controller T, datetime time.Time, timeout time.Duration) (SetTimeResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetTimeRequest(id, datetime)
	}

	return exec[T, SetTimeResponse](u, controller, f, timeout)
}

// GetListener retrieves the access controller event listener IPv4 address:port and auto-send interval.
func GetListener[T TController](u Uhppoted, controller T, timeout time.Duration) (GetListenerResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetListenerRequest(id)
	}

	return exec[T, GetListenerResponse](u, controller, f, timeout)
}

// SetListener sets the access controller event listener IPv4 address:port and auto-send interval.
func SetListener[T TController](u Uhppoted, controller T, listener netip.AddrPort, interval uint8, timeout time.Duration) (SetListenerResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetListenerAddressPortRequest(id, listener, interval)
	}

	return exec[T, SetListenerResponse](u, controller, f, timeout)
}

// GetDoor retrieves the control mode and unlock delay time for an access controller door.
func GetDoor[T TController](u Uhppoted, controller T, door uint8, timeout time.Duration) (GetDoorResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetDoorRequest(id, door)
	}

	return exec[T, GetDoorResponse](u, controller, f, timeout)
}

// SetDoor sets the control mode and unlock delay time for an access controller door.
func SetDoor[T TController](u Uhppoted, controller T, door uint8, mode uint8, delay uint8, timeout time.Duration) (SetDoorResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetDoorRequest(id, door, mode, delay)
	}

	return exec[T, SetDoorResponse](u, controller, f, timeout)
}

// Sets up to 4 passcodes for a controller door.
func SetDoorPasscodes[T TController](u Uhppoted, controller T, door uint8, passcode1 uint32, passcode2 uint32, passcode3 uint32, passcode4 uint32, timeout time.Duration) (SetDoorPasscodesResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.SetDoorPasscodesRequest(id, door, passcode1, passcode2, passcode3, passcode4)
	}

	return exec[T, SetDoorPasscodesResponse](u, controller, f, timeout)
}

// Unlocks a door controlled by an access controller.
func OpenDoor[T TController](u Uhppoted, controller T, door uint8, timeout time.Duration) (OpenDoorResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.OpenDoorRequest(id, door)
	}

	return exec[T, OpenDoorResponse](u, controller, f, timeout)
}

// GetStatus retrieves the system status from an access controller.
func GetStatus[T TController](u Uhppoted, controller T, timeout time.Duration) (GetStatusResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetStatusRequest(id)
	}

	return exec[T, GetStatusResponse](u, controller, f, timeout)
}

// Retrieves the number of cards stored on an access controller.
func GetCards[T TController](u Uhppoted, controller T, timeout time.Duration) (GetCardsResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetCardsRequest(id)
	}

	return exec[T, GetCardsResponse](u, controller, f, timeout)
}

// Retrieves the card information for a card number from an access controller.
func GetCard[T TController](u Uhppoted, controller T, cardnumber uint32, timeout time.Duration) (GetCardResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetCardRequest(id, cardnumber)
	}

	return exec[T, GetCardResponse](u, controller, f, timeout)
}

// Retrieves card record stored at the index.
func GetCardAtIndex[T TController](u Uhppoted, controller T, index uint32, timeout time.Duration) (GetCardAtIndexResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetCardAtIndexRequest(id, index)
	}

	return exec[T, GetCardAtIndexResponse](u, controller, f, timeout)
}

// Adds or updates an access controller card record.
func PutCard[T TController](u Uhppoted, controller T, card uint32, startdate time.Time, enddate time.Time, door1 uint8, door2 uint8, door3 uint8, door4 uint8, PIN uint32, timeout time.Duration) (PutCardResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.PutCardRequest(id, card, startdate, enddate, door1, door2, door3, door4, PIN)
	}

	return exec[T, PutCardResponse](u, controller, f, timeout)
}

// DeleteCard removes a card record stored on a controller.
func DeleteCard[T TController](u Uhppoted, controller T, cardnumber uint32, timeout time.Duration) (DeleteCardResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.DeleteCardRequest(id, cardnumber)
	}

	return exec[T, DeleteCardResponse](u, controller, f, timeout)
}

// Deletes all card records stored on an access controller.
func DeleteAllCards[T TController](u Uhppoted, controller T, timeout time.Duration) (DeleteAllCardsResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.DeleteCardsRequest(id)
	}

	return exec[T, DeleteAllCardsResponse](u, controller, f, timeout)
}

// Retrieves an event record stored on an access controller.
func GetEvent[T TController](u Uhppoted, controller T, eventindex uint32, timeout time.Duration) (GetEventResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetEventRequest(id, eventindex)
	}

	return exec[T, GetEventResponse](u, controller, f, timeout)
}

func GetEventIndex[T TController](u Uhppoted, controller T, timeout time.Duration) (GetEventIndexResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetEventIndexRequest(id)
	}

	return exec[T, GetEventIndexResponse](u, controller, f, timeout)
}
