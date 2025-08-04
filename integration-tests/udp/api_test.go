package uhppoted

import (
	"net/netip"
	"reflect"
	"testing"

	test "github.com/uhppoted/uhppoted-lib-go/integration-tests"
	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func TestGetController(t *testing.T) {

	response, err := lib.GetController(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetController) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetController, response)
	}
}

func TestSetIPv4(t *testing.T) {
	address := netip.MustParseAddr("192.168.1.125")
	netmask := netip.MustParseAddr("255.255.255.0")
	gateway := netip.MustParseAddr("192.168.1.1")

	response, err := lib.SetIPv4(u, controller, address, netmask, gateway, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.SetIPv4) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.SetIPv4, response)
	}
}

func TestGetStatus(t *testing.T) {

	response, err := lib.GetStatus(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetStatus) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetStatus, response)
	}
}

func TestGetTime(t *testing.T) {

	response, err := lib.GetTime(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetTime) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetTime, response)
	}
}

func TestSetTime(t *testing.T) {
	datetime := string2datetime("2024-11-04 12:34:56")

	response, err := lib.SetTime(u, controller, datetime, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.SetTime) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.SetTime, response)
	}
}

func TestGetListener(t *testing.T) {

	response, err := lib.GetListener(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetListener) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetListener, response)
	}
}

func TestSetListener(t *testing.T) {
	listener := netip.MustParseAddrPort("192.168.1.100:60001")
	interval := uint8(17)

	response, err := lib.SetListener(u, controller, listener, interval, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.SetListener) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.SetListener, response)
	}
}

func TestGetDoor(t *testing.T) {
	door := uint8(3)

	response, err := lib.GetDoor(u, controller, door, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetDoor) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetDoor, response)
	}
}

func TestSetDoor(t *testing.T) {
	door := uint8(3)
	mode := uint8(2)
	delay := uint8(17)

	response, err := lib.SetDoor(u, controller, door, mode, delay, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.SetDoor) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.SetDoor, response)
	}
}

func TestSetDoorPasscodes(t *testing.T) {
	door := uint8(3)
	passcode1 := uint32(12345)
	passcode2 := uint32(54321)
	passcode3 := uint32(999999)
	passcode4 := uint32(0)

	response, err := lib.SetDoorPasscodes(u, controller, door, passcode1, passcode2, passcode3, passcode4, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.SetDoorPasscodes) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.SetDoorPasscodes, response)
	}
}

func TestOpenDoor(t *testing.T) {
	door := uint8(3)

	response, err := lib.OpenDoor(u, controller, door, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.OpenDoor) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.OpenDoor, response)
	}
}
