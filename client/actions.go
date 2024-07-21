package main

import (
	"context"
	"crypto"
	"encoding/pem"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xadaemon/libprisma/cryptoutil/pkcrypto"
	"github.com/xadaemon/shell_vault/client/clientutils"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ssh"
	"os"
	"strings"
	"time"
)

func refreshCaCertToFile(c *cli.Context) error {
	username, password := clientutils.GetCredentials(c)

	ctx, cancel := context.WithTimeout(context.Background(), c.Duration("timeout"))
	defer cancel()

	client := clientutils.NewClientConnection()
	err := client.TryConnect(c.String("addr"))
	if err != nil {
		return err
	}
	defer client.Close()

	err = client.Authenticate(ctx, username, password)
	if err != nil {
		return err
	}

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
				cert, ttl, err := client.GetCurrentCert(ctx)
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
	username, password := clientutils.GetCredentials(c)

	ctx, cancel := context.WithTimeout(context.Background(), c.Duration("timeout"))
	defer cancel()

	client := clientutils.NewClientConnection()
	err := client.TryConnect(c.String("addr"))
	if err != nil {
		return err
	}
	defer client.Close()

	err = client.Authenticate(ctx, username, password)
	if err != nil {
		return err
	}

	cert, ttl, err := client.GetCurrentCert(ctx)
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
	username, password := clientutils.GetCredentials(c)

	ctx, cancel := context.WithTimeout(context.Background(), c.Duration("timeout"))
	defer cancel()

	client := clientutils.NewClientConnection()
	err := client.TryConnect(c.String("addr"))
	if err != nil {
		return err
	}
	defer client.Close()

	err = client.Authenticate(ctx, username, password)
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

	certBytes, err := client.RequestUserCert(ctx, key.Marshal())
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

func requestHostCert(c *cli.Context) error {
	username, password := clientutils.GetCredentials(c)

	ctx, cancel := context.WithTimeout(context.Background(), c.Duration("timeout"))
	defer cancel()

	client := clientutils.NewClientConnection()
	err := client.TryConnect(c.String("addr"))
	if err != nil {
		return err
	}

	err = client.Authenticate(ctx, username, password)
	if err != nil {
		return err
	}

	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	keys, err := loadHostKeys(c)
	if err != nil {
		return err
	}

	var cancellations []context.CancelFunc
	for _, key := range keys {
		ctx, cancel := context.WithTimeout(context.Background(), c.Duration("timeout"))
		cancellations = append(cancellations, cancel)

		certBytes, err := client.RequestHostCert(ctx, key.Marshal(), []string{hostname})
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
	}

	for _, cf := range cancellations {
		cf()
	}

	return nil
}

func loadHostKeys(c *cli.Context) ([]ssh.PublicKey, error) {
	keys := make([]ssh.PublicKey, 0)
	for _, path := range c.StringSlice("host-keys") {
		keyBytes, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		key, _, _, _, err := ssh.ParseAuthorizedKey(keyBytes)
		if err != nil {
			return nil, err
		}

		keys = append(keys, key)
	}

	return keys, nil
}
