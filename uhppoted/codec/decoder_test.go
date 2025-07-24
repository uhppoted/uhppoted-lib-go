package codec

import (
	"fmt"
	"net/netip"
	"reflect"
	"testing"
	"time"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/types"
)

func TestDecodeGetControllerResponse(t *testing.T) {
	packet := []byte{
		0x17, 0x94, 0x00, 0x00, 0x78, 0x37, 0x2a, 0x18, 0xc0, 0xa8, 0x01, 0x64, 0xff, 0xff, 0xff, 0x00,
		0xc0, 0xa8, 0x01, 0x01, 0x00, 0x12, 0x23, 0x34, 0x45, 0x56, 0x08, 0x92, 0x20, 0x18, 0x11, 0x05,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	expected := types.GetControllerResponse{
		Controller: 405419896,
		IpAddress:  netip.MustParseAddr("192.168.1.100"),
		SubnetMask: netip.MustParseAddr("255.255.255.0"),
		Gateway:    netip.MustParseAddr("192.168.1.1"),
		MACAddress: "00:12:23:34:45:56",
		Version:    "v8.92",
		Date:       string2date("2018-11-05"),
	}

	if response, err := Decode[types.GetControllerResponse](packet); err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrectly decoded response:\n   expected: %#v\n   got:      %#v", expected, response)
	}
}

func string2date(v string) time.Time {
	if d, err := time.ParseInLocation("2006-01-02", v, time.Local); err != nil {
		panic(fmt.Sprintf("invalid date (%v)", v))
	} else {
		return d
	}
}
