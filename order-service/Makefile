.Phony: all

SHELL := /bin/bash # Use bash syntax
SOURCE_DIRS=$(shell go list ./... | grep -v /vendor | grep -v /out | cut -d "/" -f2 | sort | uniq)

IN_EXE="./cmd"
OUT_EXE="./bin/order-service"

build:
	@GO111MODULE=on go build -o $(OUT_EXE) $(IN_EXE)

start-server: build
	@$(OUT_EXE)

test:
	@GO111MODULE=on go test ./...
