.PHONY: all

current_dir = $(shell pwd)

all: gettools generate format build

gettools:
	go install golang.org/x/tools/cmd/goimports@v0.1.5

generate:
	go generate ./...

format:
	goimports -w .
	go mod tidy

download:
	go mod download

build: download
	go build ./cmd/generateblock/

build-release: download
	go build -ldflags "-s -w" ./cmd/generateblock/
