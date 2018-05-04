
# Makefile
# Gustavo Paulo <gustavo.paulo.segura@gmail.com>

VERSION := $(shell git describe --always --all)

.PHONY : build install

build:
	go build -o dist/gh0std ./cmd/gh0std

install:
	go install ./cmd/gh0std

default: build
