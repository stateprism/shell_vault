package main

import (
	_ "embed"
	"fmt"
	"net"
	"os"

	pb "github.com/stateprism/prisma_ca/protocol"
	"github.com/stateprism/prisma_ca/providers"
	"github.com/stateprism/prisma_ca/server"
	jsonprovider "github.com/stateprism/prisma_ca/tomlprovider"
	"github.com/stateprism/prisma_ca/zaploggerprovider"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Embed the banner text
//
//go:embed banner.txt
var banner string

func Run() error {
	var logger providers.LogProvider
	envProvider := providers.NewEnvProvider("PRISMA_CA_")
	if envProvider.IsEnvEqual("ENV", "PROD") {
		zapLogger, err := zap.NewProduction()
		if err != nil {
			return err
		}
		logger = zaploggerprovider.NewZapLoggerProvider(zapLogger, nil)
	} else {
		zapLogger, err := zap.NewDevelopment()
		if err != nil {
			return err
		}
		logger = zaploggerprovider.NewZapLoggerProvider(nil, zapLogger.Sugar())
	}

	envMode := envProvider.GetEnvOrDefault("ENV", "DEV")

	defer logger.Flush()

	configFile := os.Args[1]
	provider, err := jsonprovider.New(configFile)
	if err != nil {
		return err
	}

	addr, err := provider.GetString("ca_host")
	if err != nil {
		return err
	}

	listen, err := net.Listen("tcp", addr)
	fmt.Printf("Listening on %s\n", addr)

	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	if envMode == "DEV" {
		reflection.Register(srv)
	}

	pb.RegisterPrismaCaServer(srv, server.NewCAServer(provider, envProvider, logger, envMode))
	srv.Serve(listen)
	return nil
}

func main() {
	fmt.Println(banner)

	fmt.Println("Starting CA server")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
