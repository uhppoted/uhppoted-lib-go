package uhppoted

import (
	"fmt"
	"net"
	"net/netip"
	"time"
)

type tcp struct {
	bindAddr *net.TCPAddr
	debug    bool
}

func (t tcp) sendTo(request []byte, dest netip.AddrPort, timeout time.Duration) ([]byte, error) {
	addr := net.TCPAddrFromAddrPort(dest)

	if socket, err := net.DialTCP("tcp", t.bindAddr, addr); err != nil {
		return nil, err
	} else {
		defer socket.Close()

		if N, err := socket.Write(request); err != nil {
			return nil, err
		} else if t.debug {
			dump("tcp", fmt.Sprintf("sent %v bytes to %v", N, dest), request)
		}

		// ... read until reply, timeout or error
		b := make(chan []byte)
		e := make(chan error)

		go func() {
			for {
				buffer := make([]byte, 1024)
				if N, err := socket.Read(buffer); err != nil {
					e <- err
				} else if N == 64 {
					b <- buffer[0:64]

					if t.debug {
						dump("tcp", fmt.Sprintf("received %v bytes from %v", N, dest), buffer[0:64])
					}
				}
			}
		}()

		select {
		case reply := <-b:
			return reply, nil

		case <-time.After(timeout):
			return nil, nil

		case err := <-e:
			return nil, err
		}
	}
}
