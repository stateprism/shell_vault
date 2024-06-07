BUILD_COMMIT:=$(shell git log --pretty=format:'%cs.%H' -n 1)
BUILD_DATE:=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
VERSION:=$(shell cat version.txt)

build:
	go build \
	-ldflags="-w -s -X main.Version=$(VERSION) -X main.BuildDate=$(BUILD_DATE) -X main.CommitInfo=$(BUILD_COMMIT)" \
	-o ../bin/ ./...