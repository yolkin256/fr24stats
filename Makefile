.PHONY: all

all: install_deps fmt build

build:
	cd cmd/cli && go build -o ../../bin/app

lint:
	golangci-lint run -v ./...

fmt:
	go fmt ./...

install_deps:
	go mod tidy
	go mod vendor
	go get -v ./...
