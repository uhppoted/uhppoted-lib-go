CMD   = ./bin/cli --debug --bind 0.0.0.0 --broadcast 255.255.255.255:60000 --listen 0.0.0.0:60001

.DEFAULT_GOAL := build

clean:
	go clean
	rm -rf bin
	rm -rf dist

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

debug: build
