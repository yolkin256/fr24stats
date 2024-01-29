PWD=$(shell pwd)
USER=$(shell id -u)

.PHONY: all

all: install_deps fmt build

build:
	cd cmd/cli && go build -o ../../bin/app

migration: # создает файл БД (если ещё не создан), запускает миграции
	bin/app migration --file=data/database.sqlite

lint:
	golangci-lint run -v ./...

fmt:
	go fmt ./...

install_deps:
	go mod tidy
	go mod vendor
	go get -v ./...
