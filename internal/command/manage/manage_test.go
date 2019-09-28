package manage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCleanTags(t *testing.T) {
	items, err := parseItems("#test1 #test2", tagRegExp)
	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 2, len(items))
}

func TestParseNonTags(t *testing.T) {
	items, err := parseItems("test1#test2", tagRegExp)
	assert.Nil(t, err)
	assert.Nil(t, items)
}

func TestParseCleanContext(t *testing.T) {
	items, err := parseItems("@test1 @test2", ctxRegExp)
	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 2, len(items))
}
