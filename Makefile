.PHONY: regenerate
.PHONY: integration-tests

.DEFAULT_GOAL := build

regenerate:
	cd .codegen          && make build
	cd src               && make regenerate
	cd integration-tests && make regenerate

update:
	cd .codegen          && make update
	cd integration-tests && make update
	cd examples/cli      && make update

update-release:
	cd .codegen && make update-release
	go mod tidy

build: 
	cd src && make build

test: 
	cd src && make test

integration-tests:
	cd integration-tests && make test

vet:
	cd src && make vet

lint:
	cd src && make lint

vuln:
	cd src && make vuln

build-all:
	cd .codegen     && make build	
	cd src          && make build-all
	cd examples/cli && make build

# release: clean build-all

# publish: release
# 	echo "Releasing version $(VERSION)"
# 	gh release create "$(VERSION)" --draft --prerelease --title "$(VERSION)-beta" --notes-file release-notes.md

# godoc:
# 	godoc -http=:80	-index_interval=60s

debug: regenerate test integration-tests
	cd src && make debug

help:
	cd examples/cli && make help

find-controllers: build
	cd examples/cli && make find-controllers

get-controller: build
	cd examples/cli && make get-controller
	cd examples/cli && make get-controller DEST="--dest 127.0.0.1"
	cd examples/cli && make get-controller DEST="--dest 127.0.0.1" TCP="--tcp"

set-IPv4: build
	cd examples/cli && make set-IPv4
	cd examples/cli && make set-IPv4 DEST="--dest 127.0.0.1"
	cd examples/cli && make set-IPv4 DEST="--dest 127.0.0.1" TCP="--tcp"

get-time: build
	cd examples/cli && make get-time
	cd examples/cli && make get-time DEST="--dest 127.0.0.1"
	cd examples/cli && make get-time DEST="--dest 127.0.0.1" TCP="--tcp"

set-time: build
	cd examples/cli && make set-time
	cd examples/cli && make set-time DEST="--dest 127.0.0.1"
	cd examples/cli && make set-time DEST="--dest 127.0.0.1" TCP="--tcp"

get-listener: build
	cd examples/cli && make get-listener
	cd examples/cli && make get-listener DEST="--dest 127.0.0.1"
	cd examples/cli && make get-listener DEST="--dest 127.0.0.1" TCP="--tcp"

set-listener: build
	cd examples/cli && make set-listener
	cd examples/cli && make set-listener DEST="--dest 127.0.0.1"
	cd examples/cli && make set-listener DEST="--dest 127.0.0.1" TCP="--tcp"

get-door: build
	cd examples/cli && make get-door
	cd examples/cli && make get-door DEST="--dest 127.0.0.1"
	cd examples/cli && make get-door DEST="--dest 127.0.0.1" TCP="--tcp"

set-door: build
	cd examples/cli && make set-door
	cd examples/cli && make set-door DEST="--dest 127.0.0.1"
	cd examples/cli && make set-door DEST="--dest 127.0.0.1" TCP="--tcp"

set-door-passcodes: build
	cd examples/cli && make set-door-passcodes
	cd examples/cli && make set-door-passcodes DEST="--dest 127.0.0.1"
	cd examples/cli && make set-door-passcodes DEST="--dest 127.0.0.1" TCP="--tcp"

open-door: build
	cd examples/cli && make open-door
	cd examples/cli && make open-door DEST="--dest 127.0.0.1"
	cd examples/cli && make open-door DEST="--dest 127.0.0.1" TCP="--tcp"

get-status: build
	cd examples/cli && make get-status
	cd examples/cli && make get-status DEST="--dest 127.0.0.1"
	cd examples/cli && make get-status DEST="--dest 127.0.0.1" TCP="--tcp"

get-cards: build
	cd examples/cli && make get-cards
	cd examples/cli && make get-cards DEST="--dest 127.0.0.1"
	cd examples/cli && make get-cards DEST="--dest 127.0.0.1" TCP="--tcp"

get-card: build
	cd examples/cli && make get-card
	cd examples/cli && make get-card DEST="--dest 127.0.0.1"
	cd examples/cli && make get-card DEST="--dest 127.0.0.1" TCP="--tcp"

get-card-at-index: build
	cd examples/cli && make get-card-at-index
	cd examples/cli && make get-card-at-index DEST="--dest 127.0.0.1"
	cd examples/cli && make get-card-at-index DEST="--dest 127.0.0.1" TCP="--tcp"

put-card: build
	cd examples/cli && make put-card
	cd examples/cli && make put-card DEST="--dest 127.0.0.1"
	cd examples/cli && make put-card DEST="--dest 127.0.0.1" TCP="--tcp"

delete-card: build
	cd examples/cli && make delete-card
	cd examples/cli && make delete-card DEST="--dest 127.0.0.1"
	cd examples/cli && make delete-card DEST="--dest 127.0.0.1" TCP="--tcp"

