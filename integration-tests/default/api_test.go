package uhppoted

import (
	"net/netip"
	"reflect"
	"testing"

	test "github.com/uhppoted/uhppoted-lib-go/integration-tests"
	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func TestFindControllers(t *testing.T) {
	response, err := lib.FindControllers(u, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.FindControllers) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.FindControllers, response)
	}
}

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

func TestGetStatus(t *testing.T) {

	response, err := lib.GetStatus(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetStatus) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetStatus, response)
	}
}

func TestGetCards(t *testing.T) {

	response, err := lib.GetCards(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCards) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCards, response)
	}
}

func TestGetCard(t *testing.T) {
	card := uint32(10058400)

	response, err := lib.GetCard(u, controller, card, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCard) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCard, response)
	}
}

func TestGetCardNotFound(t *testing.T) {
	card := uint32(10058401)

	response, err := lib.GetCard(u, controller, card, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCardNotFound) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCardNotFound, response)
	}
}

func TestGetCardAtIndex(t *testing.T) {
	index := uint32(135)

	response, err := lib.GetCardAtIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCardAtIndex) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCardAtIndex, response)
	}
}

func TestGetCardAtIndexNotFound(t *testing.T) {
	index := uint32(136)

	response, err := lib.GetCardAtIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCardAtIndexNotFound) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCardAtIndexNotFound, response)
	}
}

func TestGetCardAtIndexDeleted(t *testing.T) {
	index := uint32(137)

	response, err := lib.GetCardAtIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCardAtIndexDeleted) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCardAtIndexDeleted, response)
	}
}

func TestPutCard(t *testing.T) {
	card := uint32(10058400)
	startDate := string2date("2025-01-01")
	endDate := string2date("2025-12-31")
	door1 := uint8(1)
	door2 := uint8(0)
	door3 := uint8(17)
	door4 := uint8(1)
	PIN := uint32(999999)

	response, err := lib.PutCard(u, controller, card, startDate, endDate, door1, door2, door3, door4, PIN, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.PutCard) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.PutCard, response)
	}
}

func TestDeleteCard(t *testing.T) {
	card := uint32(10058400)

	response, err := lib.DeleteCard(u, controller, card, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.DeleteCard) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.DeleteCard, response)
	}
}

func TestDeleteAllCards(t *testing.T) {

	response, err := lib.DeleteAllCards(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.DeleteAllCards) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.DeleteAllCards, response)
	}
}
