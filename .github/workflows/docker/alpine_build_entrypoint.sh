#!/bin/sh -l

cd /github/workspace/generate-block-for-testing

go mod download

go build -ldflags "-s -w" -o generateblock-alpine ./cmd/generateblock/
