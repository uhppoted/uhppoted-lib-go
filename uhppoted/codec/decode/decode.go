package decode

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net/netip"
	"regexp"
	"time"
)

// Message constants for the UHPPOTE request/response protocol.
const SOM byte = 0x17
const SOM_v6_62 byte = 0x19

const GetStatus byte = 0x20
const SetTime byte = 0x30
const GetTime byte = 0x32
const OpenDoor byte = 0x40
const PUT_CARD byte = 0x50
const DELETE_CARD byte = 0x52
const DELETE_ALL_CARDS byte = 0x54
const GET_CARDS byte = 0x58
const GET_CARD byte = 0x5A
const GET_CARD_AT_INDEX byte = 0x5C
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

func unpackYYYYMMDDHHMMSS(packet []byte, offset uint8) time.Time {
	bcd := bcd2string(packet[offset : offset+7])

	if datetime, err := time.ParseInLocation("20060102150405", bcd, time.Local); err != nil {
		return time.Time{}
	} else {
		return datetime
	}
}

func unpackYYYYMMDD(packet []byte, offset uint8) time.Time {
	bcd := bcd2string(packet[offset : offset+4])

	if date, err := time.ParseInLocation("20060102", bcd, time.Local); err != nil {
		return time.Time{}
	} else {
		return date
	}
}

func unpackYYMMDD(packet []byte, offset uint8) time.Time {
	bcd := bcd2string(packet[offset : offset+3])

	if date, err := time.ParseInLocation("20060102", "20"+bcd, time.Local); err != nil {
		return time.Time{}
	} else {
		return date
	}
}

func unpackHHMMSS(packet []byte, offset uint8) time.Time {
	bcd := bcd2string(packet[offset : offset+3])

	if date, err := time.ParseInLocation("150405", bcd, time.Local); err != nil {
		return time.Time{}
	} else {
		return date
	}
}

func bcd2string(bytes []byte) string {
	BCD := hex.EncodeToString(bytes)

	if matched, err := regexp.MatchString(`^[0-9]*$`, BCD); err != nil || !matched {
		panic(fmt.Sprintf("invalid BCD value (%v)", bytes))
	}

	return BCD
}
