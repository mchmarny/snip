package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserDirExists(t *testing.T) {
	ok, err := userDirExists()
	assert.Nil(t, err)
	assert.True(t, ok)
}
