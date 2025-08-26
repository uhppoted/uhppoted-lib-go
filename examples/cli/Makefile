CMD   = ./bin/cli --debug --bind 0.0.0.0 --broadcast 255.255.255.255:60000 --listen 0.0.0.0:60001
CONTROLLER = 405419896
DOOR = 3
CARD = 10058400
DEST = 127.0.0.1

.DEFAULT_GOAL := build

clean:
	go clean
	rm -rf bin
	rm -rf dist

update:
	go get -u github.com/uhppoted/uhppote-lib-go
	go mod tidy

format: 
	go fmt ./...

build: format
	mkdir -p bin
	go build -trimpath -o bin  ./... 

test: build
	go test ./...

benchmark: build
	go test -bench=.  ./... 

coverage: build
	go test -cover ./...

vet:
	go vet ./...

lint:
	env GOOS=darwin  GOARCH=amd64 staticcheck ./...
	env GOOS=linux   GOARCH=amd64 staticcheck ./...
	env GOOS=windows GOARCH=amd64 staticcheck ./...

vuln:
	govulncheck ./...

build-all: test vet lint
	mkdir -p dist/$(DIST)/linux
	mkdir -p dist/$(DIST)/arm
	mkdir -p dist/$(DIST)/arm7
	mkdir -p dist/$(DIST)/arm6
	mkdir -p dist/$(DIST)/darwin-x64
	mkdir -p dist/$(DIST)/darwin-arm64
	mkdir -p dist/$(DIST)/windows
	env GOOS=linux   GOARCH=amd64         GOWORK=off go build -trimpath -o dist/$(DIST)/linux        ./...
	env GOOS=linux   GOARCH=arm64         GOWORK=off go build -trimpath -o dist/$(DIST)/arm          ./...
	env GOOS=linux   GOARCH=arm   GOARM=7 GOWORK=off go build -trimpath -o dist/$(DIST)/arm7         ./...
	env GOOS=linux   GOARCH=arm   GOARM=6 GOWORK=off go build -trimpath -o dist/$(DIST)/arm6         ./...
	env GOOS=darwin  GOARCH=amd64         GOWORK=off go build -trimpath -o dist/$(DIST)/darwin-x64   ./...
	env GOOS=darwin  GOARCH=arm64         GOWORK=off go build -trimpath -o dist/$(DIST)/darwin-arm64 ./...
	env GOOS=windows GOARCH=amd64         GOWORK=off go build -trimpath -o dist/$(DIST)/windows      ./...

godoc:
	godoc -http=:80	-index_interval=60s

debug: build
	go test -v ./... 

help: build
	$(CMD) help

find-controllers: build
	$(CMD) find-controllers

get-controller: build
	$(CMD) get-controller --controller $(CONTROLLER)

get-controller-udp: build
	$(CMD) get-controller --controller $(CONTROLLER) --dest $(DEST)

get-controller-tcp: build
	$(CMD) get-controller --controller $(CONTROLLER) --dest $(DEST) --tcp

set-IPv4: build
	$(CMD) set-IPv4 --controller $(CONTROLLER) --address 192.168.1.100 --netmask 255.255.255.0 --gateway 192.168.1.1

set-IPv4-udp: build
	$(CMD) set-IPv4 --controller $(CONTROLLER) --dest $(DEST) --address 192.168.1.100 --netmask 255.255.255.0 --gateway 192.168.1.1

set-IPv4-tcp: build
	$(CMD) set-IPv4 --controller $(CONTROLLER) --dest $(DEST) --tcp --address 192.168.1.100 --netmask 255.255.255.0 --gateway 192.168.1.1

get-time: build
	$(CMD) get-time --controller $(CONTROLLER)

get-time-udp: build
	$(CMD) get-time --controller $(CONTROLLER) --dest $(DEST)

get-time-tcp: build
	$(CMD) get-time --controller $(CONTROLLER) --dest $(DEST) --tcp

set-time: build
	$(CMD) set-time --controller $(CONTROLLER) --datetime "2022-01-02 12:34:56"

