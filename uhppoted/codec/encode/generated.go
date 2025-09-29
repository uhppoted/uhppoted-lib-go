// generated code - ** DO NOT EDIT **

package encode

import (
	"net/netip"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/entities"
)

// Encodes a get-controller-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//
//	Returns:
//	    64 byte packet.
func GetControllerRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 148

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a set-ipv4-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    address  (IPv4)  controller IPv4 address
//	    netmask  (IPv4)  controller IPv4 subnet mask
//	    gateway  (IPv4)  controller IPv4 gateway address
//	      (magic)  'magic' word
//
//	Returns:
//	    64 byte packet.
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

// Encodes a get-status-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//
//	Returns:
//	    64 byte packet.
func GetStatusRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 32

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a get-time-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//
//	Returns:
//	    64 byte packet.
func GetTimeRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 50

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a set-time-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    datetime  (datetime)  date/time
//
//	Returns:
//	    64 byte packet.
func SetTimeRequest(controller uint32, datetime entities.DateTime) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 48

	packUint32(controller, packet, 4)
	packDateTime(datetime, packet, 8)

	return packet, nil
}

// Encodes a get-listener-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//
//	Returns:
//	    64 byte packet.
func GetListenerRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 146

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a set-listener-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    address  (IPv4)  event listener IPv4 address
//	    port  (uint16)  event listener IPv4 port
//	    interval  (uint8)  status auto-send interval (seconds)
//
//	Returns:
//	    64 byte packet.
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

// Encodes a get-listener-addr-port-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//
//	Returns:
//	    64 byte packet.
func GetListenerAddrPortRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 146

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a set-listener-addr:port-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    listener  (address:port)  event listener IPv4 address:port
//	    interval  (uint8)  status auto-send interval (seconds)
//
//	Returns:
//	    64 byte packet.
func SetListenerAddrPortRequest(controller uint32, listener netip.AddrPort, interval uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 144

	packUint32(controller, packet, 4)
	packAddrPort(listener, packet, 8)
	packUint8(interval, packet, 14)

	return packet, nil
}

// Encodes a get-door-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    door  (uint8)  door ID ([1..4])
//
//	Returns:
//	    64 byte packet.
func GetDoorRequest(controller uint32, door uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 130

	packUint32(controller, packet, 4)
	packUint8(door, packet, 8)

	return packet, nil
}

// Encodes a set-door-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    door  (uint8)
//	    mode  (uint8)
//	    delay  (uint8)
//
//	Returns:
//	    64 byte packet.
func SetDoorRequest(controller uint32, door uint8, mode uint8, delay uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 128

	packUint32(controller, packet, 4)
	packUint8(door, packet, 8)
	packUint8(mode, packet, 9)
	packUint8(delay, packet, 10)

	return packet, nil
}

// Encodes a set-door-passcodes-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    door  (uint8)
//	    passcode 1  (pin)
//	    passcode 2  (pin)
//	    passcode 3  (pin)
//	    passcode 4  (pin)
//
//	Returns:
//	    64 byte packet.
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

// Encodes a open-door-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    door  (uint8)
//
//	Returns:
//	    64 byte packet.
func OpenDoorRequest(controller uint32, door uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 64

	packUint32(controller, packet, 4)
	packUint8(door, packet, 8)

	return packet, nil
}

// Encodes a get-cards-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//
//	Returns:
//	    64 byte packet.
func GetCardsRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 88

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a get-card-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    card number  (uint32)
//
//	Returns:
//	    64 byte packet.
func GetCardRequest(controller uint32, cardnumber uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 90

	packUint32(controller, packet, 4)
	packUint32(cardnumber, packet, 8)

	return packet, nil
}

// Encodes a get-card-at-index-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    index  (uint32)
//
//	Returns:
//	    64 byte packet.
func GetCardAtIndexRequest(controller uint32, index uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 92

	packUint32(controller, packet, 4)
	packUint32(index, packet, 8)

	return packet, nil
}

// Encodes a put-card-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    card  (uint32)
//	    start date  (date)
//	    end date  (date)
//	    door 1  (uint8)
//	    door 2  (uint8)
//	    door 3  (uint8)
//	    door 4  (uint8)
//	    PIN  (pin)
//
//	Returns:
//	    64 byte packet.
func PutCardRequest(controller uint32, card uint32, startdate entities.Date, enddate entities.Date, door1 uint8, door2 uint8, door3 uint8, door4 uint8, PIN uint32) ([]byte, error) {
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

// Encodes a delete-card-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    card number  (uint32)
//
//	Returns:
//	    64 byte packet.
func DeleteCardRequest(controller uint32, cardnumber uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 82

	packUint32(controller, packet, 4)
	packUint32(cardnumber, packet, 8)

	return packet, nil
}

// Encodes a delete-cards-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	      (magic)
//
//	Returns:
//	    64 byte packet.
func DeleteCardsRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 84

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)

	return packet, nil
}

// Encodes a get-event-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    event index  (uint32)
//
//	Returns:
//	    64 byte packet.
func GetEventRequest(controller uint32, eventindex uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 176

	packUint32(controller, packet, 4)
	packUint32(eventindex, packet, 8)

	return packet, nil
}

