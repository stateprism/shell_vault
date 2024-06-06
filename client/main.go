package main

import (
	"context"
	"crypto"
	"crypto/ed25519"
	"encoding/pem"
	"fmt"
	"github.com/stateprism/libprisma/cryptoutil/pkcrypto"
	"github.com/stateprism/prisma_ca/client/clientutils"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh"
	"os"
	"time"
)

func main() {
	app := &cli.App{
		Name:  "Prisma CA Client",
		Usage: "Client for the Prisma Certificate Authority",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "addr",
				Aliases: []string{"a"},
				Value:   "localhost:5000",
				Usage:   "Address of the Prisma CA server",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "new-cert",
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "localhost:5000",
						Usage:   "Where place the output files",
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("Requesting certificate")

					ctx := context.Background()
					client := clientutils.NewClientConnection(ctx)
					err := client.TryConnect(c.String("addr"))
					if err != nil {
						return err
					}
					defer client.Close()

					// authenticate
					// ask for username and password
					fmt.Println("Enter username:")
					var username string
					_, err = fmt.Scan(&username)
					if err != nil {
						return err
					}

					fmt.Println("Enter password:")
					var password string
					_, err = fmt.Scan(&password)
					if err != nil {
						return err
					}

					err = client.Authenticate(username, password)
					if err != nil {
						return err
					}

					k := pkcrypto.Ed25519.NewKey()
					pk := k.(ed25519.PrivateKey)
					var pub crypto.PublicKey
					pub = pk.Public()

					key, err := ssh.NewPublicKey(pub)
					if err != nil {
						return err
					}

					pKeyPem, err := ssh.MarshalPrivateKey(pk, fmt.Sprintf("Created for: %s", c.String("addr")))
					if err != nil {
						return err
					}

					err = os.WriteFile(c.Path("output"), pem.EncodeToMemory(pKeyPem), 0600)
					if err != nil {
						return err
					}

					certBytes, err := client.RequestCert(key.Marshal())
					if err != nil {
						return err
					}

					err = os.WriteFile(c.Path("output")+".pub", ssh.MarshalAuthorizedKey(key), 0600)
					if err != nil {
						return err
					}

					var certK ssh.PublicKey
					certK, _, _, _, err = ssh.ParseAuthorizedKey(certBytes)
					if err != nil {
						return err
					}
					cert := certK.(*ssh.Certificate)
					err = os.WriteFile(c.Path("output")+".cert.pub", certBytes, 0600)
					if err != nil {
						return err
					}

					fmt.Printf("Got certificate valid for: %s\n", time.Unix(int64(cert.ValidBefore), 0).Sub(time.Now()))

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
