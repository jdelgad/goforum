package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
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
	users, err := openPasswdFile("fakePasswd")
	assert.Nil(t, users)
	assert.Error(t, err)
}

func TestBlankPasswordFile(t *testing.T) {
	users, err := openPasswdFile("blankPasswd")
	assert.Empty(t, users)
	assert.NoError(t, err)
}

func TestOpenPasswordFile(t *testing.T) {
	users, err := openPasswdFile("passwd")
	assert.NotEmpty(t, users)
	assert.Equal(t, len(users), 2)
	assert.NoError(t, err)

	v, ok := users["jdelgad"]
	assert.NotNil(t, ok)
	assert.Equal(t, v, "pass")
}

func TestAuthenticate(t *testing.T) {
	users, err := openPasswdFile("passwd")
	if err != nil {
		assert.True(t, false)
	}

	for user,pass := range users {
		assert.True(t, Authenticate(user, pass, users))
	}

	assert.False(t, Authenticate("foo", "bar", users))
}