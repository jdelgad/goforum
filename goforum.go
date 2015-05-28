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
	role     string
	password string
}

type Session struct {
	user   User
	active bool
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

func validateUsername(n string, users map[string]User) bool {
	_, ok := users[n]

	return ok
}

func validatePassword(b []byte, up string) bool {
	if string(b) == up {
		return true
	}

	return false
}

func Authenticate(name, pass string, users map[string]User) (Session, error) {
	user, ok := users[name]
	if !ok {
		return Session{}, errors.New("user does not exist")
	}

	var session Session
	if users[name].password == pass {
		session = Session{user: user, active: true}
	} else {
		session = Session{user: user, active: false}
	}

	return session, nil
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

func promptUser() int32 {
	var c int32
	fmt.Println("Menu")
	fmt.Println("===========")
	fmt.Println("1. Logout")
	fmt.Scanf("%d", &c)
	return c
}

func isLoggedIn(name string, session Session) bool {
	return session.user.username == name && session.active
}

func main() {
	users, err := openPasswordFile("passwd")

	if err != nil {
		panic("Could not open password file")
	}

	validUsername := false
	validPassword := false
	for !validUsername || !validPassword {
		var u string
		fmt.Print("Username: ")
		fmt.Scanf("%s", &u)

		fmt.Print("Enter password: ")
		pass, err := terminal.ReadPassword(0)
		fmt.Println()

		if err != nil {
			panic("Could not obtain password")
		}

		validUsername = validateUsername(u, users)

		user, ok := users[u]

		if !ok {
			validPassword = false
		} else {
			validPassword = validatePassword([]byte(pass), user.password)
		}
	}

	sel := promptUser()
	for sel != 1 {
		sel = promptUser()
	}
}
