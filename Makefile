.PHONY: all build fmt lint test vet clean install 
TARGET := $(shell basename $(shell pwd))

all: clean build fmt lint test vet

build: get-deps
	@echo "+ $@"
	@go build .

get-deps:
	@echo "+ $@"
	@go get -t ./...
	@go get github.com/golang/lint/golint

fmt:
	@echo "+ $@"
	@gofmt -s -l . | tee /dev/stderr

lint:
	@echo "+ $@"
	@golint ./... | tee /dev/stderr

test: build fmt lint vet
	@echo "+ $@"
	@go test -v ./...

vet:
	@echo "+ $@"
	@go vet ./...

clean:
	@echo "+ $@"
	@rm -rf $(TARGET)

install:
	@echo "+ $@"
	@go install .

bump:
	@echo "+ $@"
	@./scripts/bump.sh
	@git diff

release:
	@echo "+ $@"
	@./scripts/release.sh > /dev/random 2>&1
	@ls -lh *.7z

