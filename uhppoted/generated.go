package uhppoted

import (
	"time"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/encode"
)

// GetController retrieves the system information from an access controller.
func GetController[T TController](u Uhppoted, controller T, timeout time.Duration) (GetControllerResponse, error) {
	f := func(id uint32) ([]byte, error) {
		return encode.GetControllerRequest(id)
	}

	return exec[T, GetControllerResponse](u, controller, f, timeout)
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
