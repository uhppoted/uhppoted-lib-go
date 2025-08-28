package uhppoted

import (
	"net/netip"
	"reflect"
	"testing"

	test "github.com/uhppoted/uhppoted-lib-go/integration-tests"
	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func TestGetController(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	response, err := lib.GetController(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetController) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetController, response)
	}
}

func TestSetIPv4(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

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
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	response, err := lib.GetTime(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetTime) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetTime, response)
	}
}

func TestSetTime(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	dateTime := string2datetime("2025-11-04 12:34:56")

	response, err := lib.SetTime(u, controller, dateTime, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.SetTime) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.SetTime, response)
	}
}

func TestGetListener(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	response, err := lib.GetListener(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetListener) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetListener, response)
	}
}

func TestSetListener(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

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

func TestGetDoor(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	door := uint8(4)

	response, err := lib.GetDoor(u, controller, door, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetDoor) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetDoor, response)
	}
}

func TestSetDoor(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	door := uint8(4)
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
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

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
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	door := uint8(4)

	response, err := lib.OpenDoor(u, controller, door, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.OpenDoor) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.OpenDoor, response)
	}
}

func TestGetStatus(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	response, err := lib.GetStatus(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetStatus) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetStatus, response)
	}
}

func TestGetStatusNoEvent(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419897),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	response, err := lib.GetStatus(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetStatusNoEvent) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetStatusNoEvent, response)
	}
}

func TestGetCards(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	response, err := lib.GetCards(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCards) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCards, response)
	}
}

func TestGetCard(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	card := uint32(10058400)

	response, err := lib.GetCard(u, controller, card, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCard) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCard, response)
	}
}

func TestGetCardNotFound(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	card := uint32(10058401)

	response, err := lib.GetCard(u, controller, card, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCardNotFound) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCardNotFound, response)
	}
}

func TestGetCardAtIndex(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	index := uint32(135)

	response, err := lib.GetCardAtIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCardAtIndex) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCardAtIndex, response)
	}
}

func TestGetCardAtIndexNotFound(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	index := uint32(136)

	response, err := lib.GetCardAtIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCardAtIndexNotFound) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCardAtIndexNotFound, response)
	}
}

func TestGetCardAtIndexDeleted(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	index := uint32(137)

	response, err := lib.GetCardAtIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetCardAtIndexDeleted) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetCardAtIndexDeleted, response)
	}
}

func TestPutCard(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

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
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	card := uint32(10058400)

	response, err := lib.DeleteCard(u, controller, card, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.DeleteCard) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.DeleteCard, response)
	}
}

func TestDeleteAllCards(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	response, err := lib.DeleteAllCards(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.DeleteAllCards) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.DeleteAllCards, response)
	}
}

func TestGetEvent(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	eventIndex := uint32(13579)

	response, err := lib.GetEvent(u, controller, eventIndex, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetEvent) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetEvent, response)
	}
}

func TestGetEventNotFound(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	eventIndex := uint32(24680)

	response, err := lib.GetEvent(u, controller, eventIndex, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetEventNotFound) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetEventNotFound, response)
	}
}

func TestGetEventOverwritten(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	eventIndex := uint32(98765)

	response, err := lib.GetEvent(u, controller, eventIndex, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetEventOverwritten) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetEventOverwritten, response)
	}
}

func TestGetEventIndex(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	response, err := lib.GetEventIndex(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetEventIndex) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetEventIndex, response)
	}
}

func TestSetEventIndex(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	index := uint32(13579)

	response, err := lib.SetEventIndex(u, controller, index, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.SetEventIndex) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.SetEventIndex, response)
	}
}

func TestRecordSpecialEvents(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	enabled := true

	response, err := lib.RecordSpecialEvents(u, controller, enabled, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.RecordSpecialEvents) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.RecordSpecialEvents, response)
	}
}

func TestGetTimeProfile(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	profile := uint8(37)

	response, err := lib.GetTimeProfile(u, controller, profile, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.GetTimeProfile) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.GetTimeProfile, response)
	}
}

func TestSetTimeProfile(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	profile := uint8(37)
	startDate := string2date("2025-11-26")
	endDate := string2date("2025-12-29")
	monday := true
	tuesday := true
	wednesday := false
	thursday := true
	friday := false
	saturday := true
	sunday := true
	segment1Start := string2HHmm("8:30")
	segment1End := string2HHmm("9:45")
	segment2Start := string2HHmm("11:35")
	segment2End := string2HHmm("13:15")
	segment3Start := string2HHmm("14:01")
	segment3End := string2HHmm("17:59")
	linkedProfileId := uint8(19)

	response, err := lib.SetTimeProfile(u, controller, profile, startDate, endDate, monday, tuesday, wednesday, thursday, friday, saturday, sunday, segment1Start, segment1End, segment2Start, segment2End, segment3Start, segment3End, linkedProfileId, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.SetTimeProfile) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.SetTimeProfile, response)
	}
}

func TestClearTimeProfiles(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	response, err := lib.ClearTimeProfiles(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.ClearTimeProfiles) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.ClearTimeProfiles, response)
	}
}

func TestAddTask(t *testing.T) {
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	task := uint8(2)
	startDate := string2date("2025-01-01")
	endDate := string2date("2025-12-31")
	monday := true
	tuesday := true
	wednesday := false
	thursday := true
	friday := false
	saturday := true
	sunday := true
	startTime := string2HHmm("08:45")
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
	controller := lib.Controller{
		ID:       uint32(405419896),
		Address:  netip.MustParseAddrPort("127.0.0.1:50003"),
		Protocol: "tcp",
	}

	response, err := lib.RefreshTaskList(u, controller, timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, test.Expected.RefreshTaskList) {
		t.Errorf("incorrect response\n   expected:%#v\n   got:     %#v", test.Expected.RefreshTaskList, response)
	}
}
