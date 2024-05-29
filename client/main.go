package main

import (
	"context"
	"fmt"
	"net"
	"os"

	pb "github.com/stateprism/prisma_ca/protocol"
	jsonprovider "github.com/stateprism/prisma_ca/tomlprovider"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	configFile := os.Args[1]
	clientConfig, err := jsonprovider.New(configFile)
	if err != nil {
		fmt.Println("Error reading client configuration:", err)
		panic(err)
	}

	caAddr, err := clientConfig.GetString("ca_host")
	if err != nil {
		fmt.Println("Error reading CA address from configuration:", err)
		panic(err)
	}

	conn, err := grpc.NewClient(caAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error creating connection to CA:", err)
		panic(err)
	}

	defer conn.Close()
	c := pb.NewPrismaCaClient(conn)

	authReply, err := c.Authenticate(context.Background(), &pb.AuthRequest{
		Username:    pb.StringValue(os.Args[2]),
		Password:    pb.StringValue(os.Args[3]),
		AuthpTicket: nil,
	})

	if err != nil {
		fmt.Println("Error authenticating:", err)
		panic(err)
	}

	if !authReply.Success {
		fmt.Println("Authentication failed")
		return
	}

	certReply, err := c.RequestCert(context.Background(), &pb.CertRequest{
		RequestedValidity:             3600,
		ExtendedValidityJustification: pb.StringValue("I need it"),
		AuthToken:                     authReply.AuthToken,
	})

	if err != nil {
		fmt.Println("Error requesting certificate:", err)
		panic(err)
	}

	fmt.Println("Certificate received:", certReply)
}

func ResolveDnsToIp(dns string) (string, error) {
	ips, err := net.LookupIP(dns)
	if err != nil {
		return "", err
	}

	return ips[0].String(), nil
}
