module github.com/stateprism/shell_vault/client

replace (
	github.com/stateprism/libprisma => ../libprisma
	github.com/stateprism/shell_vault/rpc => ../rpc
)

go 1.22.3

require (
	github.com/stateprism/libprisma v0.0.0-20240531192245-981a7ab3f1f0
	github.com/stateprism/shell_vault/rpc v0.0.0-00010101000000-000000000000
	github.com/urfave/cli/v2 v2.27.2
	golang.org/x/crypto v0.24.0
	golang.org/x/term v0.21.0
	google.golang.org/grpc v1.64.0
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20240312152122-5f08fbb34913 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)
