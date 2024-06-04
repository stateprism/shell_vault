module github.com/stateprism/prisma_ca/server

go 1.22.3

replace (
	github.com/stateprism/libprisma => ../../libprisma
	github.com/stateprism/prisma_ca/rpc => ../rpc
)

require (
	github.com/google/uuid v1.6.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0
	github.com/msteinert/pam v1.2.0
	github.com/pelletier/go-toml/v2 v2.2.2
	github.com/stateprism/libprisma v0.0.0-00010101000000-000000000000
	github.com/urfave/cli/v2 v2.27.2
	go.uber.org/fx v1.21.1
	golang.org/x/crypto v0.23.0
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.1
)

require (
	github.com/amazon-ion/ion-go v1.4.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20240312152122-5f08fbb34913 // indirect
	go.uber.org/dig v1.17.1 // indirect
)

require (
	github.com/spf13/afero v1.11.0
	github.com/stateprism/prisma_ca/rpc v0.0.0-00010101000000-000000000000
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
)
