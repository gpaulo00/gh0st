
# Makefile
# Gustavo Paulo <gustavo.paulo.segura@gmail.com>

VERSION := $(shell git describe --always --all)
PACKAGE := github.com/gpaulo00/gh0st
LDFLAGS := -ldflags "-X=$(PACKAGE)/models.Version=$(VERSION)"

.PHONY : build install

build:
	go build $(LDFLAGS) -o dist/gh0std ./cmd/gh0std
	go build $(LDFLAGS) -o dist/gh0st ./cmd/gh0st

install:
	go install $(LDFLAGS) ./cmd/gh0std
	go install $(LDFLAGS) ./cmd/gh0st

default: build
