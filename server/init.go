package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"github.com/xadaemon/shell_vault/server/authproviders/integratedprovider"
	"github.com/xadaemon/shell_vault/server/configproviders/tomlprovider"
	"github.com/xadaemon/shell_vault/server/middleware/auth"
	"github.com/xadaemon/shell_vault/server/plugins"
	"github.com/xadaemon/shell_vault/server/providers"
	"github.com/xadaemon/shell_vault/server/services"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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

	// Set the key paths to the configuration provider
	_ = configProvider.Set("paths.config", p.Paths.Config)
	_ = configProvider.Set("paths.data", p.Paths.Data)
	_ = configProvider.Set("paths.log", p.Paths.Log)

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

	var config zap.Config

	runMode := p.Env.GetEnvOrDefault("ENV", "PROD")
	if runMode == "DEV" || runMode == "MAINTENANCE" {
		config = zap.NewDevelopmentConfig()
		if providers.GetOrDefault(configProvider, "logging.put_to_files", false) {
			config.OutputPaths = append(config.ErrorOutputPaths, path.Join(p.Paths.Log, "server.log"))
			config.ErrorOutputPaths = append(config.ErrorOutputPaths, path.Join(p.Paths.Log, "server_error.log"))
		}
	} else {
		config = zap.NewProductionConfig()
		if providers.GetOrDefault(configProvider, "logging.put_to_files", false) {
			config.OutputPaths = append(config.ErrorOutputPaths, path.Join(p.Paths.Log, "server.log"))
			config.ErrorOutputPaths = append(config.ErrorOutputPaths, path.Join(p.Paths.Log, "server_error.log"))
		}
	}

	logger := zap.Must(config.Build())

	logger.Info("Starting server in", zap.String("env", string(runMode)))
	logger.Info("config path", zap.String("path", p.Paths.Config))

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

	plugs, err := plugins.NewProvider(plugins.ProviderParams{
		Config: configProvider,
		Logger: logger,
	})
	if err != nil {
		return empty, err
	}

	authProvider, err := NewAuthProvider(plugs, configProvider, logger)
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
	CaServer    *services.CaServer
	AdminServer *services.AdminServer
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
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}

	var aS *grpc.Server
	if p.RunMode == "DEV" {
		aS = grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				logging.UnaryServerInterceptor(InterceptorLogger(p.Log), opts...),
				selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(p.Auth.Authorize), selector.MatchFunc(skipAuthService)),
			),
		)
	} else {
		aS = grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(p.Auth.Authorize), selector.MatchFunc(skipAuthService)),
				recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(recoverPanic)),
			),
		)
	}
	var cS *grpc.Server
	if p.RunMode == "DEV" {
		cS = grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(p.Auth.Authorize), selector.MatchFunc(skipAuthService)),
				logging.UnaryServerInterceptor(InterceptorLogger(p.Log), opts...),
			),
		)
	} else {
		cS = grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(p.Auth.Authorize), selector.MatchFunc(skipAuthService)),
				logging.UnaryServerInterceptor(InterceptorLogger(p.Log), opts...),
				recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(recoverPanic)),
			),
		)
	}

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

func recoverPanic(p any) (err error) {
	return status.Error(codes.Internal, "Internal server error")
}

func skipAuthService(_ context.Context, c interceptors.CallMeta) bool {
	exemptList := []string{
		"/CertificateAuthority/Authenticate",
		//"/CertificateAuthority/GetCurrentKey",
	}
	shouldSkip := slices.Contains(exemptList, c.FullMethod())
	// skips on false
	return !shouldSkip
}

func NewAuthProvider(plugins *plugins.Provider, config providers.ConfigurationProvider, logger *zap.Logger) (providers.AuthProvider, error) {
	providerName, err := config.GetString("providers.auth_provider.realm")
	if err != nil {
		return nil, err
	}
	switch providerName {
	case "local":
		logger.Info("Setup authentication with local provider")
		return integratedprovider.New(plugins, config, logger)
	default:
		return nil, fmt.Errorf("unknown auth provider")
	}
}

func InterceptorLogger(l *zap.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make([]zap.Field, 0, len(fields)/2)

		for i := 0; i < len(fields); i += 2 {
			key := fields[i]
			value := fields[i+1]

			switch v := value.(type) {
			case string:
				f = append(f, zap.String(key.(string), v))
			case int:
				f = append(f, zap.Int(key.(string), v))
			case bool:
				f = append(f, zap.Bool(key.(string), v))
			default:
				f = append(f, zap.Any(key.(string), v))
			}
		}

		logger := l.WithOptions(zap.AddCallerSkip(1)).With(f...)

		switch lvl {
		case logging.LevelDebug:
			logger.Debug(msg)
		case logging.LevelInfo:
			logger.Info(msg)
		case logging.LevelWarn:
			logger.Warn(msg)
		case logging.LevelError:
			logger.Error(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}