set-time-udp: build
	$(CMD) set-time --controller $(CONTROLLER) --dest $(DEST)

set-time-tcp: build
	$(CMD) set-time --controller $(CONTROLLER) --dest $(DEST) --tcp

get-listener: build
	$(CMD) get-listener --controller $(CONTROLLER)

get-listener-udp: build
	$(CMD) get-listener --controller $(CONTROLLER) --dest $(DEST)

get-listener-tcp: build
	$(CMD) get-listener --controller $(CONTROLLER) --dest $(DEST) --tcp

set-listener: build
	$(CMD) set-listener --controller $(CONTROLLER) --listener 192.168.1.100:60001 --interval 5

set-listener-udp: build
	$(CMD) set-listener --controller $(CONTROLLER) --dest $(DEST) --listener 192.168.1.100:60001 --interval 5

set-listener-tcp: build
	$(CMD) set-listener --controller $(CONTROLLER) --dest $(DEST) --tcp  --listener 192.168.1.100:60001 --interval 5

get-door: build
	$(CMD) get-door --controller $(CONTROLLER) --door $(DOOR)

get-door-udp: build
	$(CMD) get-door --controller $(CONTROLLER) --dest $(DEST) --door $(DOOR)

get-door-tcp: build
	$(CMD) get-door --controller $(CONTROLLER) --dest $(DEST) --tcp --door $(DOOR)

set-door: build
	$(CMD) set-door --controller $(CONTROLLER) --door $(DOOR) --mode 2 --delay 9

set-door-udp: build
	$(CMD) set-door --controller $(CONTROLLER) --dest $(DEST) --door $(DOOR) --mode 2 --delay 9

set-door-tcp: build
	$(CMD) set-door --controller $(CONTROLLER) --dest $(DEST) --tcp --door $(DOOR) --mode 2 --delay 9

set-door-passcodes: build
	$(CMD) set-door-passcodes --controller $(CONTROLLER) --door $(DOOR) --passcodes 12345,54321,999999

set-door-passcodes-udp: build
	$(CMD) set-door-passcodes --controller $(CONTROLLER) --dest $(DEST) --door $(DOOR) --passcodes 12345,54321,999999

set-door-passcodes-tcp: build
	$(CMD) set-door-passcodes --controller $(CONTROLLER) --dest $(DEST) --tcp --door $(DOOR) --passcodes 12345,54321,999999

open-door: build
	$(CMD) open-door --controller $(CONTROLLER) --door $(DOOR)

open-door-udp: build
	$(CMD) open-door --controller $(CONTROLLER) --dest $(DEST) --door $(DOOR)

open-door-tcp: build
	$(CMD) open-door --controller $(CONTROLLER) --dest $(DEST) --tcp --door $(DOOR)

get-status: build
	$(CMD) get-status --controller $(CONTROLLER)

get-status-udp: build
	$(CMD) get-status --controller $(CONTROLLER) --dest $(DEST)

get-status-tcp: build
	$(CMD) get-status --controller $(CONTROLLER) --dest $(DEST) --tcp

get-cards: build
	$(CMD) get-cards --controller $(CONTROLLER)

get-cards-udp: build
	$(CMD) get-cards --controller $(CONTROLLER) --dest $(DEST)

get-cards-tcp: build
	$(CMD) get-cards --controller $(CONTROLLER) --dest $(DEST) --tcp

get-card: build
	$(CMD) get-card --controller $(CONTROLLER) --card $(CARD)

get-card-udp: build
	$(CMD) get-card --controller $(CONTROLLER) --dest $(DEST) --card $(CARD)

get-card-tcp: build
	$(CMD) get-card --controller $(CONTROLLER) --dest $(DEST) --tcp --card $(CARD)

get-card-at-index: build
	$(CMD) get-card-at-index --controller $(CONTROLLER) --index 3

get-card-at-index-udp: build
	$(CMD) get-card-at-index --controller $(CONTROLLER) --dest $(DEST) --index 3

