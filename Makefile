SHELL := /bin/bash
BUILD_DATE := `date +%Y%m%d%H%M`
GREEN := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
RESET := $(shell tput -Txterm sgr0)

.PHONY: help

help: ## Show this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "${YELLOW}%-16s${GREEN}%s${RESET}\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## Run the tests of the project.
	go test ./...

deps: ## Download dependencies
	go mod download

build: ## Build binary for local operating system
	go generate ./...
	go build -ldflags "-s -w -X github.com/debeando/lightflow/cli.BuildTime=$(BUILD_DATE)" -o lightflow main.go

build-linux: ## Build binary for linux operating system
	go generate ./...
	GOOS=linux go build -ldflags "-s -w -X github.com/debeando/lightflow/cli.BuildTime=$(BUILD_DATE)" -o lightflow main.go

clean: ## Remove build related file
	go clean

release: ## Create release
	scripts/release.sh
