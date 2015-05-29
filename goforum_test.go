package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordFailure(t *testing.T) {
	assert.False(t, isPasswordValid([]byte("testing"), "pow"))
}

func TestPasswordSuccess(t *testing.T) {
	assert.True(t, isPasswordValid([]byte("testing"), "testing"))
}

func TestUsernameFailure(t *testing.T) {
	user := User{username: "bad"}
	users := make(map[string]User, 1)
	users["bad"] = user
	assert.False(t, isRegisteredUser("jdelgad", users))
}

func TestUsernameSuccess(t *testing.T) {
	user := User{username: "jdelgad"}
	users := make(map[string]User, 1)
	users["jdelgad"] = user
	assert.True(t, isRegisteredUser("jdelgad", users))
}

func TestPasswordFileDoesNotExist(t *testing.T) {
	users, err := readPasswordFile("fakePasswd")
	assert.Nil(t, users)
	assert.Error(t, err)
}

func TestBlankPasswordFile(t *testing.T) {
	users, err := readPasswordFile("blankPasswd")
	assert.Empty(t, users)
	assert.NoError(t, err)
}

func TestOpenPasswordFile(t *testing.T) {
	users, err := readPasswordFile("passwd")
	assert.NotEmpty(t, users)
	assert.Equal(t, len(users), 2)
	assert.NoError(t, err)

	v, ok := users["jdelgad"]
	assert.NotNil(t, ok)
	assert.Equal(t, v.username, "jdelgad")
	assert.Equal(t, v.password, "pass")
	assert.Equal(t, v.role, "Admin")
}

func TestAuthenticate(t *testing.T) {
	users, err := readPasswordFile("passwd")
	if err != nil {
		assert.True(t, false)
	}

	for name, user := range users {
		_, ok := createSession(name, user.password, users)
		assert.Nil(t, ok)
	}

	_, ok := createSession("foo", "bar", users)
	assert.NotNil(t, ok)
}

func TestRegularUser(t *testing.T) {
	users, err := readPasswordFile("passwd")

	if err != nil {
		assert.True(t, false)
	}

	v, err := isRegularUser("jdelgad", users)
	assert.False(t, v)
	assert.Nil(t, err)

	v, err = isRegularUser("newUser", users)
	assert.True(t, v)
	assert.Nil(t, err)

	v, err = isRegularUser("noSuchUser", users)
	assert.False(t, v)
	assert.NotNil(t, err)
}

func TestAdminUser(t *testing.T) {
	users, err := readPasswordFile("passwd")

	if err != nil {
		assert.True(t, false)
	}

	v, err := isAdminUser("jdelgad", users)
	assert.True(t, v)
	assert.Nil(t, err)

	v, err = isAdminUser("newUser", users)
	assert.False(t, v)
	assert.Nil(t, err)
}

func ExamplePromptUser() {
	promptUser()
	// Output:
	// Menu
	// ===========
	// 1. Logout
}

func TestIsLoggedIn(t *testing.T) {
	users, err := readPasswordFile("passwd")
	if err != nil {
		assert.True(t, false)
	}

	session, err := createSession("jdelgad", "pass", users)
	v := isLoggedIn("jdelgad", session)
	assert.True(t, v)
	assert.Nil(t, err)

	session, err = createSession("newUser", "pass2", users)
	v = isLoggedIn("newUser", session)
	assert.True(t, v)
	assert.Nil(t, err)

	v = isLoggedIn("jdelgad", session)
	assert.False(t, v)
	assert.Nil(t, err)
}

func ExampleLoginPrompt() {
	loginPrompt()
	// Output:
	// Menu
	// ===========
	// 1. Sign in
	// 2. Create a new account
	// 3. Quit
}

func TestCreateUser(t *testing.T) {
	v, err := createUser("newestUser")
	assert.True(t, v)
	assert.NoError(t, err)

	v, err = createUser("jdelgad")
	assert.False(t, v)
	assert.Error(t, err)
}

func TestCreateUserPassword(t *testing.T) {
	createUserPassword("newestUser", "password")

	users, err := readPasswordFile("passwd")
	if err != nil {
		assert.True(t, false)
	}

	v := isRegisteredUser("newestUser", users)

	assert.True(t, v)
}

func ExampleInitialChoice() {
	initialChoice(1)
	// Output:
	// Menu
	// ===========
	// 1. Logout
}