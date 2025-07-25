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
	go mod tidy

update-release:
	go mod tidy

build: format
	go build -trimpath ./...

debug: build
	cd examples/cli && make get-listener

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

get-all-controllers: build
	cd examples/cli && make get-all-controllers

get-controller: build
	cd examples/cli && make get-controller
	cd examples/cli && make get-controller-udp
	cd examples/cli && make get-controller-tcp

set-IPv4: build
	cd examples/cli && make set-IPv4
	cd examples/cli && make set-IPv4-udp
	cd examples/cli && make set-IPv4-tcp

get-status: build
	cd examples/cli && make get-status
	cd examples/cli && make get-status-udp
	cd examples/cli && make get-status-tcp

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
