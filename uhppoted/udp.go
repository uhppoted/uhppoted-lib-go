package uhppoted

import (
	"fmt"
	"net"
	"time"
)

type udp struct {
	bindAddr      *net.UDPAddr
	broadcastAddr *net.UDPAddr
	listenAddr    *net.UDPAddr
	debug         bool
}

func (u udp) broadcast(request []byte, timeout time.Duration) ([][]byte, error) {
	replies := [][]byte{}

	if socket, err := net.ListenUDP("udp", u.bindAddr); err != nil {
		return replies, err
	} else {
		defer socket.Close()

		endpoint := socket.LocalAddr().(*net.UDPAddr)
		if endpoint.Port == u.broadcastAddr.Port {
			return replies, fmt.Errorf("invalid UDP bind address: port %d reserved for broadcast", endpoint.Port)
		}

		if _, err := socket.WriteToUDP(request, u.broadcastAddr); err != nil {
			return nil, err
		} else if u.debug {
			dump(request)
		}

		// ... read until timeout or error
		e := make(chan error)

		go func() {
			for {
				buffer := make([]byte, 1024)
				if N, _, err := socket.ReadFromUDP(buffer); err != nil {
					e <- err
				} else if N == 64 {
					replies = append(replies, buffer[0:64])
					if u.debug {
						dump(buffer[0:64])
					}
				}
			}
		}()

		select {
		case <-time.After(timeout):
			return replies, nil

		case err := <-e:
			return replies, err
		}

	}
}

func dump(packet []byte) {
	hex := "%02x %02x %02x %02x %02x %02x %02x %02x"

	for i := 0; i < 4; i++ {
		offset := i * 16
		u := packet[offset : offset+8]
		v := packet[offset+8 : offset+16]

		p := fmt.Sprintf(hex, u[0], u[1], u[2], u[3], u[4], u[5], u[6], u[7])
		q := fmt.Sprintf(hex, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7])

		fmt.Printf("   %08x  %v  %v\n", offset, p, q)
	}

	fmt.Println()
}

