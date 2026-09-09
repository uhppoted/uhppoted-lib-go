package net

import (
	"fmt"
	"net"
	"net/netip"
	"time"
)

type UDP struct {
	bindAddr      *net.UDPAddr
	broadcastAddr *net.UDPAddr
	listenAddr    *net.UDPAddr
	debug         bool
}

func MakeUDP(bind, broadcast, listen netip.AddrPort, debug bool) UDP {
	return UDP{
		bindAddr:      net.UDPAddrFromAddrPort(bind),
		broadcastAddr: net.UDPAddrFromAddrPort(broadcast),
		listenAddr:    net.UDPAddrFromAddrPort(listen),
		debug:         debug,
	}
}

func (u UDP) Broadcast(request []byte, timeout time.Duration) ([][]byte, error) {
	sockets := []*net.UDPConn{}

	defer func() {
		for _, sock := range sockets {
			sock.Close()
		}
	}()

	socket, err := net.ListenUDP("udp", u.bindAddr)
	if err != nil {
		return [][]byte{}, err
	}

	sockets = append(sockets, socket)

	// broadcast-to-self
	//
	// If the bind port is 0, retries (up to 5 times) until the OS binds to a port that is not the
	// destination port. The interim sockets are retained to prevent the OS from reusing the invalid
	// port, which it will happily do if the socket is closed immediately (hence the defer())
	endpoint := socket.LocalAddr().(*net.UDPAddr)
	if endpoint.Port == u.broadcastAddr.Port && u.bindAddr.Port != 0 {
		return [][]byte{}, fmt.Errorf("UDP bind address:port %v reserved for broadcast", endpoint.Port)
	}

	retries := 0
	for endpoint.Port == u.broadcastAddr.Port {
		if retries >= 5 {
			return [][]byte{}, fmt.Errorf("OS bound UDP socket to port %v reserved for broadcast", endpoint.Port)
		} else if socket, err = net.ListenUDP("udp", u.bindAddr); err != nil {
			return [][]byte{}, err
		} else {
			sockets = append(sockets, socket)

			endpoint = socket.LocalAddr().(*net.UDPAddr)
		}

		retries++
	}

	// ... good to go!
	if N, err := socket.WriteToUDP(request, u.broadcastAddr); err != nil {
		return nil, err
	} else if u.debug {
		dump("udp", fmt.Sprintf("sent %v bytes to %v", N, u.broadcastAddr), request)
	}

	// ... read until timeout or error
	e := make(chan error)
	replies := [][]byte{}

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

func (u UDP) BroadcastTo(request []byte, timeout time.Duration) ([]byte, error) {
	sockets := []*net.UDPConn{}

	defer func() {
		for _, sock := range sockets {
			sock.Close()
		}
	}()

	socket, err := net.ListenUDP("udp", u.bindAddr)
	if err != nil {
		return nil, err
	}

	sockets = append(sockets, socket)

	// broadcast-to-self
	//
	// If the bind port is 0, retries (up to 5 times) until the OS binds to a port that is not the
	// destination port. The interim sockets are retained to prevent the OS from reusing the invalid
	// port, which it will happily do if the socket is closed immediately (hence the defer())
	endpoint := socket.LocalAddr().(*net.UDPAddr)
	if endpoint.Port == u.broadcastAddr.Port && u.bindAddr.Port != 0 {
		return nil, fmt.Errorf("UDP bind address:port %v reserved for broadcast", endpoint.Port)
	}

	retries := 0
	for endpoint.Port == u.broadcastAddr.Port {
		if retries >= 5 {
			return nil, fmt.Errorf("OS bound UDP socket to port %v reserved for broadcast", endpoint.Port)
		} else if socket, err = net.ListenUDP("udp", u.bindAddr); err != nil {
			return nil, err
		} else {
			sockets = append(sockets, socket)

			endpoint = socket.LocalAddr().(*net.UDPAddr)
		}

		retries++
	}

	// ... good to go!
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

func (u UDP) SendTo(request []byte, dest netip.AddrPort, timeout time.Duration) ([]byte, error) {
	sockets := []*net.UDPConn{}

	defer func() {
		for _, sock := range sockets {
			sock.Close()
		}
	}()

	addr := net.UDPAddrFromAddrPort(dest)

	socket, err := net.DialUDP("udp", u.bindAddr, addr)
	if err != nil {
		return nil, err
	}

	sockets = append(sockets, socket)

	// broadcast-to-self
	//
	// If the bind port is 0, retries (up to 5 times) until the OS binds to a port that is not the
	// destination port. The interim sockets are retained to prevent the OS from reusing the invalid
	// port, which it will happily do if the socket is closed immediately (hence the defer())
	//
	// This should check the destination address as well as the port but multiple network cards, etc
	// make that complicated.
	endpoint := socket.LocalAddr().(*net.UDPAddr)
	if endpoint.Port == u.broadcastAddr.Port && u.bindAddr.Port != 0 {
		return nil, fmt.Errorf("UDP bind address:port %v reserved for broadcast", endpoint.Port)
	}

	retries := 0
	for endpoint.Port == u.broadcastAddr.Port {
		if retries >= 5 {
			return nil, fmt.Errorf("OS bound UDP socket to port %v reserved for broadcast", endpoint.Port)
		} else if socket, err := net.DialUDP("udp", u.bindAddr, addr); err != nil {
			return nil, err
		} else {
			sockets = append(sockets, socket)

			endpoint = socket.LocalAddr().(*net.UDPAddr)
		}

		retries++
	}

	// ... good to go!
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

func (u UDP) Listen(ch chan []byte) error {
	socket, err := net.ListenUDP("udp", u.listenAddr)
	if err != nil {
		return err
	}

	defer socket.Close()

	events := make(chan []byte)
	e := make(chan error)

	go func() {
		buffer := make([]byte, 1024)

		for {
			if N, addr, err := socket.ReadFromUDP(buffer); err != nil {
				e <- err
			} else if N == 64 {
				m := make([]byte, 64)
				copy(m, buffer[0:N])
				events <- m

				if u.debug {
					dump("udp", fmt.Sprintf("received %v bytes from %v", N, addr), buffer[0:64])
				}
			}
		}
	}()

	for {
		select {
		case m := <-events:
			ch <- m

		case err := <-e:
			return err
		}
	}
}
