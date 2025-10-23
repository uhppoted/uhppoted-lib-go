package uhppoted

import (
	"fmt"
	"strings"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/internal/log"
)

func dump(tag string, description string, packet []byte) {
	lines := []string{description}
	hex := "%02x %02x %02x %02x %02x %02x %02x %02x"

	for i := 0; i < 4; i++ {
		offset := i * 16
		u := packet[offset : offset+8]
		v := packet[offset+8 : offset+16]

		p := fmt.Sprintf(hex, u[0], u[1], u[2], u[3], u[4], u[5], u[6], u[7])
		q := fmt.Sprintf(hex, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7])

		lines = append(lines, fmt.Sprintf("   %08x  %v  %v", offset, p, q))
	}

	lines = append(lines, fmt.Sprintln())

	log.Debugf(tag, "%v", strings.Join(lines, "\n"))
}
