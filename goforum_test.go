package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPasswordFailure(t *testing.T) {
	assert := assert.New(t)
	assert.False(validatePassword([]byte("testing"), "pow"))
}

func TestPasswordSuccess(t *testing.T) {
	assert := assert.New(t)
	assert.True(validatePassword([]byte("testing"), "testing"))
}
