# cross parameters
SHELL:=/bin/bash -O extglob
BINARY=bin/asciiserver
VERSION=0.1.0

LDFLAGS=-ldflags "-X main.Version=${VERSION}"

# Build step, generates the binary.
build:
	@go build ${LDFLAGS} -o ${BINARY} cmd/*.go