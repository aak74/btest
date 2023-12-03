#!/usr/bin/make

.DEFAULT_GOAL := build

build:
	go build -v ./cmd/btest

run:
	go run ./cmd/btest

test:
	go test -v -race -timeout 30s ./...

up:
	docker-compose up -d
