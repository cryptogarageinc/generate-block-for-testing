#!/bin/sh -l

export PATH=$PATH:/usr/local/go/bin

go version

cd /github/workspace

go mod download

go build -ldflags "-s -w" -o generateblock-alpine ./cmd/generateblock/
