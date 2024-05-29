# Default target
all: build

build: proto
	go build -o bin/ ./...

proto:
	PATH=$$PATH:$(shell go env GOPATH)/bin \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --experimental_allow_proto3_optional protocol/*.proto

serve: build
	bin/prisma_ca config.json

clean:
	rm -rf bin/
