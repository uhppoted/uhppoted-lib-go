// UHPPOTE request packet encoder.
//
// Encodes a UHPPOTE access controller request as a 64 byte UDP packet:
//
// - uint8, uint16, uint24 and uint32 values are encoded as little endian unsigned integers
// - datetime, date and time values are encoded as BCD
// - boolean values are encoded as 0 (False) or 1 (True)
package encode

import (
	"encoding/binary"
	"fmt"
	"net/netip"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"
)

// Message constants for the UHPPOTE request/response protocol.
const SOM byte = 0x17
const SOM_v6_62 byte = 0x19

// Packs a boolean value 'in-place' as a 1-byte true (0x01) or false (0x00) value
// into the packet at the offset.
//
//	Parameters:
//	   v      (uint8)     uint8 value to encode.
//	   packet (bytearray)  64 byte array.
//	   offset (int)        Value location in array.
func packBool(v bool, packet []byte, offset int) {
	if v {
		packet[offset] = 0x01
	} else {
		packet[offset] = 0x00
	}
}

// Packs a uint8 value 'in-place' as a 1-byte value into the packet at the offset.
//
//	Parameters:
//	   v      (uint8)     uint8 value to encode.
//	   packet (bytearray)  64 byte array.
//	   offset (int)        Value location in array.
func packUint8(v uint8, packet []byte, offset int) {
	packet[offset] = v
}

// Packs a uint16 value 'in-place' as a 2-byte little endian value into the packet
// at the offset.
//
//	Parameters:
//	   v      (uint16)     uint16 value to encode.
//	   packet (bytearray)  64 byte array.
//	   offset (int)        Value location in array.
func packUint16(v uint16, packet []byte, offset int) {
	binary.LittleEndian.PutUint16(packet[offset:offset+2], v)
}

// Packs a uint32 value 'in-place' as a 4-byte little endian value into the packet
// at the offset.
//
//	Parameters:
//	   v      (uint32)     uint32 value to encode.
//	   packet (bytearray)  64 byte array.
//	   offset (int)        Value location in array.
func packUint32(v uint32, packet []byte, offset int) {
	binary.LittleEndian.PutUint32(packet[offset:offset+4], v)
}

// Packs a netip.Addr IPv4 value 'in-place' as a 4-byte value into the packet at the offset.
//
//	Parameters:
//	   v      (netip.Addr) IPv4 address.
//	   packet (bytearray)  64 byte array.
//	   offset (int)        Value location in array.
func packIPv4(v netip.Addr, packet []byte, offset int) {
	addr := v.As4()

	copy(packet[offset:], addr[:])
}

// Packs a netip.AddrPort IPv4 address:port value 'in-place' as a 6-byte value into the packet at the offset.
//
//	Parameters:
//	   v      (netip.AddrPort) IPv4 address:port.
//	   packet (bytearray)      64 byte array.
//	   offset (int)            Value location in array.
func packAddrPort(v netip.AddrPort, packet []byte, offset int) {
	addr := v.Addr().As4()
	port := v.Port()

	copy(packet[offset:], addr[:])
	binary.LittleEndian.PutUint16(packet[offset+4:offset+6], port)
}

// Packs a date/time value 'in-place' as a 7-byte BCD value into the packet at the offset.
//
//	Parameters:
//	   v      (time.Time)  date/time.
//	   packet (bytearray)  64 byte array.
//	   offset (int)        Value location in array.
func packDateTime(v types.DateTime, packet []byte, offset int) {
	s := fmt.Sprintf("%04v%02v%02v%02v%02v%02v", v.Year(), uint8(v.Month()), v.Day(), v.Hour(), v.Minute(), v.Second())

	bytes := string2bcd(s)
	copy(packet[offset:], bytes)
}

