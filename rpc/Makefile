# Ensure that GOPATH is in the PATH
PATH := $(PATH):$(shell go env GOPATH)/bin
CWD := $(shell pwd)

build:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --experimental_allow_proto3_optional caproto/*.proto
	protoc -I . -I .. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --experimental_allow_proto3_optional adminproto/*.proto
