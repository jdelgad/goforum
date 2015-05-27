package main

import (
	"golang.org/x/crypto/ssh/terminal"
	"fmt"
)

func validatePassword(b []byte, up string) (good bool) {
	if string(b) == up {
		return true
	}

	return false
}

func main() {
	up := "test"
	pass, err := terminal.ReadPassword(0)

	if err != nil {
		panic("Could not obtain password")
	}

	good := validatePassword(pass, up)
	if good {
		fmt.Println("Success")
	} else {
		fmt.Println("Failure")
	}
}
