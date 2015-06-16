package main

import (
	"fmt"
	"github.com/jdelgad/goforum/authenticator"
	"golang.org/x/crypto/ssh/terminal"
	"log"
)

func mainPrompt() int32 {
	var c int32
	fmt.Println("Menu")
	fmt.Println("===========")
	fmt.Println("1. Sign in")
	fmt.Println("2. Create a new account")
	fmt.Println("3. Quit")
	fmt.Scanf("%d", &c)
	return c
}

func loggedInPrompt() int32 {
	var c int32
	fmt.Println("Menu")
	fmt.Println("===========")
	fmt.Println("1. Logout")
	fmt.Scanf("%d", &c)
	return c
}

func main() {

	s := authenticator.SetupClientSocket("tcp://127.0.0.1:13000")
	defer s.Close()
	verified := false

	for !verified {
		var u string
		fmt.Print("Username: ")
		fmt.Scanf("%s", &u)

		fmt.Print("Enter password: ")
		p, err := terminal.ReadPassword(0)
		fmt.Println()

		if err != nil {
			panic("Could not obtain password")
		}

		req := authenticator.CreateLoginRequest()
		req.Username = &u
		pw, err := authenticator.EncryptPassword(p)

		if err != nil {
			fmt.Println("Could not encrypt password. Try again...")
			continue
		}

		req.Password = pw

		authenticator.SendLoginRequest(req, s)

		verified, err = authenticator.ServiceLoginReply(s)

		if err != nil {
			log.Fatal("Could not determine login status")
		}

		if !verified {
			fmt.Println("Login attempt failed")
		}
	}

	fmt.Println("\nLogin successful!\n")

	sel := loggedInPrompt()
	for sel != 1 {
		sel = loggedInPrompt()
	}
}
