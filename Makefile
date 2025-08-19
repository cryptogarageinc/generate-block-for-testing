current_dir = $(shell pwd)

goimports_version = v0.36.0
yamlfmt_version = v0.17.2
golangci_version = v2.4.0
actionlint_version = v1.7.7
ghalint_version = v1.5.3
pinact_version = v3.4.1

.PHONY: all
all: generate format build

generate:
	go generate ./...

format:
	go run golang.org/x/tools/cmd/goimports@${goimports_version} -w .
	go run github.com/google/yamlfmt/cmd/yamlfmt@${yamlfmt_version}
	go mod tidy

lint:
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@${golangci_version} run

lint-fix:
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@${golangci_version} run --fix

lint-actions:
	go run github.com/rhysd/actionlint/cmd/actionlint@${actionlint_version}
	go run github.com/suzuki-shunsuke/ghalint/cmd/ghalint@${ghalint_version} run

.PHONY: lint-actions-fix
lint-actions-fix: pinact lint-actions

.PHONY: pinact
pinact:
	go run github.com/suzuki-shunsuke/pinact/v3/cmd/pinact@${pinact_version} run

.PHONY: update-actions
update-actions:
	go run github.com/suzuki-shunsuke/pinact/v3/cmd/pinact@${pinact_version} run -u

download:
	go mod download

build: download
	go build ./cmd/generateblock/

build-release: download
	go build -ldflags "-s -w" ./cmd/generateblock/