delete-all-cards: build
	cd examples/cli && make delete-all-cards
	cd examples/cli && make delete-all-cards DEST="--dest 127.0.0.1"
	cd examples/cli && make delete-all-cards DEST="--dest 127.0.0.1" TCP="--tcp"

get-event: build
	cd examples/cli && make get-event
	cd examples/cli && make get-event DEST="--dest 127.0.0.1"
	cd examples/cli && make get-event DEST="--dest 127.0.0.1" TCP="--tcp"

get-event-record: build
	cd examples/cli && make get-event-record
	cd examples/cli && make get-event-record DEST="--dest 127.0.0.1"
	cd examples/cli && make get-event-record DEST="--dest 127.0.0.1" TCP="--tcp"

get-event-index: build
	cd examples/cli && make get-event-index
	cd examples/cli && make get-event-index DEST="--dest 127.0.0.1"
	cd examples/cli && make get-event-index DEST="--dest 127.0.0.1" TCP="--tcp"

set-event-index: build
	cd examples/cli && make set-event-index
	cd examples/cli && make set-event-index DEST="--dest 127.0.0.1"
	cd examples/cli && make set-event-index DEST="--dest 127.0.0.1" TCP="--tcp"

record-special-events: build
	cd examples/cli && make record-special-events
	cd examples/cli && make record-special-events DEST="--dest 127.0.0.1"
	cd examples/cli && make record-special-events DEST="--dest 127.0.0.1" TCP="--tcp"

get-time-profile: build
	cd examples/cli && make get-time-profile
	cd examples/cli && make get-time-profile DEST="--dest 127.0.0.1"
	cd examples/cli && make get-time-profile DEST="--dest 127.0.0.1" TCP="--tcp"

set-time-profile: build
	cd examples/cli && make set-time-profile
	cd examples/cli && make set-time-profile DEST="--dest 127.0.0.1"
	cd examples/cli && make set-time-profile DEST="--dest 127.0.0.1" TCP="--tcp"

clear-time-profiles: build
	cd examples/cli && make clear-time-profiles
	cd examples/cli && make clear-time-profiles DEST="--dest 127.0.0.1"
	cd examples/cli && make clear-time-profiles DEST="--dest 127.0.0.1" TCP="--tcp"

add-task: build
	cd examples/cli && make add-task
	cd examples/cli && make add-task DEST="--dest 127.0.0.1"
	cd examples/cli && make add-task DEST="--dest 127.0.0.1" TCP="--tcp"

add-task-record: build
	cd examples/cli && make add-task-record
	cd examples/cli && make add-task-record DEST="--dest 127.0.0.1"
	cd examples/cli && make add-task-record DEST="--dest 127.0.0.1" TCP="--tcp"

refresh-tasklist: build
	cd examples/cli && make refresh-tasklist
	cd examples/cli && make refresh-tasklist DEST="--dest 127.0.0.1"
	cd examples/cli && make refresh-tasklist DEST="--dest 127.0.0.1" TCP="--tcp"

clear-tasklist: build
	cd examples/cli && make clear-tasklist
	cd examples/cli && make clear-tasklist DEST="--dest 127.0.0.1"
	cd examples/cli && make clear-tasklist DEST="--dest 127.0.0.1" TCP="--tcp"

set-pc-control: build
	cd examples/cli && make set-pc-control
	cd examples/cli && make set-pc-control DEST="--dest 127.0.0.1"
	cd examples/cli && make set-pc-control DEST="--dest 127.0.0.1" DEST="--dest 127.0.0.1" TCP="--tcp""

set-interlock: build
	cd examples/cli && make set-interlock
	cd examples/cli && make set-interlock DEST="--dest 127.0.0.1"
	cd examples/cli && make set-interlock DEST="--dest 127.0.0.1" DEST="--dest 127.0.0.1" TCP="--tcp"

activate-keypads: build
	cd examples/cli && make activate-keypads
	cd examples/cli && make activate-keypads DEST="--dest 127.0.0.1"
	cd examples/cli && make activate-keypads DEST="--dest 127.0.0.1" DEST="--dest 127.0.0.1" TCP="--tcp"

get-antipassback: build
	cd examples/cli && make get-antipassback
	cd examples/cli && make get-antipassback DEST="--dest 127.0.0.1"
	cd examples/cli && make get-antipassback DEST="--dest 127.0.0.1" DEST="--dest 127.0.0.1" TCP="--tcp"

set-antipassback: build
	cd examples/cli && make set-antipassback
	cd examples/cli && make set-antipassback DEST="--dest 127.0.0.1"
	cd examples/cli && make set-antipassback DEST="--dest 127.0.0.1" DEST="--dest 127.0.0.1" TCP="--tcp"

restore-default-parameters: build
	cd examples/cli && make restore-default-parameters
	cd examples/cli && make restore-default-parameters DEST="--dest 127.0.0.1"
	cd examples/cli && make restore-default-parameters DEST="--dest 127.0.0.1" DEST="--dest 127.0.0.1" TCP="--tcp"

listen: build
	cd examples/cli && make listen
