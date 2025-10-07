// Package uhppoted provides a Go wrapper for the UHPPOTE TCP/IP access controller API.
//
// It enables interaction with UHPPOTE controllers on a local LAN using either UDP
// broadcasts, UDP connected sockets or TCP/IP. The package supports querying controller
// information, managing access cards, and controlling doors.
//
// Typical usage:
//
//	client := uhppoted.Uhppoted{
//	    BindAddr:      netip.MustParseAddrPort("0.0.0.0:0"),
//	    BroadcastAddr: netip.MustParseAddrPort("255.255.255.255:60000"),
//	    ListenAddr:    netip.MustParseAddrPort("0.0.0.0:60001"),
//	    Debug:         false,
//	}
//
//	if controllers, err := client.GetAllControllers(2 * time.Second); err != nil {
//	    log.Fatal(err)
//	} else {
//	    for _, c := range controllers {
//	        fmt.Printf("Controller: %v\n", c.SerialNumber)
//	    }
//	}
package uhppoted
