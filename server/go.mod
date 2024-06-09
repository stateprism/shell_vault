module github.com/stateprism/shell_vault/server

go 1.22.3

replace (
	github.com/stateprism/libprisma => ../libprisma
	github.com/stateprism/shell_vault/rpc => ../rpc
)

require (
	github.com/google/uuid v1.6.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0
	github.com/joho/godotenv v1.5.1
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/pelletier/go-toml/v2 v2.2.2
	github.com/stateprism/libprisma v0.0.0-00010101000000-000000000000
	go.uber.org/fx v1.22.0
	golang.org/x/crypto v0.24.0
	google.golang.org/grpc v1.64.0
)

require (
	github.com/BurntSushi/toml v1.4.0 // indirect
	github.com/expr-lang/expr v1.16.9 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	go.uber.org/dig v1.17.1 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)

require (
	github.com/stateprism/shell_vault/rpc v0.0.0-00010101000000-000000000000
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
)
