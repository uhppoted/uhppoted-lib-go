// generated code - ** DO NOT EDIT **

// UHPPOTE request packet encoder.
//
// Encodes a UHPPOTE access controller request as a 64 byte UDP packet:
//
// - uint8, uint16, uint24 and uint32 values are encoded as little endian unsigned integers
// - datetime, date and time values are encoded as BCD
// - boolean values are encoded as 0 (False) or 1 (True)
package encode

import (
	"net/netip"
	"time"
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
func SetTimeRequest(controller uint32, datetime time.Time) ([]byte, error) {
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

// Encodes a set-listener-address:port-request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//	    listener  (address:port)  event listener IPv4 address:port
//	    interval  (uint8)  status auto-send interval (seconds)
//
//	Returns:
//	    64 byte packet.
func SetListenerAddressPortRequest(controller uint32, listener netip.AddrPort, interval uint8) ([]byte, error) {
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
func PutCardRequest(controller uint32, card uint32, startdate time.Time, enddate time.Time, door1 uint8, door2 uint8, door3 uint8, door4 uint8, PIN uint32) ([]byte, error) {
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
func SetTimeProfileRequest(controller uint32, profile uint8, startdate time.Time, enddate time.Time, monday bool, tuesday bool, wednesday bool, thursday bool, friday bool, saturday bool, sunday bool, segment1start time.Time, segment1end time.Time, segment2start time.Time, segment2end time.Time, segment3start time.Time, segment3end time.Time, linkedprofileid uint8) ([]byte, error) {
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
