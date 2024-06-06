package main

import (
	_ "embed"
	"fmt"
	"github.com/stateprism/shell_vault/server/localkeychain"
	"time"

	"github.com/stateprism/shell_vault/server/providers"
	"github.com/stateprism/shell_vault/server/servers"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Embed the banner text
//
//go:embed banner.txt
var banner string

type RunEnv string

func main() {
	fmt.Printf(banner, time.Now().Year())

	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			Bootstrap,
			servers.NewCAServer,
			servers.NewAdminServer,
			localkeychain.NewLocalKeychain,
			GrpcListen,
		),
		fx.Invoke(func(providers.KeychainProvider) {}),
		fx.Invoke(func([]*grpc.Server) {}),
	).Run()
}
