prefix?=$(GOPATH)/bin
root?=$(shell pwd)
name?=$(notdir $(root))
srcpath?=$(root)
builddir?=$(root)/bin



goos?=$(shell go env GOOS)
arch?=$(shell go env GOARCH)
ldflags?=-ldflags "-s -w"
formatter?=$(shell which gofumpt)



mansrc?=$(root)/README.md
ifeq ($(shell uname),Linux)
	mandir?=/usr/local/man/man1
endif
ifeq ($(shell uname),Darwin)
	mandir?=/usr/local/share/man/man1
endif
ifeq ($(OS),Windows_NT)
	ext?=.exe
endif



.PHONY: tidy build man install pre post
default: all
all: $(prefix) pre tidy install post



pre: $(shell which go) $($(root)/go.mod) $(formatter)
	$(formatter) -e -w $(root)
	go mod tidy
	go clean
	mkdir -p $(builddir)
post:
	rm -rf $(builddir)



build: $(prefix) pre $(builddir)
	GOOS=$(goos) GOARCH=$(arch) go build $(ldflags) -o $(builddir)/$(name)$(ext) $(srcpath)

install: build 
	mv $(builddir)/$(name)$(ext) $(prefix)/
uninstall:
	rm -f $(prefix)/$(name)$(ext)



test:
	go test ./... -coverprofile=cover.out
	curl -Ls https://coverage.codacy.com/get.sh -o codacy.sh && \
	bash ./codacy.sh report -s --force-coverage-parser go -r cover.out -t ${CODACY_PROJECT_TOKEN}
lint:
	golangci-lint run ./... -c ./.golangci.yml
