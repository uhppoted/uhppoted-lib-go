![build](https://github.com/uhppoted/uhppoted-lib-go/workflows/build/badge.svg)

# uhppoted-lib-go

** IN DEVELOPMENT **

Standalone Go library for the UHPPOTE access controllers.

_uhppoted-lib-go_ supersedes [_uhppote-core_](https://github.com/uhppoted/uhppote-core) for external use:

- _uhppote-core_ was developed a **long** time ago, long before generics and other modern constructs were available in Go and is 
starting to show its age
- the _uhppoted-lib-go_ API conforms to the informal conventions of the other _uhppoted-lib-xxx_ projects
- the _uhppoted-lib-go_ implementation is considerably simpler and easier to understand and maintain

A basic example CLI illustrating the use of the library can be found in the [examples](https://github.com/uhppoted/uhppoted-lib-go/tree/main/examples)
folder.

## Installation


## Release Notes

#### Current Release

** IN DEVELOPMENT **


## API

Invoking an API function requires an instance of the `Uhppoted` struct initialised with the information required
to access a controller:

```
        u := lib.NewUhppoted(bind, broadcast, listen, debug)

where:

bind        IPv4 address to which to bind the UDP socket
broadcast   IPv4 address:port for broadcast UDP packets
listen      IPv4 address:port for events from controller
debug       Displays the controller requests/responses if true.
```

e.g.:
```
import (
    "fmt"
    "net/netip"

    lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

...
func main() {
    bind := netip.MustParseAddr("0.0.0.0")
    broadcast := netip.MustParseAddrPort("255.255.255.255:60000")
    listen := netip.MustParseAddrPort("0.0.0.0:60001")
    debug := true

    u := lib.NewUhppoted(bind, broadcast, listen, options.debug)
    
    if controller, err := lib.GetController(u, 405419896); err != nil {
        fmt.Printf("ERROR %v\n", err)
    } else {
        fmt.Printf("%v\n", controller)
    }
}
```

### Notes
1. All API functions return an error if the call fails for any reason whatsoever.
2. All API functions (other than `get_all_controllers` and `listen`) take a `controller` that may be either:
   - a _uint32_ controller serial number (legacy)
   - a Controller struct
```
type Controller struct {
    ID       uint32
    Address  netip.AddrPort
    Protocol string
}

where:
- ID        controller serial number
- Address   controller IPv4 address and port. 
            (optional - defaults to UDP broadcast if not provided)
- Protocol  either "udp" or "tcp". 
            (optional - defaults to "udp")
```
   e.g.:
```
   lib.GetController(u, 405419896, TIMEOUT)
   lib.GetController(u, lib.Controller{
                           405419896, 
                           netip.MustParseAddrPort('192.168.1.100:60000'), 
                           'tcp'}, TIMEOUT)
```

3. All API functions (other than `listen`) take a `timeout` parameter that sets the time limit for the request, 
   e.g.:
```
   GetController(u, 405419896, 750*time.Millisecond)
```

**API**

- [`GetAllControllers`](#get_all_controllers)
- [`GetController`](#get_controller)


### `GetAllControllers`
Returns a list of all controllers that responded to a _get-controller_ request within the timeout.

```
GetAllControllers(u Uhppoted, timeout time.Duration) ([]GetControllerResponse, error)

where:
u        Uhppoted struct initialised with the bind address, broadcast address, etc
timeout  maximum time to wait for a response from a controller

Returns an array of `GetControllerResponse`.

```

### `GetController`
Returns the sytem information for the requested access controller.

```
GetController(u Uhppoted, controller TController, timeout time.Duration)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
timeout     maximum time to wait for a response from a controller

Returns a `GetControllerResponse`.
```

### `SetIPv4`
Sets the controller IPv4 address, netmask and gateway address.

```
SetIPv4(self, ID, address, netmask, gateway)

controller  uint32|tuple  controller serial number or (id, address, protocol) tuple
address     netip.Addr    controller IPv4 address
subnet      netip.Addr    controller IPv4 subnet mask
gateway     netip.Addr    controller gateway IPv4 address

Returns a `SetIPv4Response`.
```

## Types

### `GetControllerResponse`

Container class for the decoded response from a _GetController_ request.
```
type GetControllerResponse struct {
    Controller uint32     `json:"controller"`    // controller serial number
    IpAddress  netip.Addr `json:"ip-address"`    // IPv4 address
    SubnetMask netip.Addr `json:"subnet-mask"`   // IPv4 netmask
    Gateway    netip.Addr `json:"gateway"`       // gateway IP v4address
    MACAddress string     `json:"MAC-address"`   // MAC address (XX:XX:XX:XX:XX:XX)
    Version    string     `json:"version"`       // firmware version (vN.NN)
    Date       time.Time  `json:"date"`          // release date (YYYY-MM-DD)
}
```

### `setIPv4Response`

Container class for the decoded response from a _SetIPv4_ request.
```
type SetIPv4Response struct {
    Controller uint32     `json:"controller"`    // controller serial number
    Ok         bool       `json:"ok"`            // succeeded/failed
}
```

