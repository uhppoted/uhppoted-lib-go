// generated code - ** DO NOT EDIT **

package uhppoted

import (
	"net/netip"
	"reflect"
	"testing"

	test "integration-tests"
	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/entities"
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
	controller := uint32(405419896)

	response, err := lib.GetController(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetController) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetController, response)
	}
}

func TestSetIPv4(t *testing.T) {
	controller := uint32(405419896)

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
	controller := uint32(405419896)

	response, err := lib.GetTime(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetTime) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetTime, response)
	}
}

func TestSetTime(t *testing.T) {
	controller := uint32(405419896)

	dateTime := entities.MustParseDateTime("2025-11-04 12:34:56")

	response, err := lib.SetTime(u, controller, dateTime, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.SetTime) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.SetTime, response)
	}
}

func TestGetListener(t *testing.T) {
	controller := uint32(405419896)

	response, err := lib.GetListener(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetListener) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetListener, response)
	}
}

func TestSetListener(t *testing.T) {
	controller := uint32(405419896)

	address := netip.MustParseAddr("192.168.1.100")
	port := uint16(60001)
	interval := uint8(17)

	response, err := lib.SetListener(u, controller, address, port, interval, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.SetListener) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.SetListener, response)
	}
}

func TestGetListenerAddrPort(t *testing.T) {
	controller := uint32(405419897)

	response, err := lib.GetListenerAddrPort(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetListenerAddrPort) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetListenerAddrPort, response)
	}
}

func TestSetListenerAddrPort(t *testing.T) {
	controller := uint32(405419897)

	listener := netip.MustParseAddrPort("192.168.1.100:60001")
	interval := uint8(17)

	response, err := lib.SetListenerAddrPort(u, controller, listener, interval, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.SetListenerAddrPort) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.SetListenerAddrPort, response)
	}
}

func TestGetDoor(t *testing.T) {
	controller := uint32(405419896)

	door := uint8(4)

	response, err := lib.GetDoor(u, controller, door, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetDoor) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetDoor, response)
	}
}

func TestSetDoor(t *testing.T) {
	controller := uint32(405419896)

	door := uint8(4)
	mode := entities.DoorMode(2)
	delay := uint8(17)

	response, err := lib.SetDoor(u, controller, door, mode, delay, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.SetDoor) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.SetDoor, response)
	}
}

func TestSetDoorPasscodes(t *testing.T) {
	controller := uint32(405419896)

	door := uint8(4)
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
	controller := uint32(405419896)

	door := uint8(4)

	response, err := lib.OpenDoor(u, controller, door, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.OpenDoor) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.OpenDoor, response)
	}
}

func TestGetStatus(t *testing.T) {
	controller := uint32(405419896)

	response, err := lib.GetStatus(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetStatus) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetStatus, response)
	}
}

func TestGetStatusNoEvent(t *testing.T) {
	controller := uint32(405419897)

	response, err := lib.GetStatus(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetStatusNoEvent) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetStatusNoEvent, response)
	}
}

func TestGetCards(t *testing.T) {
	controller := uint32(405419896)

	response, err := lib.GetCards(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCards) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCards, response)
	}
}

func TestGetCard(t *testing.T) {
	controller := uint32(405419896)

	card := uint32(10058400)

	response, err := lib.GetCard(u, controller, card, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCard) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCard, response)
	}
}

func TestGetCardNotFound(t *testing.T) {
	controller := uint32(405419896)

	card := uint32(10058401)

	response, err := lib.GetCard(u, controller, card, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCardNotFound) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCardNotFound, response)
	}
}

func TestGetCardAtIndex(t *testing.T) {
	controller := uint32(405419896)

	index := uint32(135)

	response, err := lib.GetCardAtIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCardAtIndex) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCardAtIndex, response)
	}
}

func TestGetCardAtIndexNotFound(t *testing.T) {
	controller := uint32(405419896)

	index := uint32(136)

	response, err := lib.GetCardAtIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCardAtIndexNotFound) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCardAtIndexNotFound, response)
	}
}

func TestGetCardAtIndexDeleted(t *testing.T) {
	controller := uint32(405419896)

	index := uint32(137)

	response, err := lib.GetCardAtIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCardAtIndexDeleted) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCardAtIndexDeleted, response)
	}
}

func TestPutCard(t *testing.T) {
	controller := uint32(405419896)

	card := uint32(10058400)
	startDate := entities.MustParseDate("2025-01-01")
	endDate := entities.MustParseDate("2025-12-31")
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
	controller := uint32(405419896)

	card := uint32(10058400)

	response, err := lib.DeleteCard(u, controller, card, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.DeleteCard) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.DeleteCard, response)
	}
}

func TestDeleteAllCards(t *testing.T) {
	controller := uint32(405419896)

	response, err := lib.DeleteAllCards(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.DeleteAllCards) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.DeleteAllCards, response)
	}
}

func TestGetEvent(t *testing.T) {
	controller := uint32(405419896)

	eventIndex := uint32(13579)

	response, err := lib.GetEvent(u, controller, eventIndex, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetEvent) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetEvent, response)
	}
}

