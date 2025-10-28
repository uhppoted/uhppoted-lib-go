// generated code - ** DO NOT EDIT **

package decode

import (
	"fmt"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"
)

// Decodes a GetControllerResponse from a 64 byte response packet.
func GetControllerResponse(packet []byte) (responses.GetController, error) {
	if len(packet) != 64 {
		return responses.GetController{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetController{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x94 {
		return responses.GetController{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetController{
		Controller: unpackUint32(packet, 4),
		IpAddress:  unpackIPv4(packet, 8),
		SubnetMask: unpackIPv4(packet, 12),
		Gateway:    unpackIPv4(packet, 16),
		MACAddress: unpackMAC(packet, 20),
		Version:    unpackVersion(packet, 26),
		Date:       unpackDate(packet, 28),
	}, nil
}

// Decodes a SetIPv4Response from a 64 byte response packet.
func SetIPv4Response(packet []byte) (responses.SetIPv4, error) {
	if len(packet) != 64 {
		return responses.SetIPv4{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetIPv4{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x96 {
		return responses.SetIPv4{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetIPv4{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a GetStatusResponse from a 64 byte response packet.
func GetStatusResponse(packet []byte) (responses.GetStatus, error) {
	if len(packet) != 64 {
		return responses.GetStatus{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetStatus{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x20 {
		return responses.GetStatus{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetStatus{
		Controller:         unpackUint32(packet, 4),
		SystemDate:         unpackShortDate(packet, 51),
		SystemTime:         unpackTime(packet, 37),
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
		EventType:          unpackEventType(packet, 12),
		EventAccessGranted: unpackBool(packet, 13),
		EventDoor:          unpackUint8(packet, 14),
		EventDirection:     unpackDirection(packet, 15),
		EventCard:          unpackUint32(packet, 16),
		EventTimestamp:     unpackOptionalDateTime(packet, 20),
		EventReason:        unpackReason(packet, 27),
		SequenceNo:         unpackUint32(packet, 40),
	}, nil
}

// Decodes a GetTimeResponse from a 64 byte response packet.
func GetTimeResponse(packet []byte) (responses.GetTime, error) {
	if len(packet) != 64 {
		return responses.GetTime{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetTime{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x32 {
		return responses.GetTime{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetTime{
		Controller: unpackUint32(packet, 4),
		DateTime:   unpackDateTime(packet, 8),
	}, nil
}

// Decodes a SetTimeResponse from a 64 byte response packet.
func SetTimeResponse(packet []byte) (responses.SetTime, error) {
	if len(packet) != 64 {
		return responses.SetTime{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetTime{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x30 {
		return responses.SetTime{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetTime{
		Controller: unpackUint32(packet, 4),
		DateTime:   unpackDateTime(packet, 8),
	}, nil
}

// Decodes a GetListenerResponse from a 64 byte response packet.
func GetListenerResponse(packet []byte) (responses.GetListener, error) {
	if len(packet) != 64 {
		return responses.GetListener{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetListener{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x92 {
		return responses.GetListener{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetListener{
		Controller: unpackUint32(packet, 4),
		Address:    unpackIPv4(packet, 8),
		Port:       unpackUint16(packet, 12),
		Interval:   unpackUint8(packet, 14),
	}, nil
}

// Decodes a SetListenerResponse from a 64 byte response packet.
func SetListenerResponse(packet []byte) (responses.SetListener, error) {
	if len(packet) != 64 {
		return responses.SetListener{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetListener{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x90 {
		return responses.SetListener{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetListener{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a GetListenerAddrPortResponse from a 64 byte response packet.
func GetListenerAddrPortResponse(packet []byte) (responses.GetListenerAddrPort, error) {
	if len(packet) != 64 {
		return responses.GetListenerAddrPort{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetListenerAddrPort{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x92 {
		return responses.GetListenerAddrPort{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetListenerAddrPort{
		Controller: unpackUint32(packet, 4),
		Listener:   unpackAddrPort(packet, 8),
		Interval:   unpackUint8(packet, 14),
	}, nil
}

// Decodes a SetListenerAddrPortResponse from a 64 byte response packet.
func SetListenerAddrPortResponse(packet []byte) (responses.SetListenerAddrPort, error) {
	if len(packet) != 64 {
		return responses.SetListenerAddrPort{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetListenerAddrPort{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x90 {
		return responses.SetListenerAddrPort{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetListenerAddrPort{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a GetDoorResponse from a 64 byte response packet.
func GetDoorResponse(packet []byte) (responses.GetDoor, error) {
	if len(packet) != 64 {
		return responses.GetDoor{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetDoor{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x82 {
		return responses.GetDoor{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetDoor{
		Controller: unpackUint32(packet, 4),
		Door:       unpackUint8(packet, 8),
		Mode:       unpackMode(packet, 9),
		Delay:      unpackUint8(packet, 10),
	}, nil
}

// Decodes a SetDoorResponse from a 64 byte response packet.
func SetDoorResponse(packet []byte) (responses.SetDoor, error) {
	if len(packet) != 64 {
		return responses.SetDoor{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetDoor{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x80 {
		return responses.SetDoor{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetDoor{
		Controller: unpackUint32(packet, 4),
		Door:       unpackUint8(packet, 8),
		Mode:       unpackUint8(packet, 9),
		Delay:      unpackUint8(packet, 10),
	}, nil
}

// Decodes a SetDoorPasscodesResponse from a 64 byte response packet.
func SetDoorPasscodesResponse(packet []byte) (responses.SetDoorPasscodes, error) {
	if len(packet) != 64 {
		return responses.SetDoorPasscodes{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetDoorPasscodes{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x8c {
		return responses.SetDoorPasscodes{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetDoorPasscodes{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a OpenDoorResponse from a 64 byte response packet.
func OpenDoorResponse(packet []byte) (responses.OpenDoor, error) {
	if len(packet) != 64 {
		return responses.OpenDoor{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.OpenDoor{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x40 {
		return responses.OpenDoor{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.OpenDoor{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a GetCardsResponse from a 64 byte response packet.
func GetCardsResponse(packet []byte) (responses.GetCards, error) {
	if len(packet) != 64 {
		return responses.GetCards{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetCards{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x58 {
		return responses.GetCards{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetCards{
		Controller: unpackUint32(packet, 4),
		Cards:      unpackUint32(packet, 8),
	}, nil
}

// Decodes a GetCardResponse from a 64 byte response packet.
func GetCardResponse(packet []byte) (responses.GetCard, error) {
	if len(packet) != 64 {
		return responses.GetCard{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetCard{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x5a {
		return responses.GetCard{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetCard{
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

// Decodes a GetCardAtIndexResponse from a 64 byte response packet.
func GetCardAtIndexResponse(packet []byte) (responses.GetCardAtIndex, error) {
	if len(packet) != 64 {
		return responses.GetCardAtIndex{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetCardAtIndex{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x5c {
		return responses.GetCardAtIndex{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetCardAtIndex{
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

// Decodes a PutCardResponse from a 64 byte response packet.
func PutCardResponse(packet []byte) (responses.PutCard, error) {
	if len(packet) != 64 {
		return responses.PutCard{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.PutCard{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x50 {
		return responses.PutCard{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.PutCard{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a DeleteCardResponse from a 64 byte response packet.
func DeleteCardResponse(packet []byte) (responses.DeleteCard, error) {
	if len(packet) != 64 {
		return responses.DeleteCard{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.DeleteCard{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x52 {
		return responses.DeleteCard{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.DeleteCard{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a DeleteAllCardsResponse from a 64 byte response packet.
func DeleteAllCardsResponse(packet []byte) (responses.DeleteAllCards, error) {
	if len(packet) != 64 {
		return responses.DeleteAllCards{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.DeleteAllCards{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x54 {
		return responses.DeleteAllCards{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.DeleteAllCards{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a GetEventResponse from a 64 byte response packet.
func GetEventResponse(packet []byte) (responses.GetEvent, error) {
	if len(packet) != 64 {
		return responses.GetEvent{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetEvent{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xb0 {
		return responses.GetEvent{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetEvent{
		Controller:    unpackUint32(packet, 4),
		Index:         unpackUint32(packet, 8),
		EventType:     unpackEventType(packet, 12),
		AccessGranted: unpackBool(packet, 13),
		Door:          unpackUint8(packet, 14),
		Direction:     unpackDirection(packet, 15),
		Card:          unpackUint32(packet, 16),
		Timestamp:     unpackOptionalDateTime(packet, 20),
		Reason:        unpackReason(packet, 27),
	}, nil
}

// Decodes a GetEventIndexResponse from a 64 byte response packet.
func GetEventIndexResponse(packet []byte) (responses.GetEventIndex, error) {
	if len(packet) != 64 {
		return responses.GetEventIndex{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetEventIndex{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xb4 {
		return responses.GetEventIndex{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetEventIndex{
		Controller: unpackUint32(packet, 4),
		Index:      unpackUint32(packet, 8),
	}, nil
}

// Decodes a SetEventIndexResponse from a 64 byte response packet.
func SetEventIndexResponse(packet []byte) (responses.SetEventIndex, error) {
	if len(packet) != 64 {
		return responses.SetEventIndex{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetEventIndex{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xb2 {
		return responses.SetEventIndex{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetEventIndex{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a RecordSpecialEventsResponse from a 64 byte response packet.
func RecordSpecialEventsResponse(packet []byte) (responses.RecordSpecialEvents, error) {
	if len(packet) != 64 {
		return responses.RecordSpecialEvents{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.RecordSpecialEvents{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x8e {
		return responses.RecordSpecialEvents{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.RecordSpecialEvents{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a GetTimeProfileResponse from a 64 byte response packet.
func GetTimeProfileResponse(packet []byte) (responses.GetTimeProfile, error) {
	if len(packet) != 64 {
		return responses.GetTimeProfile{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetTimeProfile{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x98 {
		return responses.GetTimeProfile{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetTimeProfile{
		Controller:    unpackUint32(packet, 4),
		Profile:       unpackUint8(packet, 8),
		StartDate:     unpackOptionalDate(packet, 9),
		EndDate:       unpackOptionalDate(packet, 13),
		Monday:        unpackBool(packet, 17),
		Tuesday:       unpackBool(packet, 18),
		Wednesday:     unpackBool(packet, 19),
		Thursday:      unpackBool(packet, 20),
		Friday:        unpackBool(packet, 21),
		Saturday:      unpackBool(packet, 22),
		Sunday:        unpackBool(packet, 23),
		Segment1Start: unpackHHmm(packet, 24),
		Segment1End:   unpackHHmm(packet, 26),
		Segment2Start: unpackHHmm(packet, 28),
		Segment2End:   unpackHHmm(packet, 30),
		Segment3Start: unpackHHmm(packet, 32),
		Segment3End:   unpackHHmm(packet, 34),
		LinkedProfile: unpackUint8(packet, 36),
	}, nil
}

// Decodes a SetTimeProfileResponse from a 64 byte response packet.
func SetTimeProfileResponse(packet []byte) (responses.SetTimeProfile, error) {
	if len(packet) != 64 {
		return responses.SetTimeProfile{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetTimeProfile{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x88 {
		return responses.SetTimeProfile{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetTimeProfile{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a ClearTimeProfilesResponse from a 64 byte response packet.
func ClearTimeProfilesResponse(packet []byte) (responses.ClearTimeProfiles, error) {
	if len(packet) != 64 {
		return responses.ClearTimeProfiles{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.ClearTimeProfiles{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x8a {
		return responses.ClearTimeProfiles{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.ClearTimeProfiles{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a AddTaskResponse from a 64 byte response packet.
func AddTaskResponse(packet []byte) (responses.AddTask, error) {
	if len(packet) != 64 {
		return responses.AddTask{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.AddTask{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xa8 {
		return responses.AddTask{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.AddTask{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a RefreshTaskListResponse from a 64 byte response packet.
func RefreshTaskListResponse(packet []byte) (responses.RefreshTaskList, error) {
	if len(packet) != 64 {
		return responses.RefreshTaskList{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.RefreshTaskList{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xac {
		return responses.RefreshTaskList{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.RefreshTaskList{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a ClearTaskListResponse from a 64 byte response packet.
func ClearTaskListResponse(packet []byte) (responses.ClearTaskList, error) {
	if len(packet) != 64 {
		return responses.ClearTaskList{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.ClearTaskList{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xa6 {
		return responses.ClearTaskList{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.ClearTaskList{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a SetPCControlResponse from a 64 byte response packet.
func SetPCControlResponse(packet []byte) (responses.SetPCControl, error) {
	if len(packet) != 64 {
		return responses.SetPCControl{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetPCControl{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xa0 {
		return responses.SetPCControl{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetPCControl{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a SetInterlockResponse from a 64 byte response packet.
func SetInterlockResponse(packet []byte) (responses.SetInterlock, error) {
	if len(packet) != 64 {
		return responses.SetInterlock{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetInterlock{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xa2 {
		return responses.SetInterlock{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetInterlock{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a ActivateKeypadsResponse from a 64 byte response packet.
func ActivateKeypadsResponse(packet []byte) (responses.ActivateKeypads, error) {
	if len(packet) != 64 {
		return responses.ActivateKeypads{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.ActivateKeypads{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xa4 {
		return responses.ActivateKeypads{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.ActivateKeypads{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a GetAntiPassbackResponse from a 64 byte response packet.
func GetAntiPassbackResponse(packet []byte) (responses.GetAntiPassback, error) {
	if len(packet) != 64 {
		return responses.GetAntiPassback{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetAntiPassback{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x86 {
		return responses.GetAntiPassback{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetAntiPassback{
		Controller:   unpackUint32(packet, 4),
		Antipassback: unpackAntiPassback(packet, 8),
	}, nil
}

// Decodes a SetAntiPassbackResponse from a 64 byte response packet.
func SetAntiPassbackResponse(packet []byte) (responses.SetAntiPassback, error) {
	if len(packet) != 64 {
		return responses.SetAntiPassback{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetAntiPassback{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x84 {
		return responses.SetAntiPassback{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetAntiPassback{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a RestoreDefaultParametersResponse from a 64 byte response packet.
func RestoreDefaultParametersResponse(packet []byte) (responses.RestoreDefaultParameters, error) {
	if len(packet) != 64 {
		return responses.RestoreDefaultParameters{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.RestoreDefaultParameters{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xc8 {
		return responses.RestoreDefaultParameters{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.RestoreDefaultParameters{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

// Decodes a ListenerEvent from a 64 byte response packet.
func ListenerEvent(packet []byte) (responses.ListenerEvent, error) {
	if len(packet) != 64 {
		return responses.ListenerEvent{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.ListenerEvent{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x20 {
		return responses.ListenerEvent{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.ListenerEvent{
		Controller:         unpackUint32(packet, 4),
		SystemDate:         unpackShortDate(packet, 51),
		SystemTime:         unpackTime(packet, 37),
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
		EventType:          unpackEventType(packet, 12),
		EventAccessGranted: unpackBool(packet, 13),
		EventDoor:          unpackUint8(packet, 14),
		EventDirection:     unpackDirection(packet, 15),
		EventCard:          unpackUint32(packet, 16),
		EventTimestamp:     unpackOptionalDateTime(packet, 20),
		EventReason:        unpackReason(packet, 27),
		SequenceNo:         unpackUint32(packet, 40),
	}, nil
}
