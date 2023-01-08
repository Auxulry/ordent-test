debug:
	go run cmd/server/main.go
.PHONY: debug

build:
	go build -o deploy/engine ./cmd
.PHONY: build

dependency:
	go mod download
.PHONY: dependency

lint:
	golangci-lint run ./...
.PHONY: lint

test:
	go test -v ./...
.PHONY: test

tidy:
	go mod tidy
.PHONY: tidy

wire_gen:
	wire gen github.com/MochamadAkbar/ordent-test/injector
.PHONY: wire_gen