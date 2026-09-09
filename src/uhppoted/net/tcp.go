package net

import (
	"fmt"
	"net"
	"net/netip"
	"time"
)

type TCP struct {
	bindAddr *net.TCPAddr
	debug    bool
}

func MakeTCP(bind netip.AddrPort, debug bool) TCP {
	return TCP{
		bindAddr: net.TCPAddrFromAddrPort(bind),
		debug:    debug,
	}
}

func (t TCP) SendTo(request []byte, dest netip.AddrPort, timeout time.Duration) ([]byte, error) {
	sockets := []*net.TCPConn{}

	defer func() {
		for _, sock := range sockets {
			sock.Close()
		}
	}()

	addr := net.TCPAddrFromAddrPort(dest)

	if socket, err := net.DialTCP("tcp", t.bindAddr, addr); err != nil {
		return nil, err
	} else {
		sockets = append(sockets, socket)

		// broadcast-to-self
		//
		// If the bind port is 0, retries (up to 5 times) until the OS binds to a port that is not the
		// destination port. The interim sockets are retained to prevent the OS from reusing the invalid
		// port, which it will happily do if the socket is closed immediately (hence the defer())
		//
		// This should check the destination address as well as the port but multiple network cards, etc
		// make that complicated.
		endpoint := socket.LocalAddr().(*net.TCPAddr)
		if endpoint.Port == int(dest.Port()) && t.bindAddr.Port != 0 {
			return nil, fmt.Errorf("TCP bind address:port %v reserved for broadcast", endpoint.Port)
		}

		retries := 0
		for endpoint.Port == int(dest.Port()) {
			if retries >= 5 {
				return nil, fmt.Errorf("OS bound TCP socket to port %v reserved for broadcast", endpoint.Port)
			} else if socket, err := net.DialTCP("tcp", t.bindAddr, addr); err != nil {
				return nil, err
			} else {
				sockets = append(sockets, socket)

				endpoint = socket.LocalAddr().(*net.TCPAddr)
			}

			retries++
		}

		// ... good to go!
		if N, err := socket.Write(request); err != nil {
			return nil, err
		} else if t.debug {
			dump("tcp", fmt.Sprintf("sent %v bytes to %v", N, dest), request)
		}

		// NTS: set-ip does not return a response
		if request[1] == 0x96 {
			reply := make([]byte, 64)

			reply[0] = 0x17
			reply[1] = 0x96
			copy(reply[4:8], request[4:8])
			reply[8] = 0x01

			return reply, nil
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
