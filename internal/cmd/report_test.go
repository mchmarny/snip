package cmd

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWeekPeriodStart(t *testing.T) {
	lastSunday := getWeekPeriodStart(0)
	assert.Equal(t, time.Sunday, lastSunday.Weekday())

	sundayWeekAgo := getWeekPeriodStart(1)
	assert.Equal(t, time.Sunday, sundayWeekAgo.Weekday())
}
