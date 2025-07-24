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
		Controller: unpackUint32(packet, 4),
		IpAddress:  unpackIPv4(packet, 8),
		SubnetMask: unpackIPv4(packet, 12),
		Gateway:    unpackIPv4(packet, 16),
		MACAddress: unpackMAC(packet, 20),
		Version:    unpackVersion(packet, 26),
		Date:       unpackYYYYMMDD(packet, 28),
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
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
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
		Controller:         unpackUint32(packet, 4),
		SystemDate:         unpackYYMMDD(packet, 51),
		SystemTime:         unpackHHMMSS(packet, 37),
		Door1Open:          unpackBool(packet, 28),
		Door2Open:          unpackBool(packet, 29),
		Door3Open:          unpackBool(packet, 30),
		Door4Open:          unpackBool(packet, 31),
		Door1Button:        unpackBool(packet, 32),
		Door2Button:        unpackBool(packet, 33),
		Door3Button:        unpackBool(packet, 34),
		Door4Button:        unpackBool(packet, 35),
		Relays:             unpackUint8(packet, 49),
		Inputs:             unpackUint8(packet, 50),
		SystemError:        unpackUint8(packet, 36),
		SpecialInfo:        unpackUint8(packet, 48),
		EventIndex:         unpackUint32(packet, 8),
		EventType:          unpackUint8(packet, 12),
		EventAccessGranted: unpackBool(packet, 13),
		EventDoor:          unpackUint8(packet, 14),
		EventDirection:     unpackUint8(packet, 15),
		EventCard:          unpackUint32(packet, 16),
		EventTimestamp:     unpackYYYYMMDDHHMMSS(packet, 20),
		EventReason:        unpackUint8(packet, 27),
		SequenceNo:         unpackUint32(packet, 40),
	}, nil
}

// Decodes a get-time response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - GetTimeResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func GetTimeResponse(packet []byte) (codec.GetTimeResponse, error) {
	if len(packet) != 64 {
		return codec.GetTimeResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != codec.SOM {
		return codec.GetTimeResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != codec.GetTime {
		return codec.GetTimeResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return codec.GetTimeResponse{
		Controller: unpackUint32(packet, 4),
		DateTime:   unpackYYYYMMDDHHMMSS(packet, 8),
	}, nil
}

// Decodes a set-time response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - SetTimeResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func SetTimeResponse(packet []byte) (codec.SetTimeResponse, error) {
	if len(packet) != 64 {
		return codec.SetTimeResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != codec.SOM {
		return codec.SetTimeResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != codec.SetTime {
		return codec.SetTimeResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return codec.SetTimeResponse{
		Controller: unpackUint32(packet, 4),
		DateTime:   unpackYYYYMMDDHHMMSS(packet, 8),
	}, nil
}

// Decodes a get-listener response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - GetListenerResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func GetListenerResponse(packet []byte) (codec.GetListenerResponse, error) {
	if len(packet) != 64 {
		return codec.GetListenerResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != codec.SOM {
		return codec.GetListenerResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != codec.GetListener {
		return codec.GetListenerResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return codec.GetListenerResponse{
		Controller: unpackUint32(packet, 4),
		Address:    unpackAddrPort(packet, 8),
		Interval:   unpackUint8(packet, 14),
	}, nil
}
