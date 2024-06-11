# Default target
always: FORCE server client

server: FORCE
	$(MAKE) -C server

client: FORCE server
	$(MAKE) -C client

proto: FORCE
	$(MAKE) -C rpc

clean: FORCE
	rm -rf bin/

FORCE:
