.DEFAULT_GOAL := build

all: test      \
	 benchmark \
     coverage

clean:
	go clean
	rm -rf bin
	rm -rf dist

regenerate:
	cd .codegen && make build
	go generate ./...

format: 
	go fmt ./...

update:
	cd .codegen && make update
	go mod tidy

update-release:
	cd .codegen && make update-release
	go mod tidy

build: format
	go build -trimpath ./...

debug: build
	cd examples/cli && make set-interlock

test: build
	go test ./uhppoted/...

integration-tests: build
	go test ./integration-tests/...

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

build-all: regenerate test integration-tests vet lint
	env GOOS=linux   GOARCH=amd64       GOWORK=off go build -trimpath ./...
	env GOOS=linux   GOARCH=arm GOARM=7 GOWORK=off go build -trimpath ./...
	env GOOS=linux   GOARCH=arm GOARM=6 GOWORK=off go build -trimpath ./...
	env GOOS=darwin  GOARCH=amd64       GOWORK=off go build -trimpath ./...
	env GOOS=windows GOARCH=amd64       GOWORK=off go build -trimpath ./...

release: clean build-all

publish: release
	echo "Releasing version $(VERSION)"
	gh release create "$(VERSION)" --draft --prerelease --title "$(VERSION)-beta" --notes-file release-notes.md

godoc:
	godoc -http=:80	-index_interval=60s

help:
	cd examples/cli && make help

find-controllers: build
	cd examples/cli && make find-controllers

get-controller: build
	cd examples/cli && make get-controller
	cd examples/cli && make get-controller DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make get-controller DEST="--dest 127.0.0.1" TCP="- DEST="--dest 127.0.0.1" TCP="--tcp""

set-IPv4: build
	cd examples/cli && make set-IPv4
	cd examples/cli && make set-IPv4 DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make set-IPv4 DEST="--dest 127.0.0.1" TCP="- DEST="--dest 127.0.0.1" TCP="--tcp""

get-time: build
	cd examples/cli && make get-time
	cd examples/cli && make get-time DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make get-time DEST="--dest 127.0.0.1" TCP="--tcp"

set-time: build
	cd examples/cli && make set-time
	cd examples/cli && make set-time DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make set-time DEST="--dest 127.0.0.1" TCP="--tcp"

get-listener: build
	cd examples/cli && make get-listener
	cd examples/cli && make get-listener DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make get-listener DEST="--dest 127.0.0.1" TCP="--tcp"

set-listener: build
	cd examples/cli && make set-listener
	cd examples/cli && make set-listener DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make set-listener DEST="--dest 127.0.0.1" TCP="--tcp"

get-door: build
	cd examples/cli && make get-door
	cd examples/cli && make get-door DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make get-door DEST="--dest 127.0.0.1" TCP="--tcp"

set-door: build
	cd examples/cli && make set-door
	cd examples/cli && make set-door DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make set-door DEST="--dest 127.0.0.1" TCP="--tcp"

set-door-passcodes: build
	cd examples/cli && make set-door-passcodes
	cd examples/cli && make set-door-passcodes DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make set-door-passcodes DEST="--dest 127.0.0.1" TCP="--tcp"

open-door: build
	cd examples/cli && make open-door
	cd examples/cli && make open-door DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make open-door DEST="--dest 127.0.0.1" TCP="--tcp"

get-status: build
	cd examples/cli && make get-status
	cd examples/cli && make get-status DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make get-status DEST="--dest 127.0.0.1" TCP="--tcp"

get-cards: build
	cd examples/cli && make get-cards
	cd examples/cli && make get-cards DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make get-cards DEST="--dest 127.0.0.1" TCP="--tcp"

get-card: build
	cd examples/cli && make get-card
	cd examples/cli && make get-card DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make get-card DEST="--dest 127.0.0.1" TCP="--tcp"

get-card-at-index: build
	cd examples/cli && make get-card-at-index
	cd examples/cli && make get-card-at-index DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make get-card-at-index DEST="--dest 127.0.0.1" TCP="--tcp"

put-card: build
	cd examples/cli && make put-card
	cd examples/cli && make put-card DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make put-card DEST="--dest 127.0.0.1" TCP="--tcp"

delete-card: build
	cd examples/cli && make delete-card
	cd examples/cli && make delete-card DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make delete-card DEST="--dest 127.0.0.1" TCP="--tcp"

delete-all-cards: build
	cd examples/cli && make delete-all-cards
	cd examples/cli && make delete-all-cards DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make delete-all-cards DEST="--dest 127.0.0.1" TCP="--tcp"

get-event: build
	cd examples/cli && make get-event
	cd examples/cli && make get-event DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make get-event DEST="--dest 127.0.0.1" TCP="--tcp"

get-event-index: build
	cd examples/cli && make get-event-index
	cd examples/cli && make get-event-index DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make get-event-index DEST="--dest 127.0.0.1" TCP="--tcp"

set-event-index: build
	cd examples/cli && make set-event-index
	cd examples/cli && make set-event-index DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make set-event-index DEST="--dest 127.0.0.1" TCP="--tcp"

record-special-events: build
	cd examples/cli && make record-special-events
	cd examples/cli && make record-special-events DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make record-special-events DEST="--dest 127.0.0.1" TCP="--tcp"

get-time-profile: build
	cd examples/cli && make get-time-profile
	cd examples/cli && make get-time-profile DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make get-time-profile DEST="--dest 127.0.0.1" TCP="--tcp"

set-time-profile: build
	cd examples/cli && make set-time-profile
	cd examples/cli && make set-time-profile DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make set-time-profile DEST="--dest 127.0.0.1" TCP="--tcp"

clear-time-profiles: build
	cd examples/cli && make clear-time-profiles
	cd examples/cli && make clear-time-profiles DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make clear-time-profiles DEST="--dest 127.0.0.1" TCP="--tcp"

add-task: build
	cd examples/cli && make add-task
	cd examples/cli && make add-task DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make add-task DEST="--dest 127.0.0.1" TCP="--tcp"

refresh-tasklist: build
	cd examples/cli && make refresh-tasklist
	cd examples/cli && make refresh-tasklist DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make refresh-tasklist DEST="--dest 127.0.0.1" TCP="--tcp"

clear-tasklist: build
	cd examples/cli && make clear-tasklist
	cd examples/cli && make clear-tasklist DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make clear-tasklist DEST="--dest 127.0.0.1" TCP="--tcp"

set-pc-control: build
	cd examples/cli && make set-pc-control
	cd examples/cli && make set-pc-control DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make set-pc-control DEST="--dest 127.0.0.1" TCP="- DEST="--dest 127.0.0.1" TCP="--tcp""

set-interlock: build
	cd examples/cli && make set-interlock
	cd examples/cli && make set-interlock DEST="--dest 127.0.0.1" TCP=""
	cd examples/cli && make set-interlock DEST="--dest 127.0.0.1" TCP="- DEST="--dest 127.0.0.1" TCP="--tcp""