func TestGetEventNotFound(t *testing.T) {
	controller := uint32(405419896)

	eventIndex := uint32(24680)

	response, err := lib.GetEvent(u, controller, eventIndex, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetEventNotFound) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetEventNotFound, response)
	}
}

func TestGetEventOverwritten(t *testing.T) {
	controller := uint32(405419896)

	eventIndex := uint32(98765)

	response, err := lib.GetEvent(u, controller, eventIndex, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetEventOverwritten) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetEventOverwritten, response)
	}
}

func TestGetEventIndex(t *testing.T) {
	controller := uint32(405419896)

	response, err := lib.GetEventIndex(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetEventIndex) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetEventIndex, response)
	}
}

func TestSetEventIndex(t *testing.T) {
	controller := uint32(405419896)

	index := uint32(13579)

	response, err := lib.SetEventIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.SetEventIndex) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.SetEventIndex, response)
	}
}

func TestRecordSpecialEvents(t *testing.T) {
	controller := uint32(405419896)

	enabled := true

	response, err := lib.RecordSpecialEvents(u, controller, enabled, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.RecordSpecialEvents) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.RecordSpecialEvents, response)
	}
}

func TestGetTimeProfile(t *testing.T) {
	controller := uint32(405419896)

	profile := uint8(37)

	response, err := lib.GetTimeProfile(u, controller, profile, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetTimeProfile) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetTimeProfile, response)
	}
}

func TestSetTimeProfile(t *testing.T) {
	controller := uint32(405419896)

	profile := uint8(37)
	startDate := entities.MustParseDate("2025-11-26")
	endDate := entities.MustParseDate("2025-12-29")
	monday := true
	tuesday := true
	wednesday := false
	thursday := true
	friday := false
	saturday := true
	sunday := true
	segment1Start := entities.MustParseHHmm("8:30")
	segment1End := entities.MustParseHHmm("9:45")
	segment2Start := entities.MustParseHHmm("11:35")
	segment2End := entities.MustParseHHmm("13:15")
	segment3Start := entities.MustParseHHmm("14:01")
	segment3End := entities.MustParseHHmm("17:59")
	linkedProfileId := uint8(19)

	response, err := lib.SetTimeProfile(u, controller, profile, startDate, endDate, monday, tuesday, wednesday, thursday, friday, saturday, sunday, segment1Start, segment1End, segment2Start, segment2End, segment3Start, segment3End, linkedProfileId, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.SetTimeProfile) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.SetTimeProfile, response)
	}
}

func TestClearTimeProfiles(t *testing.T) {
	controller := uint32(405419896)

	response, err := lib.ClearTimeProfiles(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.ClearTimeProfiles) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.ClearTimeProfiles, response)
	}
}

func TestAddTask(t *testing.T) {
	controller := uint32(405419896)

	task := entities.TaskType(2)
	startDate := entities.MustParseDate("2025-01-01")
	endDate := entities.MustParseDate("2025-12-31")
	monday := true
	tuesday := true
	wednesday := false
	thursday := true
	friday := false
	saturday := true
	sunday := true
	startTime := entities.MustParseHHmm("08:45")
	door := uint8(3)
	moreCards := uint8(7)

	response, err := lib.AddTask(u, controller, task, startDate, endDate, monday, tuesday, wednesday, thursday, friday, saturday, sunday, startTime, door, moreCards, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.AddTask) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.AddTask, response)
	}
}

func TestRefreshTaskList(t *testing.T) {
	controller := uint32(405419896)

	response, err := lib.RefreshTaskList(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.RefreshTaskList) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.RefreshTaskList, response)
	}
}

func TestClearTaskList(t *testing.T) {
	controller := uint32(405419896)

	response, err := lib.ClearTaskList(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.ClearTaskList) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.ClearTaskList, response)
	}
}

func TestSetPcControl(t *testing.T) {
	controller := uint32(405419896)

	enabled := true

	response, err := lib.SetPCControl(u, controller, enabled, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.SetPcControl) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.SetPcControl, response)
	}
}

func TestSetInterlock(t *testing.T) {
	controller := uint32(405419896)

	interlock := entities.Interlock(8)

	response, err := lib.SetInterlock(u, controller, interlock, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.SetInterlock) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.SetInterlock, response)
	}
}

func TestActivateKeypads(t *testing.T) {
	controller := uint32(405419896)

	reader1 := true
	reader2 := true
	reader3 := false
	reader4 := true

	response, err := lib.ActivateKeypads(u, controller, reader1, reader2, reader3, reader4, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.ActivateKeypads) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.ActivateKeypads, response)
	}
}

func TestGetAntipassback(t *testing.T) {
	controller := uint32(405419896)

	response, err := lib.GetAntiPassback(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetAntipassback) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetAntipassback, response)
	}
}

func TestSetAntipassback(t *testing.T) {
	controller := uint32(405419896)

	antipassback := uint8(2)

	response, err := lib.SetAntiPassback(u, controller, antipassback, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.SetAntipassback) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.SetAntipassback, response)
	}
}

func TestRestoreDefaultParameters(t *testing.T) {
	controller := uint32(405419896)

	response, err := lib.RestoreDefaultParameters(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.RestoreDefaultParameters) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.RestoreDefaultParameters, response)
	}
}
