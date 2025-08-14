package codec

import (
	"fmt"

	decoder "github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/decode"
)

func Decode[R any](packet []byte) (R, error) {
	var zero R

	if v, err := decode(packet); err != nil {
		return zero, fmt.Errorf("invalid packet")
	} else if response, ok := v.(R); !ok {
		return zero, fmt.Errorf("invalid packet")
	} else {
		return response, nil
	}
}

func decode(packet []byte) (any, error) {
	if len(packet) != 64 {
		return nil, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return nil, fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}

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

	default:
		return nil, fmt.Errorf("unknown message type (%02x)", packet[1])
	}
}
