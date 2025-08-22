# build imgbase binary in the current directory with the version set to the git tag
# or commit hash if there is no tag.
.PHONY: build
build:
	go build -ldflags "-s -w -X main.Version=$(shell git describe --match 'v[0-9]*' --dirty='.m' --always --tags)"