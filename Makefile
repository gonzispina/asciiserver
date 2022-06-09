# cross parameters
SHELL:=/bin/bash -O extglob
BINARY=bin/asciiserver
VERSION=0.1.0

LDFLAGS=-ldflags "-X main.Version=${VERSION}"

# Build step, generates the binary.
run:
	@docker-compose up -d
	@go build ${LDFLAGS} -o ${BINARY} cmd/*.go
	@./bin/asciiserver

# Run the test for all the directories.
test:
	@go test -v ./...