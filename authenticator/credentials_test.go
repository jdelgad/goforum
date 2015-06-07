package authenticator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
}

func TestPasswordFailure(t *testing.T) {
	m := make(map[string]string)
	m["user"] = "badpass"
	assert.False(t, IsPasswordValid("testing", "user", m))
}

func TestPasswordSuccess(t *testing.T) {
	m := make(map[string]string)
	m["user"] = "testing"
	assert.True(t, IsPasswordValid("testing", "user", m))
}

func TestUsernameFailure(t *testing.T) {
	user := User{Username: "bad"}
	users := make(map[string]User, 1)
	users["bad"] = user
	assert.False(t, IsRegisteredUser("jdelgad", users))
}

func TestUsernameSuccess(t *testing.T) {
	user := User{Username: "jdelgad"}
	users := make(map[string]User, 1)
	users["jdelgad"] = user
	assert.True(t, IsRegisteredUser("jdelgad", users))
}

func TestPasswordFileDoesNotExist(t *testing.T) {
	users, err := GetUserPasswordList("fakePasswd")
	assert.Nil(t, users)
	assert.Error(t, err)
}

func TestBlankPasswordFile(t *testing.T) {
	users, err := GetUserPasswordList("blankPasswd")
	assert.Empty(t, users)
	assert.NoError(t, err)
}

func TestOpenPasswordFile(t *testing.T) {
	users, err := GetUserPasswordList("passwd")
	assert.NotEmpty(t, users)
	assert.Equal(t, len(users), 2)
	assert.NoError(t, err)

	v, ok := users["jdelgad"]
	assert.NotNil(t, ok)
	assert.Equal(t, v.Username, "jdelgad")
	assert.Equal(t, v.Password, "pass")
	assert.Equal(t, v.Role, "Admin")
}

func TestAuthenticate(t *testing.T) {
	users, err := GetUserPasswordList("passwd")
	if err != nil {
		assert.True(t, false)
	}

	for name, user := range users {
		_, ok := openSession(name, user.Password, users)
		assert.Nil(t, ok)
	}

	_, ok := openSession("foo", "bar", users)
	assert.NotNil(t, ok)
}

func TestRegularUser(t *testing.T) {
	users, err := GetUserPasswordList("passwd")

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
	users, err := GetUserPasswordList("passwd")

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
	LoggedInPrompt()
	// Output:
	// Menu
	// ===========
	// 1. Logout
}

func TestIsLoggedIn(t *testing.T) {
	users, err := GetUserPasswordList("passwd")
	if err != nil {
		assert.True(t, false)
	}

	session, err := openSession("jdelgad", "pass", users)
	v := isLoggedIn("jdelgad", session)
	assert.True(t, v)
	assert.Nil(t, err)

	session, err = openSession("newUser", "pass2", users)
	v = isLoggedIn("newUser", session)
	assert.True(t, v)
	assert.Nil(t, err)

	v = isLoggedIn("jdelgad", session)
	assert.False(t, v)
	assert.Nil(t, err)
}

func ExampleLoginPrompt() {
	mainPrompt()
	// Output:
	// Menu
	// ===========
	// 1. Sign in
	// 2. Create a new account
	// 3. Quit
}

func TestCreateUser(t *testing.T) {
	v, err := isValidUsername("newestUser")
	assert.True(t, v)
	assert.NoError(t, err)

	v, err = isValidUsername("jdelgad")
	assert.False(t, v)
	assert.Error(t, err)
}

func TestRegisterUser(t *testing.T) {
	registerUser("newestUser", "password")

	users, err := GetUserPasswordList("passwd")
	if err != nil {
		assert.True(t, false)
	}

	v := IsRegisteredUser("newestUser", users)

	assert.True(t, v)
}

func TestDeleteUser(t *testing.T) {
	registerUser("newestUser", "pass3")
	err := deleteUser("newestUser")

	assert.Nil(t, err)

	users, err := GetUserPasswordList("passwd")
	_, ok := users["newestUser"]

	assert.False(t, ok)
}

func ExampleInitialChoice() {
	initialChoice(1)
	// Output:
	// Menu
	// ===========
	// 1. Logout
}
