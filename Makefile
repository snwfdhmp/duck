#
# Makefile
#
VERSION = snapshot
GHRFLAGS =
.PHONY: build release

default: build

build:
	go build -o pkg/$(VERSION) duck.go

release:
	ghr  -u snwfdhmp  $(GHRFLAGS) v$(VERSION) pkg/$(VERSION)
