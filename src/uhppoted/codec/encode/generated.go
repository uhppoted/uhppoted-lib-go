// generated code - ** DO NOT EDIT **

package encode

import (
	"net/netip"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"
)

// Encodes a GetControllerRequest request to a 64 byte packet.
func GetControllerRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x94

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a SetIPv4Request request to a 64 byte packet.
func SetIPv4Request(controller uint32, address netip.Addr, netmask netip.Addr, gateway netip.Addr) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x96

	packUint32(controller, packet, 4)
	packIPv4(address, packet, 8)
	packIPv4(netmask, packet, 12)
	packIPv4(gateway, packet, 16)
	packUint32(0x55aaaa55, packet, 20)

	return packet, nil
}

// Encodes a GetStatusRequest request to a 64 byte packet.
func GetStatusRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x20

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a GetTimeRequest request to a 64 byte packet.
func GetTimeRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x32

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a SetTimeRequest request to a 64 byte packet.
func SetTimeRequest(controller uint32, datetime types.DateTime) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x30

	packUint32(controller, packet, 4)
	packDateTime(datetime, packet, 8)

	return packet, nil
}

// Encodes a GetListenerRequest request to a 64 byte packet.
func GetListenerRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x92

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a SetListenerRequest request to a 64 byte packet.
func SetListenerRequest(controller uint32, address netip.Addr, port uint16, interval uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x90

	packUint32(controller, packet, 4)
	packIPv4(address, packet, 8)
	packUint16(port, packet, 12)
	packUint8(interval, packet, 14)

	return packet, nil
}

// Encodes a GetListenerAddrPortRequest request to a 64 byte packet.
func GetListenerAddrPortRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x92

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a SetListenerAddrPortRequest request to a 64 byte packet.
func SetListenerAddrPortRequest(controller uint32, listener netip.AddrPort, interval uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x90

	packUint32(controller, packet, 4)
	packAddrPort(listener, packet, 8)
	packUint8(interval, packet, 14)

	return packet, nil
}

// Encodes a GetDoorRequest request to a 64 byte packet.
func GetDoorRequest(controller uint32, door uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x82

	packUint32(controller, packet, 4)
	packUint8(door, packet, 8)

	return packet, nil
}

// Encodes a SetDoorRequest request to a 64 byte packet.
func SetDoorRequest(controller uint32, door uint8, mode types.DoorMode, delay uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x80

	packUint32(controller, packet, 4)
	packUint8(door, packet, 8)
	packMode(mode, packet, 9)
	packUint8(delay, packet, 10)

	return packet, nil
}

// Encodes a SetDoorPasscodesRequest request to a 64 byte packet.
func SetDoorPasscodesRequest(controller uint32, door uint8, passcode1 uint32, passcode2 uint32, passcode3 uint32, passcode4 uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x8c

	packUint32(controller, packet, 4)
	packUint8(door, packet, 8)
	packPIN(passcode1, packet, 12)
	packPIN(passcode2, packet, 16)
	packPIN(passcode3, packet, 20)
	packPIN(passcode4, packet, 24)

	return packet, nil
}

// Encodes a OpenDoorRequest request to a 64 byte packet.
func OpenDoorRequest(controller uint32, door uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x40

	packUint32(controller, packet, 4)
	packUint8(door, packet, 8)

	return packet, nil
}

// Encodes a GetCardsRequest request to a 64 byte packet.
func GetCardsRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x58

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a GetCardRequest request to a 64 byte packet.
func GetCardRequest(controller uint32, cardnumber uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x5a

	packUint32(controller, packet, 4)
	packUint32(cardnumber, packet, 8)

	return packet, nil
}

// Encodes a GetCardAtIndexRequest request to a 64 byte packet.
func GetCardAtIndexRequest(controller uint32, index uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x5c

	packUint32(controller, packet, 4)
	packUint32(index, packet, 8)

	return packet, nil
}

// Encodes a PutCardRequest request to a 64 byte packet.
func PutCardRequest(controller uint32, card uint32, startdate types.Date, enddate types.Date, door1 uint8, door2 uint8, door3 uint8, door4 uint8, PIN uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x50

	packUint32(controller, packet, 4)
	packUint32(card, packet, 8)
	packDate(startdate, packet, 12)
	packDate(enddate, packet, 16)
	packUint8(door1, packet, 20)
	packUint8(door2, packet, 21)
	packUint8(door3, packet, 22)
	packUint8(door4, packet, 23)
	packPIN(PIN, packet, 24)

	return packet, nil
}

// Encodes a DeleteCardRequest request to a 64 byte packet.
func DeleteCardRequest(controller uint32, cardnumber uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x52

	packUint32(controller, packet, 4)
	packUint32(cardnumber, packet, 8)

	return packet, nil
}

// Encodes a DeleteCardsRequest request to a 64 byte packet.
func DeleteCardsRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x54

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)

	return packet, nil
}

// Encodes a GetEventRequest request to a 64 byte packet.
func GetEventRequest(controller uint32, eventindex uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0xb0

	packUint32(controller, packet, 4)
	packUint32(eventindex, packet, 8)

	return packet, nil
}

