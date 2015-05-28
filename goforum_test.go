package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordFailure(t *testing.T) {
	assert.False(t, validatePassword([]byte("testing"), "pow"))
}

func TestPasswordSuccess(t *testing.T) {
	assert.True(t, validatePassword([]byte("testing"), "testing"))
}

func TestUsernameFailure(t *testing.T) {
	assert.False(t, validateUsername("jdelgad", []string{"bad"}))
}

func TestUsernameSuccess(t *testing.T) {
	assert.True(t, validateUsername("jdelgad", []string{"jdelgad"}))
}

func TestPasswordFileDoesNotExist(t *testing.T) {
	users, err := openPasswordFile("fakePasswd")
	assert.Nil(t, users)
	assert.Error(t, err)
}

func TestBlankPasswordFile(t *testing.T) {
	users, err := openPasswordFile("blankPasswd")
	assert.Empty(t, users)
	assert.NoError(t, err)
}

func TestOpenPasswordFile(t *testing.T) {
	users, err := openPasswordFile("passwd")
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
	users, err := openPasswordFile("passwd")
	if err != nil {
		assert.True(t, false)
	}

	for name, user := range users {
		assert.True(t, Authenticate(name, user.password, users))
	}

	assert.False(t, Authenticate("foo", "bar", users))
}

func TestRegularUser(t *testing.T) {
	users, err := openPasswordFile("passwd")

	if err!= nil {
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
	users, err := openPasswordFile("passwd")

	if err!= nil {
		assert.True(t, false)
	}

	v, err := isAdminUser("jdelgad", users)
	assert.True(t, v)
	assert.Nil(t, err)

	v, err = isAdminUser("newUser", users)
	assert.False(t, v)
	assert.Nil(t, err)
}
