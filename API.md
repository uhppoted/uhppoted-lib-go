# API

- [`FindControllers`](#findcontrollers)
- [`GetController`](#getcontroller)
- [`SetIPv4`](#setipv4)
- [`GetTime`](#gettime)
- [`SetTime`](#settime)
- [`GetListener`](#getlistener)
- [`SetListener`](#setlistener)
- [`GetListenerAddrPort`](#getlisteneraddrport)
- [`SetListenerAddrPort`](#setlisteneraddrport)
- [`GetDoor`](#getdoor)
- [`SetDoor`](#setdoor)
- [`SetDoorPasscodes`](#setdoorpasscodes)
- [`OpenDoor`](#opendoor)
- [`GetStatus`](#getstatus)
- [`GetCards`](#getcards)
- [`GetCard`](#getcard)
- [`GetCardAtIndex`](#getcardatindex)
- [`PutCard`](#putcard)
- [`DeleteCard`](#deletecard)
- [`DeleteAllCards`](#deleteallcards)
- [`GetEvent`](#getevent)
- [`GetEventIndex`](#geteventindex)
- [`SetEventIndex`](#seteventindex)
- [`RecordSpecialEvents`](#recordspecialevents)
- [`GetTimeProfile`](#gettimeprofile)
- [`SetTimeProfile`](#settimeprofile)
- [`ClearTimeProfiles`](#cleartimeprofiles)
- [`AddTask`](#addtask)
- [`RefreshTaskList`](#refreshtasklist)
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

Returns a `GetListenerResponse` with the configured event listener IPv4 address and port and the auto-send 
interval.
```

### `SetListener`
```
SetListener(u Uhppoted, controller TController, address netip.Addr, port uint16, interval uint8, timeout time.Duration) (SetListenerResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
address     IPv4 address of host to receive controller events
port        UDP port of host for controller events
interval    status auto-send interval (seconds). A 0 interval disables auto-send.
timeout     maximum time to wait for a response from a controller

Returns a `SetListenerResponse`. 
interval.
```

### `GetListenerAddrPort`
```
GetListenerAddrPort(u Uhppoted, controller TController, timeout time.Duration) (GetListenerAddrPortResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
timeout     maximum time to wait for a response from a controller

Returns a `GetListenerAddrPortResponse` with the configured event listener IPv4 address:port and the auto-send 
interval.
```

### `SetListenerAddrPort`
```
SetListenerAddrPort(u Uhppoted, controller TController, listener netip.AddrPort, interval uint8, timeout time.Duration) (SetListenerAddrPortResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
listener    IPv4 address:port of host to receive controller events
interval    status auto-send interval (seconds). A 0 interval disables auto-send.
timeout     maximum time to wait for a response from a controller

Returns a `SetListenerAddrPortResponse`.
interval.
```

### `GetDoor`
```
GetDoor(u Uhppoted, controller TController, door uint8, timeout time.Duration) (GetDoorResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
door        door ID ([1..4]
timeout     maximum time to wait for a response from a controller

Returns a `GetDoorResponse` with the door control mode and unlock delay.
```

### `SetDoor`
```
SetDoor(u Uhppoted, controller TController, door uint8, mode uint8, delay uint8, timeout time.Duration) (SetDoorResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
door        door ID ([1..4]
mode        control mode (1:normally open, 2:normally closed. 3:controlled)
delay       unlock delay(seconds)
timeout     maximum time to wait for a response from a controller

Returns a `SetDoorResponse` with the door control mode and unlock delay.
```

### `SetDoorPasscodes`
```
SetDoorPasscodes(u Uhppoted, controller TController, door uint8, passcode1, passcode2, passcode3, passcode4 uint32, timeout time.Duration) (SetDoorPasscodesResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
door        door ID ([1..4]
passcode1   supervisor passcode ([0..99999]), 0 for 'none'
passcode2   supervisor passcode ([0..99999]), 0 for 'none'
passcode3   supervisor passcode ([0..99999]), 0 for 'none'
passcode4   supervisor passcode ([0..99999]), 0 for 'none'
timeout     maximum time to wait for a response from a controller

Returns a `SetDoorPasscodesResponse`.
```

### `OpenDoor`
```
OpenDoor(u Uhppoted, controller TController, door uint8, timeout time.Duration) (OpenDoorResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
door        door ID ([1..4]
timeout     maximum time to wait for a response from a controller

Returns an `OpenDoorResponse`.
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

### `GetCards`
```
GetCards(u Uhppoted, controller TController, timeout time.Duration) (GetCardsResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
timeout     maximum time to wait for a response from a controller

Returns a `GetCardsResponse` with the number of cards stored on the controller.
```

### `GetCard`
```
GetCard(u Uhppoted, controller TController, card uint32, timeout time.Duration) (GetCardResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
card        card number
timeout     maximum time to wait for a response from a controller

Returns a `GetCardResponse` with the card information for the requested card. The card number is 0 if the
card is not stored on the controller.
```

### `GetCardAtIndex`
```
GetCardAtIndex(u Uhppoted, controller TController, index uint32, timeout time.Duration) (GetCardResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
index       card record index
timeout     maximum time to wait for a response from a controller

Returns a `GetCardAtIndexResponse` with the card information for the requested card. The card number is 0 if the
card is not stored on the controller and 0xffffffff if the record has been deleted.
```

### `PutCard`
```
PutCard(u Uhppoted, controller TController, card uint32, startDate, endDate time.Time, door1, door2, door3, door4 uint8, PIN uint32, timeout time.Duration) (GetCardResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
card        card number
startDate   date from which card is valid
endDate     date until  which card is valid (inclusive)
door1       access permissions for door 1 (0: no access, 1: 24/7 access, [2..253] time profile)
door2       access permissions for door 2 (0: no access, 1: 24/7 access, [2..253] time profile)
door3       access permissions for door 3 (0: no access, 1: 24/7 access, [2..253] time profile)
door4       access permissions for door 4 (0: no access, 1: 24/7 access, [2..253] time profile)
PIN         optional PIN code [0..999999] (0 for none)
timeout     maximum time to wait for a response from a controller

Returns a `PutCardResponse` with the card add/update result.
```

### `DeleteCard`
```
DeleteCard(u Uhppoted, controller TController, card uint32, timeout time.Duration) (DeleteCardResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
card        card number
timeout     maximum time to wait for a response from a controller

Returns a `DeleteCardResponse` with `ok` set to `true` if the card record was deleted from the controller.
```

### `DeleteAllCards`
```
DeleteAllCards(u Uhppoted, controller TController, timeout time.Duration) (DeleteCardResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
timeout     maximum time to wait for a response from a controller

Returns a `DeleteAllCardsResponse` with `ok` set to `true` if all card records were deleted from the controller.
```

### `GetEvent`
```
GetEvent(u Uhppoted, controller TController, index uint32, timeout time.Duration) (GetEventResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
index       index of event to retrieve
timeout     maximum time to wait for a response from a controller

Returns a `GetEventResponse`.
```

### `GetEventIndex`
```
GetEventIndex(u Uhppoted, controller TController, timeout time.Duration) (GetEventResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
timeout     maximum time to wait for a response from a controller

Returns a `GetEventIndexResponse`.
```

### `SetEventIndex`
```
SetEventIndex(u Uhppoted, controller TController, index uint32, timeout time.Duration) (SetEventResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
index       Downloaded event index
timeout     maximum time to wait for a response from a controller

Returns a `SetEventIndexResponse`.
```

### `RecordSpecialEvents`
```
RecordSpecialEvents(u Uhppoted, controller TController, enabled bool, timeout time.Duration) (RecordSpecialEventsResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
enabled     Enables door opened, door closed and button pressed events if true.
timeout     maximum time to wait for a response from a controller

Returns a `RecordSpecialEventsResponse`.
```

### `GetTimeProfile`
```
GetTimeProfile(u Uhppoted, controller TController, profile uint8, timeout time.Duration) (GetTimeProfileResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
profile     Profile Id ([2..254] to fetch.
timeout     maximum time to wait for a response from a controller

Returns a `GetTimeProfileResponse`.
```

### `SetTimeProfile`
```
SetTimeProfile(u Uhppoted, controller     TController, 
                           profile        uint8, 
                           startDate      time.Time,
                           endDate        time.Time,
                           monday         bool,
                           tuesday        bool,
                           wednesday      bool,
                           thursday       bool,
                           friday         bool,
                           saturday       bool,
                           sunday         bool,
                           segment1Start  time.Time,
                           segment1End    time.Time,
                           segment2Start  time.Time,
                           segment2End    time.Time,
                           segment3Start  time.Time,
                           segment3End    time.Time,
                           linkedProfile  uint8,
                           timeout        time.Duration) (SetTimeProfileResponse,error)

u              Uhppoted struct initialised with the bind address, broadcast address, etc
controller     uint32|Controller controller serial number or {id, address, protocol} Controller struct
profile        uint8             profile Id ([2..254] to create/update
startDate      time.Time         date from which profile is valid (inclusive)
endDate        time.Time         date after which profile is invalid
monday         bool              profile enabled on Monday if true
tuesday        bool              profile enabled on Tuesday if true
wednesday      bool              profile enabled on Wednesday if true
thursday       bool              profile enabled on Thursday if true
friday         bool              profile enabled on Monday if true
saturday       bool              profile enabled on Friday if true
sunday         bool              profile enabled on Sunday if true
segment1Start  time.Time         start time for first time segment
segment1End    time.Time         end time for first time segment
segment2Start  time.Time         start time for second time segment
segment2End    time.Time         end time for second time segment
segment3Start  time.Time         start time for third time segment
segment3End    time.Time         end time for third time segment
linkedProfile  uint8             ID of linked profile (0 if not linked)
timeout        maximum time to wait for a response from a controller

Returns a `SetTimeProfileResponse`.
```

### `ClearTimeProfiles`
```
ClearTimeProfiles(u Uhppoted, controller TController, timeout time.Duration) (ClearTimeProfilesResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
timeout     maximum time to wait for a response from a controller

Returns a `ClearTimeProfilesResponse`.
```

### `AddTask`
```
AddTask(u Uhppoted, controller     TController, 
                    task           uint8, 
                    startDate      time.Time,
                    endDate        time.Time,
                    monday         bool,
                    tuesday        bool,
                    wednesday      bool,
                    thursday       bool,
                    friday         bool,
                    saturday       bool,
                    sunday         bool,
                    startTime      time.Time,
                    door           uint8,
                    moreCards      uint8,
                    timeout        time.Duration) (SetTimeProfileResponse,error)

u              Uhppoted struct initialised with the bind address, broadcast address, etc
controller     uint32|Controller controller serial number or {id, address, protocol} Controller struct
task           uint8             task type
startDate      time.Time         date from which task is valid (inclusive)
endDate        time.Time         date after which task is invalid
monday         bool              task enabled on Monday if true
tuesday        bool              task enabled on Tuesday if true
wednesday      bool              task enabled on Wednesday if true
thursday       bool              task enabled on Thursday if true
friday         bool              task enabled on Monday if true
saturday       bool              task enabled on Friday if true
sunday         bool              task enabled on Sunday if true
startTime      time.Time         time at which task is scheduled
door           uint8             door ([1..4] to which task is linked
moreCards      uint8             number of 'more cards' for the MORE CARDS task type
timeout        maximum time to wait for a response from a controller

Returns an `AddTaskResponse`.
```
Task types:
```
0:  control door
1:  unlock door
2:  lock door
3:  disable time profiles
4:  enable time profiles
5:  enable card, no password
6:  enable card+IN password
7:  enable card+password
8:  enable more cards
9:  disable more cards
10: trigger once
11: disable pushbutton
12: enable pushbutton
```

### `RefreshTaskList`
```
RefreshTaskList(u Uhppoted, controller TController, timeout time.Duration) (RefreshTaskListResponse,error)

u           Uhppoted struct initialised with the bind address, broadcast address, etc
controller  uint32|Controller controller serial number or {id, address, protocol} Controller struct
timeout     maximum time to wait for a response from a controller

Returns a `RefreshTaskListResponse`.
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

### `SetIPv4Response`

Container class for the decoded response from a _SetIPv4_ request.
```
type SetIPv4Response struct {
    Controller uint32     `json:"controller"`    // controller serial number
    Ok         bool       `json:"ok"`            // succeeded/failed
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
type GetListenerResponse struct {
  Controller  uint32      `json:"controller"` // controller serial number
  Address     netip.Addr  `json:"address"`    // event listener IPv4 address
  Port        uint16      `json:"port"`       // event listener UDP port
  Interval    uint8       `json:"interval"`   // auto-send interval (seconds)
}
```

### `SetListenerResponse`

Container class for the decoded response from a _SetListener_ request.
```
type SetListenerResponse struct {
  Controller  uint32  `json:"controller"` // controller serial number
  Ok          bool    `json:"ok"`         // true if request succeeded
}
```

### `GetListenerAddrPortResponse`

Container class for the decoded response from a _GetListener_ request.
```
type GetListenerAddrPortResponse struct {
  Controller  uint32          `json:"controller"` // controller serial number
  Listener    netip.AddrPort  `json:"listener"`   // event listener IPv4 address:port
  Interval    uint8           `json:"interval"`   // auto-send interval (seconds)
}
```

### `SetListenerAddrPortResponse`

Container class for the decoded response from a _SetListenerAddrPort_ request.
```
type SetListenerAddrPortResponse struct {
  Controller  uint32  `json:"controller"` // controller serial number
  Ok          bool    `json:"ok"`         // true if request succeeded
}
```

### `GetDoorResponse`

Container class for the decoded response from a _GetDoor_ request.
```
type GetDoorResponse struct {
  Controller  uint32  `json:"controller"` // controller serial number
  Door        uint8   `json:"door"`       // door ID ([1..4])
  Mode        uint8   `json:"mode"`       // control mode (1:normally open, 2:normally closed. 3:controlled)
  Delay       uint8   `json:"delay"`      // unlock delay(seconds)
}
```

### `SetDoorResponse`

Container class for the decoded response from a _SetDoor_ request.
```
type SetDoorResponse struct {
  Controller  uint32  `json:"controller"` // controller serial number
  Door        uint8   `json:"door"`       // door ID ([1..4])
  Mode        uint8   `json:"mode"`       // control mode (1:normally open, 2:normally closed. 3:controlled)
  Delay       uint8   `json:"delay"`      // unlock delay(seconds)
}
```

### `SetDoorPasscodesResponse`

Container class for the decoded response from a _SetDoorPasscodes_ request.
```
type SetDoorPasscodesResponse struct {
  Controller  uint32  `json:"controller"` // controller serial number
  Ok          bool    `json:"ok"`         // true if request succeeded
}
```

### `OpenDoorResponse`

Container class for the decoded response from an _OpenDoor_ request.
```
type OpenDoorResponse struct {
  Controller  uint32  `json:"controller"` // controller serial number
  Ok          bool    `json:"ok"`         // true if request succeeded
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

### `GetCardsResponse`

Container class for the decoded response from a _GetCards_ request.
```
type GetStatusResponse struct {
  Controller  uint32     `json:"controller"` // controller serial number
  Cards       uint32     `json:"cards"`      // number of stored cards
}
```

### `GetCardResponse`

Container class for the decoded response from a _GetCard_ request.
```
type GetCardResponse struct {
  Controller  uint32     `json:"controller"`  // controller serial number
  Card        uint32     `json:"card"`        // card number
  StartDate   time.Time  `json:"start-date"`  // 'valid from' date
  EndDate     time.Time  `json:"end-date"`    // 'valid until' date
  Door1       uint8      `json:"door-1"`      // access permissions for door 1
  Door2       uint8      `json:"door-2"`      // access permissions for door 2
  Door3       uint8      `json:"door-3"`      // access permissions for door 3
  Door4       uint8      `json:"door-4"`      // access permissions for door 4
  PIN         uint32     `json:"PIN"`         // (optional) PIN code [0..999999], 0 for none
}
```

### `GetCardAtIndexResponse`

Container class for the decoded response from a _GetCardAtIndex_ request.
```
type GetCardAtIndexResponse struct {
  Controller  uint32     `json:"controller"`  // controller serial number
  Card        uint32     `json:"card"`        // card number
  StartDate   time.Time  `json:"start-date"`  // 'valid from' date
  EndDate     time.Time  `json:"end-date"`    // 'valid until' date
  Door1       uint8      `json:"door-1"`      // access permissions for door 1
  Door2       uint8      `json:"door-2"`      // access permissions for door 2
  Door3       uint8      `json:"door-3"`      // access permissions for door 3
  Door4       uint8      `json:"door-4"`      // access permissions for door 4
  PIN         uint32     `json:"PIN"`         // (optional) PIN code [0..999999], 0 for none
}

Card is 0 if there is no record at the index, 0xffffffff if the record has been deleted.
```

### `PutCardResponse`

Container class for the decoded response from a _PutCard_ request.
```
type PutCardResponse struct {
  Controller  uint32  `json:"controller"` // controller serial number
  Ok          bool    `json:"ok"`         // true if request succeeded
}
```

### `DeleteCardResponse`

Container class for the decoded response from a _DeleteCard_ request.
```
type DeleteCardResponse struct {
  Controller  uint32  `json:"controller"` // controller serial number
  Ok          bool    `json:"ok"`         // true if request succeeded
}
```

### `DeleteAllCardsResponse`

Container class for the decoded response from a _DeleteAllCards_ request.
```
type DeleteAllCardsResponse struct {
  Controller  uint32  `json:"controller"` // controller serial number
  Ok          bool    `json:"ok"`         // true if request succeeded
}
```

### `GetEventResponse`

Container class for the decoded response from a _GetEvent_ request.
```
type GetEventResponse struct {
  Controller     uint32     `json:"controller"`     // controller serial number
  Index          uint32     `json:"index"`          // event index
  Timestamp      time.Time  `json:"timestamp"`      // event timestamp
  EventType      uint8      `json:"event-type"`     // event type 
  AccessGranted  bool       `json:"acces-granted"`  // true if the door was unlocked
  Door           uint8      `json:"door"`           // door no. for card and door events
  Direction      uint8      `json:"direction"`      // direction for card and door events
  Card           uint32     `json:"card"`           // card number (for card events)
  Reason         uint8      `json:"reason"`         // reason code
}

Event types:
- 0:   unknown
- 1:   card
- 2:   door
- 3:   alarm
- 4:   system
- 255: overwritten

Direction:
- 1: in
- 2: out

Reasons:
0:      unknown
1:      card ok
5:      card denied (PC control)
6:      card denied (no access)
7:      card denied (invalid password)
8:      card denied (anti-passback)
9:      card denied (more cards)
10:     card denied (first card open)
11:     card denied (door normally closed)
12:     card denied (door interlock)
13:     card denied (limited times)
15:     card denied (invalid timezone)
18:     card denied
20:     door pushbutton
23:     door open
24:     door closed
25:     door supervisor password open
28:     controller power on
29:     controller reset
30:     pushbutton denied (disabled by task)
31:     pushbutton denied (forced lock)
32:     pushbutton denied (not online)
33:     pushbutton denied (door interlock
34:     alarm (threat)
37:     alarm (open too long)
38:     alarm (forced open)
39:     alarm (fire)
40:     alarm (forced close)
41:     alarm (tamper detect)
42:     alarm (24x7 zone)
43:     alarm (emergency call)
44:     remote open door
45:     remote open door (USB reader)
```

### `GetEventIndexResponse`

Container class for the decoded response from a _GetEventIndex_ request.
```
type GetEventResponse struct {
  Controller  uint32  `json:"controller"` // controller serial number
  Index       uint32  `json:"index"`      // event index
}
```

### `GetEventIndexResponse`

Container class for the decoded response from a _SetEventIndex_ request.
```
type SetEventResponse struct {
  Controller  uint32  `json:"controller"` // controller serial number
  Ok          bool    `json:"ok"`         // succeeded/failed
}
```

### `RecordSpecialEventsResponse`

Container class for the decoded response from a _RecordSpecialEvents_ request.
```
type RecordSpecialEventsResponse struct {
  Controller  uint32  `json:"controller"` // controller serial number
  Ok          bool    `json:"ok"`         // succeeded/failed
}
```

### `GetTimeProfileResponse`

Container class for the decoded response from a _GetTimeProfile_ request.
```
type GetTimeProfileResponse struct {
  Controller     uint32     `json:"controller"`      // controller serial number
  Profile        uint8      `json:"profile"`         // profile ID [2..254]
  StartDate      time.Time  `json:"start-date"`      // date from which profile is valid (inclusive)
  EndDate        time.Time  `json:"end-date"`        // date after which profile is invalid
  Monday         bool       `json:"monday"`          // profile enabled on Monday if true
  Tuesday        bool       `json:"tuesday"`         // profile enabled on Tuesday if true
  Wednesday      bool       `json:"wednesday"`       // profile enabled on Wednesday if true
  Thursday       bool       `json:"thursday"`        // profile enabled on Thursday if true
  Friday         bool       `json:"friday"`          // profile enabled on Monday if true
  Saturday       bool       `json:"saturday"`        // profile enabled on Friday if true
  Sunday         bool       `json:"sunday"`          // profile enabled on Sunday if true
  Segment1Start  time.Time  `json:"segment1-start"`  // start time for first time segment
  Segment1End    time.Time  `json:"segment1-end"`    // end time for first time segment
  Segment2Start  time.Time  `json:"segment2-start"`  // start time for second time segment
  Segment2End    time.Time  `json:"segment2-end"`    // end time for second time segment
  Segment3Start  time.Time  `json:"segment3-start"`  // start time for third time segment
  Segment3End    time.Time  `json:"segment3-end"`    // end time for third time segment
  LinkedProfile  uint8      `json:"linked-profile"`  // ID of linked profile (0 if not linked)
}
```

### `SetTimeProfileResponse`

Container class for the decoded response from a _SetTimeProfile_ request.
```
type SetTimeProfileResponse struct {
  Controller  uint32  `json:"controller"` // controller serial number
  Ok          bool    `json:"ok"`         // succeeded/failed
}
```

### `ClearTimeProfilesResponse`

Container class for the decoded response from a _ClearTimeProfiles_ request.
```
type ClearTimeProfilesResponse struct {
  Controller  uint32  `json:"controller"` // controller serial number
  Ok          bool    `json:"ok"`         // succeeded/failed
}
```

### `AddTaskResponse`

Container class for the decoded response from an _AddTask_ request.
```
type AddTaskResponse struct {
  Controller  uint32  `json:"controller"` // controller serial number
  Ok          bool    `json:"ok"`         // succeeded/failed
}
```

### `RefreshTasklistResponse`

Container class for the decoded response from a _RefreshTasklist_ request.
```
type RefreshTasklistResponse struct {
  Controller  uint32  `json:"controller"` // controller serial number
  Ok          bool    `json:"ok"`         // succeeded/failed
}
```

