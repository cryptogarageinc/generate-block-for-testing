.PHONY: all

current_dir = $(shell pwd)

all: generate format build

generate:
	go generate ./...

format:
	go run golang.org/x/tools/cmd/goimports@v0.7.0 -w .
	go mod tidy

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.1 run

lint-fix:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.1 run --fix

download:
	go mod download

build: download
	go build ./cmd/generateblock/

build-release: download
	go build -ldflags "-s -w" ./cmd/generateblock/
