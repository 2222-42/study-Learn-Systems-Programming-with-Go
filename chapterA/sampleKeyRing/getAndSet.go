package main

import (
	"fmt"
	"github.com/tmc/keyring"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

func main() {
	secretValue, err := keyring.Get("progo-keyring-test", "password")
	if err == keyring.ErrNotFound {
		fmt.Printf("Secret Value is not found. Please type: ")
		pw, err := terminal.ReadPassword(syscall.Stdin)
		if err != nil {
			panic(err)
		}
		err = keyring.Set("progo-keyring-test", "password", string(pw))
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("Secret Value: %s\n", secretValue)
	}
}
