# Development setup
To work on this project follow the steps outlined in this file

## Requisites 
Get protoc and the go codegen
```bash
sudo apt install protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

# Building and running
1. Make sure that the go path is in your PATH
2. Run `make` on the root folder of the repo
3. Run `devserver.sh` for a quick run of the server in dev mode or with goland use the usual run methods
