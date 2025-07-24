CMD   = ./bin/cli --debug --bind 0.0.0.0 --broadcast 255.255.255.255:60000 --listen 0.0.0.0:60001
CONTROLLER = 405419896
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

get-all-controllers: build
	$(CMD) get-all-controllers

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

get-status: build
	$(CMD) get-status --controller $(CONTROLLER)

get-status-udp: build
	$(CMD) get-status --controller $(CONTROLLER) --dest $(DEST)

get-status-tcp: build
	$(CMD) get-status --controller $(CONTROLLER) --dest $(DEST) --tcp

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
