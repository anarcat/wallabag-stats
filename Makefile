.PHONY: all build fmt lint test vet clean install 
TARGET := $(shell basename $(shell pwd))

all: clean build fmt lint test vet

build: get-deps
	@echo "+ $@"
	@go build .

get-deps: govendor
	@echo "+ $@"
	@go get github.com/golang/lint/golint

fmt:
	@echo "+ $@"
	@gofmt -s -l . | tee /dev/stderr

lint:
	@echo "+ $@"
	@for d in `govendor list -no-status +local | sed 's/github.com.Strubbl.wallabag-stats/./' | grep -v wallabag-stats` ; do \
		if [ "`golint $$d | tee /dev/stderr`"  ]; then \
			echo "^ golint errors!" && echo && exit 1; \
		fi \
	done

test: build fmt lint vet
	@echo "+ $@"
	govendor test +local

vet:
	@echo "+ $@"
	@if [ "`govendor vet +local | tee /dev/stderr`"  ]; then \
		echo "^ go vet errors!" && echo && exit 1; \
	fi

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

govendor:
	@echo "+ $@"
	@go get -u github.com/kardianos/govendor
	@go install github.com/kardianos/govendor
	@govendor sync github.com/Strubbl/wallabag-stats

update-vendor:
	@echo "+ $@"
	@govendor fetch -v +vendor

