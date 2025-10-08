.DEFAULT_GOAL := build

all: test      \
	 benchmark \
     coverage

clean:
	go clean
	rm -rf bin
	rm -rf dist

format: 
	go fmt ./...

build: format
	go build -trimpath ./...

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
