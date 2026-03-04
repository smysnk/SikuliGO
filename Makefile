SHELL := /usr/bin/env bash

.DEFAULT_GOAL := build

ROOT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
API_DIR := $(ROOT_DIR)/packages/api

GO ?= go
NPM ?= npm
YARN ?= yarn
PYTHON ?= python3
BENCH_HELPER_LISTEN ?= :8091
BENCH_HELPER_SPEC ?= packages/api/internal/grpcv1/testdata/find-bench-assets/regions.json
BENCH_HELPER_MANIFEST ?= docs/bench/find-on-screen-scenarios.example.json
BENCH_HELPER_OPEN_BROWSER ?= 1

# Set to 0 to skip dependency install steps.
NPM_INSTALL ?= 1
PYTHON_INSTALL ?= 1
GO_OPENCV_TAGS := gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d

# macOS Homebrew OCR/OpenCV build flags (requested fixed versions)
OS_NAME := $(shell uname -s)
MACOS_CGO_CXXFLAGS := -I/opt/homebrew/Cellar/leptonica/1.87.0/include -I/opt/homebrew/Cellar/tesseract/5.5.2/include
MACOS_CGO_LDFLAGS := -L/opt/homebrew/Cellar/leptonica/1.87.0/lib -L/opt/homebrew/Cellar/tesseract/5.5.2/lib

.PHONY: help build build-all build-go build-go-api build-go-monitor build-go-benchmark-helper benchmark-helper build-stubs \
	build-grpc-stubs build-node-stubs build-python-stubs build-lua-descriptor \
	build-node build-node-binaries build-node-client build-python local-verify integration-verify real-desktop-e2e benchmark-find-on-screen-e2e clean

help:
	@echo "Targets:"
	@echo "  make build            Build all project outputs"
	@echo "  make build-go         Build sikuligo + sikuligo-monitor + benchmark-helper binaries"
	@echo "  make benchmark-helper Launch benchmark-helper web editor"
	@echo "  make build-stubs      Generate Go/Node/Python/Lua protocol artifacts"
	@echo "  make build-node       Build Node SDK + platform binaries"
	@echo "  make build-python     Build Python distribution artifacts"
	@echo "  make local-verify     Run local end-to-end CLI/client smoke verification"
	@echo "  make integration-verify  Run full local integration verification (RPC surface + API flows + Node/Python E2E)"
	@echo "  make real-desktop-e2e Optional manual real desktop E2E (FindOnScreen + OCR)"
	@echo "  make benchmark-find-on-screen-e2e  Benchmark FindOnScreen engines across size/orientation scenarios"
	@echo "  make clean            Remove build outputs"
	@echo ""
	@echo "Options:"
	@echo "  NPM_INSTALL=0         Skip npm install in Node build"
	@echo "  PYTHON_INSTALL=0      Skip pip install in Python build"
	@echo "  BENCH_HELPER_LISTEN=:8091      benchmark-helper listen address"
	@echo "  BENCH_HELPER_SPEC=<path>       benchmark-helper regions spec path"
	@echo "  BENCH_HELPER_MANIFEST=<path>   benchmark-helper scenario manifest path"
	@echo "  BENCH_HELPER_OPEN_BROWSER=1    open browser automatically when launching benchmark-helper"

build: build-all

build-all: build-stubs build-go build-node build-python

build-go: build-go-api build-go-monitor build-go-benchmark-helper

build-go-api:
	cd "$(API_DIR)" && \
	$(if $(filter Darwin,$(OS_NAME)),CGO_CXXFLAGS='$(MACOS_CGO_CXXFLAGS)' CGO_LDFLAGS='$(MACOS_CGO_LDFLAGS)',) \
	$(GO) build -tags "$(GO_OPENCV_TAGS)" -trimpath -ldflags="-s -w" -o "$(ROOT_DIR)/sikuligo" ./cmd/sikuligrpc

build-go-monitor:
	cd "$(API_DIR)" && \
	$(if $(filter Darwin,$(OS_NAME)),CGO_CXXFLAGS='$(MACOS_CGO_CXXFLAGS)' CGO_LDFLAGS='$(MACOS_CGO_LDFLAGS)',) \
	$(GO) build -tags "$(GO_OPENCV_TAGS)" -trimpath -ldflags="-s -w" -o "$(ROOT_DIR)/sikuligo-monitor" ./cmd/sikuligo-monitor

build-go-benchmark-helper:
	cd "$(API_DIR)" && \
	$(GO) build -trimpath -ldflags="-s -w" -o "$(ROOT_DIR)/benchmark-helper" ./cmd/benchmark-helper

benchmark-helper:
	cd "$(API_DIR)" && \
	listen="$(BENCH_HELPER_LISTEN)"; \
	hostport="$$listen"; \
	case "$$hostport" in \
		:*) hostport="127.0.0.1$$hostport" ;; \
		0.0.0.0:*) hostport="127.0.0.1:$${hostport#0.0.0.0:}" ;; \
		\[*\]:*) hostport="127.0.0.1:$${hostport##*:}" ;; \
	esac; \
	url="http://$$hostport"; \
	if [[ "$(BENCH_HELPER_OPEN_BROWSER)" == "1" ]]; then \
		( sleep 1; \
			if command -v open >/dev/null 2>&1; then \
				open "$$url"; \
			elif command -v xdg-open >/dev/null 2>&1; then \
				xdg-open "$$url"; \
			elif command -v powershell.exe >/dev/null 2>&1; then \
				powershell.exe -NoProfile -Command "Start-Process '$$url'"; \
			fi ) >/dev/null 2>&1 & \
		echo "benchmark-helper opening $$url"; \
	fi; \
	$(GO) run ./cmd/benchmark-helper \
		-listen "$$listen" \
		-spec "$(BENCH_HELPER_SPEC)" \
		-manifest "$(BENCH_HELPER_MANIFEST)"

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
	cd "$(ROOT_DIR)" && \
	if [[ "$(NPM_INSTALL)" == "1" ]]; then $(YARN) install; fi && \
	$(YARN) workspace @sikuligo/sikuligo build && \
	cd "$(ROOT_DIR)/packages/client-node" && \
	$(NPM) pack --dry-run

build-python:
	cd "$(ROOT_DIR)" && \
	SKIP_INSTALL=$$( [[ "$(PYTHON_INSTALL)" == "1" ]] && echo 0 || echo 1 ) \
	./scripts/clients/release-python-client.sh

local-verify:
	cd "$(ROOT_DIR)" && ./scripts/clients/local-verify.sh

integration-verify:
	cd "$(ROOT_DIR)" && ./scripts/clients/integration-verify.sh

real-desktop-e2e:
	cd "$(ROOT_DIR)" && REAL_DESKTOP_E2E=1 ./scripts/clients/real-desktop-e2e.sh

benchmark-find-on-screen-e2e:
	cd "$(ROOT_DIR)" && ./scripts/clients/benchmark-find-on-screen-e2e.sh

clean:
	cd "$(ROOT_DIR)" && \
	rm -rf \
	benchmark-helper \
	packages/client-node/dist \
	packages/client-node/generated \
	packages/client-node/packages/bin-*/bin/sikuligo \
	packages/client-node/packages/bin-*/bin/sikuligo.exe \
	packages/client-node/packages/checksums.txt \
	packages/client-python/dist \
	packages/client-python/build \
	packages/client-python/*.egg-info \
	packages/api/internal/grpcv1/pb \
	packages/api/sikuligo \
	packages/api/sikuligo-monitor
