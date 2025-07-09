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
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec"
)

// Encodes a get-controller request.
//
//	Parameters:
//	    controller  (uint32)  controller serial number
//
//	Returns:
//	    64 byte packet.
func GetControllerRequest(controller uint32) ([]byte, error) {
	packet := make([]byte, 64)

	packet[0] = codec.SOM
	packet[1] = 148

	packUint32(controller, packet, 4)

	return packet, nil
}
