package uhppoted

import (
    "fmt"
    "net/netip"
    "time"

    "github.com/uhppoted/uhppoted-lib-go/uhppoted/responses"
    "github.com/uhppoted/uhppoted-lib-go/uhppoted/entities"
)

var Expected = struct { 
    FindControllers []responses.GetControllerResponse
    GetController responses.GetControllerResponse
    SetIPv4 responses.SetIPv4Response
    GetTime responses.GetTimeResponse
    SetTime responses.SetTimeResponse
    GetListener responses.GetListenerResponse
    SetListener responses.SetListenerResponse
    GetListenerAddrPort responses.GetListenerAddrPortResponse
    SetListenerAddrPort responses.SetListenerAddrPortResponse
    GetDoor responses.GetDoorResponse
    SetDoor responses.SetDoorResponse
    SetDoorPasscodes responses.SetDoorPasscodesResponse
    OpenDoor responses.OpenDoorResponse
    GetStatus responses.GetStatusResponse
    GetStatusNoEvent responses.GetStatusResponse
    GetCards responses.GetCardsResponse
    GetCard responses.GetCardResponse
    GetCardNotFound responses.GetCardResponse
    GetCardAtIndex responses.GetCardAtIndexResponse
    GetCardAtIndexNotFound responses.GetCardAtIndexResponse
    GetCardAtIndexDeleted responses.GetCardAtIndexResponse
    PutCard responses.PutCardResponse
    DeleteCard responses.DeleteCardResponse
    DeleteAllCards responses.DeleteAllCardsResponse
    GetEvent responses.GetEventResponse
    GetEventNotFound responses.GetEventResponse
    GetEventOverwritten responses.GetEventResponse
    GetEventIndex responses.GetEventIndexResponse
    SetEventIndex responses.SetEventIndexResponse
    RecordSpecialEvents responses.RecordSpecialEventsResponse
    GetTimeProfile responses.GetTimeProfileResponse
    SetTimeProfile responses.SetTimeProfileResponse
    ClearTimeProfiles responses.ClearTimeProfilesResponse
    AddTask responses.AddTaskResponse
    RefreshTaskList responses.RefreshTaskListResponse
    ClearTaskList responses.ClearTasklistResponse
    SetPcControl responses.SetPCControlResponse
    SetInterlock responses.SetInterlockResponse
    ActivateKeypads responses.ActivateKeypadsResponse
}{ 
    FindControllers: []responses.GetControllerResponse{
        responses.GetControllerResponse{
            Controller: 201020304,
            IpAddress:  netip.MustParseAddr("192.168.1.101"),
            SubnetMask: netip.MustParseAddr("255.255.255.0"),
            Gateway:    netip.MustParseAddr("192.168.1.1"),
            MACAddress: "52:fd:fc:07:21:82",
            Version:    "v6.62",
            Date:       entities.MustParseDate("2020-01-01"),
        },
        responses.GetControllerResponse{
            Controller: 303986753,
            IpAddress:  netip.MustParseAddr("192.168.1.100"),
            SubnetMask: netip.MustParseAddr("255.255.255.0"),
            Gateway:    netip.MustParseAddr("192.168.1.1"),
            MACAddress: "52:fd:fc:07:21:82",
            Version:    "v8.92",
            Date:       entities.MustParseDate("2019-08-15"),
        },
        responses.GetControllerResponse{
            Controller: 405419896,
            IpAddress:  netip.MustParseAddr("192.168.1.100"),
            SubnetMask: netip.MustParseAddr("255.255.255.0"),
            Gateway:    netip.MustParseAddr("192.168.1.1"),
            MACAddress: "00:12:23:34:45:56",
            Version:    "v8.92",
            Date:       entities.MustParseDate("2018-11-05"),
        },
    },

    GetController: responses.GetControllerResponse {
        Controller: 405419896,
        IpAddress: IPv4("192.168.1.100"),
        SubnetMask: IPv4("255.255.255.0"),
        Gateway: IPv4("192.168.1.1"),
        MACAddress: "00:12:23:34:45:56",
        Version: "v8.92",
        Date: entities.MustParseDate("2018-11-05"),
     },

    SetIPv4: responses.SetIPv4Response {
        Controller: 405419896,
        Ok: true,
     },

    GetTime: responses.GetTimeResponse {
        Controller: 405419896,
        DateTime: string2datetime("2025-11-01 12:34:56"),
     },

    SetTime: responses.SetTimeResponse {
        Controller: 405419896,
        DateTime: string2datetime("2025-11-01 12:34:56"),
     },

    GetListener: responses.GetListenerResponse {
        Controller: 405419896,
        Address: IPv4("192.168.1.100"),
        Port: 60001,
        Interval: 13,
     },

    SetListener: responses.SetListenerResponse {
        Controller: 405419896,
        Ok: true,
     },

    GetListenerAddrPort: responses.GetListenerAddrPortResponse {
        Controller: 405419897,
        Listener: addrport("192.168.1.100:60001"),
        Interval: 13,
     },

    SetListenerAddrPort: responses.SetListenerAddrPortResponse {
        Controller: 405419897,
        Ok: true,
     },

    GetDoor: responses.GetDoorResponse {
        Controller: 405419896,
        Door: 4,
        Mode: 3,
        Delay: 7,
     },

    SetDoor: responses.SetDoorResponse {
        Controller: 405419896,
        Door: 4,
        Mode: 2,
        Delay: 17,
     },

    SetDoorPasscodes: responses.SetDoorPasscodesResponse {
        Controller: 405419896,
        Ok: true,
     },

    OpenDoor: responses.OpenDoorResponse {
        Controller: 405419896,
        Ok: true,
     },

    GetStatus: responses.GetStatusResponse {
        Controller: 405419896,
        SystemDate: string2date("2022-08-23"),
        SystemTime: string2time("09:49:39"),
        Door1Open: false,
        Door2Open: true,
        Door3Open: false,
        Door4Open: false,
        Door1Button: false,
        Door2Button: false,
        Door3Button: false,
        Door4Button: true,
        Relays: 7,
        Inputs: 9,
        SystemError: 3,
        SpecialInfo: 39,
        EventIndex: 78,
        EventType: 2,
        EventAccessGranted: true,
        EventDoor: 3,
        EventDirection: 1,
        EventCard: 8165537,
        EventTimestamp: string2datetime("2022-08-23 09:47:06"),
        EventReason: 44,
        SequenceNo: 0,
     },

    GetStatusNoEvent: responses.GetStatusResponse {
        Controller: 405419897,
        SystemDate: string2date("2025-11-23"),
        SystemTime: string2time("14:37:53"),
        Door1Open: true,
        Door2Open: false,
        Door3Open: true,
        Door4Open: true,
        Door1Button: true,
        Door2Button: true,
        Door3Button: false,
        Door4Button: true,
        Relays: 7,
        Inputs: 9,
        SystemError: 27,
        SpecialInfo: 39,
        EventIndex: 0,
        EventType: 0,
        EventAccessGranted: false,
        EventDoor: 0,
        EventDirection: 0,
        EventCard: 0,
        EventTimestamp: string2datetime("0001-01-01 00:00:00"),
        EventReason: 0,
        SequenceNo: 21987,
     },

    GetCards: responses.GetCardsResponse {
        Controller: 405419896,
        Cards: 13579,
     },

    GetCard: responses.GetCardResponse {
        Controller: 405419896,
        Card: 10058400,
        StartDate: string2date("2025-01-01"),
        EndDate: string2date("2025-12-31"),
        Door1: 1,
        Door2: 0,
        Door3: 17,
        Door4: 1,
        PIN: 7531,
     },

    GetCardNotFound: responses.GetCardResponse {
        Controller: 405419896,
        Card: 0,
        StartDate: string2date("0001-01-01"),
        EndDate: string2date("0001-01-01"),
        Door1: 0,
        Door2: 0,
        Door3: 0,
        Door4: 0,
        PIN: 0,
     },

    GetCardAtIndex: responses.GetCardAtIndexResponse {
        Controller: 405419896,
        Card: 10058400,
        StartDate: string2date("2025-01-01"),
        EndDate: string2date("2025-12-31"),
        Door1: 1,
        Door2: 0,
        Door3: 17,
        Door4: 1,
        PIN: 7531,
     },

    GetCardAtIndexNotFound: responses.GetCardAtIndexResponse {
        Controller: 405419896,
        Card: 0,
        StartDate: string2date("0001-01-01"),
        EndDate: string2date("0001-01-01"),
        Door1: 0,
        Door2: 0,
        Door3: 0,
        Door4: 0,
        PIN: 0,
     },

    GetCardAtIndexDeleted: responses.GetCardAtIndexResponse {
        Controller: 405419896,
        Card: 4294967295,
        StartDate: string2date("0001-01-01"),
        EndDate: string2date("0001-01-01"),
        Door1: 0,
        Door2: 0,
        Door3: 0,
        Door4: 0,
        PIN: 0,
     },

    PutCard: responses.PutCardResponse {
        Controller: 405419896,
        Ok: true,
     },

    DeleteCard: responses.DeleteCardResponse {
        Controller: 405419896,
        Ok: true,
     },

    DeleteAllCards: responses.DeleteAllCardsResponse {
        Controller: 405419896,
        Ok: true,
     },

    GetEvent: responses.GetEventResponse {
        Controller: 405419896,
        Index: 13579,
        Timestamp: string2datetime("2025-11-17 12:34:56"),
        EventType: 2,
        AccessGranted: true,
        Door: 4,
        Direction: 2,
        Card: 10058400,
        Reason: 21,
     },

    GetEventNotFound: responses.GetEventResponse {
        Controller: 405419896,
        Index: 24680,
        Timestamp: string2datetime("0001-01-01 00:00:00"),
        EventType: 0,
        AccessGranted: false,
        Door: 0,
        Direction: 0,
        Card: 0,
        Reason: 0,
     },

    GetEventOverwritten: responses.GetEventResponse {
        Controller: 405419896,
        Index: 98765,
        Timestamp: string2datetime("0001-01-01 00:00:00"),
        EventType: 255,
        AccessGranted: false,
        Door: 0,
        Direction: 0,
        Card: 0,
        Reason: 0,
     },

    GetEventIndex: responses.GetEventIndexResponse {
        Controller: 405419896,
        Index: 13579,
     },

    SetEventIndex: responses.SetEventIndexResponse {
        Controller: 405419896,
        Ok: true,
     },

    RecordSpecialEvents: responses.RecordSpecialEventsResponse {
        Controller: 405419896,
        Ok: true,
     },

    GetTimeProfile: responses.GetTimeProfileResponse {
        Controller: 405419896,
        Profile: 37,
        StartDate: string2date("2025-11-26"),
        EndDate: string2date("2025-12-29"),
        Monday: true,
        Tuesday: true,
        Wednesday: false,
        Thursday: true,
        Friday: false,
        Saturday: true,
        Sunday: true,
        Segment1Start: string2HHmm("08:30"),
        Segment1End: string2HHmm("09:45"),
        Segment2Start: string2HHmm("11:35"),
        Segment2End: string2HHmm("13:15"),
        Segment3Start: string2HHmm("14:01"),
        Segment3End: string2HHmm("17:59"),
        LinkedProfile: 19,
     },

    SetTimeProfile: responses.SetTimeProfileResponse {
        Controller: 405419896,
        Ok: true,
     },

    ClearTimeProfiles: responses.ClearTimeProfilesResponse {
        Controller: 405419896,
        Ok: true,
     },

    AddTask: responses.AddTaskResponse {
        Controller: 405419896,
        Ok: true,
     },

    RefreshTaskList: responses.RefreshTaskListResponse {
        Controller: 405419896,
        Ok: true,
     },

    ClearTaskList: responses.ClearTasklistResponse {
        Controller: 405419896,
        Ok: true,
     },

    SetPcControl: responses.SetPCControlResponse {
        Controller: 405419896,
        Ok: true,
     },

    SetInterlock: responses.SetInterlockResponse {
        Controller: 405419896,
        Ok: true,
     },

    ActivateKeypads: responses.ActivateKeypadsResponse {
        Controller: 405419896,
        Ok: true,
     },

}

func IPv4(v string) netip.Addr {
    return netip.MustParseAddr(v)
}

func addrport(v string) netip.AddrPort {
    return netip.MustParseAddrPort(v)
}

func string2datetime(v string) time.Time {
    if d, err := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local); err != nil {
        panic(fmt.Sprintf("invalid datetime (%v)", v))
    } else {
        return d
    }
}

func string2date(v string) time.Time {
    if d, err := time.ParseInLocation("2006-01-02", v, time.Local); err != nil {
        panic(fmt.Sprintf("invalid date (%v)", v))
    } else {
        return d
    }
}

func string2time(v string) time.Time {
    if d, err := time.ParseInLocation("15:04:05", v, time.Local); err != nil {
        panic(fmt.Sprintf("invalid time (%v)", v))
    } else {
        return d
    }
}

func string2HHmm(v string) time.Time {
    if d, err := time.ParseInLocation("15:04", v, time.Local); err != nil {
        panic(fmt.Sprintf("invalid time (%v)", v))
    } else {
        return d
    }
}
