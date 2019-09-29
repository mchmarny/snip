package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserDirExists(t *testing.T) {
	err := initDir()
	assert.Nil(t, err)
	ok, err := userDirExists()
	assert.Nil(t, err)
	assert.True(t, ok)
}
