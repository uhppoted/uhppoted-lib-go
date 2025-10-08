CMD   = ./bin/cli --debug --bind 0.0.0.0 --broadcast 255.255.255.255:60000 --listen 0.0.0.0:60001
CONTROLLER = 405419896
DOOR = 3
CARD = 10058400
DEST ?= 
TCP  ?= 

.DEFAULT_GOAL := build

clean:
	go clean
	rm -rf bin
	rm -rf dist

update:
	go get -u github.com/uhppoted/uhppoted-lib-go/src/uhppoted
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
	$(CMD) get-controller --controller $(CONTROLLER) $(DEST) $(TCP)

set-IPv4: build
	$(CMD) set-IPv4 --controller $(CONTROLLER) $(DEST) $(TCP) --address 192.168.1.100 --netmask 255.255.255.0 --gateway 192.168.1.1

get-time: build
	$(CMD) get-time --controller $(CONTROLLER) $(DEST) $(TCP)

set-time: build
	$(CMD) set-time --controller $(CONTROLLER) $(DEST) $(TCP) --datetime "2025-10-02 12:34:56"

get-listener: build
	$(CMD) get-listener --controller $(CONTROLLER) $(DEST) $(TCP)

set-listener: build
	$(CMD) set-listener --controller $(CONTROLLER) $(DEST) $(TCP) --listener 192.168.1.100:60001 --interval 5

get-door: build
	$(CMD) get-door --controller $(CONTROLLER) $(DEST) $(TCP) --door $(DOOR)

set-door: build
	$(CMD) set-door --controller $(CONTROLLER) $(DEST) $(TCP) --door $(DOOR) --mode 2 --delay 9

set-door-passcodes: build
	$(CMD) set-door-passcodes --controller $(CONTROLLER) $(DEST) $(TCP) --door $(DOOR) --passcodes 12345,54321,999999

open-door: build
	$(CMD) open-door --controller $(CONTROLLER) $(DEST) $(TCP) --door $(DOOR)

get-status: build
	$(CMD) get-status --controller $(CONTROLLER) $(DEST) $(TCP)

get-status-record: build
	$(CMD) get-status-record --controller $(CONTROLLER) $(DEST) $(TCP)

get-cards: build
	$(CMD) get-cards --controller $(CONTROLLER) $(DEST) $(TCP)

get-card: build
	$(CMD) get-card-record --controller $(CONTROLLER) $(DEST) $(TCP) --card $(CARD)

get-card-at-index: build
	$(CMD) get-card-record-at-index --controller $(CONTROLLER) $(DEST) $(TCP) --index 3

put-card: build
	$(CMD) put-card-record --controller $(CONTROLLER) $(DEST) $(TCP) --card $(CARD) --start-date "2025-01-01" --end-date "2025-12-31" --doors 1,3:17,4 --PIN 987654

delete-card: build
	$(CMD) delete-card --controller $(CONTROLLER) $(DEST) $(TCP) --card $(CARD)

delete-all-cards: build
	$(CMD) delete-all-cards --controller $(CONTROLLER) $(DEST) $(TCP)

get-event: build
	$(CMD) get-event --controller $(CONTROLLER) $(DEST) $(TCP) --index 17

get-event-record: build
	$(CMD) get-event-record --controller $(CONTROLLER) $(DEST) $(TCP) --index 17

get-event-index: build
	$(CMD) get-event-index --controller $(CONTROLLER) $(DEST) $(TCP)

set-event-index: build
	$(CMD) set-event-index --controller $(CONTROLLER) $(DEST) $(TCP) --index 29

record-special-events: build
	$(CMD) record-special-events --controller $(CONTROLLER) $(DEST) $(TCP) --enabled true

get-time-profile: build
	$(CMD) get-time-profile --controller $(CONTROLLER) $(DEST) $(TCP) --profile 29

get-time-profile-record: build
	$(CMD) get-time-profile-record --controller $(CONTROLLER) $(DEST) $(TCP) --profile 29

set-time-profile: build
	$(CMD) set-time-profile --controller $(CONTROLLER) $(DEST) $(TCP) --profile 29

set-time-profile-record: build
	$(CMD) set-time-profile-record --controller $(CONTROLLER) $(DEST) $(TCP) --profile 29

clear-time-profiles: build
	$(CMD) clear-time-profiles --controller $(CONTROLLER) $(DEST) $(TCP) 

add-task: build
	$(CMD) add-task --controller $(CONTROLLER) $(DEST) $(TCP) 

add-task-record: build
	$(CMD) add-task-record --controller $(CONTROLLER) $(DEST) $(TCP) 

refresh-tasklist: build
	$(CMD) refresh-tasklist --controller $(CONTROLLER) $(DEST) $(TCP) 

clear-tasklist: build
	$(CMD) clear-tasklist --controller $(CONTROLLER) $(DEST) $(TCP) 

set-pc-control: build
	$(CMD) set-pc-control --controller $(CONTROLLER) $(DEST) $(TCP) --enabled true

set-interlock: build
	$(CMD) set-interlock --controller $(CONTROLLER) $(DEST) $(TCP) --interlock "1&2,3&4"

activate-keypads: build
	$(CMD) activate-keypads --controller $(CONTROLLER) $(DEST) $(TCP) --keypads "12,4"

get-antipassback: build
	$(CMD) get-antipassback --controller $(CONTROLLER) $(DEST) $(TCP)

set-antipassback: build
	$(CMD) set-antipassback --controller $(CONTROLLER) $(DEST) $(TCP) --anti-passback "(1:2);(3:4)"

restore-default-parameters: build
	$(CMD) restore-default-parameters --controller $(CONTROLLER) $(DEST) $(TCP)

listen: build
	$(CMD) listen
