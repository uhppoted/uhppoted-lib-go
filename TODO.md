# TODO

- [ ] FIXME: default broadcast port
```
./bin/cli --debug set-door -controller 405419896 --door 1 --mode 3 --delay 5
2025-09-25 11:06:43 DEBUG  udp      sent 64 bytes to 255.255.255.255:60001
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
- [x] door mode field type
- [x] event type field type
- [ ] event direction field type
- [ ] event reason field type
- [x] card record
- [x] event record
- [x] get-time-profile-record
- [x] set-time-profile-record
- [ ] task record
- [ ] status record
- [ ] listener interface
- [ ] godoc/examples
- [x] logging
    - [ ] set default handler
    - https://www.dash0.com/guides/logging-in-go-with-slog

- [ ] integration tests: use 'test' domain
         - https://github.com/golang/go/issues/37641

- [ ] code generation
      - [ ] Use 'local' domain
         - https://github.com/golang/go/issues/37641
      - (?) remove Request/Response suffixes

   - [x] decode_AST
   - [x] decoder
       - [ ] decode ListenEvent
       - [ ] decode entities

   - [x] encode
   - [x] encode_test
   - [x] decode_test
   - [x] integration tests

