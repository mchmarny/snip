package cmd

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseCleanObjectives(t *testing.T) {
	items, err := parseItems("^obj1 ^obj2", objectiveRegExp)
	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 2, len(items))
	assert.Equal(t, "^obj1", items[0])
}

func TestParseNonObjectives(t *testing.T) {
	items, err := parseItems("obj1^obj2", objectiveRegExp)
	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 0, len(items))
}

func TestParseCleanContext(t *testing.T) {
	items, err := parseItems("@person @place", contextRegExp)
	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 2, len(items))
	assert.Equal(t, "@person", items[0])
}

func TestParseFullSnippet(t *testing.T) {

	txt := "did this and that with @person1 in @place1"
	raw := txt + " ^obj1"

	item, err := parseSnippet(raw)
	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, raw, item.Raw)
	assert.Equal(t, txt, item.Text)
}

func TestIDGetter(t *testing.T) {
	id := getID(time.Now())
	assert.NotNil(t, id)
}
