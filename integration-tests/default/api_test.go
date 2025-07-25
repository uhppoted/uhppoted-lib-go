package uhppoted

import (
	"net/netip"
	"reflect"
	"testing"

	test "github.com/uhppoted/uhppoted-lib-go/integration-tests"
	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func TestGetAllControllers(t *testing.T) {
	response, err := lib.GetAllControllers(u, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetAllControllers) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetAllControllers, response)
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
