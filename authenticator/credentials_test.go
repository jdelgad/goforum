package authenticator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordFailureNoUser(t *testing.T) {
	assert.False(t, IsValidUserPass("user", []byte("testing")))
}

func TestPasswordFailure(t *testing.T) {
	assert.False(t, IsValidUserPass("user", []byte("testing")))
}

func TestPasswordSuccess(t *testing.T) {
	assert.True(t, IsValidUserPass("jdelgad", []byte("pass")))
}

func TestUsernameFailure(t *testing.T) {
	assert.False(t, IsRegisteredUser("fakeUser"))
}

func TestUsernameSuccess(t *testing.T) {
	assert.True(t, IsRegisteredUser("jdelgad"))
}

func TestPasswordFileDoesNotExist(t *testing.T) {
	users, err := getUserPasswordList("fakePasswd")
	assert.Nil(t, users)
	assert.Error(t, err)
}

func TestBlankPasswordFile(t *testing.T) {
	users, err := getUserPasswordList("blankPasswd")
	assert.Empty(t, users)
	assert.NoError(t, err)
}

func TestOpenPasswordFile(t *testing.T) {
	users, err := getUserPasswordList("passwd")
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
	users, err := getUserPasswordList("passwd")
	if err != nil {
		assert.True(t, false)
	}

	for name, user := range users {
		_, ok := OpenSession(name, user.Password, users)
		assert.Nil(t, ok)
	}

	_, ok := OpenSession("foo", "bar", users)
	assert.NotNil(t, ok)
}

func TestRegularUser(t *testing.T) {
	users, err := getUserPasswordList("passwd")

	if err != nil {
		assert.True(t, false)
	}

	v, err := IsRegularUser("jdelgad", users)
	assert.False(t, v)
	assert.Nil(t, err)

	v, err = IsRegularUser("newUser", users)
	assert.True(t, v)
	assert.Nil(t, err)

	v, err = IsRegularUser("noSuchUser", users)
	assert.False(t, v)
	assert.NotNil(t, err)
}

func TestAdminUser(t *testing.T) {
	users, err := getUserPasswordList("passwd")

	if err != nil {
		assert.True(t, false)
	}

	v, err := IsAdminUser("jdelgad", users)
	assert.True(t, v)
	assert.Nil(t, err)

	v, err = IsAdminUser("newUser", users)
	assert.False(t, v)
	assert.Nil(t, err)
}

func TestIsLoggedIn(t *testing.T) {
	users, err := getUserPasswordList("passwd")
	if err != nil {
		assert.True(t, false)
	}

	session, err := OpenSession("jdelgad", "pass", users)
	v := IsLoggedIn("jdelgad", session)
	assert.True(t, v)
	assert.Nil(t, err)

	session, err = OpenSession("newUser", "pass2", users)
	v = IsLoggedIn("newUser", session)
	assert.True(t, v)
	assert.Nil(t, err)

	v = IsLoggedIn("jdelgad", session)
	assert.False(t, v)
	assert.Nil(t, err)
}

func TestCreateUser(t *testing.T) {
	v, err := IsValidNewUsername("newestUser")
	assert.True(t, v)
	assert.NoError(t, err)

	v, err = IsValidNewUsername("jdelgad")
	assert.False(t, v)
	assert.Error(t, err)
}

func TestRegisterUser(t *testing.T) {
	RegisterUser("newestUser", "password")

	v := IsRegisteredUser("newestUser")

	assert.True(t, v)
}

func TestDeleteUser(t *testing.T) {
	RegisterUser("newestUser", "pass3")
	err := DeleteUser("newestUser")

	assert.Nil(t, err)

	users, err := getUserPasswordList("passwd")
	_, ok := users["newestUser"]

	assert.False(t, ok)
}
