![build](https://github.com/uhppoted/uhppoted-lib-go/workflows/build/badge.svg)

# uhppoted-lib-go

**IN DEVELOPMENT**

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

**API**

- [`FindControllers`](API.md#findcontrollers)
- [`GetController`](API.md#getcontroller)
- [`SetIPv4`](API.md#setipv4)
- [`GetTime`](API.md#gettime)
- [`SetTime`](API.md#settime)
- [`GetListener`](API.md#getlistener)
- [`SetListener`](API.md#setlistener)
- [`GetListenerAddrPort`](API.md#getlisteneraddrport)
- [`SetListenerAddrPort`](API.md#setlisteneraddrport)
- [`GetDoor`](API.md#getdoor)
- [`SetDoor`](API.md#setdoor)
- [`SetDoorPasscodes`](API.md#setdoorpasscodes)
- [`OpenDoor`](API.md#opendoor)
- [`GetStatus`](API.md#getstatus)
- [`GetCards`](API.md#getcards)
- [`GetCard`](API.md#getcard)
- [`GetCardAtIndex`](API.md#getcardatindex)
- [`PutCard`](API.md#putcard)
- [`DeleteCard`](API.md#deletecard)
- [`DeleteAllCards`](API.md#deleteallcards)
- [`GetEvent`](API.md#getevent)
- [`GetEventIndex`](API.md#geteventindex)
- [`SetEventIndex`](API.md#seteventindex)
- [`RecordSpecialEvents`](API.md#recordspecialevents)
- [`GetTimeProfile`](API.md#gettimeprofile)
- [`SetTimeProfile`](API.md#settimeprofile)
- [`ClearTimeProfiles`](API.md#cleartimeprofiles)
- [`AddTask`](API.md#addtask)
- [`RefreshTaskList`](API.md#refreshtasklist)
- [`ClearTaskList`](API.md#cleartasklist)
- [`SetPCControl`](API.md#setpccontrol)
- [`SetInterlock`](API.md#setinterlock)
- [`ActivateKeypads`](API.md#activatekeypads)
