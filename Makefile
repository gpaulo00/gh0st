
# Makefile
# Gustavo Paulo <gustavo.paulo.segura@gmail.com>

VERSION := $(shell git describe --always --all)
PACKAGE := github.com/gpaulo00/gh0st
LDFLAGS := -ldflags "-X=$(PACKAGE)/models.Version=$(VERSION)"

.PHONY : build install

build:
	go build $(LDFLAGS) -o dist/gh0st ./main.go

install:
	go install $(LDFLAGS) ./main.go

default: build
