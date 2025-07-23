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
