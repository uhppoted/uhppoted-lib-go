package decode

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net/netip"
	"regexp"
	"time"

	"github.com/uhppoted/uhppoted-lib-go/uhppoted/entities"
)

// Message constants for the UHPPOTE request/response protocol.
const SOM byte = 0x17
const SOM_v6_62 byte = 0x19

const GetStatus byte = 0x20
const SetTime byte = 0x30
const GetTime byte = 0x32
const OpenDoor byte = 0x40
const PutCard byte = 0x50
const DeleteCard byte = 0x52
const DeleteAllCards byte = 0x54
const GetCards byte = 0x58
const GetCard byte = 0x5a
const GetCardAtIndex byte = 0x5c
const SetDoor byte = 0x80
const GetDoor byte = 0x82
const SET_ANTIPASSBACK byte = 0x84
const GET_ANTIPASSBACK byte = 0x86
const SET_TIME_PROFILE byte = 0x88
const CLEAR_TIME_PROFILES byte = 0x8A
const SetDoorPasscodes byte = 0x8c
const RECORD_SPECIAL_EVENTS byte = 0x8E
const SetListener byte = 0x90
const GetListener byte = 0x92
const GetController byte = 0x94
const SetIPv4 byte = 0x96
const GET_TIME_PROFILE byte = 0x98
const SET_PC_CONTROL byte = 0xA0
const SET_INTERLOCK byte = 0xA2
const ACTIVATE_KEYPADS byte = 0xA4
const CLEAR_TASKLIST byte = 0xA6
const ADD_TASK byte = 0xA8
const SET_FIRST_CARD byte = 0xAA
const REFRESH_TASKLIST byte = 0xAC
const GET_EVENT byte = 0xB0
const SET_EVENT_INDEX byte = 0xB2
const GET_EVENT_INDEX byte = 0xB4
const RESTORE_DEFAULT_PARAMETERS byte = 0xC8
const LISTEN_EVENT byte = 0x20

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
		return entities.NewDateTime(1, time.January, 1, 0, 0, 0)
	} else {
		return date
	}
}

func unpackOptionalDateTime(packet []byte, offset uint8) time.Time {
	bcd := bcd2string(packet[offset : offset+7])

	if d, err := time.ParseInLocation("20060102150405", bcd, time.Local); err != nil {
		return time.Date(1, time.January, 1, 0, 0, 0, 0, time.Local)
	} else {
		return d
	}
}

func unpackDate(packet []byte, offset uint8) entities.Date {
	bcd := bcd2string(packet[offset : offset+4])
	s := bcd[:4] + "-" + bcd[4:6] + "-" + bcd[6:]

	if date, err := entities.ParseDate(s); err != nil {
		return entities.NewDate(1, time.January, 1)
	} else {
		return date
	}
}

func unpackShortDate(packet []byte, offset uint8) entities.Date {
	bcd := bcd2string(packet[offset : offset+3])
	s := "20" + bcd[:2] + "-" + bcd[2:4] + "-" + bcd[4:]

	if date, err := entities.ParseDate(s); err != nil {
		return entities.NewDate(1, time.January, 1)
	} else {
		return date
	}
}

func unpackOptionalDate(packet []byte, offset uint8) entities.Date {
	bcd := bcd2string(packet[offset : offset+4])
	s := bcd[:4] + "-" + bcd[4:6] + "-" + bcd[6:]

	if date, err := entities.ParseDate(s); err != nil {
		return entities.NewDate(1, time.January, 1)
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

func bcd2string(bytes []byte) string {
	BCD := hex.EncodeToString(bytes)

	if matched, err := regexp.MatchString(`^[0-9]*$`, BCD); err != nil || !matched {
		panic(fmt.Sprintf("invalid BCD value (%v)", bytes))
	}

	return BCD
}