get-card-at-index-tcp: build
	$(CMD) get-card-at-index --controller $(CONTROLLER) --dest $(DEST) --tcp --index 3

put-card: build
	$(CMD) put-card --controller $(CONTROLLER) --card $(CARD) --start-date "2025-01-01" --end-date "2025-12-31" --doors 1,3:17,4 --PIN 987654

put-card-udp: build
	$(CMD) put-card --controller $(CONTROLLER) --dest $(DEST) --card $(CARD) --start-date "2025-01-01" --end-date "2025-12-31" --doors 1,3:17,4 --PIN 987654

put-card-tcp: build
	$(CMD) put-card --controller $(CONTROLLER) --dest $(DEST) --tcp --card $(CARD) --start-date "2025-01-01" --end-date "2025-12-31" --doors 1,3:17,4 --PIN 987654

delete-card: build
	$(CMD) delete-card --controller $(CONTROLLER) --card $(CARD)

delete-card-udp: build
	$(CMD) delete-card --controller $(CONTROLLER) --dest $(DEST) --card $(CARD)

delete-card-tcp: build
	$(CMD) delete-card --controller $(CONTROLLER) --dest $(DEST) --tcp --card $(CARD)

delete-all-cards: build
	$(CMD) delete-all-cards --controller $(CONTROLLER)

delete-all-cards-udp: build
	$(CMD) delete-all-cards --controller $(CONTROLLER) --dest $(DEST)

delete-all-cards-tcp: build
	$(CMD) delete-all-cards --controller $(CONTROLLER) --dest $(DEST) --tcp

get-event: build
	$(CMD) get-event --controller $(CONTROLLER) --index 17

get-event-udp: build
	$(CMD) get-event --controller $(CONTROLLER) --dest $(DEST) --index 17

get-event-tcp: build
	$(CMD) get-event --controller $(CONTROLLER) --dest $(DEST) --tcp --index 17

get-event-index: build
	$(CMD) get-event-index --controller $(CONTROLLER)

get-event-index-udp: build
	$(CMD) get-event-index --controller $(CONTROLLER) --dest $(DEST)

get-event-index-tcp: build
	$(CMD) get-event-index --controller $(CONTROLLER) --dest $(DEST) --tcp

set-event-index: build
	$(CMD) set-event-index --controller $(CONTROLLER) --index 29

set-event-index-udp: build
	$(CMD) set-event-index --controller $(CONTROLLER) --dest $(DEST) --index 29

set-event-index-tcp: build
	$(CMD) set-event-index --controller $(CONTROLLER) --dest $(DEST) --tcp --index 29

record-special-events: build
	$(CMD) record-special-events --controller $(CONTROLLER) --enabled true

record-special-events-udp: build
	$(CMD) record-special-events --controller $(CONTROLLER) --dest $(DEST) --enabled true

record-special-events-tcp: build
	$(CMD) record-special-events --controller $(CONTROLLER) --dest $(DEST) --tcp --enabled true

get-time-profile: build
	$(CMD) get-time-profile --controller $(CONTROLLER) --profile 29

get-time-profile-udp: build
	$(CMD) get-time-profile --controller $(CONTROLLER) --dest $(DEST) --profile 29

get-time-profile-tcp: build
	$(CMD) get-time-profile --controller $(CONTROLLER) --dest $(DEST) --tcp --profile 29

set-time-profile: build
	$(CMD) set-time-profile --controller $(CONTROLLER) --profile 29

set-time-profile-udp: build
	$(CMD) set-time-profile --controller $(CONTROLLER) --dest $(DEST) --profile 29

set-time-profile-tcp: build
	$(CMD) set-time-profile --controller $(CONTROLLER) --dest $(DEST) --tcp --profile 29

clear-time-profiles: build
	$(CMD) clear-time-profiles --controller $(CONTROLLER)

clear-time-profiles-udp: build
	$(CMD) clear-time-profiles --controller $(CONTROLLER) --dest $(DEST)

clear-time-profiles-tcp: build
	$(CMD) clear-time-profiles --controller $(CONTROLLER) --dest $(DEST) --tcp

