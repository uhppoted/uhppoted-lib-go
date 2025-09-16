package responses

import (
	"fmt"
	
	"github.com/uhppoted/uhppoted-lib-go/uhppoted/responses"
)

func GetControllerResponse(packet []byte) (responses.GetControllerResponse, error) {
	if len(packet) != 64 {
		return responses.GetControllerResponse(), fmt.Errorf("invalid reply packet length (%v)", len(packet))
	}
	
	if packet[0] != SOM && (packet[0] != SOM_v6_62 || packet[1] != 0x20) {
		return responses.GetControllerResponse(), fmt.Errorf("invalid reply start of message byte (%02x)", packet[0])
	}
	
	if packet[1] != 0x94 {
		return responses.GetControllerResponse(), fmt.Errorf("invalid reply function code (%02x)", packet[1])
	}
	
	return responses.GetControllerResponse{Controller: unpackUint32(packet, 4)}, nil
}
