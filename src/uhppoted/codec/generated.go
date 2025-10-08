// generated code - ** DO NOT EDIT **

package codec

import (
	"fmt"

	decoder "github.com/uhppoted/uhppoted-lib-go/src/uhppoted/codec/decode"
)

func decode(packet []byte) (any, error) {
	switch packet[1] {
	case 0x94:
		return decoder.GetControllerResponse(packet)

	case 0x96:
		return decoder.SetIPv4Response(packet)

	case 0x20:
		return decoder.GetStatusResponse(packet)

	case 0x32:
		return decoder.GetTimeResponse(packet)

	case 0x30:
		return decoder.SetTimeResponse(packet)

	case 0x92:
		return decoder.GetListenerResponse(packet)

	case 0x90:
		return decoder.SetListenerResponse(packet)

	case 0x82:
		return decoder.GetDoorResponse(packet)

	case 0x80:
		return decoder.SetDoorResponse(packet)

	case 0x8c:
		return decoder.SetDoorPasscodesResponse(packet)

	case 0x40:
		return decoder.OpenDoorResponse(packet)

	case 0x58:
		return decoder.GetCardsResponse(packet)

	case 0x5a:
		return decoder.GetCardResponse(packet)

	case 0x5c:
		return decoder.GetCardAtIndexResponse(packet)

	case 0x50:
		return decoder.PutCardResponse(packet)

	case 0x52:
		return decoder.DeleteCardResponse(packet)

	case 0x54:
		return decoder.DeleteAllCardsResponse(packet)

	case 0xb0:
		return decoder.GetEventResponse(packet)

	case 0xb4:
		return decoder.GetEventIndexResponse(packet)

	case 0xb2:
		return decoder.SetEventIndexResponse(packet)

	case 0x8e:
		return decoder.RecordSpecialEventsResponse(packet)

	case 0x98:
		return decoder.GetTimeProfileResponse(packet)

	case 0x88:
		return decoder.SetTimeProfileResponse(packet)

	case 0x8a:
		return decoder.ClearTimeProfilesResponse(packet)

	case 0xa8:
		return decoder.AddTaskResponse(packet)

	case 0xac:
		return decoder.RefreshTaskListResponse(packet)

	case 0xa6:
		return decoder.ClearTaskListResponse(packet)

	case 0xa0:
		return decoder.SetPCControlResponse(packet)

	case 0xa2:
		return decoder.SetInterlockResponse(packet)

	case 0xa4:
		return decoder.ActivateKeypadsResponse(packet)

	case 0x86:
		return decoder.GetAntiPassbackResponse(packet)

	case 0x84:
		return decoder.SetAntiPassbackResponse(packet)

	case 0xc8:
		return decoder.RestoreDefaultParametersResponse(packet)

	default:
		return nil, fmt.Errorf("unknown message type (0x%02x)", packet[1])
	}
}
