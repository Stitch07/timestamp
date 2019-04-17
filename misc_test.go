package timestamp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	first := Time{time.Date(2009, time.April, 1, 0, 0, 0, 0, time.UTC)}
	second := Time{time.Date(2009, time.April, 3, 0, 0, 0, 0, time.UTC)}
	// the value being compared against is first, so comparing to second will be in the future
	assert.Equal(t, "in 2 days", Diff(first, second))
	// now first will be in the past
	assert.Equal(t, "2 days ago", Diff(second, first))
}
