// Package uhppoted provides a Go wrapper for the UHPPOTE TCP/IP access controller API.
//
// It enables interaction with UHPPOTE controllers using either UDP broadcasts (on a local LAN),
// UDP connected sockets or TCP/IP. The package supports controller configuration, access card
// management and door control.
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
