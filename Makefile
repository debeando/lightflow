SHELL := /bin/bash
BUILD_DATE := `date +%Y%m%d%H%M`

.PHONY: help

help: ## Show this help.
		@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build binary for local operating system
	go generate ./...
	go build -ldflags "-s -w -X github.com/debeando/lightflow/cli.BuildTime=$(BUILD_DATE)" -o lightflow main.go

build-linux: ## Build binary for linux operating system
	go generate ./...
	GOOS=linux go build -ldflags "-s -w -X github.com/debeando/lightflow/cli.BuildTime=$(BUILD_DATE)" -o lightflow main.go
