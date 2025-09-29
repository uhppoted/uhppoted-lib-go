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
- [ ] door mode arg type
- [ ] door mode field type
- [x] card record
- [ ] event record
- [ ] profile record
- [ ] task record
- [ ] godoc/examples
- [x] logging
    - [ ] set default handler

- [ ] integration tests: use 'test' domain
         - https://github.com/golang/go/issues/37641

- [ ] code generation
      - [ ] use 'local' domain
      - [ ] Replace TestArg with Arg
      - [ ] Use 'local' domain
         - https://github.com/golang/go/issues/37641
      - (?) remove Request/Response suffixes

   - [x] response structs
   - [x] decoder
   - [x] decoder_test
   - [x] README
   - [x] API.md
   - [x] decode_AST
   - [x] decode
       - [ ] decode Get/SetAddrPort

   - [x] encode
   - [x] encode_test
   - [x] decode_test
   - [x] integration tests

