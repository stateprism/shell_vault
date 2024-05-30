package main

import (
	"context"
	_ "embed"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/spf13/afero"
	"github.com/stateprism/prisma_ca/server/configproviders/tomlprovider"
	"github.com/stateprism/prisma_ca/server/providers"
	"github.com/stateprism/prisma_ca/server/servers"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Embed the banner text
//
//go:embed banner.txt
var banner string

type RunEnv string

type BootStrapResult struct {
	fx.Out
	Config providers.ConfigurationProvider
	Auth   providers.AuthProvider
	Loger  *zap.Logger
	Env    *providers.EnvProvider
	REnv   RunEnv
}

type GrpcServerParams struct {
	fx.In
	Lc          fx.Lifecycle
	Config      providers.ConfigurationProvider
	Log         *zap.Logger
	CaServer    *servers.CaServer
	AdminServer *servers.AdminServer
	REnv        RunEnv
}

func NewConfig(configStr string) (providers.ConfigurationProvider, error) {
	configPath := configStr
	var config providers.ConfigurationProvider
	if strings.HasPrefix(configPath, "file://") {
		fs := afero.NewOsFs()
		configPath = strings.TrimPrefix(configPath, "file://")
		configTemp, err := tomlprovider.New(fs, configPath)
		if err != nil {
			return nil, err
		}
		config = configTemp
	} else {
		return nil, fmt.Errorf("invalid configuration source")
	}
	return config, nil
}

func GrpcListen(p GrpcServerParams) []*grpc.Server {
	s := make([]*grpc.Server, 0)
	aS := grpc.NewServer()
	cS := grpc.NewServer()

	p.Lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			listenCS, err := net.Listen("tcp", p.CaServer.Listen)
			if err != nil {
				return err
			}
			listenAS, err := net.Listen("tcp", p.AdminServer.Listen)
			if err != nil {
				return err
			}

			p.CaServer.RegisterServer(cS)
			p.AdminServer.RegisterServer(aS)
			if p.REnv == "DEV" {
				p.Log.Debug("Starting server in debug mode with reflection")
				reflection.Register(cS)
				reflection.Register(aS)
			}

			go cS.Serve(listenCS)
			p.Log.Info("Server started at", zap.String("listen", p.CaServer.Listen))
			go aS.Serve(listenAS)
			p.Log.Info("Admin server started at", zap.String("listen", p.AdminServer.Listen))
			s = append(s, cS, aS)
			return nil
		},
		OnStop: func(context.Context) error {
			for _, server := range s {
				server.GracefulStop()
			}
			return nil
		},
	})
	return s
}

func Bootstrap() BootStrapResult {
	configParams := BootStrapResult{}
	app := &cli.App{
		Name:  "prisma_ca",
		Usage: "Prisma Certificate Authority",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Required: true,
				Usage:    "Path to the configuration file",
				Aliases:  []string{"c"},
			},
		},
		Action: func(c *cli.Context) error {
			configPath := c.String("config")
			configProvider, err := NewConfig(configPath)
			if err != nil {
				return err
			}
			envProvider := providers.NewEnvProvider("PRISMA_CA_")
			authProvider, err := providers.NewAuthProvider(configProvider)
			if err != nil {
				return err
			}

			var logger *zap.Logger
			var rEnv RunEnv
			if envProvider.GetEnvOrDefault("ENV", "PROD") == "DEV" {
				l, _ := zap.NewDevelopment()
				rEnv = "DEV"
				logger = l
			} else {
				l, _ := zap.NewProduction()
				rEnv = "PROD"
				logger = l
			}

			configParams = BootStrapResult{
				Config: configProvider,
				Env:    envProvider,
				Loger:  logger,
				REnv:   rEnv,
				Auth:   authProvider,
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}

	return configParams
}

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
			GrpcListen,
		),
		fx.Invoke(func([]*grpc.Server) {}),
	).Run()
}
