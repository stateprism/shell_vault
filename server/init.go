package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"github.com/spf13/afero"
	"github.com/stateprism/libprisma/memkv"
	"github.com/stateprism/prisma_ca/server/authproviders/integratedprovider"
	"github.com/stateprism/prisma_ca/server/configproviders/tomlprovider"
	"github.com/stateprism/prisma_ca/server/middleware"
	"github.com/stateprism/prisma_ca/server/providers"
	"github.com/stateprism/prisma_ca/server/servers"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"path"
	"path/filepath"
	"slices"
)

type BootStrapResult struct {
	fx.Out
	Config providers.ConfigurationProvider
	Auth   providers.AuthProvider
	Logger *zap.Logger
	Env    *providers.EnvProvider
	REnv   RunEnv
	MemKV  *memkv.MemKV
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

func NewConfig(configPath string) (providers.ConfigurationProvider, error) {
	var config providers.ConfigurationProvider
	fs := afero.NewOsFs()
	p, _ := filepath.Abs(path.Join(configPath, "config.toml"))
	configTemp, err := tomlprovider.New(fs, p)
	if err != nil {
		return nil, fmt.Errorf("error setting up config provider with %s: %s", configPath, err)
	}
	config = configTemp
	return config, nil
}

func GrpcListen(p GrpcServerParams) []*grpc.Server {
	s := make([]*grpc.Server, 0)
	aS := grpc.NewServer()
	cS := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			selector.UnaryServerInterceptor(middleware.UnaryServerInterceptor(p.CaServer.Auth.GetSession), selector.MatchFunc(skipAuthService)),
			recovery.UnaryServerInterceptor(),
		),
	)

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

func skipAuthService(_ context.Context, c interceptors.CallMeta) bool {
	exemptList := []string{
		"/PrismaCa/Authenticate",
		"/PrismaCa/GetCurrentKey",
	}
	return !slices.Contains(exemptList, c.FullMethod())
}

func NewAuthProvider(config providers.ConfigurationProvider, logger *zap.Logger, kv *memkv.MemKV) (providers.AuthProvider, error) {
	providerName, err := config.GetString("providers.auth_provider.realm")
	if err != nil {
		return nil, err
	}
	switch providerName {
	case "local":
		logger.Info("Setup authentication with local provider")
		return integratedprovider.New(config, kv, logger)
	default:
		return nil, fmt.Errorf("unknown auth provider")
	}
}

func Bootstrap() (BootStrapResult, error) {
	configParams := BootStrapResult{}
	app := &cli.App{
		Name:  "prisma_ca",
		Usage: "Prisma Certificate Authority",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:     "config",
				Required: true,
				Usage:    "Path to the configuration dir",
				Aliases:  []string{"c"},
			},
			&cli.BoolFlag{
				Name:  "clean-local-dbs",
				Usage: "Clean local databases, this will reset all the user databases, needs env SHELL_VAULT_ENV=DEV",
			},
		},
		Action: func(c *cli.Context) error {
			configPath := c.String("config")
			configProvider, err := NewConfig(configPath)
			if err != nil {
				return err
			}
			envProvider := providers.NewEnvProvider("SHELL_VAULT_")
			memKV := memkv.NewMemKV(".", &memkv.Opts{CaseInsensitive: true})

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

			if rEnv == "DEV" && c.Bool("clean-local-dbs") {
				logger.Info("Cleaning local databases")
				localStore := configProvider.GetLocalStore()
				dbPath := path.Join(localStore, "users.db")
				if err := os.Remove(dbPath); err != nil {
					logger.Error("Failed to remove db", zap.String("path", dbPath), zap.Error(err))
				}
			} else if c.Bool("clean-local-dbs") && rEnv != "DEV" {
				return fmt.Errorf("clean-local-dbs can only be used in DEV mode current mode is: %s", rEnv)
			}

			authProvider, err := NewAuthProvider(configProvider, logger, memKV)
			if err != nil {
				return err
			}

			configParams = BootStrapResult{
				Config: configProvider,
				Env:    envProvider,
				Logger: logger,
				REnv:   rEnv,
				Auth:   authProvider,
				MemKV:  memKV,
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		return BootStrapResult{}, err
	}

	return configParams, nil
}
