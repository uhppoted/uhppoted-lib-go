// generated code - ** DO NOT EDIT **

package encode

import (
	"net/netip"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"
)

// Encodes a GetControllerRequest.
func GetControllerRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 148

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a SetIPv4Request.
func SetIPv4Request(controller uint32, address netip.Addr, netmask netip.Addr, gateway netip.Addr) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 150

	packUint32(controller, packet, 4)
	packIPv4(address, packet, 8)
	packIPv4(netmask, packet, 12)
	packIPv4(gateway, packet, 16)
	packUint32(0x55aaaa55, packet, 20)

	return packet, nil
}

// Encodes a GetStatusRequest.
func GetStatusRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 32

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a GetTimeRequest.
func GetTimeRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 50

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a SetTimeRequest.
func SetTimeRequest(controller uint32, datetime types.DateTime) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 48

	packUint32(controller, packet, 4)
	packDateTime(datetime, packet, 8)

	return packet, nil
}

// Encodes a GetListenerRequest.
func GetListenerRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 146

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a SetListenerRequest.
func SetListenerRequest(controller uint32, address netip.Addr, port uint16, interval uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 144

	packUint32(controller, packet, 4)
	packIPv4(address, packet, 8)
	packUint16(port, packet, 12)
	packUint8(interval, packet, 14)

	return packet, nil
}

// Encodes a GetListenerAddrPortRequest.
func GetListenerAddrPortRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 146

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a SetListenerAddrPortRequest.
func SetListenerAddrPortRequest(controller uint32, listener netip.AddrPort, interval uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 144

	packUint32(controller, packet, 4)
	packAddrPort(listener, packet, 8)
	packUint8(interval, packet, 14)

	return packet, nil
}

// Encodes a GetDoorRequest.
func GetDoorRequest(controller uint32, door uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 130

	packUint32(controller, packet, 4)
	packUint8(door, packet, 8)

	return packet, nil
}

// Encodes a SetDoorRequest.
func SetDoorRequest(controller uint32, door uint8, mode types.DoorMode, delay uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 128

	packUint32(controller, packet, 4)
	packUint8(door, packet, 8)
	packMode(mode, packet, 9)
	packUint8(delay, packet, 10)

	return packet, nil
}

// Encodes a SetDoorPasscodesRequest.
func SetDoorPasscodesRequest(controller uint32, door uint8, passcode1 uint32, passcode2 uint32, passcode3 uint32, passcode4 uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 140

	packUint32(controller, packet, 4)
	packUint8(door, packet, 8)
	packPIN(passcode1, packet, 12)
	packPIN(passcode2, packet, 16)
	packPIN(passcode3, packet, 20)
	packPIN(passcode4, packet, 24)

	return packet, nil
}

// Encodes a OpenDoorRequest.
func OpenDoorRequest(controller uint32, door uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 64

	packUint32(controller, packet, 4)
	packUint8(door, packet, 8)

	return packet, nil
}

// Encodes a GetCardsRequest.
func GetCardsRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 88

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a GetCardRequest.
func GetCardRequest(controller uint32, cardnumber uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 90

	packUint32(controller, packet, 4)
	packUint32(cardnumber, packet, 8)

	return packet, nil
}

// Encodes a GetCardAtIndexRequest.
func GetCardAtIndexRequest(controller uint32, index uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 92

	packUint32(controller, packet, 4)
	packUint32(index, packet, 8)

	return packet, nil
}

// Encodes a PutCardRequest.
func PutCardRequest(controller uint32, card uint32, startdate types.Date, enddate types.Date, door1 uint8, door2 uint8, door3 uint8, door4 uint8, PIN uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 80

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

// Encodes a DeleteCardRequest.
func DeleteCardRequest(controller uint32, cardnumber uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 82

	packUint32(controller, packet, 4)
	packUint32(cardnumber, packet, 8)

	return packet, nil
}

// Encodes a DeleteCardsRequest.
func DeleteCardsRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 84

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)

	return packet, nil
}

// Encodes a GetEventRequest.
func GetEventRequest(controller uint32, eventindex uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 176

	packUint32(controller, packet, 4)
	packUint32(eventindex, packet, 8)

	return packet, nil
}

// Encodes a GetEventIndexRequest.
func GetEventIndexRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 180

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a SetEventIndexRequest.
func SetEventIndexRequest(controller uint32, eventindex uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 178

	packUint32(controller, packet, 4)
	packUint32(eventindex, packet, 8)
	packUint32(0x55aaaa55, packet, 12)

	return packet, nil
}

// Encodes a RecordSpecialEventsRequest.
func RecordSpecialEventsRequest(controller uint32, enabled bool) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 142

	packUint32(controller, packet, 4)
	packBool(enabled, packet, 8)

	return packet, nil
}

// Encodes a GetTimeProfileRequest.
func GetTimeProfileRequest(controller uint32, profile uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 152

	packUint32(controller, packet, 4)
	packUint8(profile, packet, 8)

	return packet, nil
}

// Encodes a SetTimeProfileRequest.
func SetTimeProfileRequest(controller uint32, profile uint8, startdate types.Date, enddate types.Date, monday bool, tuesday bool, wednesday bool, thursday bool, friday bool, saturday bool, sunday bool, segment1start types.HHmm, segment1end types.HHmm, segment2start types.HHmm, segment2end types.HHmm, segment3start types.HHmm, segment3end types.HHmm, linkedprofileid uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 136

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

// Encodes a ClearTimeProfilesRequest.
func ClearTimeProfilesRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 138

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)

	return packet, nil
}

// Encodes a AddTaskRequest.
func AddTaskRequest(controller uint32, task types.TaskType, startdate types.Date, enddate types.Date, monday bool, tuesday bool, wednesday bool, thursday bool, friday bool, saturday bool, sunday bool, starttime types.HHmm, door uint8, morecards uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 168

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

// Encodes a RefreshTaskListRequest.
func RefreshTaskListRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 172

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)

	return packet, nil
}

// Encodes a ClearTasklistRequest.
func ClearTasklistRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 166

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)

	return packet, nil
}

// Encodes a SetPCControlRequest.
func SetPCControlRequest(controller uint32, enabled bool) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 160

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)
	packBool(enabled, packet, 12)

	return packet, nil
}

// Encodes a SetInterlockRequest.
func SetInterlockRequest(controller uint32, interlock types.Interlock) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 162

	packUint32(controller, packet, 4)
	packInterlock(interlock, packet, 8)

	return packet, nil
}

// Encodes a ActivateKeypadsRequest.
func ActivateKeypadsRequest(controller uint32, reader1 bool, reader2 bool, reader3 bool, reader4 bool) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 164

	packUint32(controller, packet, 4)
	packBool(reader1, packet, 8)
	packBool(reader2, packet, 9)
	packBool(reader3, packet, 10)
	packBool(reader4, packet, 11)

	return packet, nil
}

// Encodes a GetAntipassbackRequest.
func GetAntipassbackRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 134

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a SetAntipassbackRequest.
func SetAntipassbackRequest(controller uint32, antipassback uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 132

	packUint32(controller, packet, 4)
	packUint8(antipassback, packet, 8)

	return packet, nil
}

// Encodes a RestoreDefaultParametersRequest.
func RestoreDefaultParametersRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 200

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)

	return packet, nil
}
