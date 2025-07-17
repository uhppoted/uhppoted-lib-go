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
