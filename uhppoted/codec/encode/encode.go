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
	"net/netip"
	"time"
)

// Message constants for the UHPPOTE request/response protocol.
const SOM byte = 0x17
const SOM_v6_62 byte = 0x19

const GetStatus byte = 0x20
const SetTime byte = 0x30
const GetTime byte = 0x32
const OPEN_DOOR byte = 0x40
const PUT_CARD byte = 0x50
const DELETE_CARD byte = 0x52
const DELETE_ALL_CARDS byte = 0x54
const GET_CARDS byte = 0x58
const GET_CARD byte = 0x5A
const GET_CARD_AT_INDEX byte = 0x5C
const SET_DOOR byte = 0x80
const GET_DOOR byte = 0x82
const SET_ANTIPASSBACK byte = 0x84
const GET_ANTIPASSBACK byte = 0x86
const SET_TIME_PROFILE byte = 0x88
const CLEAR_TIME_PROFILES byte = 0x8A
const SET_DOOR_PASSCODES byte = 0x8C
const RECORD_SPECIAL_EVENTS byte = 0x8E
const SET_LISTENER byte = 0x90
const GetListener byte = 0x92
const GetController byte = 0x94
const SetIPv4 byte = 0x96
const GET_TIME_PROFILE byte = 0x98
const SET_PC_CONTROL byte = 0xA0
const SET_INTERLOCK byte = 0xA2
const ACTIVATE_KEYPADS byte = 0xA4
const CLEAR_TASKLIST byte = 0xA6
const ADD_TASK byte = 0xA8
const SET_FIRST_CARD byte = 0xAA
const REFRESH_TASKLIST byte = 0xAC
const GET_EVENT byte = 0xB0
const SET_EVENT_INDEX byte = 0xB2
const GET_EVENT_INDEX byte = 0xB4
const RESTORE_DEFAULT_PARAMETERS byte = 0xC8
const LISTEN_EVENT byte = 0x20

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

// Packs a date/time value 'in-place' as a 7-byte value into the packet at the offset.
//
//	Parameters:
//	   v      (time.Time)  date/time.
//	   packet (bytearray)  64 byte array.
//	   offset (int)        Value location in array.
func packDateTime(v time.Time, packet []byte, offset int) {
	s := v.Format("20060102150405")

	bytes := string2bcd(s)
	copy(packet[offset:], bytes)
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
