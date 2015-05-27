package main

import (
	"golang.org/x/crypto/ssh/terminal"
	"fmt"
	"os"
	"encoding/csv"
	"log"
	"errors"
)

// exists returns whether the given file or directory exists or not
func exists(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	return false
}

func openPasswd(file string) (map[string]string, error) {
	passwd := "passwd"
	if !exists(passwd) {
		log.Fatal("password file does not exist")
		return nil, errors.New("password file does not exist")
	}

	csvfile, err := os.Open("passwd")

	if err != nil {
		return nil, errors.New("could not open password file")
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1 // see the Reader struct information below

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		return nil, errors.New("could not read password file")
	}

	var userPass = make(map[string]string)

	// sanity check, display to standard output
	for _, each := range rawCSVdata {
		fmt.Println(each)
		userPass[each[0]] = each[1]
	}

	return userPass, nil
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
