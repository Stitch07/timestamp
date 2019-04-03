package timestamp

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestVerbs(t *testing.T) {
	date := Time{time.Date(2008, time.January, 7, 2, 0, 0, 0, time.UTC)}
	t.Run("full weekday", func(t *testing.T) {
		verb := verbs['A']
		assert.Equal(t, "Monday", verb(date))
	})
	t.Run("abbreviated weekday", func(t *testing.T) {
		verb := verbs['a']
		assert.Equal(t, "Mon", verb(date))
	})
	t.Run("weekday as decimal number", func(t *testing.T) {
		verb := verbs['w']
		assert.Equal(t, "1", verb(date))
	})
	t.Run("zero padded date", func(t *testing.T) {
		verb := verbs['d']
		assert.Equal(t, "7", verb(date))
	})
	t.Run("abbreviated month", func(t *testing.T) {
		verb := verbs['b']
		assert.Equal(t, "Jan", verb(date))
	})
	t.Run("full month", func(t *testing.T) {
		verb := verbs['B']
		assert.Equal(t, "January", verb(date))
	})
	t.Run("month number", func(t *testing.T) {
		verb := verbs['m']
		assert.Equal(t, "1", verb(date))
	})
}
