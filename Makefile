# info
VERSION=v0.0.1

PROJECT_NAME=$(shell basename "$(PWD)")
GO_VERSION=$(shell go version  | awk '{print $$3}')
BUILD_TIME=$(shell date +%FT%T%z)
OS_ARCH=$(shell go version  | awk '{print $$4}')
GIT_COMMIT=$(shell git rev-parse HEAD)
CGO_ENABLED=0
GO_BIN=$(GOBIN)

GO_MODULE=$(shell sed -n '/module/p'  go.mod | awk '{print $$2}')
GO_IMPORT_PATH=$(GO_MODULE)/cmd

LDFLAGS=-X $(GO_IMPORT_PATH)._version=$(VERSION) -X $(GO_IMPORT_PATH)._goVersion=$(GO_VERSION) -X $(GO_IMPORT_PATH)._buildTime=$(BUILD_TIME) -X $(GO_IMPORT_PATH)._osArch=$(OS_ARCH) -X $(GO_IMPORT_PATH)._gitCommit=$(GIT_COMMIT)

## help: Help for this project
help: Makefile
	@echo "Usage:\n  make [command]"
	@echo
	@echo "Available Commands:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	
## build: Compile the binary.
build:
	@go generate
	@go build -o $(PROJECT_NAME) -ldflags "$(LDFLAGS)"

## install: build and install.
install:
	@go generate
	@go build -o $(PROJECT_NAME) -ldflags "$(LDFLAGS)" 
	@mv $(PROJECT_NAME) $(GO_BIN)

## build-linux: Compile the linux binary.
build-linux:
	@go generate
	@CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 go build -o $(PROJECT_NAME) -ldflags "$(LDFLAGS)"

## run: Build and run
run: build
	@go generate
	@./$(PROJECT_NAME)

## clean: Clean build files.
clean:
	rm -f $(PROJECT_NAME)
	
