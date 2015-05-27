package main

import (
	"golang.org/x/crypto/ssh/terminal"
	"fmt"
)

func validateUsername(n string, uns []string) bool {
	found := false
	for i := range uns {
		if n == uns[i] {
			found = true
			break
		}
	}
	return found
}

func validatePassword(b []byte, up string) bool {
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
