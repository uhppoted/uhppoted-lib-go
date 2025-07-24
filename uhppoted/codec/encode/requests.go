// generated code - ** DO NOT EDIT **

// UHPPOTE request packet encoder.
//
// Encodes a UHPPOTE access controller request as a 64 byte UDP packet:
//
// - uint8, uint16, uint24 and uint32 values are encoded as little endian unsigned integers
// - datetime, date and time values are encoded as BCD
// - boolean values are encoded as 0 (False) or 1 (True)
package encode

import (
    "net/netip"
    "time"
)


// Encodes a get-controller request.
//
//  Parameters:
//      controller  (uint32)  controller serial number
//
//  Returns:
//      64 byte packet.
func GetControllerRequest(controller uint32) ([]byte, error) {
    packet := make([]byte, 64)

    packet[0] = SOM
    packet[1] = 148

    packUint32(controller, packet, 4)

    return packet, nil
}

// Encodes a set-ipv4 request.
//
//  Parameters:
//      controller  (uint32)  controller serial number
//      address  (IPv4)  controller IPv4 address
//      netmask  (IPv4)  controller IPv4 subnet mask
//      gateway  (IPv4)  controller IPv4 gateway address
//        (magic)  'magic' word
//
//  Returns:
//      64 byte packet.
func SetIPv4Request(controller uint32, address netip.Addr, netmask netip.Addr, gateway netip.Addr) ([]byte, error) {
    packet := make([]byte, 64)

    packet[0] = SOM
    packet[1] = 150

    packUint32(controller, packet, 4)
    packIPv4(address, packet, 8)
    packIPv4(netmask, packet, 12)
    packIPv4(gateway, packet, 16)
    packUint32(0x55aaaa55, packet, 20)

    return packet, nil
}

// Encodes a get-status request.
//
//  Parameters:
//      controller  (uint32)  controller serial number
//
//  Returns:
//      64 byte packet.
func GetStatusRequest(controller uint32) ([]byte, error) {
    packet := make([]byte, 64)

    packet[0] = SOM
    packet[1] = 32

    packUint32(controller, packet, 4)

    return packet, nil
}

// Encodes a get-time request.
//
//  Parameters:
//      controller  (uint32)  controller serial number
//
//  Returns:
//      64 byte packet.
func GetTimeRequest(controller uint32) ([]byte, error) {
    packet := make([]byte, 64)

    packet[0] = SOM
    packet[1] = 50

    packUint32(controller, packet, 4)

    return packet, nil
}

// Encodes a set-time request.
//
//  Parameters:
//      controller  (uint32)  controller serial number
//      datetime  (datetime)  controller serial number
//
//  Returns:
//      64 byte packet.
func SetTimeRequest(controller uint32, datetime time.Time) ([]byte, error) {
    packet := make([]byte, 64)

    packet[0] = SOM
    packet[1] = 48

    packUint32(controller, packet, 4)
    packDateTime(datetime, packet, 8)

    return packet, nil
}

// Encodes a get-listener request.
//
//  Parameters:
//      controller  (uint32)  controller serial number
//
//  Returns:
//      64 byte packet.
func GetListenerRequest(controller uint32) ([]byte, error) {
    packet := make([]byte, 64)

    packet[0] = SOM
    packet[1] = 146

    packUint32(controller, packet, 4)

    return packet, nil
}
