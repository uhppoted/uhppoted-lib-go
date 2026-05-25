// generated code - ** DO NOT EDIT **

package uhppoted

import (
	"net/netip"
	"reflect"
	"testing"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"

	test "integration-tests"
)

func TestGetController(t *testing.T) {
	expected := test.Expected.GetController

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	response, err := uhppoted.GetController(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetIPv4(t *testing.T) {
	expected := test.Expected.SetIPv4

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	address := netip.MustParseAddr("192.168.1.125")
	netmask := netip.MustParseAddr("255.255.255.0")
	gateway := netip.MustParseAddr("192.168.1.1")

	response, err := uhppoted.SetIPv4(u, controller, address, netmask, gateway, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetTime(t *testing.T) {
	expected := test.Expected.GetTime

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	response, err := uhppoted.GetTime(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetTime(t *testing.T) {
	expected := test.Expected.SetTime

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	dateTime := types.MustParseDateTime("2025-11-04 12:34:56")

	response, err := uhppoted.SetTime(u, controller, dateTime, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetListener(t *testing.T) {
	expected := test.Expected.GetListener

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	response, err := uhppoted.GetListener(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetListener(t *testing.T) {
	expected := test.Expected.SetListener

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	address := netip.MustParseAddr("192.168.1.100")
	port := uint16(60001)
	interval := uint8(17)

	response, err := uhppoted.SetListener(u, controller, address, port, interval, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetListenerAddrPort(t *testing.T) {
	expected := test.Expected.GetListenerAddrPort

	controller := uhppoted.Controller{
		ID:       405419897,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	response, err := uhppoted.GetListenerAddrPort(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetListenerAddrPort(t *testing.T) {
	expected := test.Expected.SetListenerAddrPort

	controller := uhppoted.Controller{
		ID:       405419897,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	listener := netip.MustParseAddrPort("192.168.1.100:60001")
	interval := uint8(17)

	response, err := uhppoted.SetListenerAddrPort(u, controller, listener, interval, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetDoor(t *testing.T) {
	expected := test.Expected.GetDoor

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	door := uint8(4)

	response, err := uhppoted.GetDoor(u, controller, door, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetDoor(t *testing.T) {
	expected := test.Expected.SetDoor

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	door := uint8(4)
	mode := types.DoorMode(2)
	delay := uint8(17)

	response, err := uhppoted.SetDoor(u, controller, door, mode, delay, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetDoorPasscodes(t *testing.T) {
	expected := test.Expected.SetDoorPasscodes

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	door := uint8(4)
	passcode1 := uint32(12345)
	passcode2 := uint32(54321)
	passcode3 := uint32(999999)
	passcode4 := uint32(0)

	response, err := uhppoted.SetDoorPasscodes(u, controller, door, passcode1, passcode2, passcode3, passcode4, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestOpenDoor(t *testing.T) {
	expected := test.Expected.OpenDoor

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	door := uint8(4)

	response, err := uhppoted.OpenDoor(u, controller, door, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetStatus(t *testing.T) {
	expected := test.Expected.GetStatus

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	response, err := uhppoted.GetStatus(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetStatusNoEvent(t *testing.T) {
	expected := test.Expected.GetStatusNoEvent

	controller := uhppoted.Controller{
		ID:       405419897,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	response, err := uhppoted.GetStatus(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetCards(t *testing.T) {
	expected := test.Expected.GetCards

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	response, err := uhppoted.GetCards(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetCard(t *testing.T) {
	expected := test.Expected.GetCard

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	card := uint32(10058400)

	response, err := uhppoted.GetCard(u, controller, card, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetCardNotFound(t *testing.T) {
	expected := test.Expected.GetCardNotFound

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	card := uint32(10058401)

	response, err := uhppoted.GetCard(u, controller, card, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetCardAtIndex(t *testing.T) {
	expected := test.Expected.GetCardAtIndex

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	index := uint32(135)

	response, err := uhppoted.GetCardAtIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetCardAtIndexNotFound(t *testing.T) {
	expected := test.Expected.GetCardAtIndexNotFound

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	index := uint32(136)

	response, err := uhppoted.GetCardAtIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetCardAtIndexDeleted(t *testing.T) {
	expected := test.Expected.GetCardAtIndexDeleted

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	index := uint32(137)

	response, err := uhppoted.GetCardAtIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestPutCard(t *testing.T) {
	expected := test.Expected.PutCard

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	card := uint32(10058400)
	startDate := types.MustParseDate("2025-01-01")
	endDate := types.MustParseDate("2025-12-31")
	door1 := uint8(1)
	door2 := uint8(0)
	door3 := uint8(17)
	door4 := uint8(1)
	PIN := uint32(999999)

	response, err := uhppoted.PutCard(u, controller, card, startDate, endDate, door1, door2, door3, door4, PIN, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestDeleteCard(t *testing.T) {
	expected := test.Expected.DeleteCard

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	card := uint32(10058400)

	response, err := uhppoted.DeleteCard(u, controller, card, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestDeleteAllCards(t *testing.T) {
	expected := test.Expected.DeleteAllCards

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	response, err := uhppoted.DeleteAllCards(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetEvent(t *testing.T) {
	expected := test.Expected.GetEvent

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	eventIndex := uint32(13579)

	response, err := uhppoted.GetEvent(u, controller, eventIndex, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetEventNotFound(t *testing.T) {
	expected := test.Expected.GetEventNotFound

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	eventIndex := uint32(24680)

	response, err := uhppoted.GetEvent(u, controller, eventIndex, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetEventOverwritten(t *testing.T) {
	expected := test.Expected.GetEventOverwritten

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	eventIndex := uint32(98765)

	response, err := uhppoted.GetEvent(u, controller, eventIndex, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetEventIndex(t *testing.T) {
	expected := test.Expected.GetEventIndex

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	response, err := uhppoted.GetEventIndex(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetEventIndex(t *testing.T) {
	expected := test.Expected.SetEventIndex

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	index := uint32(13579)

	response, err := uhppoted.SetEventIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestRecordSpecialEvents(t *testing.T) {
	expected := test.Expected.RecordSpecialEvents

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	enabled := true

	response, err := uhppoted.RecordSpecialEvents(u, controller, enabled, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetTimeProfile(t *testing.T) {
	expected := test.Expected.GetTimeProfile

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	profile := uint8(37)

	response, err := uhppoted.GetTimeProfile(u, controller, profile, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetTimeProfile(t *testing.T) {
	expected := test.Expected.SetTimeProfile

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	profile := uint8(37)
	startDate := types.MustParseDate("2025-11-26")
	endDate := types.MustParseDate("2025-12-29")
	monday := true
	tuesday := true
	wednesday := false
	thursday := true
	friday := false
	saturday := true
	sunday := true
	segment1Start := types.MustParseHHmm("8:30")
	segment1End := types.MustParseHHmm("9:45")
	segment2Start := types.MustParseHHmm("11:35")
	segment2End := types.MustParseHHmm("13:15")
	segment3Start := types.MustParseHHmm("14:01")
	segment3End := types.MustParseHHmm("17:59")
	linkedProfileId := uint8(19)

	response, err := uhppoted.SetTimeProfile(u, controller, profile, startDate, endDate, monday, tuesday, wednesday, thursday, friday, saturday, sunday, segment1Start, segment1End, segment2Start, segment2End, segment3Start, segment3End, linkedProfileId, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestClearTimeProfiles(t *testing.T) {
	expected := test.Expected.ClearTimeProfiles

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	response, err := uhppoted.ClearTimeProfiles(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestAddTask(t *testing.T) {
	expected := test.Expected.AddTask

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	task := types.TaskType(2)
	startDate := types.MustParseDate("2025-01-01")
	endDate := types.MustParseDate("2025-12-31")
	monday := true
	tuesday := true
	wednesday := false
	thursday := true
	friday := false
	saturday := true
	sunday := true
	startTime := types.MustParseHHmm("08:45")
	door := uint8(3)
	moreCards := uint8(7)

	response, err := uhppoted.AddTask(u, controller, task, startDate, endDate, monday, tuesday, wednesday, thursday, friday, saturday, sunday, startTime, door, moreCards, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestRefreshTaskList(t *testing.T) {
	expected := test.Expected.RefreshTaskList

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	response, err := uhppoted.RefreshTaskList(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestClearTaskList(t *testing.T) {
	expected := test.Expected.ClearTaskList

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	response, err := uhppoted.ClearTaskList(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetPcControl(t *testing.T) {
	expected := test.Expected.SetPcControl

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	enabled := true

	response, err := uhppoted.SetPCControl(u, controller, enabled, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetInterlock(t *testing.T) {
	expected := test.Expected.SetInterlock

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	interlock := types.Interlock(8)

	response, err := uhppoted.SetInterlock(u, controller, interlock, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestActivateKeypads(t *testing.T) {
	expected := test.Expected.ActivateKeypads

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	reader1 := true
	reader2 := true
	reader3 := false
	reader4 := true

	response, err := uhppoted.ActivateKeypads(u, controller, reader1, reader2, reader3, reader4, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetAntipassback(t *testing.T) {
	expected := test.Expected.GetAntipassback

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	response, err := uhppoted.GetAntiPassback(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetAntipassback(t *testing.T) {
	expected := test.Expected.SetAntipassback

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	antipassback := types.AntiPassback(2)

	response, err := uhppoted.SetAntiPassback(u, controller, antipassback, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestRestoreDefaultParameters(t *testing.T) {
	expected := test.Expected.RestoreDefaultParameters

	controller := uhppoted.Controller{
		ID:       405419896,
		Address:  netip.MustParseAddrPort("127.0.0.1:50002"),
		Protocol: "udp",
	}

	response, err := uhppoted.RestoreDefaultParameters(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}
