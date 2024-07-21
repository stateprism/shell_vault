package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/xadaemon/shell_vault/server/consts"
	"github.com/xadaemon/shell_vault/server/localkeychain"
	"github.com/xadaemon/shell_vault/server/providers"
	"github.com/xadaemon/shell_vault/server/services"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"os"
	"path"
	"runtime"
	"text/template"
	"time"
)

// Embed the banner text
//
//go:embed banner.txt
var banner string

var Version string
var CommitInfo string
var BuildDate string

var flagFirstSetup bool
var configPathGlobal string
var overrideRoot string
var cleanLocalDBs bool
var startServer bool

func main() {
	flag.BoolVar(&flagFirstSetup, "first-setup", false, "Run the first setup")
	flag.StringVar(&configPathGlobal, "config", "", "Path to the configurations folder")
	flag.BoolVar(&cleanLocalDBs, "clean-local-dbs", false, "Clean local databases")
	flag.BoolVar(&startServer, "run", false, "Start the server")
	flag.StringVar(&overrideRoot, "override-root", "", "Override the root path for any relative paths in the configurations")
	showVersion := flag.Bool("version", false, "Show version information")
	flag.Parse()

	if flagFirstSetup {
		fmt.Printf(banner, Version, time.Now().Year())
		firstSetup()
		os.Exit(0)
	}

	if overrideRoot != "" {
		fmt.Printf("Overriding root path to: %s\n", overrideRoot)
		err := os.Chdir(overrideRoot)
		if err != nil {
			fmt.Println("Failed to change directory to:", overrideRoot)
			os.Exit(1)
		}
	}

	if *showVersion {
		fmt.Printf("Version: %s\nCommit: %s\nBuild Date: %s\n", Version, CommitInfo, BuildDate)
		os.Exit(0)
	}

	if startServer {
		fmt.Printf(banner, Version, time.Now().Year())
		fmt.Printf("\n\nstarting server\n\n")
		fx.New(
			fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
				return &fxevent.ZapLogger{Logger: log}
			}),
			fx.Provide(
				ProvideFlags,
				Bootstrap,
				services.NewCAServer,
				services.NewAdminServer,
				localkeychain.NewLocalKeychain,
				GrpcListen,
			),
			fx.Invoke(func(providers.KeychainProvider) {}),
			fx.Invoke(func([]*grpc.Server) {}),
		).Run()
	} else {
		fmt.Println("No action specified, see help bellow:")
		flag.PrintDefaults()
	}
}

func setupAndProvidePaths() KeyPaths {
	configPath := configPathGlobal
	dataPath := consts.UNIXVarFolder
	logPath := consts.UNIXVarLogFolder
	if configPathGlobal == "" {
		if runtime.GOOS == "windows" {
			configPath = consts.WindowsEtcFolder
			dataPath = consts.WindowsVarFolder
			logPath = consts.WindowsVarLogFolder
		} else {
			configPath = consts.UNIXEtcFolder
		}
	}

	if overrideRoot != "" {
		if runtime.GOOS == "windows" {
			configPath = consts.WindowsEtcFolder[3:]
			dataPath = consts.WindowsVarFolder[3:]
			logPath = consts.WindowsVarLogFolder[3:]
		} else {
			configPath = consts.UNIXEtcFolder[1:]
			dataPath = consts.UNIXVarFolder[1:]
			logPath = consts.UNIXVarLogFolder[1:]
		}
	}

	toCreate := []string{configPath, dataPath, logPath}
	for _, p := range toCreate {
		err := os.MkdirAll(p, 0755)
		if err != nil {
			fmt.Println("Failed to create path:", p)
			os.Exit(1)
		}
	}

	return KeyPaths{
		Root:   overrideRoot,
		Config: configPath,
		Data:   dataPath,
		Log:    logPath,
	}
}

func ProvideFlags() ProvideFlagsOut {
	paths := setupAndProvidePaths()
	err := godotenv.Load(path.Join(paths.Config, "server.env"))
	if err != nil {
		fmt.Println("Did not load env from server.env")
	}

	envProvider := providers.NewEnvProvider("SHELL_VAULT_")

	return ProvideFlagsOut{
		Paths:         paths,
		CleanLocalDbs: cleanLocalDBs,
		Env:           envProvider,
	}
}

func firstSetup() {
	fmt.Println("Running first setup")

	paths := setupAndProvidePaths()

	// Create the configurations file
	configFile := path.Join(paths.Config, consts.DefaultConfigFile)
	confTemplate, err := template.New("config").Parse(consts.ConfigTemplate)
	if err != nil {
		fmt.Println("Failed to parse configuration template, contact support")
		os.Exit(1)
	}
	fd, err := os.OpenFile(configFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to create configuration file", configFile, err)
		os.Exit(1)
	}

	err = confTemplate.Execute(fd, paths)
	if err != nil {
		fmt.Println("Failed to write configuration file", err)
		os.Exit(1)
	}

	// Create the environment file
	envFile := path.Join(paths.Config, consts.DefaultServerEnvFile)
	err = os.WriteFile(envFile, []byte(consts.ServerEnv), 0600)
	if err != nil {
		fmt.Println("Failed to create environment file", envFile)
		os.Exit(1)
	}

	fmt.Println(
		fmt.Sprintf(
			"Setup is now completed, check %s, for the configuration files and make any changes to your liking.",
			paths.Config,
		))

	fmt.Println()
}
