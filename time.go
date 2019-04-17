// Timestamp - A general purpose time library for Go.
// Available at https://github.com/Soumil07/timestamp

// Copyright 2019 Soumil07. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package timestamp

import (
	"math"
	"time"
)

const (
	RFC822   = "%d %b %y %H:%M %Z"
	RFC822Z  = "%d %b %y %H:%M %z"
	RFC850   = "%A, %d-%b-%y %H:%M:%S %Z"
	RFC1123  = "%a, %d %b %Y %H:%M:%S %Z"
	RFC1123Z = "%a, %d %b %Y %H:%M:%S %z"
	RFC3339  = "%Y-%m-%dT%H:%M:%SZ"
	Kitchen  = "%-H:%M%p"
)

// Time is a generic wrapper over stdlib's time.Time to provide consistent methods
type Time struct {
	time.Time
}

// New returns a wrapper over the given time.Time
func New(t time.Time) Time {
	return Time{t}
}

func Now() Time {
	return Time{time.Now()}
}

// Microsecond returns the microsecond offset within the nanoseconds specified by t.
func (t Time) Microsecond() int {
	return int(math.Round(float64(t.Nanosecond() / 1000)))
}

// Millisecond returns the millisecond offset within the nanoseconds specified by t.
func (t Time) Millisecond() int {
	return int(math.Round(float64(t.Nanosecond() / 1e6)))
}
