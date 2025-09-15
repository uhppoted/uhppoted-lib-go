package codec

// Message constants for the UHPPOTE request/response protocol.
const SOM byte = 0x17
const SOM_v6_62 byte = 0x19

//go:generate ../../.codegen/bin/codegen codec
