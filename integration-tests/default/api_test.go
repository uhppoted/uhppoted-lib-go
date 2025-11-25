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

func TestFindControllers(t *testing.T) {
	expected := test.Expected.FindControllers

	response, err := uhppoted.FindControllers(u, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetController(t *testing.T) {
	controller := uint32(405419896)

	expected := test.Expected.GetController

	response, err := uhppoted.GetController(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetIPv4(t *testing.T) {
	controller := uint32(405419896)

	address := netip.MustParseAddr("192.168.1.125")
	netmask := netip.MustParseAddr("255.255.255.0")
	gateway := netip.MustParseAddr("192.168.1.1")

	expected := test.Expected.SetIPv4

	response, err := uhppoted.SetIPv4(u, controller, address, netmask, gateway, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetTime(t *testing.T) {
	controller := uint32(405419896)

	expected := test.Expected.GetTime

	response, err := uhppoted.GetTime(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetTime(t *testing.T) {
	controller := uint32(405419896)

	dateTime := types.MustParseDateTime("2025-11-04 12:34:56")

	expected := test.Expected.SetTime

	response, err := uhppoted.SetTime(u, controller, dateTime, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetListener(t *testing.T) {
	controller := uint32(405419896)

	expected := test.Expected.GetListener

	response, err := uhppoted.GetListener(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetListener(t *testing.T) {
	controller := uint32(405419896)

	address := netip.MustParseAddr("192.168.1.100")
	port := uint16(60001)
	interval := uint8(17)

	expected := test.Expected.SetListener

	response, err := uhppoted.SetListener(u, controller, address, port, interval, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetListenerAddrPort(t *testing.T) {
	controller := uint32(405419897)

	expected := test.Expected.GetListenerAddrPort

	response, err := uhppoted.GetListenerAddrPort(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetListenerAddrPort(t *testing.T) {
	controller := uint32(405419897)

	listener := netip.MustParseAddrPort("192.168.1.100:60001")
	interval := uint8(17)

	expected := test.Expected.SetListenerAddrPort

	response, err := uhppoted.SetListenerAddrPort(u, controller, listener, interval, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetDoor(t *testing.T) {
	controller := uint32(405419896)

	door := uint8(4)

	expected := test.Expected.GetDoor

	response, err := uhppoted.GetDoor(u, controller, door, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetDoor(t *testing.T) {
	controller := uint32(405419896)

	door := uint8(4)
	mode := types.DoorMode(2)
	delay := uint8(17)

	expected := test.Expected.SetDoor

	response, err := uhppoted.SetDoor(u, controller, door, mode, delay, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetDoorPasscodes(t *testing.T) {
	controller := uint32(405419896)

	door := uint8(4)
	passcode1 := uint32(12345)
	passcode2 := uint32(54321)
	passcode3 := uint32(999999)
	passcode4 := uint32(0)

	expected := test.Expected.SetDoorPasscodes

	response, err := uhppoted.SetDoorPasscodes(u, controller, door, passcode1, passcode2, passcode3, passcode4, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestOpenDoor(t *testing.T) {
	controller := uint32(405419896)

	door := uint8(4)

	expected := test.Expected.OpenDoor

	response, err := uhppoted.OpenDoor(u, controller, door, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetStatus(t *testing.T) {
	controller := uint32(405419896)

	expected := test.Expected.GetStatus

	response, err := uhppoted.GetStatus(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetStatusNoEvent(t *testing.T) {
	controller := uint32(405419897)

	expected := test.Expected.GetStatusNoEvent

	response, err := uhppoted.GetStatus(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetCards(t *testing.T) {
	controller := uint32(405419896)

	expected := test.Expected.GetCards

	response, err := uhppoted.GetCards(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetCard(t *testing.T) {
	controller := uint32(405419896)

	card := uint32(10058400)

	expected := test.Expected.GetCard

	response, err := uhppoted.GetCard(u, controller, card, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetCardNotFound(t *testing.T) {
	controller := uint32(405419896)

	card := uint32(10058401)

	expected := test.Expected.GetCardNotFound

	response, err := uhppoted.GetCard(u, controller, card, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetCardAtIndex(t *testing.T) {
	controller := uint32(405419896)

	index := uint32(135)

	expected := test.Expected.GetCardAtIndex

	response, err := uhppoted.GetCardAtIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetCardAtIndexNotFound(t *testing.T) {
	controller := uint32(405419896)

	index := uint32(136)

	expected := test.Expected.GetCardAtIndexNotFound

	response, err := uhppoted.GetCardAtIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetCardAtIndexDeleted(t *testing.T) {
	controller := uint32(405419896)

	index := uint32(137)

	expected := test.Expected.GetCardAtIndexDeleted

	response, err := uhppoted.GetCardAtIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestPutCard(t *testing.T) {
	controller := uint32(405419896)

	card := uint32(10058400)
	startDate := types.MustParseDate("2025-01-01")
	endDate := types.MustParseDate("2025-12-31")
	door1 := uint8(1)
	door2 := uint8(0)
	door3 := uint8(17)
	door4 := uint8(1)
	PIN := uint32(999999)

	expected := test.Expected.PutCard

	response, err := uhppoted.PutCard(u, controller, card, startDate, endDate, door1, door2, door3, door4, PIN, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestDeleteCard(t *testing.T) {
	controller := uint32(405419896)

	card := uint32(10058400)

	expected := test.Expected.DeleteCard

	response, err := uhppoted.DeleteCard(u, controller, card, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestDeleteAllCards(t *testing.T) {
	controller := uint32(405419896)

	expected := test.Expected.DeleteAllCards

	response, err := uhppoted.DeleteAllCards(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetEvent(t *testing.T) {
	controller := uint32(405419896)

	eventIndex := uint32(13579)

	expected := test.Expected.GetEvent

	response, err := uhppoted.GetEvent(u, controller, eventIndex, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetEventNotFound(t *testing.T) {
	controller := uint32(405419896)

	eventIndex := uint32(24680)

	expected := test.Expected.GetEventNotFound

	response, err := uhppoted.GetEvent(u, controller, eventIndex, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetEventOverwritten(t *testing.T) {
	controller := uint32(405419896)

	eventIndex := uint32(98765)

	expected := test.Expected.GetEventOverwritten

	response, err := uhppoted.GetEvent(u, controller, eventIndex, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetEventIndex(t *testing.T) {
	controller := uint32(405419896)

	expected := test.Expected.GetEventIndex

	response, err := uhppoted.GetEventIndex(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetEventIndex(t *testing.T) {
	controller := uint32(405419896)

	index := uint32(13579)

	expected := test.Expected.SetEventIndex

	response, err := uhppoted.SetEventIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestRecordSpecialEvents(t *testing.T) {
	controller := uint32(405419896)

	enabled := true

	expected := test.Expected.RecordSpecialEvents

	response, err := uhppoted.RecordSpecialEvents(u, controller, enabled, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetTimeProfile(t *testing.T) {
	controller := uint32(405419896)

	profile := uint8(37)

	expected := test.Expected.GetTimeProfile

	response, err := uhppoted.GetTimeProfile(u, controller, profile, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetTimeProfile(t *testing.T) {
	controller := uint32(405419896)

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

	expected := test.Expected.SetTimeProfile

	response, err := uhppoted.SetTimeProfile(u, controller, profile, startDate, endDate, monday, tuesday, wednesday, thursday, friday, saturday, sunday, segment1Start, segment1End, segment2Start, segment2End, segment3Start, segment3End, linkedProfileId, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestClearTimeProfiles(t *testing.T) {
	controller := uint32(405419896)

	expected := test.Expected.ClearTimeProfiles

	response, err := uhppoted.ClearTimeProfiles(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestAddTask(t *testing.T) {
	controller := uint32(405419896)

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

	expected := test.Expected.AddTask

	response, err := uhppoted.AddTask(u, controller, task, startDate, endDate, monday, tuesday, wednesday, thursday, friday, saturday, sunday, startTime, door, moreCards, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestRefreshTaskList(t *testing.T) {
	controller := uint32(405419896)

	expected := test.Expected.RefreshTaskList

	response, err := uhppoted.RefreshTaskList(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestClearTaskList(t *testing.T) {
	controller := uint32(405419896)

	expected := test.Expected.ClearTaskList

	response, err := uhppoted.ClearTaskList(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetPcControl(t *testing.T) {
	controller := uint32(405419896)

	enabled := true

	expected := test.Expected.SetPcControl

	response, err := uhppoted.SetPCControl(u, controller, enabled, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetInterlock(t *testing.T) {
	controller := uint32(405419896)

	interlock := types.Interlock(8)

	expected := test.Expected.SetInterlock

	response, err := uhppoted.SetInterlock(u, controller, interlock, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestActivateKeypads(t *testing.T) {
	controller := uint32(405419896)

	reader1 := true
	reader2 := true
	reader3 := false
	reader4 := true

	expected := test.Expected.ActivateKeypads

	response, err := uhppoted.ActivateKeypads(u, controller, reader1, reader2, reader3, reader4, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestGetAntipassback(t *testing.T) {
	controller := uint32(405419896)

	expected := test.Expected.GetAntipassback

	response, err := uhppoted.GetAntiPassback(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestSetAntipassback(t *testing.T) {
	controller := uint32(405419896)

	antipassback := types.AntiPassback(2)

	expected := test.Expected.SetAntipassback

	response, err := uhppoted.SetAntiPassback(u, controller, antipassback, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}

func TestRestoreDefaultParameters(t *testing.T) {
	controller := uint32(405419896)

	expected := test.Expected.RestoreDefaultParameters

	response, err := uhppoted.RestoreDefaultParameters(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", expected, response)
	}
}
