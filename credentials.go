package credentials

import (
	"errors"
	"fmt"
	"log"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// UsernameAndPassword asks the user for the credentials
func UsernameAndPassword() (username, password string) {
	fmt.Print("Username: ")
	fmt.Scan(&username)

	fmt.Print("Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(errors.New("Error reading password: " + err.Error()))
	}
	fmt.Println()
	password = string(bytePassword)

	return username, password
}
