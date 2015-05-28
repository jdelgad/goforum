package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

type User struct {
	username string
	role string
	password string
}

// exists returns whether the given file or directory exists or not
func exists(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	return false
}

func openPasswordFile(file string) (map[string]User, error) {
	if !exists(file) {
		return nil, errors.New("password file does not exist")
	}

	csvfile, err := os.Open(file)

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

	var userPass = make(map[string]User)

	for _, each := range rawCSVdata {
		var username, password, role string
		if len(each) == 3 {
			username, password, role = each[0], each[1], each[2]
		} else {
			username, password, role = each[0], each[1], "Regular"
		}

		userInfo := User{username: username, password: password, role: role}
		userPass[username] = userInfo
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

func Authenticate(name, pass string, users map[string]User) bool {
	_, ok := users[name]
	if !ok {
		return false
	}

	return users[name].password == pass
}

func isRegularUser(name string, users map[string]User) (bool, error) {
	user, ok := users[name]

	if !ok {
		return false, errors.New("user not found")
	}

	return user.role == "Regular", nil
}

func isAdminUser(name string, users map[string]User) (bool, error) {
	user, ok := users[name]

	if !ok {
		return false, errors.New("user not found")
	}

	return user.role == "Admin", nil
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
