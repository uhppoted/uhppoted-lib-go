// UHPPOTE request packet decoder.
//
// Decodes a UHPPOTE access controller 64 byte UDP response packet into the
// equivalent Python object.
package decode

import (
	"fmt"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/codec"
)

// Decodes a get-controller response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - GetControllerResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func GetControllerResponse(packet []byte) (codec.GetControllerResponse, error) {
	if len(packet) != 64 {
		return codec.GetControllerResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != codec.SOM {
		return codec.GetControllerResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != codec.GetController {
		return codec.GetControllerResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return codec.GetControllerResponse{
		unpackUint32(packet, 4),
		unpackIPv4(packet, 8),
		unpackIPv4(packet, 12),
		unpackIPv4(packet, 16),
		unpackMAC(packet, 20),
		unpackVersion(packet, 26),
		unpackYYYYMMDD(packet, 28),
	}, nil
}

// Decodes a set-ipv4 response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - SetIPv4Response initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func SetIPv4Response(packet []byte) (codec.SetIPv4Response, error) {
	if len(packet) != 64 {
		return codec.SetIPv4Response{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != codec.SOM {
		return codec.SetIPv4Response{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != codec.SetIPv4 {
		return codec.SetIPv4Response{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return codec.SetIPv4Response{
		unpackUint32(packet, 4),
		unpackBool(packet, 8),
	}, nil
}

// Decodes a get-status response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - GetStatusResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func GetStatusResponse(packet []byte) (codec.GetStatusResponse, error) {
	if len(packet) != 64 {
		return codec.GetStatusResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != codec.SOM {
		return codec.GetStatusResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != codec.GetStatus {
		return codec.GetStatusResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return codec.GetStatusResponse{
		unpackUint32(packet, 4),
		unpackYYMMDD(packet, 51),
		unpackHHMMSS(packet, 37),
		unpackBool(packet, 28),
		unpackBool(packet, 29),
		unpackBool(packet, 30),
		unpackBool(packet, 31),
		unpackBool(packet, 32),
		unpackBool(packet, 33),
		unpackBool(packet, 34),
		unpackBool(packet, 35),
		unpackUint8(packet, 49),
		unpackUint8(packet, 50),
		unpackUint8(packet, 36),
		unpackUint8(packet, 48),
		unpackUint32(packet, 8),
		unpackUint8(packet, 12),
		unpackBool(packet, 13),
		unpackUint8(packet, 14),
		unpackUint8(packet, 15),
		unpackUint32(packet, 16),
		unpackYYYYMMDDHHMMSS(packet, 20),
		unpackUint8(packet, 27),
		unpackUint32(packet, 40),
	}, nil
}
