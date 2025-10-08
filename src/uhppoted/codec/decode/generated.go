// generated code - ** DO NOT EDIT **

package decode

import (
	"fmt"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/responses"
)

func GetControllerResponse(packet []byte) (responses.GetControllerResponse, error) {
	if len(packet) != 64 {
		return responses.GetControllerResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetControllerResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x94 {
		return responses.GetControllerResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetControllerResponse{
		Controller: unpackUint32(packet, 4),
		IpAddress:  unpackIPv4(packet, 8),
		SubnetMask: unpackIPv4(packet, 12),
		Gateway:    unpackIPv4(packet, 16),
		MACAddress: unpackMAC(packet, 20),
		Version:    unpackVersion(packet, 26),
		Date:       unpackDate(packet, 28),
	}, nil
}

func SetIPv4Response(packet []byte) (responses.SetIPv4Response, error) {
	if len(packet) != 64 {
		return responses.SetIPv4Response{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetIPv4Response{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x96 {
		return responses.SetIPv4Response{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetIPv4Response{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func GetStatusResponse(packet []byte) (responses.GetStatusResponse, error) {
	if len(packet) != 64 {
		return responses.GetStatusResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetStatusResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x20 {
		return responses.GetStatusResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetStatusResponse{
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

func GetTimeResponse(packet []byte) (responses.GetTimeResponse, error) {
	if len(packet) != 64 {
		return responses.GetTimeResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetTimeResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x32 {
		return responses.GetTimeResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetTimeResponse{
		Controller: unpackUint32(packet, 4),
		DateTime:   unpackDateTime(packet, 8),
	}, nil
}

func SetTimeResponse(packet []byte) (responses.SetTimeResponse, error) {
	if len(packet) != 64 {
		return responses.SetTimeResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetTimeResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x30 {
		return responses.SetTimeResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetTimeResponse{
		Controller: unpackUint32(packet, 4),
		DateTime:   unpackDateTime(packet, 8),
	}, nil
}

func GetListenerResponse(packet []byte) (responses.GetListenerResponse, error) {
	if len(packet) != 64 {
		return responses.GetListenerResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetListenerResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x92 {
		return responses.GetListenerResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetListenerResponse{
		Controller: unpackUint32(packet, 4),
		Address:    unpackIPv4(packet, 8),
		Port:       unpackUint16(packet, 12),
		Interval:   unpackUint8(packet, 14),
	}, nil
}

func SetListenerResponse(packet []byte) (responses.SetListenerResponse, error) {
	if len(packet) != 64 {
		return responses.SetListenerResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetListenerResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x90 {
		return responses.SetListenerResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetListenerResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func GetListenerAddrPortResponse(packet []byte) (responses.GetListenerAddrPortResponse, error) {
	if len(packet) != 64 {
		return responses.GetListenerAddrPortResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetListenerAddrPortResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x92 {
		return responses.GetListenerAddrPortResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetListenerAddrPortResponse{
		Controller: unpackUint32(packet, 4),
		Listener:   unpackAddrPort(packet, 8),
		Interval:   unpackUint8(packet, 14),
	}, nil
}

func SetListenerAddrPortResponse(packet []byte) (responses.SetListenerAddrPortResponse, error) {
	if len(packet) != 64 {
		return responses.SetListenerAddrPortResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetListenerAddrPortResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x90 {
		return responses.SetListenerAddrPortResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetListenerAddrPortResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func GetDoorResponse(packet []byte) (responses.GetDoorResponse, error) {
	if len(packet) != 64 {
		return responses.GetDoorResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetDoorResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x82 {
		return responses.GetDoorResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetDoorResponse{
		Controller: unpackUint32(packet, 4),
		Door:       unpackUint8(packet, 8),
		Mode:       unpackMode(packet, 9),
		Delay:      unpackUint8(packet, 10),
	}, nil
}

func SetDoorResponse(packet []byte) (responses.SetDoorResponse, error) {
	if len(packet) != 64 {
		return responses.SetDoorResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetDoorResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x80 {
		return responses.SetDoorResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetDoorResponse{
		Controller: unpackUint32(packet, 4),
		Door:       unpackUint8(packet, 8),
		Mode:       unpackUint8(packet, 9),
		Delay:      unpackUint8(packet, 10),
	}, nil
}

func SetDoorPasscodesResponse(packet []byte) (responses.SetDoorPasscodesResponse, error) {
	if len(packet) != 64 {
		return responses.SetDoorPasscodesResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetDoorPasscodesResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x8c {
		return responses.SetDoorPasscodesResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetDoorPasscodesResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func OpenDoorResponse(packet []byte) (responses.OpenDoorResponse, error) {
	if len(packet) != 64 {
		return responses.OpenDoorResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.OpenDoorResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x40 {
		return responses.OpenDoorResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.OpenDoorResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func GetCardsResponse(packet []byte) (responses.GetCardsResponse, error) {
	if len(packet) != 64 {
		return responses.GetCardsResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetCardsResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x58 {
		return responses.GetCardsResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetCardsResponse{
		Controller: unpackUint32(packet, 4),
		Cards:      unpackUint32(packet, 8),
	}, nil
}

func GetCardResponse(packet []byte) (responses.GetCardResponse, error) {
	if len(packet) != 64 {
		return responses.GetCardResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetCardResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x5a {
		return responses.GetCardResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetCardResponse{
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

func GetCardAtIndexResponse(packet []byte) (responses.GetCardAtIndexResponse, error) {
	if len(packet) != 64 {
		return responses.GetCardAtIndexResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetCardAtIndexResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x5c {
		return responses.GetCardAtIndexResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetCardAtIndexResponse{
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

func PutCardResponse(packet []byte) (responses.PutCardResponse, error) {
	if len(packet) != 64 {
		return responses.PutCardResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.PutCardResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x50 {
		return responses.PutCardResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.PutCardResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func DeleteCardResponse(packet []byte) (responses.DeleteCardResponse, error) {
	if len(packet) != 64 {
		return responses.DeleteCardResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.DeleteCardResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x52 {
		return responses.DeleteCardResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.DeleteCardResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func DeleteAllCardsResponse(packet []byte) (responses.DeleteAllCardsResponse, error) {
	if len(packet) != 64 {
		return responses.DeleteAllCardsResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.DeleteAllCardsResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x54 {
		return responses.DeleteAllCardsResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.DeleteAllCardsResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func GetEventResponse(packet []byte) (responses.GetEventResponse, error) {
	if len(packet) != 64 {
		return responses.GetEventResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetEventResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xb0 {
		return responses.GetEventResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetEventResponse{
		Controller:    unpackUint32(packet, 4),
		Index:         unpackUint32(packet, 8),
		EventType:     unpackEventType(packet, 12),
		AccessGranted: unpackBool(packet, 13),
		Door:          unpackUint8(packet, 14),
		Direction:     unpackUint8(packet, 15),
		Card:          unpackUint32(packet, 16),
		Timestamp:     unpackOptionalDateTime(packet, 20),
		Reason:        unpackUint8(packet, 27),
	}, nil
}

func GetEventIndexResponse(packet []byte) (responses.GetEventIndexResponse, error) {
	if len(packet) != 64 {
		return responses.GetEventIndexResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetEventIndexResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xb4 {
		return responses.GetEventIndexResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetEventIndexResponse{
		Controller: unpackUint32(packet, 4),
		Index:      unpackUint32(packet, 8),
	}, nil
}

func SetEventIndexResponse(packet []byte) (responses.SetEventIndexResponse, error) {
	if len(packet) != 64 {
		return responses.SetEventIndexResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetEventIndexResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xb2 {
		return responses.SetEventIndexResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetEventIndexResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func RecordSpecialEventsResponse(packet []byte) (responses.RecordSpecialEventsResponse, error) {
	if len(packet) != 64 {
		return responses.RecordSpecialEventsResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.RecordSpecialEventsResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x8e {
		return responses.RecordSpecialEventsResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.RecordSpecialEventsResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func GetTimeProfileResponse(packet []byte) (responses.GetTimeProfileResponse, error) {
	if len(packet) != 64 {
		return responses.GetTimeProfileResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetTimeProfileResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x98 {
		return responses.GetTimeProfileResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetTimeProfileResponse{
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

func SetTimeProfileResponse(packet []byte) (responses.SetTimeProfileResponse, error) {
	if len(packet) != 64 {
		return responses.SetTimeProfileResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetTimeProfileResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x88 {
		return responses.SetTimeProfileResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetTimeProfileResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func ClearTimeProfilesResponse(packet []byte) (responses.ClearTimeProfilesResponse, error) {
	if len(packet) != 64 {
		return responses.ClearTimeProfilesResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.ClearTimeProfilesResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x8a {
		return responses.ClearTimeProfilesResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.ClearTimeProfilesResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func AddTaskResponse(packet []byte) (responses.AddTaskResponse, error) {
	if len(packet) != 64 {
		return responses.AddTaskResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.AddTaskResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xa8 {
		return responses.AddTaskResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.AddTaskResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func RefreshTaskListResponse(packet []byte) (responses.RefreshTaskListResponse, error) {
	if len(packet) != 64 {
		return responses.RefreshTaskListResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.RefreshTaskListResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xac {
		return responses.RefreshTaskListResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.RefreshTaskListResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func ClearTaskListResponse(packet []byte) (responses.ClearTaskListResponse, error) {
	if len(packet) != 64 {
		return responses.ClearTaskListResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.ClearTaskListResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xa6 {
		return responses.ClearTaskListResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.ClearTaskListResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func SetPCControlResponse(packet []byte) (responses.SetPCControlResponse, error) {
	if len(packet) != 64 {
		return responses.SetPCControlResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetPCControlResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xa0 {
		return responses.SetPCControlResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetPCControlResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func SetInterlockResponse(packet []byte) (responses.SetInterlockResponse, error) {
	if len(packet) != 64 {
		return responses.SetInterlockResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetInterlockResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xa2 {
		return responses.SetInterlockResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetInterlockResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func ActivateKeypadsResponse(packet []byte) (responses.ActivateKeypadsResponse, error) {
	if len(packet) != 64 {
		return responses.ActivateKeypadsResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.ActivateKeypadsResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xa4 {
		return responses.ActivateKeypadsResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.ActivateKeypadsResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func GetAntiPassbackResponse(packet []byte) (responses.GetAntiPassbackResponse, error) {
	if len(packet) != 64 {
		return responses.GetAntiPassbackResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetAntiPassbackResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x86 {
		return responses.GetAntiPassbackResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.GetAntiPassbackResponse{
		Controller:   unpackUint32(packet, 4),
		Antipassback: unpackUint8(packet, 8),
	}, nil
}

func SetAntiPassbackResponse(packet []byte) (responses.SetAntiPassbackResponse, error) {
	if len(packet) != 64 {
		return responses.SetAntiPassbackResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.SetAntiPassbackResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0x84 {
		return responses.SetAntiPassbackResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.SetAntiPassbackResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

func RestoreDefaultParametersResponse(packet []byte) (responses.RestoreDefaultParametersResponse, error) {
	if len(packet) != 64 {
		return responses.RestoreDefaultParametersResponse{}, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.RestoreDefaultParametersResponse{}, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

	if packet[1] != 0xc8 {
		return responses.RestoreDefaultParametersResponse{}, fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}

	return responses.RestoreDefaultParametersResponse{
		Controller: unpackUint32(packet, 4),
		Ok:         unpackBool(packet, 8),
	}, nil
}

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

	return responses.ListenerEvent{Controller: unpackUint32(packet, 4), SystemDate: unpackShortDate(packet, 51), SystemTime: unpackTime(packet, 37), Door1Open: unpackBool(packet, 28), Door2Open: unpackBool(packet, 29), Door3Open: unpackBool(packet, 30), Door4Open: unpackBool(packet, 31), Door1Button: unpackBool(packet, 32), Door2Button: unpackBool(packet, 33), Door3Button: unpackBool(packet, 34), Door4Button: unpackBool(packet, 35), Relays: unpackUint8(packet, 49), Inputs: unpackUint8(packet, 50), SystemError: unpackUint8(packet, 36), SpecialInfo: unpackUint8(packet, 48), EventIndex: unpackUint32(packet, 8), EventType: unpackUint8(packet, 12), EventAccessGranted: unpackBool(packet, 13), EventDoor: unpackUint8(packet, 14), EventDirection: unpackUint8(packet, 15), EventCard: unpackUint32(packet, 16), EventTimestamp: unpackOptionalDateTime(packet, 20), EventReason: unpackUint8(packet, 27), SequenceNo: unpackUint32(packet, 40)}, nil
}
