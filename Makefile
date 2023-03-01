.ONESHELL:

.SHELL = /bin/bash
.SHELLFLAGS = -ce

.DEFAULT_GOAL = help

.PHONY: build
build: ## Compile the applications source files to produce a single binary.
	go build -ldflags="-w -s" -o bin/astro ./cmd/astro.go

.PHONY: zbuild
zbuild: ## Compress the build output to produce a smaller binary suitable for production.
	upx --best bin/astro

.PHONY: test
test: ## Run the test suite.
	maelstrom test --workload=echo --bin=./bin/astro --node-count=1 --time-limit=10

.PHONY: lint
lint: ## List this projects source files.
	revive -config revive.toml ./...

.PHONY: help
help:  ## Print this help.
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		sed 's/Makefile://' | \
		sort -d | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Silence output by default, use `VERBOSE=1 make <command>` to enable.
ifndef VERBOSE
.SILENT:
endif
