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
- [`ClearTaskList`](#cleartasklist)
- [`SetPCControl`](#setpccontrol)
- [`SetInterlock`](#setinterlock)
- [`ActivateKeypads`](#activatekeypads)
- [`GetAntiPassback`](#getantipassback)
- [`SetAntiPassback`](#setantipassback)
- [`RestoreDefaultParameters`](#restoredefaultparameters)
- [`Listen`](#listen)

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

### `FindControllers`
FindControllers retrieves a list of all UHPPOTE controllers accessible via UDP broadcast
on the local LAN.
```
FindControllers(u, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a ``:
```
type  struct { 
}
```

### `GetController`
Retrieves the system information for an access controller.
```
GetController(u, controller, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `GetControllerResponse`:
```
type GetControllerResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  IpAddress           IPv4                `json:"ip-address"`     // controller IPv4 address, e.g. 192.168.1.100
  SubnetMask          IPv4                `json:"netmask"`        // controller IPv4 netmask, e.g. 255.255.255.0
  Gateway             IPv4                `json:"gateway"`        // controller IPv4 gateway address, e.g. 192.168.1.1
  MACAddress          MAC                 `json:"MAC"`            // controller MAC address, e.g. 52:fd:fc:07:21:82
  Version             version             `json:"version"`        // controller firmware version, e.g. v6.62
  Date                date                `json:"date"`           // controller firmware release date, e.g. 2020-12-31
}
```

### `SetIPv4`
Sets the controller IPv4 address, netmask and gateway address.
```
SetIPv4(u, controller, address, netmask, gateway, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- address       IPv4            controller IPv4 address
- netmask       IPv4            controller IPv4 subnet mask
- gateway       IPv4            controller gateway IPv4 address
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `SetIPv4Response`:
```
type SetIPv4Response struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `GetTime`
Retrieves the access controller system date and time.
```
GetTime(u, controller, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `GetTimeResponse`:
```
type GetTimeResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  DateTime            datetime            `json:"date-time"`      // controller system date/time
}
```

### `SetTime`
Sets the access controller system date and time.
```
SetTime(u, controller, date-time, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- date-time     datetime        date/time to which to set controller system time
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `SetTimeResponse`:
```
type SetTimeResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  DateTime            datetime            `json:"date-time"`      // controller system date/time
}
```

### `GetListener`
Retrieves the access controller event listener IPv4 address:port and auto-send interval.
```
GetListener(u, controller, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `GetListenerResponse`:
```
type GetListenerResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Address             IPv4                `json:"address"`        // event listener IPv4 address
  Port                uint16              `json:"port"`           // event listener IPv4 port
  Interval            uint8               `json:"interval"`       // status auto-send interval (seconds)
}
```

### `SetListener`
Sets the access controller event listener IPv4 address:port and auto-send interval.
```
SetListener(u, controller, address, port, interval, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- address       IPv4            IPv4 address of host to receive controller events
- port          uint16          UDP port of host for controller events
- interval      uint8           status auto-send interval (seconds). A '0' interval disables auto-send.
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `SetListenerResponse`:
```
type SetListenerResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `GetListenerAddrPort`
Retrieves the access controller event listener IPv4 address:port and auto-send interval.
```
GetListenerAddrPort(u, controller, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `GetListenerAddrPortResponse`:
```
type GetListenerAddrPortResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Listener            address:port        `json:"listener"`       // event listener IPv4 address:port
  Interval            uint8               `json:"interval"`       // status auto-send interval (seconds)
}
```

### `SetListenerAddrPort`
Sets the access controller event listener IPv4 address:port and auto-send interval.
```
SetListenerAddrPort(u, controller, listener, interval, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- listener      address:port    IPv4 address:port of host to receive controller events
- interval      uint8           status auto-send interval (seconds). A '0'interval disables auto-send.
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `SetListenerAddrPortResponse`:
```
type SetListenerAddrPortResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `GetDoor`
Retrieves the control mode and unlock delay time for an access controller door.
```
GetDoor(u, controller, door, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- door          uint8           door ID ([1..4])
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `GetDoorResponse`:
```
type GetDoorResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Door                uint8               `json:"door"`           // door ID ([1..4]
  Mode                uint8               `json:"mode"`           // control mode (1:normally open, 2:normally closed. 3:controlled)
  Delay               uint8               `json:"delay"`          // unlock delay (seconds)
}
```

### `SetDoor`
Sets the control mode and unlock delay time for an access controller door.
```
SetDoor(u, controller, door, mode, delay, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- door          uint8           door ID ([1..4])
- mode          uint8           control mode (1:normally open, 2:normally closed. 3:controlled)
- delay         uint8           unlock delay (seconds))
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `SetDoorResponse`:
```
type SetDoorResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Door                uint8               `json:"door"`           // 
  Mode                uint8               `json:"mode"`           // 
  Delay               uint8               `json:"delay"`          // 
}
```

### `SetDoorPasscodes`
Sets up to 4 passcodes for a controller door.
```
SetDoorPasscodes(u, controller, door, passcode 1, passcode 2, passcode 3, passcode 4, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- door          uint8           door ID ([1..4])
- passcode 1    pin             supervisor passcode ([0..99999]), 0 for 'none'
- passcode 2    pin             supervisor passcode ([0..99999]), 0 for 'none'
- passcode 3    pin             supervisor passcode ([0..99999]), 0 for 'none'
- passcode 4    pin             supervisor passcode ([0..99999]), 0 for 'none'
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `SetDoorPasscodesResponse`:
```
type SetDoorPasscodesResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `OpenDoor`
Unlocks a door controlled by an access controller.
```
OpenDoor(u, controller, door, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- door          uint8           door ID ([1..4])
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns an `OpenDoorResponse`:
```
type OpenDoorResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `GetStatus`
Retrieves the system status from an access controller.
```
GetStatus(u, controller, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `GetStatusResponse`:
```
type GetStatusResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  SystemDate          shortdate           `json:"system-date"`    // controller system date, e.g. 2025-07-21
  SystemTime          time                `json:"system-time"`    // controller system time, e.g. 13:25:47
  Door1Open           bool                `json:"door-1-open"`    // door 1 open sensor
  Door2Open           bool                `json:"door-2-open"`    // door 2 open sensor
  Door3Open           bool                `json:"door-3-open"`    // door 3 open sensor
  Door4Open           bool                `json:"door-4-open"`    // door 4 open sensor
  Door1Button         bool                `json:"door-1-button"`  // door 1 button pressed
  Door2Button         bool                `json:"door-2-button"`  // door 2 button pressed
  Door3Button         bool                `json:"door-3-button"`  // door 3 button pressed
  Door4Button         bool                `json:"door-4-button"`  // door 4 button pressed
  Relays              uint8               `json:"relays"`         // bitset of door unlock relay states
  Inputs              uint8               `json:"alarm-inputs"`   // bitset of alarm inputs
  SystemError         uint8               `json:"system-error"`   // system error code
  SpecialInfo         uint8               `json:"special-info"`   // absolutely no idea
  EventIndex          uint32              `json:"event-index"`    // last event index
  EventType           uint8               `json:"event-type"`     // last event type
  EventAccessGranted  bool                `json:"event-granted"`  // last event access granted
  EventDoor           uint8               `json:"event-door"`     // last event door
  EventDirection      uint8               `json:"event-direction"` // last event door direction (0: in, 1: out)
  EventCard           uint32              `json:"event-card"`     // last event card number
  EventTimestamp      optional datetime   `json:"event-timestamp"` // last event timestamp
  EventReason         uint8               `json:"event-reason"`   // last event reason
  SequenceNo          uint32              `json:"sequence-no"`    // packet sequence number
}
```

### `GetCards`
Retrieves the number of cards stored on an access controller.
```
GetCards(u, controller, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `GetCardsResponse`:
```
type GetCardsResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Cards               uint32              `json:"cards"`          // number of stored cards
}
```

### `GetCard`
Retrieves the card record for a given card number.
```
GetCard(u, controller, card number, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- card number   uint32          card number
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `GetCardResponse`:
```
type GetCardResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Card                uint32              `json:"card"`           // card number
  StartDate           optional date       `json:"start-date"`     // 'valid from' date
  EndDate             optional date       `json:"end-date"`       // 'valid until' date
  Door1               uint8               `json:"door-1"`         // access permissions for door 1
  Door2               uint8               `json:"door-2"`         // access permissions for door 2
  Door3               uint8               `json:"door-3"`         // access permissions for door 3
  Door4               uint8               `json:"door-4"`         // access permissions for door 4
  PIN                 pin                 `json:"PIN"`            // (optional) PIN code [0..999999], 0 for none
}
```

### `GetCardAtIndex`
Retrieves the card record stored at a given index.
```
GetCardAtIndex(u, controller, index, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- index         uint32          card record index
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `GetCardAtIndexResponse`:
```
type GetCardAtIndexResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Card                uint32              `json:"card"`           // card number
  StartDate           optional date       `json:"start-date"`     // 'valid from' date
  EndDate             optional date       `json:"end-date"`       // 'valid until' date
  Door1               uint8               `json:"door-1"`         // access permissions for door 1
  Door2               uint8               `json:"door-2"`         // access permissions for door 2
  Door3               uint8               `json:"door-3"`         // access permissions for door 3
  Door4               uint8               `json:"door-4"`         // access permissions for door 4
  PIN                 pin                 `json:"PIN"`            // (optional) PIN code [0..999999], 0 for none
}
```

### `PutCard`
Creates or updates a card record stored on an access controller.
```
PutCard(u, controller, card, start date, end date, door 1, door 2, door 3, door 4, PIN, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- card          uint32          card number
- start date    date            date from which card is valid
- end date      date            date after which card is invalid
- door 1        uint8           access permissions for door 1 (0: no access, 1: 24/7 access, [2..253] time profile)
- door 2        uint8           access permissions for door 2 (0: no access, 1: 24/7 access, [2..253] time profile)
- door 3        uint8           access permissions for door 3 (0: no access, 1: 24/7 access, [2..253] time profile)
- door 4        uint8           access permissions for door 4 (0: no access, 1: 24/7 access, [2..253] time profile)
- PIN           pin             optional PIN code [0..999999] (0 for none)
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `PutCardResponse`:
```
type PutCardResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `DeleteCard`
Removes a card record stored on a controller.
```
DeleteCard(u, controller, card number, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- card number   uint32          card number
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `DeleteCardResponse`:
```
type DeleteCardResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `DeleteAllCards`
Deletes all card records stored on an access controller.
```
DeleteAllCards(u, controller, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `DeleteAllCardsResponse`:
```
type DeleteAllCardsResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `GetEvent`
Retrieves an event record stored on an access controller.
```
GetEvent(u, controller, event index, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- event index   uint32          index of event to retrieve
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `GetEventResponse`:
```
type GetEventResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Index               uint32              `json:"index"`          // event index
  EventType           uint8               `json:"event-type"`     // event type 
  AccessGranted       bool                `json:"granted"`        // true if the door was unlocked
  Door                uint8               `json:"door"`           // door no. ([1..4]) for card and door events
  Direction           uint8               `json:"direction"`      // direction (1:IN, 2:OUT) for card and door events
  Card                uint32              `json:"card"`           // card number (for card events)
  Timestamp           optional datetime   `json:"timestamp"`      // event timestamp
  Reason              uint8               `json:"reason"`         // reason code
}
```

### `GetEventIndex`
Retrieves the downloaded event index from an access controller.
```
GetEventIndex(u, controller, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `GetEventIndexResponse`:
```
type GetEventIndexResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Index               uint32              `json:"index"`          // event index
}
```

### `SetEventIndex`
Sets the downloaded event index on an access controller.
```
SetEventIndex(u, controller, event index, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- event index   uint32          downloaded event index
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `SetEventIndexResponse`:
```
type SetEventIndexResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `RecordSpecialEvents`
Enables/disables events for door opened, door closed and door button pressed.
```
RecordSpecialEvents(u, controller, enabled, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- enabled       bool            enables door opened, door closed and button pressed events if true
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `RecordSpecialEventsResponse`:
```
type RecordSpecialEventsResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `GetTimeProfile`
Retrieves the requested access time profile from a controller.
```
GetTimeProfile(u, controller, profile, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- profile       uint8           profile ID ([2..254] to retrieve
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `GetTimeProfileResponse`:
```
type GetTimeProfileResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Profile             uint8               `json:"profile"`        // profile ID [2..254]
  StartDate           optional date       `json:"start-date"`     // date from which profile is valid (inclusive)
  EndDate             optional date       `json:"end-date"`       // date after which profile is invalid
  Monday              bool                `json:"monday"`         // profile enabled on Monday if true
  Tuesday             bool                `json:"tuesday"`        // profile enabled on Tuesday if true
  Wednesday           bool                `json:"wednesday"`      // profile enabled on Wednesday if true
  Thursday            bool                `json:"thursday"`       // profile enabled on Thursday if true
  Friday              bool                `json:"friday"`         // profile enabled on Monday if true
  Saturday            bool                `json:"saturday"`       // profile enabled on Friday if true
  Sunday              bool                `json:"sunday"`         // profile enabled on Sunday if true
  Segment1Start       HHmm                `json:"segment1-start"` // start time for first time segment
  Segment1End         HHmm                `json:"segment1-end"`   // end time for first time segment
  Segment2Start       HHmm                `json:"segment2-start"` // start time for second time segment
  Segment2End         HHmm                `json:"segment2-end"`   // end time for second time segment
  Segment3Start       HHmm                `json:"segment3-start"` // start time for third time segment
  Segment3End         HHmm                `json:"segment3-end"`   // end time for third time segment
  LinkedProfile       uint8               `json:"linked-profile"` // ID of linked profile (0 if not linked)
}
```

### `SetTimeProfile`
Adds or updates an access time profile stored on a controller.
```
SetTimeProfile(u, controller, profile, start date, end date, monday, tuesday, wednesday, thursday, friday, saturday, sunday, segment 1 start, segment 1 end, segment 2 start, segment 2 end, segment 3 start, segment 3 end, linked profile id, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- profile       uint8           profile Id ([2..254] to create/update
- start date    date            date from which profile is valid (inclusive)
- end date      date            date after which profile is invalid
- monday        bool            profile valid on Monday if true
- tuesday       bool            profile valid on Tuesday if true
- wednesday     bool            profile valid on Wednesday if true
- thursday      bool            profile valid on Thursday if true
- friday        bool            profile valid on Friday if true
- saturday      bool            profile valid on Saturday if true
- sunday        bool            profile valid on Sunday if true
- segment 1 start  HHmm            start time for first time segment
- segment 1 end  HHmm            end time for first time segment
- segment 2 start  HHmm            start time for second time segment
- segment 2 end  HHmm            end time for second time segment
- segment 3 start  HHmm            start time for third time segment
- segment 3 end  HHmm            end time for third time segment
- linked profile id  uint8           ID of linked profile (0 if not linked)
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `SetTimeProfileResponse`:
```
type SetTimeProfileResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `ClearTimeProfiles`
Clears all access time profiles stored on a controller.
```
ClearTimeProfiles(u, controller, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `ClearTimeProfilesResponse`:
```
type ClearTimeProfilesResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `AddTask`
Creates a scheduled task.

Task types
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
AddTask(u, controller, task, start date, end date, monday, tuesday, wednesday, thursday, friday, saturday, sunday, start time, door, more cards, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- task          uint8           task type
- start date    date            date from which task is scheduled (inclusive)
- end date      date            date after which task no longer scheduled
- monday        bool            task enabled on Monday if true
- tuesday       bool            task enabled on Tuesday if true
- wednesday     bool            task enabled on Wednesday if true
- thursday      bool            task enabled on Thursday if true
- friday        bool            task enabled on Monday if true
- saturday      bool            task enabled on Friday if true
- sunday        bool            task enabled on Sunday if true
- start time    HHmm            time at which task is scheduled
- door          uint8           door ([1..4] to which task is linked
- more cards    uint8           number of 'more cards' for the MORE CARDS task type
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns an `AddTaskResponse`:
```
type AddTaskResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `RefreshTaskList`
Updates scheduler with newly created scheduled tasks.
```
RefreshTaskList(u, controller, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `RefreshTaskListResponse`:
```
type RefreshTaskListResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `ClearTaskList`
Removes all scheduled tasks.
```
ClearTaskList(u, controller, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `ClearTaskListResponse`:
```
type ClearTaskListResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `SetPCControl`
Enables remote access control. Remote access control will remain in effect provided the controller
receives a message from the host at least once every 30 seconds.
```
SetPCControl(u, controller, enabled, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- enabled       bool            enables/disables remote access control
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `SetPCControlResponse`:
```
type SetPCControlResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `SetInterlock`
Sets the door interlock mode on an access controller.

The following interlock modes are supported:
- 0: disabled
- 1: doors 1&2
- 2: doors 3&4
- 3: doors 1&2, doors 3&4
- 4: doors 1,2&3
- 8: doors 1,2,3&4
```
SetInterlock(u, controller, interlock, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- interlock     uint8           door interlock mode
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `SetInterlockResponse`:
```
type SetInterlockResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `ActivateKeypads`
Enables/disables door keypad readers.
```
ActivateKeypads(u, controller, reader 1, reader 2, reader 3, reader 4, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- reader 1      bool            enables/disable the keypad for reader 1
- reader 2      bool            enables/disables the keypad for reader 2
- reader 3      bool            enables/disables the keypad for reader 3
- reader 4      bool            enables/disables the keypad for reader 4
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns an `ActivateKeypadsResponse`:
```
type ActivateKeypadsResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `GetAntiPassback`
Retrieves the anti-passback mode for an access controller. The anti-passback mode
will be one of the following:
- 0: disabled
- 1: readers 1:2; 3:4 (independently)
- 2: readers (1,3):(2,4)
- 3: readers 1:(2,3)
- 4: readers 1:(2,3,4)
```
GetAntiPassback(u, controller, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `GetAntiPassbackResponse`:
```
type GetAntiPassbackResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Antipassback        uint8               `json:"antipassback"`   // anti-passback mode
}
```

### `SetAntiPassback`
Sets the access controller anti-passback mode.

The following modes are supported:
- 0: disabled
- 1: doors 1&2, doors 3&4
- 2: doors 1&3, doors 2&4
- 3: door 1 & doors 2,3
- 4: door 1 & doors 1,2,3
```
SetAntiPassback(u, controller, antipassback, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- antipassback  uint8           anti-passback mode
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `SetAntiPassbackResponse`:
```
type SetAntiPassbackResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `RestoreDefaultParameters`
Restores the controller configuration to the manufacturer defaults.
```
RestoreDefaultParameters(u, controller, timeout)

where:
- u             Uhppoted        Uhppoted struct initialised with the bind address, broadcast address, etc
- controller    controller      uint32|Controller controller serial number or {id, address, protocol} Controller struct
- timeout       time.Duration   maximum time to wait for a response from a controller
```
Returns a `RestoreDefaultParametersResponse`:
```
type RestoreDefaultParametersResponse struct { 
  Controller          uint32              `json:"controller"`     // controller serial number
  Ok                  bool                `json:"ok"`             // succeeded/failed
}
```

### `Listen`
Listens for controller events sent via UDP to the designated listener host.
```
Listen(u, events chan ListenerEvent, errors chan error, interrupt chan os.Signal) error {

where:
- u          Uhppoted            Uhppoted struct initialised with the bind address, broadcast address, etc
- events     chan ListenerEvent  events channel
- errors     chan error          errors channel
- interrupt  chan os.Signal      interrupts channel
```

- received events are posted to the _events_ channel
- non-fatal errors are posted to the _errors_ channel
- signals posted on the interrupts channel cause the event listener to close and return

```
type ListenerEvent struct { 
  Controller          uint32        `json:"controller"`      // controller serial number
  SystemDate          Date          `json:"system-date"`     // controller system date, e.g. 2025-07-21
  SystemTime          Time          `json:"system-time"`     // controller system time, e.g. 13:25:47
  Door1Open           bool          `json:"door-1-open"`     // door 1 open sensor
  Door2Open           bool          `json:"door-2-open"`     // door 2 open sensor
  Door3Open           bool          `json:"door-3-open"`     // door 3 open sensor
  Door4Open           bool          `json:"door-4-open"`     // door 4 open sensor
  Door1Button         bool          `json:"door-1-button"`   // door 1 button pressed
  Door2Button         bool          `json:"door-2-button"`   // door 2 button pressed
  Door3Button         bool          `json:"door-3-button"`   // door 3 button pressed
  Door4Button         bool          `json:"door-4-button"`   // door 4 button pressed
  Relays              uint8         `json:"relays"`          // bitset of door unlock relay states
  Inputs              uint8         `json:"alarm-inputs"`    // bitset of alarm inputs
  SystemError         uint8         `json:"system-error"`    // system error code
  SpecialInfo         uint8         `json:"special-info"`    // absolutely no idea
  EventIndex          uint32        `json:"event-index"`     // last event index
  EventType           uint8         `json:"event-type"`      // last event type
  EventAccessGranted  bool          `json:"event-granted"`   // last event access granted
  EventDoor           uint8         `json:"event-door"`      // last event door
  EventDirection      uint8         `json:"event-direction"` // last event door direction (0: in, 1: out)
  EventCard           uint32        `json:"event-card"`      // last event card number
  EventTimestamp      DateTime      `json:"event-timestamp"` // last event timestamp
  EventReason         uint8         `json:"event-reason"`    // last event reason
  SequenceNo          uint32        `json:"sequence-no"`     // packet sequence number
}
```

