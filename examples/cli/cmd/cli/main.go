package main

import (
	"flag"
	"fmt"
	"log/slog"
	"maps"
	"net/netip"
	"os"
	"slices"
	"time"

	lib "github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
)

var options = struct {
	bind      string
	broadcast string
	listen    string
	timeout   time.Duration
	debug     bool
}{
	bind:      "0.0.0.0:0",
	broadcast: "255.255.255.255:60000",
	listen:    "0.0.0.0:60001",
	timeout:   2500 * time.Millisecond,
	debug:     false,
}

func main() {
	// ... initialise log handler
	h := handler{}
	logger := slog.New(&h)

	slog.SetDefault(logger)
	slog.SetLogLoggerLevel(slog.LevelDebug)

	// ... parse command line
	flag.StringVar(&options.bind, "bind", options.bind, "Sets the local IP address and port to which to bind (e.g. 192.168.0.100)")
	flag.StringVar(&options.broadcast, "broadcast", options.broadcast, "Sets the IP address and port for UDP broadcast (e.g. 192.168.0.255:60000)")
	flag.StringVar(&options.listen, "listen", options.listen, "Sets the local IP address and port to which to bind for events (e.g. 192.168.0.100:60001)")
	flag.DurationVar(&options.timeout, "timeout", options.timeout, "Sets the timeout for a response from a controller (e.g. 2.5s)")
	flag.BoolVar(&options.debug, "debug", options.debug, "Displays internal information for diagnosing errors")
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		usage()
		os.Exit(1)
	} else if cmd, ok := commands[flag.Arg(0)]; !ok {
		fmt.Printf("*** ERROR unknown command (%v)\n\n", flag.Arg(0))
		usage()
		os.Exit(1)
	} else if bind, err := bindAddr(); err != nil {
		fmt.Printf("*** ERROR invalid bind address (%v)\n\n", err)
		os.Exit(1)
	} else if broadcast, err := broadcastAddr(); err != nil {
		fmt.Printf("*** ERROR invalid broadcast address (%v)\n\n", err)
		os.Exit(1)
	} else if listen, err := listenAddr(); err != nil {
		fmt.Printf("*** ERROR invalid listen address (%v)\n\n", err)
		os.Exit(1)
	} else {
		u := lib.NewUhppoted(bind, broadcast, listen, options.debug)

		if err := cmd(u, args[1:]); err != nil {
			fmt.Printf("*** ERROR %v\n\n", err)
			os.Exit(1)
		}
	}
}

func usage() {
	fmt.Println()
	fmt.Println("  Usage: cli <options> <command> <args>")
	fmt.Println("")

	fmt.Println("  Options:")
	fmt.Println("    --bind       IPv4 address to bind to (e.g. --bind 192.168.0.100). Default value is 0.0.0.0")
	fmt.Println("    --broadcast  IPv4 address for UDP broadcast (e.g. --broadcast 192.168.0.255:60000). Default value is 255.255.255.255:60000")
	fmt.Println("    --listen     IPv4 address for events (e.g. --listen 192.168.0.100:60001). Default value is 0.0.0.0:60001")
	fmt.Println("    --timeout    command timeout (e.g. --timeout 5s). Default value is 2.5s")
	fmt.Println("    --debug      displays internal information while executing commands")
	fmt.Println("")

	fmt.Println("  Supported commands:")
	for _, k := range slices.Sorted(maps.Keys(commands)) {
		fmt.Printf("    %v\n", k)
	}
	fmt.Println("")
}

func bindAddr() (netip.AddrPort, error) {
	if v, err := netip.ParseAddrPort(options.bind); err == nil {
		return v, nil
	} else if v, err := netip.ParseAddr(options.bind); err == nil {
		return netip.AddrPortFrom(v, 0), nil
	} else {
		return netip.AddrPort{}, err
	}
}

func broadcastAddr() (netip.AddrPort, error) {
	if v, err := netip.ParseAddrPort(options.broadcast); err == nil {
		return v, nil
	} else if v, err := netip.ParseAddr(options.broadcast); err == nil {
		return netip.AddrPortFrom(v, 60000), nil
	} else {
		return netip.AddrPort{}, err
	}
}

func listenAddr() (netip.AddrPort, error) {
	if v, err := netip.ParseAddrPort(options.listen); err == nil {
		return v, nil
	} else if v, err := netip.ParseAddr(options.listen); err == nil {
		return netip.AddrPortFrom(v, 60001), nil
	} else {
		return netip.AddrPort{}, err
	}
}
