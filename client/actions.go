package main

import (
	"context"
	"crypto"
	"encoding/pem"
	"fmt"
	"github.com/stateprism/libprisma/cryptoutil/pkcrypto"
	"github.com/stateprism/shell_vault/client/clientutils"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
	"os"
	"strings"
	"time"
)

func refreshCaCertToFile(c *cli.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.Duration("timeout"))
	defer cancel()

	client := clientutils.NewClientConnection(ctx)
	err := client.TryConnect(c.String("addr"))
	if err != nil {
		return err
	}
	defer client.Close()

	tickRate := c.Duration("refresh-rate")

	tick := make(chan time.Time)
	errChan := make(chan error)
	tick <- time.Now()
	ticker := time.NewTicker(tickRate * time.Second)
	for {
		select {
		case t := <-ticker.C:
			tick <- t
		case t := <-tick:
			go func() {
				fmt.Println("Refreshing certificate at", t)
				cert, ttl, err := client.GetCurrentCert()
				if err != nil {
					errChan <- err
				}
				if cert == "" {
					errChan <- fmt.Errorf("no certificate returned")
				}

				cert = strings.Replace(cert, "\n", "", 1)
				line := fmt.Sprintf("%s %s", cert, time.Unix(ttl, 0).UTC())
				err = os.WriteFile(c.Path("output"), []byte(line), 0600)
				if err != nil {
					errChan <- err
				}
			}()
		case err := <-errChan:
			fmt.Println("Error:", err)
		}
	}
}

func refreshCaCert(c *cli.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.Duration("timeout"))
	defer cancel()

	client := clientutils.NewClientConnection(ctx)
	err := client.TryConnect(c.String("addr"))
	if err != nil {
		return err
	}
	defer client.Close()

	cert, ttl, err := client.GetCurrentCert()
	if err != nil {
		return err
	}
	if cert == "" {
		return fmt.Errorf("no certificate")
	}

	cert = strings.Replace(cert, "\n", "", 1)
	fmt.Printf("%s %s", cert, time.Unix(ttl, 0).UTC())

	return nil
}

func requestCert(c *cli.Context) error {
	fmt.Println("Requesting certificate")

	// authenticate
	// ask for username and password
	fmt.Println("Enter username:")
	var username string
	_, err := fmt.Scan(&username)
	if err != nil {
		return err
	}

	fmt.Println("Enter password:")
	var password string
	pw, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	fmt.Println()
	password = string(pw)

	ctx, cancel := context.WithTimeout(context.Background(), c.Duration("timeout"))
	defer cancel()

	client := clientutils.NewClientConnection(ctx)
	err = client.TryConnect(c.String("addr"))
	if err != nil {
		return err
	}
	defer client.Close()

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

	certBytes, err := client.RequestUserCert(key.Marshal())
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
}
