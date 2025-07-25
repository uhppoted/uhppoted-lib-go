// generated code - ** DO NOT EDIT **

package decode

import (
	"fmt"
	"net/netip"
	"reflect"
	"testing"
	"time"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/types"
)

func TestGetControllerResponse(t *testing.T) {
	packet := []byte{
		0x17, 0x94, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0xc0, 0xa8, 0x01, 0x64, 0xff, 0xff, 0xff, 0x00,
		0xc0, 0xa8, 0x01, 0x01, 0x00, 0x12, 0x23, 0x34, 0x45, 0x56, 0x08, 0x92, 0x20, 0x18, 0x11, 0x05,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	expected := types.GetControllerResponse{
		Controller: 405419896,
		IpAddress:  IPv4("192.168.1.100"),
		SubnetMask: IPv4("255.255.255.0"),
		Gateway:    IPv4("192.168.1.1"),
		MACAddress: "00:12:23:34:45:56",
		Version:    "v8.92",
		Date:       string2date("2018-11-05"),
	}

	response, err := GetControllerResponse(packet)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrectly decoded response:\n   expected: %#v\n   got:      %#v", expected, response)
	}
}

func TestSetIPv4Response(t *testing.T) {
	packet := []byte{
		0x17, 0x96, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	expected := types.SetIPv4Response{
		Controller: 405419896,
		Ok:         true,
	}

	response, err := SetIPv4Response(packet)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrectly decoded response:\n   expected: %#v\n   got:      %#v", expected, response)
	}
}

func TestGetStatusResponse(t *testing.T) {
	packet := []byte{
		0x17, 0x20, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x4e, 0x00, 0x00, 0x00, 0x02, 0x01, 0x03, 0x01,
		0xa1, 0x98, 0x7c, 0x00, 0x20, 0x22, 0x08, 0x23, 0x09, 0x47, 0x06, 0x2c, 0x00, 0x01, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x03, 0x09, 0x49, 0x39, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x27, 0x07, 0x09, 0x22, 0x08, 0x23, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	expected := types.GetStatusResponse{
		Controller:         405419896,
		SystemDate:         string2date("2022-08-23"),
		SystemTime:         string2time("09:49:39"),
		Door1Open:          false,
		Door2Open:          true,
		Door3Open:          false,
		Door4Open:          false,
		Door1Button:        false,
		Door2Button:        false,
		Door3Button:        false,
		Door4Button:        true,
		Relays:             7,
		Inputs:             9,
		SystemError:        3,
		SpecialInfo:        39,
		EventIndex:         78,
		EventType:          2,
		EventAccessGranted: true,
		EventDoor:          3,
		EventDirection:     1,
		EventCard:          8165537,
		EventTimestamp:     string2datetime("2022-08-23 09:47:06"),
		EventReason:        44,
		SequenceNo:         0,
	}

	response, err := GetStatusResponse(packet)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrectly decoded response:\n   expected: %#v\n   got:      %#v", expected, response)
	}
}

func TestGetTimeResponse(t *testing.T) {
	packet := []byte{
		0x17, 0x32, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x20, 0x24, 0x11, 0x01, 0x12, 0x34, 0x56, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	expected := types.GetTimeResponse{
		Controller: 405419896,
		DateTime:   string2datetime("2024-11-01 12:34:56"),
	}

	response, err := GetTimeResponse(packet)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrectly decoded response:\n   expected: %#v\n   got:      %#v", expected, response)
	}
}

func TestSetTimeResponse(t *testing.T) {
	packet := []byte{
		0x17, 0x30, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0x20, 0x24, 0x11, 0x01, 0x12, 0x34, 0x56, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	expected := types.SetTimeResponse{
		Controller: 405419896,
		DateTime:   string2datetime("2024-11-01 12:34:56"),
	}

	response, err := SetTimeResponse(packet)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrectly decoded response:\n   expected: %#v\n   got:      %#v", expected, response)
	}
}

func TestGetListenerResponse(t *testing.T) {
	packet := []byte{
		0x17, 0x92, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0xc0, 0xa8, 0x01, 0x64, 0x61, 0xea, 0x11, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	expected := types.GetListenerResponse{
		Controller: 405419896,
		Address:    addrport("192.168.1.100:60001"),
		Interval:   17,
	}

	response, err := GetListenerResponse(packet)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrectly decoded response:\n   expected: %#v\n   got:      %#v", expected, response)
	}
}

func IPv4(v string) netip.Addr {
	return netip.MustParseAddr(v)
}

func addrport(v string) netip.AddrPort {
	return netip.MustParseAddrPort(v)
}

func string2datetime(v string) time.Time {
	if d, err := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local); err != nil {
		panic(fmt.Sprintf("invalid datetime (%v)", v))
	} else {
		return d
	}
}

func string2date(v string) time.Time {
	if d, err := time.ParseInLocation("2006-01-02", v, time.Local); err != nil {
		panic(fmt.Sprintf("invalid date (%v)", v))
	} else {
		return d
	}
}

func string2time(v string) time.Time {
	if d, err := time.ParseInLocation("15:04:05", v, time.Local); err != nil {
		panic(fmt.Sprintf("invalid time (%v)", v))
	} else {
		return d
	}
}
