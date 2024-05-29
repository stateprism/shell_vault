package main

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/stateprism/prisma_ca/certificate"
	pb "github.com/stateprism/prisma_ca/protocol"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	app := &cli.App{
		Name:  "prismactl",
		Usage: "CLI for managing the Prisma CA server, CA certificate rotation, and key management",
		Commands: []*cli.Command{
			{
				Name:  "status",
				Usage: "Get the status of the Prisma CA server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "server-addr",
						Usage:   "Address of the Prisma CA server",
						EnvVars: []string{"PRISMA_CA_SERVER_ADDR"},
					},
				},
				Action: func(c *cli.Context) error {
					serverAddr := c.String("server-addr")
					if serverAddr == "" {
						fmt.Println("Server address is required")
						os.Exit(1)
					}
					fmt.Printf("Getting status of Prisma CA server at %s\n", serverAddr)
					client := MakeCLient(serverAddr)
					config, err := client.GetConfig(context.Background(), &pb.ConfigRequest{
						ClientVersion:         &pb.Version{Major: 1},
						ClientProtocolVersion: &pb.Version{Major: 1},
					})
					if err != nil {
						log.Fatalf("Failed to get config: %v", err)
					}

					fmt.Printf("Server protocol: %s\n", config.GetServerProtocolVersion().Display())
					fmt.Printf("Server policy: %s\n", config.GetPolicy())

					return nil
				},
			},
			{
				Name:  "create-admin-cert",
				Usage: "Create an admin certificate for the Prisma CA server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "username",
						Usage:    "Username of the admin",
						Required: true,
						Aliases:  []string{"u"},
					},
					&cli.PathFlag{
						Name:     "output",
						Usage:    "Output file for the admin key private key",
						Required: true,
						Aliases:  []string{"o"},
					},
					&cli.PathFlag{
						Name:        "server-config",
						Usage:       "Path to the server configuration file folder",
						DefaultText: "/etc/prisma-ca/",
						Aliases:     []string{"c"},
					},
					&cli.BoolFlag{
						Name:    "password",
						Usage:   "Prompt for a password to encrypt the private key",
						Aliases: []string{"p"},
					},
				},
				Action: func(c *cli.Context) error {
					username := c.String("username")
					output := c.Path("output")
					encrypt := c.Bool("password")
					serverConfig := c.Path("server-config")

					fmt.Printf("Creating admin certificate for %s\n", username)
					// Get trully random bytes
					seedBytes := make([]byte, ed25519.SeedSize)
					n, err := rand.Read(seedBytes)
					if n != ed25519.SeedSize || err != nil {
						return fmt.Errorf("failed to generate random bytes: %v", err)
					}
					sshKey := ed25519.NewKeyFromSeed(seedBytes)
					pubKey, _ := ssh.NewPublicKey(sshKey.Public())
					var privateKey *pem.Block
					if encrypt {
						// Ask for a password
						fmt.Println("Enter a password to encrypt the private key:")
						password, err := term.ReadPassword(int(os.Stdin.Fd()))
						if err != nil {
							return fmt.Errorf("failed to read password: %v", err)
						}
						// Confirm the password
						fmt.Println("Confirm the password:")
						confirm, err := term.ReadPassword(int(os.Stdin.Fd()))
						if err != nil {
							return fmt.Errorf("failed to read password: %v", err)
						}
						if string(password) != string(confirm) {
							return fmt.Errorf("passwords do not match")
						}
						privateKey, _ = ssh.MarshalPrivateKeyWithPassphrase(sshKey, fmt.Sprintf("Admin key for %s", username), password)
					} else {
						privateKey, _ = ssh.MarshalPrivateKey(sshKey, fmt.Sprintf("Admin key for %s", username))
					}
					// Check if the file already exists
					if _, err := os.Stat(output); err == nil {
						return fmt.Errorf("output file already exists")
					}
					err = os.WriteFile(output, pem.EncodeToMemory(privateKey), 0600)
					if err != nil {
						return fmt.Errorf("failed to write private key to file: %v", err)
					}

					aup := path.Join(serverConfig, "authorized_users.dat")
					us, err := certificate.LoadUserStore(aup, true)
					if err != nil {
						return fmt.Errorf("failed to load user store: %v", err)
					}
					us.AddUser(username, string(ssh.MarshalAuthorizedKey(pubKey)))
					err = us.SaveUserStore(aup)
					if err != nil {
						return fmt.Errorf("failed to save user store: %v", err)
					}

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func MakeCLient(serverAddr string) pb.PrismaCaClient {
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	return pb.NewPrismaCaClient(conn)
}
