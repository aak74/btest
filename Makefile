#!/usr/bin/make

.DEFAULT_GOAL := build

APP_TAG := 0.1.2


help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo "\n  Allowed for overriding next properties:\n\n\
		Usage example:\n\
		make build"

build: ## Build
	go build -v ./cmd/btest

run:
	go run ./cmd/btest

test:
	go test -v ./cmd/btest

up:
	docker-compose up -d


build-ghcr:
	docker build --rm -f Dockerfile -t ghcr.io/aak74/btest:latest -t ghcr.io/aak74/btest:$(APP_TAG) .

push-ghcr:
	docker push ghcr.io/aak74/btest:latest
	docker push ghcr.io/aak74/btest:$(APP_TAG)