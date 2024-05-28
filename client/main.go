package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/netip"
	"os"

	pb "github.com/stateprism/prism_ca/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientConfig struct {
	CAHost string `json:"ca_host"`
	CAPort string `json:"ca_port"`
}

func main() {
	clientConfig, err := LoadClientConfigFromFile()
	var caAddr string
	if err != nil {
		fmt.Println("Error loading client config file:", err)
		panic(err)
	}

	_, err = netip.ParseAddr(clientConfig.CAHost)
	if err != nil {
		caAddr, err = ResolveDnsToIp(clientConfig.CAHost)
		if err != nil {
			fmt.Println("Error resolving DNS to IP:", err)
			panic(err)
		}
	}

	if caAddr == "" {
		caAddr = clientConfig.CAHost
	}

	caAddr = fmt.Sprintf("%s:%s", caAddr, clientConfig.CAPort)

	conn, err := grpc.NewClient(caAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error creating connection to CA:", err)
		panic(err)
	}

	defer conn.Close()
	c := pb.NewCAClient(conn)
}

func ResolveDnsToIp(dns string) (string, error) {
	ips, err := net.LookupIP(dns)
	if err != nil {
		return "", err
	}

	return ips[0].String(), nil
}

func LoadClientConfigFromFile() (*ClientConfig, error) {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		return nil, err
	}

	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)
	config := &ClientConfig{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
