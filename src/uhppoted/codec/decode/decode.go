// UHPPOTE response packet decoder.
//
// Decodes a UHPPOTE access controller 64 byte UDP response packet into the corresponding
// response struct.
package decode

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net/netip"
	"regexp"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/entities"
)

// Message constants for the UHPPOTE request/response protocol.
const SOM byte = 0x17
const SOM_v6_62 byte = 0x19

func unpackBool(packet []byte, offset uint8) bool {
	return packet[offset] != 0x00
}

func unpackUint8(packet []byte, offset uint8) uint8 {
	return packet[offset]
}

func unpackUint16(packet []byte, offset uint8) uint16 {
	return binary.LittleEndian.Uint16(packet[offset : offset+2])
}

func unpackUint32(packet []byte, offset uint8) uint32 {
	return binary.LittleEndian.Uint32(packet[offset : offset+4])
}

func unpackIPv4(packet []byte, offset uint8) netip.Addr {
	addr, _ := netip.AddrFromSlice(packet[offset : offset+4])

	return addr
}

func unpackAddrPort(packet []byte, offset uint8) netip.AddrPort {
	addr, _ := netip.AddrFromSlice(packet[offset : offset+4])
	port := binary.LittleEndian.Uint16(packet[offset+4 : offset+6])

	return netip.AddrPortFrom(addr, port)
}

func unpackMAC(packet []byte, offset uint8) string {
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		packet[offset],
		packet[offset+1],
		packet[offset+2],
		packet[offset+3],
		packet[offset+4],
		packet[offset+5])
}

func unpackVersion(packet []byte, offset uint8) string {
	major := packet[offset]
	minor := packet[offset+1]

	return fmt.Sprintf("v%x.%02x", major, minor)
}

func unpackDateTime(packet []byte, offset uint8) entities.DateTime {
	bcd := bcd2string(packet[offset : offset+7])
	s := bcd[:4] + "-" + bcd[4:6] + "-" + bcd[6:8] + " " + bcd[8:10] + ":" + bcd[10:12] + ":" + bcd[12:]

	if date, err := entities.ParseDateTime(s); err != nil {
		return entities.NewDateTime(1, 1, 1, 0, 0, 0)
	} else {
		return date
	}
}

func unpackOptionalDateTime(packet []byte, offset uint8) entities.DateTime {
	bcd := bcd2string(packet[offset : offset+7])
	s := bcd[:4] + "-" + bcd[4:6] + "-" + bcd[6:8] + " " + bcd[8:10] + ":" + bcd[10:12] + ":" + bcd[12:]

	if date, err := entities.ParseDateTime(s); err != nil {
		return entities.NewDateTime(1, 1, 1, 0, 0, 0)
	} else {
		return date
	}
}

func unpackDate(packet []byte, offset uint8) entities.Date {
	bcd := bcd2string(packet[offset : offset+4])
	s := bcd[:4] + "-" + bcd[4:6] + "-" + bcd[6:]

	if date, err := entities.ParseDate(s); err != nil {
		return entities.NewDate(1, 1, 1)
	} else {
		return date
	}
}

func unpackShortDate(packet []byte, offset uint8) entities.Date {
	bcd := bcd2string(packet[offset : offset+3])
	s := "20" + bcd[:2] + "-" + bcd[2:4] + "-" + bcd[4:]

	if date, err := entities.ParseDate(s); err != nil {
		return entities.NewDate(1, 1, 1)
	} else {
		return date
	}
}

func unpackOptionalDate(packet []byte, offset uint8) entities.Date {
	bcd := bcd2string(packet[offset : offset+4])
	s := bcd[:4] + "-" + bcd[4:6] + "-" + bcd[6:]

	if date, err := entities.ParseDate(s); err != nil {
		return entities.NewDate(1, 1, 1)
	} else {
		return date
	}
}

func unpackTime(packet []byte, offset uint8) entities.Time {
	bcd := bcd2string(packet[offset : offset+3])
	s := bcd[:2] + ":" + bcd[2:4] + ":" + bcd[4:]

	if t, err := entities.ParseTime(s); err != nil {
		return entities.Time{}
	} else {
		return t
	}
}

func unpackHHmm(packet []byte, offset uint8) entities.HHmm {
	bcd := bcd2string(packet[offset : offset+2])
	s := bcd[:2] + ":" + bcd[2:4]

	if t, err := entities.ParseHHmm(s); err != nil {
		return entities.HHmm{}
	} else {
		return t
	}
}

func unpackPIN(packet []byte, offset uint8) uint32 {
	b := []byte{packet[offset+0], packet[offset+1], packet[offset+2], 0x00}
	v := binary.LittleEndian.Uint32(b)

	return v
}

func unpackMode(packet []byte, offset uint8) entities.DoorMode {
	return entities.DoorMode(packet[offset])
}

func unpackEventType(packet []byte, offset uint8) entities.EventType {
	return entities.EventType(packet[offset])
}

func unpackDirection(packet []byte, offset uint8) entities.Direction {
	return entities.Direction(packet[offset])
}

func unpackReason(packet []byte, offset uint8) entities.Reason {
	return entities.Reason(packet[offset])
}

func bcd2string(bytes []byte) string {
	BCD := hex.EncodeToString(bytes)

	if matched, err := regexp.MatchString(`^[0-9]*$`, BCD); err != nil || !matched {
		panic(fmt.Sprintf("invalid BCD value (%v)", bytes))
	}

	return BCD
}
