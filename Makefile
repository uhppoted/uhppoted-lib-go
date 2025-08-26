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
	cd examples/cli && make clear-time-profiles
	# go test ./integration-tests/... --run TestGetListener

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
	cd examples/cli && make get-controller-udp
	cd examples/cli && make get-controller-tcp

set-IPv4: build
	cd examples/cli && make set-IPv4
	cd examples/cli && make set-IPv4-udp
	cd examples/cli && make set-IPv4-tcp

get-time: build
	cd examples/cli && make get-time
	cd examples/cli && make get-time-udp
	cd examples/cli && make get-time-tcp

set-time: build
	cd examples/cli && make set-time
	cd examples/cli && make set-time-udp
	cd examples/cli && make set-time-tcp

get-listener: build
	cd examples/cli && make get-listener
	cd examples/cli && make get-listener-udp
	cd examples/cli && make get-listener-tcp

set-listener: build
	cd examples/cli && make set-listener
	cd examples/cli && make set-listener-udp
	cd examples/cli && make set-listener-tcp

get-door: build
	cd examples/cli && make get-door
	cd examples/cli && make get-door-udp
	cd examples/cli && make get-door-tcp

set-door: build
	cd examples/cli && make set-door
	cd examples/cli && make set-door-udp
	cd examples/cli && make set-door-tcp

set-door-passcodes: build
	cd examples/cli && make set-door-passcodes
	cd examples/cli && make set-door-passcodes-udp
	cd examples/cli && make set-door-passcodes-tcp

open-door: build
	cd examples/cli && make open-door
	cd examples/cli && make open-door-udp
	cd examples/cli && make open-door-tcp

get-status: build
	cd examples/cli && make get-status
	cd examples/cli && make get-status-udp
	cd examples/cli && make get-status-tcp

get-cards: build
	cd examples/cli && make get-cards
	cd examples/cli && make get-cards-udp
	cd examples/cli && make get-cards-tcp

get-card: build
	cd examples/cli && make get-card
	cd examples/cli && make get-card-udp
	cd examples/cli && make get-card-tcp

get-card-at-index: build
	cd examples/cli && make get-card-at-index
	cd examples/cli && make get-card-at-index-udp
	cd examples/cli && make get-card-at-index-tcp

put-card: build
	cd examples/cli && make put-card
	cd examples/cli && make put-card-udp
	cd examples/cli && make put-card-tcp

delete-card: build
	cd examples/cli && make delete-card
	cd examples/cli && make delete-card-udp
	cd examples/cli && make delete-card-tcp

delete-all-cards: build
	cd examples/cli && make delete-all-cards
	cd examples/cli && make delete-all-cards-udp
	cd examples/cli && make delete-all-cards-tcp

get-event: build
	cd examples/cli && make get-event
	cd examples/cli && make get-event-udp
	cd examples/cli && make get-event-tcp

get-event-index: build
	cd examples/cli && make get-event-index
	cd examples/cli && make get-event-index-udp
	cd examples/cli && make get-event-index-tcp

set-event-index: build
	cd examples/cli && make set-event-index
	cd examples/cli && make set-event-index-udp
	cd examples/cli && make set-event-index-tcp

record-special-events: build
	cd examples/cli && make record-special-events
	cd examples/cli && make record-special-events-udp
	cd examples/cli && make record-special-events-tcp

get-time-profile: build
	cd examples/cli && make get-time-profile
	cd examples/cli && make get-time-profile-udp
	cd examples/cli && make get-time-profile-tcp

set-time-profile: build
	cd examples/cli && make set-time-profile
	cd examples/cli && make set-time-profile-udp
	cd examples/cli && make set-time-profile-tcp

clear-time-profiles: build
	cd examples/cli && make clear-time-profiles
	cd examples/cli && make clear-time-profiles-udp
	cd examples/cli && make clear-time-profiles-tcp
