package uhppoted

import (
    "fmt"
    "net/netip"
    "time"

    lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

var Expected = struct { 
    GetAllControllers []lib.GetControllerResponse
    GetController lib.GetControllerResponse
    SetIPv4 lib.SetIPv4Response
    GetStatus lib.GetStatusResponse
    GetTime lib.GetTimeResponse
    SetTime lib.SetTimeResponse
    GetListener lib.GetListenerResponse
}{ 
    GetAllControllers: []lib.GetControllerResponse{
        lib.GetControllerResponse{
            Controller: 201020304,
            IpAddress:  netip.MustParseAddr("192.168.1.101"),
            SubnetMask: netip.MustParseAddr("255.255.255.0"),
            Gateway:    netip.MustParseAddr("192.168.1.1"),
            MACAddress: "52:fd:fc:07:21:82",
            Version:    "v6.62",
            Date:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
        },
        lib.GetControllerResponse{
            Controller: 303986753,
            IpAddress:  netip.MustParseAddr("192.168.1.100"),
            SubnetMask: netip.MustParseAddr("255.255.255.0"),
            Gateway:    netip.MustParseAddr("192.168.1.1"),
            MACAddress: "52:fd:fc:07:21:82",
            Version:    "v8.92",
            Date:       time.Date(2019, 8, 15, 0, 0, 0, 0, time.Local),
        },
        lib.GetControllerResponse{
            Controller: 405419896,
            IpAddress:  netip.MustParseAddr("192.168.1.100"),
            SubnetMask: netip.MustParseAddr("255.255.255.0"),
            Gateway:    netip.MustParseAddr("192.168.1.1"),
            MACAddress: "00:12:23:34:45:56",
            Version:    "v8.92",
            Date:       time.Date(2018, 11, 5, 0, 0, 0, 0, time.Local),
        },
    },

    GetController: lib.GetControllerResponse {
        Controller: 405419896,
        IpAddress: IPv4("192.168.1.100"),
        SubnetMask: IPv4("255.255.255.0"),
        Gateway: IPv4("192.168.1.1"),
        MACAddress: "00:12:23:34:45:56",
        Version: "v8.92",
        Date: string2date("2018-11-05"),
     },

    SetIPv4: lib.SetIPv4Response {
        Controller: 405419896,
        Ok: true,
     },

    GetStatus: lib.GetStatusResponse {
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

    GetTime: lib.GetTimeResponse {
        Controller: 405419896,
        DateTime: string2datetime("2024-11-01 12:34:56"),
     },

    SetTime: lib.SetTimeResponse {
        Controller: 405419896,
        DateTime: string2datetime("2024-11-01 12:34:56"),
     },

    GetListener: lib.GetListenerResponse {
        Controller: 405419896,
        Address: addrport("192.168.1.100:60001"),
        Interval: 17,
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
