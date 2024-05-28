package main

import (
	_ "embed"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/stateprism/prisma_ca/jsonprovider"
	pb "github.com/stateprism/prisma_ca/protocol"
	"github.com/stateprism/prisma_ca/providers"
	"github.com/stateprism/prisma_ca/server"
	"github.com/stateprism/prisma_ca/zaploggerprovider"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Embed the banner text
//
//go:embed banner.txt
var banner string

func Run() error {
	var zapLogger *zap.Logger
	var logger providers.LogProvider
	envProvider := providers.NewEnvProvider("PRISMA_CA_")
	if envProvider.IsEnvEqual("ENV", "DEV") {
		zapLogger, _ = zap.NewDevelopment()
		logger = zaploggerprovider.NewZapLoggerProvider(nil, zapLogger.Sugar())
	} else if envProvider.IsEnvEqual("ENV", "PROD") {
		zapLogger, _ := zap.NewProduction()
		logger = zaploggerprovider.NewZapLoggerProvider(zapLogger, nil)
	}

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
	pb.RegisterPrismaCaServer(srv, server.NewCAServer(provider, envProvider, logger))
	srv.Serve(listen)
	return nil
}

func main() {
	fmt.Println(banner)

	fmt.Println("Starting CA server")
	if err := Run(); err != nil {
		log.Fatalf("Error starting CA server: %v", err)
	}
}
