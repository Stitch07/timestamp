package timestamp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	date := New(time.Date(2008, time.January, 7, 2, 30, 25, 0, time.UTC))

	assert.Equal(t, "Jan 07 2008 02:30:25", date.Format("%b %d %Y %H:%M:%S"))
	assert.Equal(t, "January 7 08 02:30:25", date.Format("%B %-d %y %H:%M:%S"))
	assert.Equal(t, "07 Jan 08 02:30 UTC", date.Format(RFC822))
	assert.Equal(t, "07 Jan 08 02:30 +0000", date.Format(RFC822Z))
	assert.Equal(t, "Monday, 07-Jan-08 02:30:25 UTC", date.Format(RFC850))
	assert.Equal(t, "Mon, 07 Jan 2008 02:30:25 UTC", date.Format(RFC1123))
	assert.Equal(t, "Mon, 07 Jan 2008 02:30:25 +0000", date.Format(RFC1123Z))
	assert.Equal(t, "2008-01-07T02:30:25Z", date.Format(RFC3339))
	assert.Equal(t, "2:30AM", date.Format(Kitchen))
}

func BenchmarkTime_Format(b *testing.B) {
	date := Time{time.Date(2008, time.January, 7, 2, 30, 25, 0, time.UTC)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		date.Format(RFC3339)
	}
}

func BenchmarkNativeTime(b *testing.B) {
	date := time.Date(2008, time.January, 7, 2, 30, 25, 0, time.UTC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		date.Format(time.RFC3339)
	}
}
