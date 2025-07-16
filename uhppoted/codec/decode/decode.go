package decode

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net/netip"
	"regexp"
	"time"
)

func unpackUint32(packet []byte, offset uint8) uint32 {
	return binary.LittleEndian.Uint32(packet[offset : offset+4])
}

func unpackIPv4(packet []byte, offset uint8) netip.Addr {
	addr, _ := netip.AddrFromSlice(packet[offset : offset+4])

	return addr
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

func unpackDate(packet []byte, offset uint8) time.Time {
	bcd := bcd2string(packet[offset : offset+4])

	if date, err := time.ParseInLocation("20060102", bcd, time.Local); err != nil {
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
