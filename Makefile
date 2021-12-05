.PHONY: build test test-unit test-cli

default: all

all: build

build:
	go build -o _output/driedel main.go 

test: test-unit test-cli

test-unit:
	go test ./...

test-cli: build
	test/bats/bin/bats test/root.bats
