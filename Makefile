.PHONY: all

current_dir = $(shell pwd)

all: generate format build

generate:
	go generate ./...

format:
	go run golang.org/x/tools/cmd/goimports@v0.6.0 -w .
	go mod tidy

download:
	go mod download

build: download
	go build ./cmd/generateblock/

build-release: download
	go build -ldflags "-s -w" ./cmd/generateblock/
