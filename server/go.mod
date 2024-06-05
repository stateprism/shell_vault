module github.com/stateprism/prisma_ca/server

go 1.22.3

replace (
	github.com/stateprism/libprisma => ../libprisma
	github.com/stateprism/prisma_ca/rpc => ../rpc
)

require (
	github.com/google/uuid v1.6.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0
	github.com/mattn/go-sqlite3 v1.14.22
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
	github.com/alexeyxo/go-grpc-middleware/v2 v2.0.3 // indirect
	github.com/amazon-ion/ion-go v1.4.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/dgraph-io/badger/v4 v4.2.0 // indirect
	github.com/dgraph-io/ristretto v0.1.1 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v1.2.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/google/flatbuffers v1.12.1 // indirect
	github.com/klauspost/compress v1.12.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20240312152122-5f08fbb34913 // indirect
	go.opencensus.io v0.24.0 // indirect
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
