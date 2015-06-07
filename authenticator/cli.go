// +build cli

package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"github.com/jdelgad/goforum/authenticator"
)


func main() {
	users, err := authenticator.GetUserPasswordList("passwd")

	if err != nil {
		panic("Could not open password file")
	}

	vu := false
	vp := false
	for !vu || !vp {
		var u string
		fmt.Print("Username: ")
		fmt.Scanf("%s", &u)

		fmt.Print("Enter password: ")
		pass, err := terminal.ReadPassword(0)
		fmt.Println()

		if err != nil {
			panic("Could not obtain password")
		}

		vu = authenticator.IsRegisteredUser(u, users)

		user, ok := users[u]

		if !ok {
			vp = false
		} else {
			password := string(pass)
			vp = authenticator.IsPasswordValid(password, user.Password, users)
		}
	}

	sel := authenticator.LoggedInPrompt()
	for sel != 1 {
		sel =  authenticator.LoggedInPrompt()
	}
}
