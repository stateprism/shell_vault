package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"github.com/stateprism/shell_vault/server/authproviders/integratedprovider"
	"github.com/stateprism/shell_vault/server/configproviders/tomlprovider"
	"github.com/stateprism/shell_vault/server/middleware"
	"github.com/stateprism/shell_vault/server/providers"
	"github.com/stateprism/shell_vault/server/servers"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"path"
	"slices"
	"strings"
)

type RunMode string

type KeyPaths struct {
	Root   string
	Config string
	Data   string
	Log    string
}

type BootStrapResult struct {
	fx.Out

	Config providers.ConfigurationProvider
	Auth   providers.AuthProvider
	Logger *zap.Logger
	Mode   RunMode
}

type ProvideFlagsOut struct {
	fx.Out

	Paths         KeyPaths
	CleanLocalDbs bool
	Env           *providers.EnvProvider
}

type BootstrapParams struct {
	fx.In

	Paths         KeyPaths
	CleanLocalDbs bool
	Env           *providers.EnvProvider
}

// Bootstrap is a function that initializes the server with a few configurations and other dependencies
func Bootstrap(p BootstrapParams) (BootStrapResult, error) {
	empty := BootStrapResult{}
	configProvider, err := NewConfig(p.Paths.Config)
	if err != nil {
		return empty, err
	}

	// Set the environment variables to the configuration provider
	for k, v := range p.Env.GetEnvMap() {
		k = strings.ToLower(k)
		k = strings.TrimPrefix(k, "shell_vault_")
		err := configProvider.Set(k, v)
		if err != nil {
			return empty, err
		}
		p.Env.UnsetEnv(k)
	}

	// Set the key paths to the configuration provider
	_ = configProvider.Set("paths.config", p.Paths.Config)
	_ = configProvider.Set("paths.data", p.Paths.Data)
	_ = configProvider.Set("paths.log", p.Paths.Log)

	var logger *zap.Logger
	runMode := p.Env.GetEnvOrDefault("ENV", "PROD")
	if runMode == "DEV" || runMode == "MAINTENANCE" {
		l, _ := zap.NewDevelopment()
		logger = l
	} else {
		l, _ := zap.NewProduction()
		logger = l
	}

	logger.Info("Starting server in", zap.String("env", string(runMode)))
	logger.Info("Config path", zap.String("path", p.Paths.Config))

	if (runMode == "DEV" || runMode == "MAINTENANCE") && p.CleanLocalDbs {
		logger.Info("Cleaning local databases")
		localStore := configProvider.GetLocalStore()
		dbPath := path.Join(localStore, "users.db")
		if err := os.Remove(dbPath); err != nil {
			logger.Error("Failed to remove db", zap.String("path", dbPath), zap.Error(err))
		}
	} else if p.CleanLocalDbs && runMode != "DEV" {
		logger.Fatal("clean-local-dbs can only be used in DEV or MAINTENANCE modes")
	}

	authProvider, err := NewAuthProvider(configProvider, logger)
	if err != nil {
		return empty, err
	}

	configParams := BootStrapResult{
		Config: configProvider,
		Logger: logger,
		Mode:   RunMode(runMode),
		Auth:   authProvider,
	}

	return configParams, nil
}

type GrpcServerParams struct {
	fx.In
	Lc          fx.Lifecycle
	Config      providers.ConfigurationProvider
	Log         *zap.Logger
	CaServer    *servers.CaServer
	AdminServer *servers.AdminServer
	Auth        providers.AuthProvider
	RunMode     RunMode
}

func NewConfig(configPath string) (providers.ConfigurationProvider, error) {
	var config providers.ConfigurationProvider
	configTemp, err := tomlprovider.New(path.Join(configPath, "config.toml"))
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
			selector.UnaryServerInterceptor(middleware.UnaryServerInterceptor(p.Auth.GetSession), selector.MatchFunc(skipAuthService)),
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
			if p.RunMode == "DEV" {
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

func NewAuthProvider(config providers.ConfigurationProvider, logger *zap.Logger) (providers.AuthProvider, error) {
	providerName, err := config.GetString("providers.auth_provider.realm")
	if err != nil {
		return nil, err
	}
	switch providerName {
	case "local":
		logger.Info("Setup authentication with local provider")
		return integratedprovider.New(config, logger)
	default:
		return nil, fmt.Errorf("unknown auth provider")
	}
}
