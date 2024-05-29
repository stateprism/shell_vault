# Default target

# Ensure that GOPATH is in the PATH
PATH := $(PATH):$(shell go env GOPATH)/bin

all: build

build: proto
	go build -o bin/ ./...

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --experimental_allow_proto3_optional protocol/*.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --experimental_allow_proto3_optional adminprotocol/*.proto
	protoc --go_out=. --go_opt=paths=source_relative --experimental_allow_proto3_optional certificate/*.proto

serve: build
	bin/prisma_ca config.json

clean:
	rm -rf bin/
