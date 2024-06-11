package clientutils

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"golang.org/x/term"
	"os"
)

func GetCredentials(c *cli.Context) (string, string) {
	if c.String("username") != "" && c.String("password") != "" {
		return c.String("username"), c.String("password")
	}
	// authenticate
	// ask for username and password
	fmt.Println("Enter username:")
	var username string
	_, err := fmt.Scan(&username)
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter password:")
	var password string
	pw, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	fmt.Println()
	password = string(pw)

	return username, password
}
