SHELL := /usr/bin/env bash

.DEFAULT_GOAL := build

ROOT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

GO ?= go
NPM ?= npm
PYTHON ?= python3

# Set to 0 to skip dependency install steps.
NPM_INSTALL ?= 1
PYTHON_INSTALL ?= 1

.PHONY: help build build-all build-go build-go-api build-go-monitor build-stubs \
	build-grpc-stubs build-node-stubs build-python-stubs build-lua-descriptor \
	build-node build-node-binaries build-node-client build-python clean

help:
	@echo "Targets:"
	@echo "  make build            Build all project outputs"
	@echo "  make build-go         Build sikuligo + sikuligo-monitor binaries"
	@echo "  make build-stubs      Generate Go/Node/Python/Lua protocol artifacts"
	@echo "  make build-node       Build Node SDK + platform binaries"
	@echo "  make build-python     Build Python distribution artifacts"
	@echo "  make clean            Remove build outputs"
	@echo ""
	@echo "Options:"
	@echo "  NPM_INSTALL=0         Skip npm install in Node build"
	@echo "  PYTHON_INSTALL=0      Skip pip install in Python build"

build: build-all

build-all: build-stubs build-go build-node build-python

build-go: build-go-api build-go-monitor

build-go-api:
	cd "$(ROOT_DIR)" && \
	$(GO) build -tags "gosseract opencv" -trimpath -ldflags="-s -w" -o sikuligo ./cmd/sikuligrpc

build-go-monitor:
	cd "$(ROOT_DIR)" && \
	$(GO) build -trimpath -ldflags="-s -w" -o sikuligo-monitor ./cmd/sikuligo-monitor

build-stubs: build-grpc-stubs build-node-stubs build-python-stubs build-lua-descriptor

build-grpc-stubs:
	cd "$(ROOT_DIR)" && ./scripts/generate-grpc-stubs.sh

build-node-stubs:
	cd "$(ROOT_DIR)" && ./scripts/clients/generate-node-stubs.sh

build-python-stubs:
	cd "$(ROOT_DIR)" && ./scripts/clients/generate-python-stubs.sh

build-lua-descriptor:
	cd "$(ROOT_DIR)" && ./scripts/clients/generate-lua-descriptor.sh

build-node: build-node-binaries build-node-client

build-node-binaries:
	cd "$(ROOT_DIR)" && ./scripts/clients/build-node-binaries.sh

build-node-client:
	cd "$(ROOT_DIR)/clients/node" && \
	if [[ "$(NPM_INSTALL)" == "1" ]]; then $(NPM) install; fi && \
	$(NPM) run build && \
	$(NPM) pack --dry-run

build-python:
	cd "$(ROOT_DIR)" && \
	SKIP_INSTALL=$$( [[ "$(PYTHON_INSTALL)" == "1" ]] && echo 0 || echo 1 ) \
	./scripts/clients/release-python-client.sh

clean:
	cd "$(ROOT_DIR)" && \
	rm -rf \
	clients/node/dist \
	clients/node/generated \
	clients/node/packages/bin-*/bin/sikuligo \
	clients/node/packages/bin-*/bin/sikuligo.exe \
	clients/node/packages/checksums.txt \
	clients/python/dist \
	clients/python/build \
	clients/python/*.egg-info \
	internal/grpcv1/pb \
	sikuligo \
	sikuligo-monitor
