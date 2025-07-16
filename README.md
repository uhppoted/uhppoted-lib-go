# uhppoted-lib-go

Standalone Go library for the UHPPOTE access controllers.

_uhppoted-lib-go_ supersedes [_uhppote-core_](https://github.com/uhppoted/uhppote-core) for external use:

- _uhppote-core_ was developed a **long** time ago, long before generics and other modern constructs were available in Go
- the _uhppoted-lib-go_ implementation is considerably simpler and easier to understand and maintain
- the _uhppoted-lib-go_ API conforms to the informal conventions of the other _uhppoted-lib-xxx_ projects

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
2. All API functions (other than `get_controllers` and `listen`) take a `controller` that may be either:
   - a _uint32_ controller serial number (legacy)
   - a Controller struct:
```
type Controller struct {
    ID       uint32
    Address  netip.AddrPort
    Protocol string
}

where:
- ID is the controller serial number
- Address is the controller IPv4 address and port. Optional - defaults to UDP broadcast if not provided.
- Protocol is either "udp" or "tcp". Optional - defaults to "udp".
```
   e.g.:
```
   lib.GetController(u, 405419896, TIMEOUT)
   lib.GetController(u, lib.Controller{405419896, netip.MustParseAddrPort('192.168.1.100:60000'), 'tcp'), TIMEOUT)
```

3. All API functions (other than `listen`) take a `timeout` parameter that sets the time limit for the request, 
   e.g.:
```
   GetController(u, 405419896, 750*time.Millisecond)
```

**API**

- [`get_controllers`](#get_controllers)
- [`get_controller`](#get_controller)

### `GetAllControllers`
```
GetAllControllers(u Uhppoted, timeout time.Duration) ([]GetControllerResponse, error)

where:
u        Uhppoted struct initialised with the bind address, broadcast address, etc
timeout  maximum time to wait for a response from a controller

Returns a list of all controllers that responded to the request within the timeout interval.
```

### `GetController`
```
GetController(u Uhppoted, controller TController, timeout time.Duration)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
timeout     maximum time to wait for a response from a controller

Returns a `GetControllerResponse` with the controller information.
```

## Types

### `GetControllerResponse`

Container class for the decoded response from a _GetController_ request.
```
    Fields:
        Controller   Controller serial number.
        IpAddress    IPv4 address.
        SubnetMask   IPv4 subnet mask.
        Gateway      Gateway IP v4address.
        MACAddress   MAC address (XX:XX:XX:XX:XX:XX).
        Version      Firmware version (vN.NN).
        Date         Release date (YYYY-MM-DD).
```
```
type GetControllerResponse struct {
    Controller uint32     `json:"controller"`
    IpAddress  netip.Addr `json:"ip-address"`
    SubnetMask netip.Addr `json:"subnet-mask"`
    Gateway    netip.Addr `json:"gateway"`
    MACAddress string     `json:"MAC-address"`
    Version    string     `json:"version"`
    Date       time.Time  `json:"date"`
}
```
