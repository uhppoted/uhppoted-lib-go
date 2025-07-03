package codec

// # Message constants for the UHPPOTE request/response protocol.

const SOM byte = 0x17
const SOM_v6_62 byte = 0x19

const GET_STATUS byte = 0x20
const SET_TIME byte = 0x30
const GET_TIME byte = 0x32
const OPEN_DOOR byte = 0x40
const PUT_CARD byte = 0x50
const DELETE_CARD byte = 0x52
const DELETE_ALL_CARDS byte = 0x54
const GET_CARDS byte = 0x58
const GET_CARD byte = 0x5A
const GET_CARD_AT_INDEX byte = 0x5C
const SET_DOOR byte = 0x80
const GET_DOOR byte = 0x82
const SET_ANTIPASSBACK byte = 0x84
const GET_ANTIPASSBACK byte = 0x86
const SET_TIME_PROFILE byte = 0x88
const CLEAR_TIME_PROFILES byte = 0x8A
const SET_DOOR_PASSCODES byte = 0x8C
const RECORD_SPECIAL_EVENTS byte = 0x8E
const SET_LISTENER byte = 0x90
const GET_LISTENER byte = 0x92
const GetController byte = 0x94
const SET_IPv4 byte = 0x96
const GET_TIME_PROFILE byte = 0x98
const SET_PC_CONTROL byte = 0xA0
const SET_INTERLOCK byte = 0xA2
const ACTIVATE_KEYPADS byte = 0xA4
const CLEAR_TASKLIST byte = 0xA6
const ADD_TASK byte = 0xA8
const SET_FIRST_CARD byte = 0xAA
const REFRESH_TASKLIST byte = 0xAC
const GET_EVENT byte = 0xB0
const SET_EVENT_INDEX byte = 0xB2
const GET_EVENT_INDEX byte = 0xB4
const RESTORE_DEFAULT_PARAMETERS byte = 0xC8
const LISTEN_EVENT byte = 0x20
