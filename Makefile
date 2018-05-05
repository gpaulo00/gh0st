
# Makefile
# Gustavo Paulo <gustavo.paulo.segura@gmail.com>

VERSION := $(shell git describe --always --all)

.PHONY : build install

build:
	go build -o dist/gh0std ./cmd/gh0std
	go build -o dist/gh0st ./cmd/gh0st

install:
	go install ./cmd/gh0std
	go install ./cmd/gh0st

default: build
