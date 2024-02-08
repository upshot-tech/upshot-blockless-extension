# Makefile for building the upshot-blockless-extension

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=allora-offchain-extension
SOURCE_FILE=main.go

# Detect the operating system and architecture
UNAME_S := $(shell uname -s)
UNAME_M := $(shell uname -m)
ifeq ($(UNAME_S),Linux)
    OS := linux
    ifeq ($(UNAME_M),x86_64)
        ARCH := x86_64
    endif
    ifeq ($(UNAME_M),aarch64)
        ARCH := aarch64
    endif
endif
ifeq ($(UNAME_S),Darwin)
    OS := macos
    ifeq ($(UNAME_M),x86_64)
        ARCH := x86_64
    endif
    ifeq ($(UNAME_M),arm64)
        ARCH := aarch64
    endif
endif
ifeq ($(OS_ARCH),windows)
    OS := windows
    ARCH := x86_64
endif

# Define the URL based on the detected OS and architecture
RUNTIME_URL := https://github.com/blocklessnetwork/runtime/releases/download/v0.3.1/blockless-runtime.$(OS)-latest.$(ARCH).tar.gz

all: build

$(BINARY_NAME): $(SOURCE_FILE)
	$(GOBUILD) -o $(BINARY_NAME) $(SOURCE_FILE)

build: $(BINARY_NAME)

clean:
	$(GOCLEAN)
	rm -rf $(BINARY_NAME)

example:
	cd allora-inference-function && npm run build:release

setup:
	@echo "\n📥 Downloading and extracting runtime...\n"
	mkdir -p /tmp/runtime
	wget -O /tmp/blockless-runtime.tar.gz $(RUNTIME_URL)
	tar -xzf /tmp/blockless-runtime.tar.gz -C /tmp/runtime
	@echo "\n✅ Done.\n"

test: example build
	export ALLORA_ARG_PARAMS=yuga BLS_LIST_VARS=ALLORA_ARG_PARAMS; \
	/tmp/runtime/bls-runtime allora-inference-function/build/release.wasm - --drivers-root-path=$(PWD)

.PHONY: all build clean example setup test
