.PHONY: build test

default: all

all: build

build:
	go build -o _output/driedel main.go 

test:
	go test ./...