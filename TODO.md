# TODO

- [x] README
- [x] github workflow
- [x] get-all-controllers
- [x] get-controller
- [x] set-IPv4
- [x] get-status
- [x] get-time
- [x] set-time
- [x] get-listener
- [x] set-listener
- [x] get-door
- [x] set-door
- [x] set-door-passcodes
- [x] open-door
- [x] get-cards
- [ ] get-card
   - [x] API
   - [x] CLI
   - [x] README
   - [x] integration test
       - [ ] card not found

- [ ] custom date/time types
- [ ] check for controller = 0
- [ ] check response.controller == request.controller
- [ ] check response.card == request.card
- [ ] v6.62 version handling
- [ ] godoc examples
- [x] logging
    - [ ] set default handler

- [ ] code generation
   - [ ] use/extend uhppoted-codegen models
      - [ ] response

   - [ ] AST
      - [x] decoder
      - [ ] decoder_test
      - [ ] API
         - [ ] godoc
         - [ ] response structs
   - [x] encode
   - [x] encode_test
   - [x] decode
   - [x] decode_test
   - [x] integration tests
   - [ ] README
