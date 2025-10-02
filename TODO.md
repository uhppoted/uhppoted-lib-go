# TODO

- [ ] FIXME: default broadcast port
```
./bin/cli --debug set-door -controller 405419896 --door 1 --mode 3 --delay 5
2025-09-25 11:06:43 DEBUG  udp      sent 64 bytes to 255.255.255.255:60001
```

- [x] FIXME: panic
```
goroutine 1 [running]:
github.com/uhppoted/uhppoted-lib-go/uhppoted/codec.Decode[...]({0x0?, 0xc0000c80d8?, 0x0?})
  github.com/uhppoted/uhppoted-lib-go/uhppoted/codec/decoder.go:28 +0x128
  ...
  github.com/uhppoted/uhppoted-lib-go/examples/cli/cmd/cli/open-door.go:31 +0x1d4
main.main()
  github.com/uhppoted/uhppoted-lib-go/examples/cli/cmd/cli/main.go:58 +0x4e8
```

- [x] custom date
- [x] custom time
- [x] custom HHmm
- [x] custom datetime
- [x] custom optional datetime
- [x] datetime args
- [x] date args
- [x] HHmm args
- [x] task arg type
- [x] interlock arg type
- [x] door mode arg type
- [ ] door mode field type
- [ ] event type field type
- [ ] event direction field type
- [ ] event reason field type
- [x] card record
- [x] event record
- [ ] profile record
- [ ] task record
- [ ] listener interface
- [ ] godoc/examples
- [x] logging
    - [ ] set default handler
    - https://www.dash0.com/guides/logging-in-go-with-slog

- [ ] integration tests: use 'test' domain
         - https://github.com/golang/go/issues/37641

- [ ] code generation
      - [x] Replace TestArg with Arg
      - [ ] Replace TestReply with Reply
      - [ ] Use 'local' domain
         - https://github.com/golang/go/issues/37641
      - (?) remove Request/Response suffixes

   - [x] decode_AST
   - [x] decoder
       - [x] test invalid packets
       - [x] decode Get/SetAddrPort
       - [x] split Decoder() and decode() into separate files
       - [ ] decode ListenEvent

   - [x] encode
   - [x] encode_test
   - [x] decode_test
   - [x] integration tests

