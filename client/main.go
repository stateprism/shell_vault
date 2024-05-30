package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/stateprism/prisma_ca/client/clientutils"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Prisma CA Client",
		Usage: "Client for the Prisma Certificate Authority",
		Commands: []*cli.Command{
			{
				Name: "test-auth",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "addr",
						Aliases: []string{"a"},
						Value:   "localhost:5000",
						Usage:   "Address of the Prisma CA server",
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("Testing authentication")

					ctx := context.Background()
					client := clientutils.NewClientConnection(ctx)
					err := client.TryConnect(c.String("addr"))
					if err != nil {
						return err
					}
					defer client.Close()

					err = client.Authenticate("test", "test")
					if err != nil {
						return err
					}

					fmt.Println(hex.EncodeToString(client.GetToken()))

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
