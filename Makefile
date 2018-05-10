
# Makefile
# Gustavo Paulo <gustavo.paulo.segura@gmail.com>

VERSION := $(shell git describe --always --all)
PACKAGE := github.com/gpaulo00/gh0st
LDFLAGS := -ldflags "-X=$(PACKAGE)/models.Version=$(VERSION)"

.PHONY : build install tools

tools:
	go get -u github.com/gobuffalo/packr/...

build:
	packr build $(LDFLAGS) -o dist/gh0st ./main.go

install:
	packr install $(LDFLAGS) ./main.go

clean:
	rm -rf dist/

default: build
