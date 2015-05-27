package main

import (
	"golang.org/x/crypto/ssh/terminal"
	"fmt"
	"os"
)

// exists returns whether the given file or directory exists or not
func exists(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	return false
}

func openPasswd(file string) error {
	return nil
}

func validateUsername(n string, uns []string) bool {
	for i := range uns {
		if n == uns[i] {
			return true
		}
	}
	return false
}

func validatePassword(b []byte, up string) bool {
	if string(b) == up {
		return true
	}

	return false
}



func main() {
	var u string
	fmt.Print("Username: ")
	fmt.Scanf("%s", &u)

	fmt.Print("Enter password: ")
	_, err := terminal.ReadPassword(0)
	fmt.Println()

	if err != nil {
		panic("Could not obtain password")
	}

}
