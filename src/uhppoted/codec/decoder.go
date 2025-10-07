package codec

import (
	"fmt"

	decoder "github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/decode"
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/responses"
)

type key[R any] struct {
	opcode byte
}

var decoders = map[any]func([]byte) (any, error){
	key[responses.GetListenerAddrPortResponse]{0x92}: func(b []byte) (any, error) {
		return decoder.GetListenerAddrPortResponse(b)
	},
	key[responses.SetListenerAddrPortResponse]{0x90}: func(b []byte) (any, error) {
		return decoder.SetListenerAddrPortResponse(b)
	},
}

func Decode[R any](packet []byte) (R, error) {
	var zero R

	if len(packet) != 64 {
		return zero, fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}

	if packet[0] != SOM {
		return zero, fmt.Errorf("invalid reply SOM byte (0x%02x)", packet[0])
	}

	k := key[R]{packet[1]}
	fn := decode
	if f, ok := decoders[k]; ok {
		fn = f
	}

	if v, err := fn(packet); err != nil {
		return zero, err
	} else if response, ok := v.(R); !ok {
		return zero, err
	} else {
		return response, nil
	}
}
