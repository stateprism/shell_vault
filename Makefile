# Default target
all: build

build: server client

server: proto
	$(MAKE) -C server

client: proto
	$(MAKE) -C client

proto:
	$(MAKE) -C rpc


clean:
	rm -rf bin/