// Encodes a get-event-index-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//
//	Returns:
//	    64 byte packet.
func GetEventIndexRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 180

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a set-event-index-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    event index  (uint32)
//	      (magic)
//
//	Returns:
//	    64 byte packet.
func SetEventIndexRequest(controller uint32, eventindex uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 178

	packUint32(controller, packet, 4)
	packUint32(eventindex, packet, 8)
	packUint32(0x55aaaa55, packet, 12)

	return packet, nil
}

// Encodes a record-special-events-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    enabled  (bool)
//
//	Returns:
//	    64 byte packet.
func RecordSpecialEventsRequest(controller uint32, enabled bool) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 142

	packUint32(controller, packet, 4)
	packBool(enabled, packet, 8)

	return packet, nil
}

// Encodes a get-time-profile-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    profile  (uint8)
//
//	Returns:
//	    64 byte packet.
func GetTimeProfileRequest(controller uint32, profile uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 152

	packUint32(controller, packet, 4)
	packUint8(profile, packet, 8)

	return packet, nil
}

// Encodes a set-time-profile-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    profile  (uint8)
//	    start date  (date)
//	    end date  (date)
//	    monday  (bool)
//	    tuesday  (bool)
//	    wednesday  (bool)
//	    thursday  (bool)
//	    friday  (bool)
//	    saturday  (bool)
//	    sunday  (bool)
//	    segment 1 start  (HHmm)
//	    segment 1 end  (HHmm)
//	    segment 2 start  (HHmm)
//	    segment 2 end  (HHmm)
//	    segment 3 start  (HHmm)
//	    segment 3 end  (HHmm)
//	    linked profile id  (uint8)
//
//	Returns:
//	    64 byte packet.
func SetTimeProfileRequest(controller uint32, profile uint8, startdate entities.Date, enddate entities.Date, monday bool, tuesday bool, wednesday bool, thursday bool, friday bool, saturday bool, sunday bool, segment1start entities.HHmm, segment1end entities.HHmm, segment2start entities.HHmm, segment2end entities.HHmm, segment3start entities.HHmm, segment3end entities.HHmm, linkedprofileid uint8) ([]byte, error) {
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

// Encodes a clear-time-profiles-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	      (magic)
//
//	Returns:
//	    64 byte packet.
func ClearTimeProfilesRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 138

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)

	return packet, nil
}

// Encodes a add-task-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    task  (task)
//	    start date  (date)
//	    end date  (date)
//	    monday  (bool)
//	    tuesday  (bool)
//	    wednesday  (bool)
//	    thursday  (bool)
//	    friday  (bool)
//	    saturday  (bool)
//	    sunday  (bool)
//	    start time  (HHmm)
//	    door  (uint8)
//	    more cards  (uint8)
//
//	Returns:
//	    64 byte packet.
func AddTaskRequest(controller uint32, task entities.TaskType, startdate entities.Date, enddate entities.Date, monday bool, tuesday bool, wednesday bool, thursday bool, friday bool, saturday bool, sunday bool, starttime entities.HHmm, door uint8, morecards uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 168

	packUint32(controller, packet, 4)
	packTask(task, packet, 26)
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

// Encodes a refresh-task-list-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	      (magic)
//
//	Returns:
//	    64 byte packet.
func RefreshTaskListRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 172

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)

	return packet, nil
}

// Encodes a clear-tasklist-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	      (magic)
//
//	Returns:
//	    64 byte packet.
func ClearTasklistRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 166

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)

	return packet, nil
}

// Encodes a set-pc-control-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	      (magic)
//	    enabled  (bool)
//
//	Returns:
//	    64 byte packet.
func SetPCControlRequest(controller uint32, enabled bool) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 160

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)
	packBool(enabled, packet, 12)

	return packet, nil
}

// Encodes a set-interlock-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    interlock  (interlock)
//
//	Returns:
//	    64 byte packet.
func SetInterlockRequest(controller uint32, interlock entities.Interlock) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 162

	packUint32(controller, packet, 4)
	packInterlock(interlock, packet, 8)

	return packet, nil
}

// Encodes a activate-keypads-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    reader 1  (bool)
//	    reader 2  (bool)
//	    reader 3  (bool)
//	    reader 4  (bool)
//
//	Returns:
//	    64 byte packet.
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

// Encodes a get-antipassback-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//
//	Returns:
//	    64 byte packet.
func GetAntipassbackRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 134

	packUint32(controller, packet, 4)

	return packet, nil
}

// Encodes a set-antipassback-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    antipassback  (uint8)
//
//	Returns:
//	    64 byte packet.
func SetAntipassbackRequest(controller uint32, antipassback uint8) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 132

	packUint32(controller, packet, 4)
	packUint8(antipassback, packet, 8)

	return packet, nil
}

// Encodes a restore-default-parameters-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	      (magic)
//
//	Returns:
//	    64 byte packet.
func RestoreDefaultParametersRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = SOM
	packet[1] = 200

	packUint32(controller, packet, 4)
	packUint32(0x55aaaa55, packet, 8)

	return packet, nil
}
