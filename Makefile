SHELL := /bin/bash

default: fmt

fmt:
	go fmt ./...

build: fmt
	go build ./cmd/main/main.go

test:
	go test -v ./test

.PHONY: test