// Packs a date value 'in-place' as a 4-byte BCD value into the packet at the offset.
//
//	Parameters:
//	   v      (time.Time)  date/time.
//	   packet (bytearray)  64 byte array.
//	   offset (int)        Value location in array.
func packDate(v types.Date, packet []byte, offset int) {
	s := fmt.Sprintf("%04v%02v%02v", v.Year(), uint8(v.Month()), v.Day())

	bytes := string2bcd(s)
	copy(packet[offset:], bytes)
}

// Packs an hour:minute value 'in-place' as a 2-byte BCD value into the packet at the offset.
//
//	Parameters:
//	   v      (time.Time)  date/time.
//	   packet (bytearray)  64 byte array.
//	   offset (int)        Value location in array.
func packHHmm(v types.HHmm, packet []byte, offset uint8) {
	s := fmt.Sprintf("%02v%02v", v.Hour(), v.Minute())
	bytes := string2bcd(s)

	copy(packet[offset:], bytes)
}

// Packs a 6 digit PIN as a 3-byte uint value into the packet at the offset.
//
//	Parameters:
//	   v      (uint32)     PIN [0..999999]
//	   packet (bytearray)  64 byte array.
//	   offset (int)        Value location in array.
func packPIN(v uint32, packet []byte, offset uint8) error {
	bytes := make([]byte, 4)

	if v < 1000000 {
		binary.LittleEndian.PutUint32(bytes, uint32(v))

		packet[offset] = bytes[0]
		packet[offset+1] = bytes[1]
		packet[offset+2] = bytes[2]
	}

	return nil
}

// Packs a DoorMode value 'in-place' as a 1-byte value into the packet at the offset.
//
//	Parameters:
//	   v      (DoorMode)   'door mode' value to encode.
//	   packet (bytearray)  64 byte array.
//	   offset (int)        Value location in array.
func packMode(v types.DoorMode, packet []byte, offset int) {
	packet[offset] = uint8(v)
}

// Packs a Task value 'in-place' as a 1-byte value into the packet at the offset.
//
//	Parameters:
//	   v      (Task)      'task' value to encode.
//	   packet (bytearray)  64 byte array.
//	   offset (int)        Value location in array.
func packTaskType(v types.TaskType, packet []byte, offset int) {
	packet[offset] = uint8(v)
}

// Packs an Interlock value 'in-place' as a 1-byte value into the packet at the offset.
//
//	Parameters:
//	   v      (Interlock)  'interlock' value to encode.
//	   packet (bytearray)  64 byte array.
//	   offset (int)        Value location in array.
func packInterlock(v types.Interlock, packet []byte, offset int) {
	packet[offset] = uint8(v)
}

// Packs an AntiPassback value 'in-place' as a 1-byte value into the packet at the offset.
//
//	Parameters:
//	   v      (AntiPassback)  'antipassback' value to encode.
//	   packet (bytearray)     64 byte array.
//	   offset (int)           Value location in array.
func packAntiPassback(v types.AntiPassback, packet []byte, offset int) {
	packet[offset] = uint8(v)
}

// Converts a string of digits to packed BCD. Invalid characters (non-digits) are silently
// discarded.
func string2bcd(s string) []byte {
	BCD := map[rune]uint8{
		'0': 0x00,
		'1': 0x01,
		'2': 0x02,
		'3': 0x03,
		'4': 0x04,
		'5': 0x05,
		'6': 0x06,
		'7': 0x07,
		'8': 0x08,
		'9': 0x09,
	}

	nibbles := []byte{}

	if len(s)%2 != 0 {
		nibbles = append(nibbles, 0)
	}

	for _, ch := range s {
		if n, ok := BCD[ch]; ok {
			nibbles = append(nibbles, n)
		}
	}

	bytes := []byte{}
	for i := 0; i < len(nibbles); i += 2 {
		msb := nibbles[i]
		lsb := nibbles[i+1]
		b := (msb << 4) | lsb

		bytes = append(bytes, b)
	}

	return bytes
}
