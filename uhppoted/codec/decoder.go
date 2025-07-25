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

	default:
		return nil, fmt.Errorf("unknown message type (%02x)", packet[1])
	}
}
