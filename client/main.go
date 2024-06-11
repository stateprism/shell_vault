package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

var Version string
var CommitInfo string
var BuildDate string

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
			&cli.DurationFlag{
				Name:  "timeout",
				Value: 10 * time.Second,
				Usage: "ttl for any network calls the client makes",
			},
			&cli.StringFlag{
				Name: "username",
				EnvVars: []string{
					"SHELL_VAULT_USERNAME",
				},
			},
			&cli.StringFlag{
				Name: "password",
				EnvVars: []string{
					"SHELL_VAULT_PASSWORD",
				},
			},
		},
		Commands: []*cli.Command{
			{
				Name: "version",
				Action: func(c *cli.Context) error {
					fmt.Printf("Version: %s\nCommit: %s\nBuild Date: %s\n", Version, CommitInfo, BuildDate)
					return nil
				},
			},
			{
				Name:   "get-key",
				Action: refreshCaCert,
			},
			{
				Name: "refresh-key",
				Flags: []cli.Flag{
					&cli.DurationFlag{
						Name:     "refresh-rate",
						Value:    2 * time.Hour,
						Required: true,
					},
					&cli.PathFlag{
						Name:     "output",
						Aliases:  []string{"o"},
						Required: true,
					},
				},
				Action: refreshCaCertToFile,
			},
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
				Action: requestCert,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
