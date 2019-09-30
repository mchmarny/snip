package cmd

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeComp(t *testing.T) {
	b1 := toByte(time.Now())
	b2 := toByte(time.Now())
	assert.True(t, bytes.Compare(b1, b2) <= 0)
}
