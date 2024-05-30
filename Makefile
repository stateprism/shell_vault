# Default target
all: build

build: proto
	go build -o bin/ ./...

proto:
	$(MAKE) -C rpc

serve: build
	bin/prisma_ca config.json

clean:
	rm -rf bin/
