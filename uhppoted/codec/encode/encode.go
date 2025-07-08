package encode

import (
	"encoding/binary"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec"
)

//go:generate ../../../.codegen/bin/codegen

// UHPPOTE request packet encoder.
//
// Encodes a UHPPOTE access controller request as a 64 byte UDP packet:
//
// - uint8, uint16, uint24 and uint32 values are encoded as little endian unsigned integers
// - datetime, date and time values are encoded as BCD
// - boolean values are encoded as 0 (False) or 1 (True)

// Encodes a get-controller request.
//
//	Parameters:
//	    controller (uint32)  Controller serial number.
//
//	Returns:
//	    64 byte UDP packet.
func GetControllerRequest(controller uint32) ([]byte, error) {

	packet := make([]byte, 64)

	packet[0] = codec.SOM
	packet[1] = codec.GetController

	pack_uint32(controller, packet, 4)

	return packet, nil
}

// Packs a uint32 value 'in-place' as a 4-byte little endian value into the packet
// at the offset.
//
//	Parameters:
//	   v      (uint32)     uint32 value to encode.
//	   packet (bytearray)  64 byte array.
//	   offset (int)        Value location in array.
func pack_uint32(v uint32, packet []byte, offset int) {
	binary.LittleEndian.PutUint32(packet[offset:offset+4], v)
}
