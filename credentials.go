package credentials

import (
	"fmt"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh/terminal"
)

// UsernameAndPassword asks the user for the credentials
func UsernameAndPassword() (username, password string, err error) {
	fmt.Print("Username: ")
	fmt.Scan(&username)

	fmt.Print("Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", errors.Wrap(err, "error reading password")
	}
	fmt.Println()
	password = string(bytePassword)

	return username, password, nil
}
