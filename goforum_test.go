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

func TestPasswordFileExists(t *testing.T) {
	b := exists("passwd")
	assert.True(t, b)
}

func TestPasswordFileDoesNotExist(t *testing.T) {
	b := exists("fakePasswd")
	assert.False(t, b)
}

