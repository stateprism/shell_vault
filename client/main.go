package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stateprism/prisma_ca/client/clientutils"
	"github.com/urfave/cli/v2"
	"golang.org/x/term"
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

					fmt.Println("Enter your username: ")
					username, err := term.ReadPassword(0)
					if err != nil {
						return err
					}
					fmt.Println("Enter your password: ")
					password, err := term.ReadPassword(0)
					if err != nil {
						return err
					}

					err = client.Authenticate(string(username), string(password))
					if err != nil {
						return err
					}

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
