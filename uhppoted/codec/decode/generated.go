// UHPPOTE request packet decoder.
//
// Decodes a UHPPOTE access controller 64 byte UDP response packet into the
// equivalent Python object.
package decode

import (
	"fmt"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/types"
)

// Decodes a get-controller-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - GetControllerResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func GetControllerResponse(packet []byte) (types.GetControllerResponse, error) {
	if len(packet) != 64 {
		return types.GetControllerResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.GetControllerResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 148 {
		return types.GetControllerResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.GetControllerResponse{
		Controller: unpackUint32(packet, 4),
		IpAddress:  unpackIPv4(packet, 8),
		SubnetMask: unpackIPv4(packet, 12),
		Gateway:    unpackIPv4(packet, 16),
		MACAddress: unpackMAC(packet, 20),
		Version:    unpackVersion(packet, 26),
		Date:       unpackDate(packet, 28),
	}, nil
}

// Decodes a set-ipv4-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - SetIPv4Response initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func SetIPv4Response(packet []byte) (types.SetIPv4Response, error) {
	if len(packet) != 64 {
		return types.SetIPv4Response{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.SetIPv4Response{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 150 {
		return types.SetIPv4Response{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.SetIPv4Response{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a get-status-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - GetStatusResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func GetStatusResponse(packet []byte) (types.GetStatusResponse, error) {
	if len(packet) != 64 {
		return types.GetStatusResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.GetStatusResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 32 {
		return types.GetStatusResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.GetStatusResponse{
		Controller:         unpackUint32(packet, 4),
		SystemDate:         unpackShortDate(packet, 51),
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
		EventTimestamp:     unpackOptionalDateTime(packet, 20),
		EventReason:        unpackUint8(packet, 27),
		SequenceNo:         unpackUint32(packet, 40),
	}, nil
}

// Decodes a get-time-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - GetTimeResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func GetTimeResponse(packet []byte) (types.GetTimeResponse, error) {
	if len(packet) != 64 {
		return types.GetTimeResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.GetTimeResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 50 {
		return types.GetTimeResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.GetTimeResponse{
		Controller: unpackUint32(packet, 4),
		DateTime:   unpackDateTime(packet, 8),
	}, nil
}

// Decodes a set-time-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - SetTimeResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func SetTimeResponse(packet []byte) (types.SetTimeResponse, error) {
	if len(packet) != 64 {
		return types.SetTimeResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.SetTimeResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 48 {
		return types.SetTimeResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.SetTimeResponse{
		Controller: unpackUint32(packet, 4),
		DateTime:   unpackDateTime(packet, 8),
	}, nil
}

// Decodes a get-listener-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - GetListenerResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func GetListenerResponse(packet []byte) (types.GetListenerResponse, error) {
	if len(packet) != 64 {
		return types.GetListenerResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.GetListenerResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 146 {
		return types.GetListenerResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.GetListenerResponse{
		Controller: unpackUint32(packet, 4),
		Listener:   unpackAddrPort(packet, 8),
		Interval:   unpackUint8(packet, 14),
	}, nil
}

// Decodes a set-listener-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - SetListenerResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func SetListenerResponse(packet []byte) (types.SetListenerResponse, error) {
	if len(packet) != 64 {
		return types.SetListenerResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.SetListenerResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 144 {
		return types.SetListenerResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.SetListenerResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a get-door-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - GetDoorResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func GetDoorResponse(packet []byte) (types.GetDoorResponse, error) {
	if len(packet) != 64 {
		return types.GetDoorResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.GetDoorResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 130 {
		return types.GetDoorResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.GetDoorResponse{
		Controller: unpackUint32(packet, 4),
		Door:       unpackUint8(packet, 8),
		Mode:       unpackUint8(packet, 9),
		Delay:      unpackUint8(packet, 10),
	}, nil
}

// Decodes a set-door-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - SetDoorResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func SetDoorResponse(packet []byte) (types.SetDoorResponse, error) {
	if len(packet) != 64 {
		return types.SetDoorResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.SetDoorResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 128 {
		return types.SetDoorResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.SetDoorResponse{
		Controller: unpackUint32(packet, 4),
		Door:       unpackUint8(packet, 8),
		Mode:       unpackUint8(packet, 9),
		Delay:      unpackUint8(packet, 10),
	}, nil
}

// Decodes a set-door-passcodes-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - SetDoorPasscodesResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func SetDoorPasscodesResponse(packet []byte) (types.SetDoorPasscodesResponse, error) {
	if len(packet) != 64 {
		return types.SetDoorPasscodesResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.SetDoorPasscodesResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 140 {
		return types.SetDoorPasscodesResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.SetDoorPasscodesResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a open-door-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - OpenDoorResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func OpenDoorResponse(packet []byte) (types.OpenDoorResponse, error) {
	if len(packet) != 64 {
		return types.OpenDoorResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.OpenDoorResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 64 {
		return types.OpenDoorResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.OpenDoorResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a get-cards-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - GetCardsResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func GetCardsResponse(packet []byte) (types.GetCardsResponse, error) {
	if len(packet) != 64 {
		return types.GetCardsResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.GetCardsResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 88 {
		return types.GetCardsResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.GetCardsResponse{
		Controller: unpackUint32(packet, 4),
		Cards:      unpackUint32(packet, 8),
	}, nil
}

// Decodes a get-card-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - GetCardResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func GetCardResponse(packet []byte) (types.GetCardResponse, error) {
	if len(packet) != 64 {
		return types.GetCardResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.GetCardResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 90 {
		return types.GetCardResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.GetCardResponse{
		Controller: unpackUint32(packet, 4),
		Card:       unpackUint32(packet, 8),
		StartDate:  unpackOptionalDate(packet, 12),
		EndDate:    unpackOptionalDate(packet, 16),
		Door1:      unpackUint8(packet, 20),
		Door2:      unpackUint8(packet, 21),
		Door3:      unpackUint8(packet, 22),
		Door4:      unpackUint8(packet, 23),
		PIN:        unpackPIN(packet, 24),
	}, nil
}

// Decodes a get-card-at-index-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - GetCardAtIndexResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func GetCardAtIndexResponse(packet []byte) (types.GetCardAtIndexResponse, error) {
	if len(packet) != 64 {
		return types.GetCardAtIndexResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.GetCardAtIndexResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 92 {
		return types.GetCardAtIndexResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.GetCardAtIndexResponse{
		Controller: unpackUint32(packet, 4),
		Card:       unpackUint32(packet, 8),
		StartDate:  unpackOptionalDate(packet, 12),
		EndDate:    unpackOptionalDate(packet, 16),
		Door1:      unpackUint8(packet, 20),
		Door2:      unpackUint8(packet, 21),
		Door3:      unpackUint8(packet, 22),
		Door4:      unpackUint8(packet, 23),
		PIN:        unpackPIN(packet, 24),
	}, nil
}

// Decodes a put-card-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - PutCardResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func PutCardResponse(packet []byte) (types.PutCardResponse, error) {
	if len(packet) != 64 {
		return types.PutCardResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.PutCardResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 80 {
		return types.PutCardResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.PutCardResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a delete-card-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - DeleteCardResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func DeleteCardResponse(packet []byte) (types.DeleteCardResponse, error) {
	if len(packet) != 64 {
		return types.DeleteCardResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.DeleteCardResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 82 {
		return types.DeleteCardResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.DeleteCardResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a delete-all-cards-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - DeleteAllCardsResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func DeleteAllCardsResponse(packet []byte) (types.DeleteAllCardsResponse, error) {
	if len(packet) != 64 {
		return types.DeleteAllCardsResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.DeleteAllCardsResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 84 {
		return types.DeleteAllCardsResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.DeleteAllCardsResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a get-event-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - GetEventResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func GetEventResponse(packet []byte) (types.GetEventResponse, error) {
	if len(packet) != 64 {
		return types.GetEventResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.GetEventResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 176 {
		return types.GetEventResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.GetEventResponse{
		Controller:    unpackUint32(packet, 4),
		Index:         unpackUint32(packet, 8),
		EventType:     unpackUint8(packet, 12),
		AccessGranted: unpackBool(packet, 13),
		Door:          unpackUint8(packet, 14),
		Direction:     unpackUint8(packet, 15),
		Card:          unpackUint32(packet, 16),
		Timestamp:     unpackOptionalDateTime(packet, 20),
		Reason:        unpackUint8(packet, 27),
	}, nil
}

// Decodes a get-event-index-response response.
//
//	Parameters:
//	    packet  (bytearray)  64 byte UDP packet.
//
//	Returns:
//	    - GetEventIndexResponse initialised from the UDP packet.
//	    - error if the packet is not 64 bytes, has an invalid start-of-message byte or has
//	               the incorrect message type.
func GetEventIndexResponse(packet []byte) (types.GetEventIndexResponse, error) {
	if len(packet) != 64 {
		return types.GetEventIndexResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return types.GetEventIndexResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 180 {
		return types.GetEventIndexResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return types.GetEventIndexResponse{
		Controller: unpackUint32(packet, 4),
		Index:      unpackUint32(packet, 8),
	}, nil
}
