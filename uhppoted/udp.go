package uhppoted

import (
	"fmt"
	"net"
	"net/netip"
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

		if N, err := socket.WriteToUDP(request, u.broadcastAddr); err != nil {
			return nil, err
		} else if u.debug {
			dump("udp", fmt.Sprintf("sent %v bytes to %v", N, u.broadcastAddr), request)
		}

		// ... read until timeout or error
		e := make(chan error)

		go func() {
			for {
				buffer := make([]byte, 1024)
				if N, addr, err := socket.ReadFromUDP(buffer); err != nil {
					e <- err
				} else if N == 64 {
					if u.debug {
						dump("udp", fmt.Sprintf("received %v bytes from %v", N, addr), buffer[0:64])
					}

					replies = append(replies, buffer[0:64])
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

func (u udp) broadcastTo(request []byte, timeout time.Duration) ([]byte, error) {
	if socket, err := net.ListenUDP("udp", u.bindAddr); err != nil {
		return nil, err
	} else {
		defer socket.Close()

		endpoint := socket.LocalAddr().(*net.UDPAddr)
		if endpoint.Port == u.broadcastAddr.Port {
			return nil, fmt.Errorf("invalid UDP bind address: port %d reserved for broadcast", endpoint.Port)
		}

		if N, err := socket.WriteToUDP(request, u.broadcastAddr); err != nil {
			return nil, err
		} else if u.debug {
			dump("udp", fmt.Sprintf("sent %v bytes to %v", N, u.broadcastAddr), request)
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
				if N, addr, err := socket.ReadFromUDP(buffer); err != nil {
					e <- err
				} else if N == 64 {
					b <- buffer[0:64]

					if u.debug {
						dump("udp", fmt.Sprintf("received %v bytes from %v", N, addr), buffer[0:64])
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

func (u udp) sendTo(request []byte, dest netip.AddrPort, timeout time.Duration) ([]byte, error) {
	addr := net.UDPAddrFromAddrPort(dest)

	if socket, err := net.DialUDP("udp", u.bindAddr, addr); err != nil {
		return nil, err
	} else {
		defer socket.Close()

		endpoint := socket.LocalAddr().(*net.UDPAddr)
		if endpoint.Port == u.broadcastAddr.Port {
			return nil, fmt.Errorf("invalid UDP bind address: port %d reserved for broadcast", endpoint.Port)
		}

		if N, err := socket.Write(request); err != nil {
			return nil, err
		} else if u.debug {
			dump("udp", fmt.Sprintf("sent %v bytes to %v", N, dest), request)
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
				if N, addr, err := socket.ReadFromUDP(buffer); err != nil {
					e <- err
				} else if N == 64 {
					b <- buffer[0:64]

					if u.debug {
						dump("udp", fmt.Sprintf("received %v bytes from %v", N, addr), buffer[0:64])
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
