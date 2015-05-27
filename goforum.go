package main

import (
	"golang.org/x/crypto/ssh/terminal"
	"fmt"
)

func main() {
	up := "test"
	pass, err := terminal.ReadPassword(0)

	if err != nil {
		panic("Could not obtain password")
	}

	if string(pass) == up {
		fmt.Println("Success!")
	} else {
		fmt.Println("Failure")
	}
}
