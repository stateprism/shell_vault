package clientutils

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

func GetCredentials() (string, string) {
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
