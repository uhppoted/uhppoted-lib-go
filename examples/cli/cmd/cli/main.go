package main

import (
	"flag"
	"fmt"
	"maps"
	"slices"
	"time"
)

var options = struct {
	bind      string
	broadcast string
	listen    string
	timeout   time.Duration
	debug     bool
}{
	bind:      "0.0.0.0:0",
	broadcast: "255.255.255.255:60001",
	listen:    "0.0.0.0:60001",
	timeout:   2500 * time.Millisecond,
	debug:     false,
}

var commands = map[string]func(args []string) error{
	"get-all-controllers": GetAllControllers,
}

func main() {
	flag.StringVar(&options.bind, "bind", options.bind, "Sets the local IP address and port to which to bind (e.g. 192.168.0.100)")
	flag.StringVar(&options.broadcast, "broadcast", options.broadcast, "Sets the IP address and port for UDP broadcast (e.g. 192.168.0.255:60000)")
	flag.StringVar(&options.listen, "listen", options.listen, "Sets the local IP address and port to which to bind for events (e.g. 192.168.0.100:60001)")
	flag.DurationVar(&options.timeout, "timeout", options.timeout, "Sets the timeout for a response from a controller (e.g. 2.5s)")
	flag.BoolVar(&options.debug, "debug", options.debug, "Displays internal information for diagnosing errors")
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		usage()
	} else if cmd, ok := commands[flag.Arg(0)]; !ok {
		fmt.Printf("*** ERROR unknown command (%v)\n\n", flag.Arg(0))
		usage()
	} else if err := cmd(args[1:]); err != nil {
		fmt.Printf("*** ERROR %v\n\n", err)
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
