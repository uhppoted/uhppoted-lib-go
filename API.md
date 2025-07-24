# API

- [`GetAllControllers`](#getallcontrollers)
- [`GetController`](#getcontroller)
- [`SetIPv4`](#setipv4)
- [`GetStatus`](#getstatus)
- [`GetTime`](#gettime)
- [`SetTime`](#settime)
- [`GetListener`](#getlistener)
---
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

## Functions

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
GetController(u Uhppoted, controller TController, timeout time.Duration) (GetControllerResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
timeout     maximum time to wait for a response from a controller

Returns a `GetControllerResponse`.
```

### `SetIPv4`
Sets the controller IPv4 address, netmask and gateway address.

```
SetIPv4(self, ID, address, netmask, gateway) (SetIPv4Response,error)

controller  uint32|tuple  controller serial number or (id, address, protocol) tuple
address     netip.Addr    controller IPv4 address
subnet      netip.Addr    controller IPv4 subnet mask
gateway     netip.Addr    controller gateway IPv4 address

Returns a `SetIPv4Response`.
```

### `GetStatus`
```
GetStatus(u Uhppoted, controller TController, timeout time.Duration) (GetStatusResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
timeout     maximum time to wait for a response from a controller

Returns a `GetStatusResponse` with the controller status information. If the response does not contain a
valid event, the event fields are set to `None`.
```

### `GetTime`
```
GetTime(u Uhppoted, controller TController, timeout time.Duration) (GetTimeResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
timeout     maximum time to wait for a response from a controller

Returns a `GetTimeResponse` with the controller system date/time.
```

### `SetTime`
```
SetTime(u Uhppoted, controller TController, datetime time.Time, timeout time.Duration) (SetTimeResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
datetime    date/time to which to set controller system time
timeout     maximum time to wait for a response from a controller

Returns a `SetTimeResponse` with the controller system date/time.
```

### `GetListener`
```
GetListener(u Uhppoted, controller TController, timeout time.Duration) (GetListenerResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
timeout     maximum time to wait for a response from a controller

Returns a `GetListenerResponse` with the configured event listener IPv4 address:port and the auto-send 
interval.
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

### `GetStatusResponse`

Container class for the decoded response from a _GetStatus_ request.
```
type GetStatusResponse struct {
  Controller          uint32     `json:"controller"`      // controller serial number
  SystemDate          time.Time  `json:"system-date"`     // controller system date
  SystemTime          time.Time  `json:"system-time"`     // controller system time
  Door1Open           bool       `json:"door-1-open"`     // door 1 locked/unlocked
  Door2Open           bool       `json:"door-2-open"`     // door 2 locked/unlocked
  Door3Open           bool       `json:"door-3-open"`     // door 3 locked/unlocked
  Door4Open           bool       `json:"door-4-open"`     // door 4 locked/unlocked
  Door1Button         bool       `json:"door-1-button"`   // pushbutton 1 pressed/released
  Door2Button         bool       `json:"door-2-button"`   // pushbutton 2 pressed/released
  Door3Button         bool       `json:"door-3-button"`   // pushbutton 3 pressed/released
  Door4Button         bool       `json:"door-4-button"`   // pushbutton 4 pressed/released
  Relays              uint8      `json:"relays"`          // bit array of relay states
  Inputs              uint8      `json:"alarm-inputs"`    // bit array of door sensor states
  SystemError         uint8      `json:"system-error"`    // system error code
  SpecialInfo         uint8      `json:"special-info"`    // (absolutely no idea)
  EventIndex          uint32     `json:"event-index"`     // index of last recorded event
  EventType           uint8      `json:"event-type"`      // type of last recorded event
  EventAccessGranted  bool       `json:"event-granted"`   // last event access granted/denied
  EventDoor           uint8      `json:"event-door"`      // last event door no. [1..4]
  EventDirection      uint8      `json:"event-direction"` // last event direction (0: in, 1: out)
  EventCard           uint32     `json:"event-card"`      // last event card number
  EventTimestamp      time.Time  `json:"event-timestamp"` // last event timestamp
  EventReason         uint8      `json:"event-reason"`    // last event access granted/denied reason code
  SequenceNo          uint32     `json:"sequence-no"`     // packet sequence number
}
```

### `GetTimeResponse`

Container class for the decoded response from a _GetTime_ request.
```
type GetTimeResponse struct {
  Controller  uint32    `json:"controller"`      // controller serial number
  DateTime    time.Time `json:"event-timestamp"` // controller date/time
}
```

### `SetTimeResponse`

Container class for the decoded response from a _SetTime_ request.
```
type SetTimeResponse struct {
  Controller  uint32    `json:"controller"`      // controller serial number
  DateTime    time.Time `json:"event-timestamp"` // controller date/time
}
```

### `GetListenerResponse`

Container class for the decoded response from a _GetListener_ request.
```
type GetTimeResponse struct {
  Controller  uint32          `json:"controller"` // controller serial number
  Address     netip.AddrPort  `json:"address"`    // event listener IPv4 address:port
  Interval    uint8           `json:"interval"`   // auto-send interval (seconds)
}
```