// Encodes a GetEventIndexRequest request to a 64 byte packet.
func GetEventIndexRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0xb4

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a SetEventIndexRequest request to a 64 byte packet.
func SetEventIndexRequest(controller uint32, eventindex uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0xb2

	packUint32(controller, packet, 4)
	packUint32(eventindex, packet, 8)
	packUint32(0x55aaaa55, packet, 12)

	return packet, nil
}

// Encodes a RecordSpecialEventsRequest request to a 64 byte packet.
func RecordSpecialEventsRequest(controller uint32, enabled bool) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x8e

	packUint32(controller, packet, 4)
	packBool(enabled, packet, 8)

	return packet, nil
}

// Encodes a GetTimeProfileRequest request to a 64 byte packet.
func GetTimeProfileRequest(controller uint32, profile uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x98

	packUint32(controller, packet, 4)
	packUint8(profile, packet, 8)

	return packet, nil
}

// Encodes a SetTimeProfileRequest request to a 64 byte packet.
func SetTimeProfileRequest(controller uint32, profile uint8, startdate types.Date, enddate types.Date, monday bool, tuesday bool, wednesday bool, thursday bool, friday bool, saturday bool, sunday bool, segment1start types.HHmm, segment1end types.HHmm, segment2start types.HHmm, segment2end types.HHmm, segment3start types.HHmm, segment3end types.HHmm, linkedprofileid uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x88

	packUint32(controller, packet, 4)
	packUint8(profile, packet, 8)
	packDate(startdate, packet, 9)
	packDate(enddate, packet, 13)
	packBool(monday, packet, 17)
	packBool(tuesday, packet, 18)
	packBool(wednesday, packet, 19)
	packBool(thursday, packet, 20)
	packBool(friday, packet, 21)
	packBool(saturday, packet, 22)
	packBool(sunday, packet, 23)
	packHHmm(segment1start, packet, 24)
	packHHmm(segment1end, packet, 26)
	packHHmm(segment2start, packet, 28)
	packHHmm(segment2end, packet, 30)
	packHHmm(segment3start, packet, 32)
	packHHmm(segment3end, packet, 34)
	packUint8(linkedprofileid, packet, 36)

	return packet, nil
}

// Encodes a ClearTimeProfilesRequest request to a 64 byte packet.
func ClearTimeProfilesRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x8a

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)

	return packet, nil
}

// Encodes a AddTaskRequest request to a 64 byte packet.
func AddTaskRequest(controller uint32, task types.TaskType, startdate types.Date, enddate types.Date, monday bool, tuesday bool, wednesday bool, thursday bool, friday bool, saturday bool, sunday bool, starttime types.HHmm, door uint8, morecards uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0xa8

	packUint32(controller, packet, 4)
	packTaskType(task, packet, 26)
	packDate(startdate, packet, 8)
	packDate(enddate, packet, 12)
	packBool(monday, packet, 16)
	packBool(tuesday, packet, 17)
	packBool(wednesday, packet, 18)
	packBool(thursday, packet, 19)
	packBool(friday, packet, 20)
	packBool(saturday, packet, 21)
	packBool(sunday, packet, 22)
	packHHmm(starttime, packet, 23)
	packUint8(door, packet, 25)
	packUint8(morecards, packet, 27)

	return packet, nil
}

// Encodes a RefreshTaskListRequest request to a 64 byte packet.
func RefreshTaskListRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0xac

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)

	return packet, nil
}

// Encodes a ClearTasklistRequest request to a 64 byte packet.
func ClearTasklistRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0xa6

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)

	return packet, nil
}

// Encodes a SetPCControlRequest request to a 64 byte packet.
func SetPCControlRequest(controller uint32, enabled bool) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0xa0

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)
	packBool(enabled, packet, 12)

	return packet, nil
}

// Encodes a SetInterlockRequest request to a 64 byte packet.
func SetInterlockRequest(controller uint32, interlock types.Interlock) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0xa2

	packUint32(controller, packet, 4)
	packInterlock(interlock, packet, 8)

	return packet, nil
}

// Encodes a ActivateKeypadsRequest request to a 64 byte packet.
func ActivateKeypadsRequest(controller uint32, reader1 bool, reader2 bool, reader3 bool, reader4 bool) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0xa4

	packUint32(controller, packet, 4)
	packBool(reader1, packet, 8)
	packBool(reader2, packet, 9)
	packBool(reader3, packet, 10)
	packBool(reader4, packet, 11)

	return packet, nil
}

// Encodes a GetAntipassbackRequest request to a 64 byte packet.
func GetAntipassbackRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x86

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a SetAntipassbackRequest request to a 64 byte packet.
func SetAntipassbackRequest(controller uint32, antipassback types.AntiPassback) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0x84

	packUint32(controller, packet, 4)
	packAntiPassback(antipassback, packet, 8)

	return packet, nil
}

// Encodes a RestoreDefaultParametersRequest request to a 64 byte packet.
func RestoreDefaultParametersRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 0xc8

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)

	return packet, nil
}
