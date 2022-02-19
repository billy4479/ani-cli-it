GO ?= go
VERSION ?= $(shell git describe --always --tags)
OUTPUT_DIR ?= ./build


all: build
.PHONY: all

build:
	mkdir -p $(OUTPUT_DIR)
	$(GO) build -o $(OUTPUT_DIR)

.PHONY: build

run: build
	$(OUTPUT_DIR)/ani-cli-it

.PHONY: run
