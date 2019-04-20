package timestamp

import (
	"fmt"
	"math"
)

// Diff returns the human representation of the difference between two Times.
// It adds the "in " prefix if the date is in the future, and the "ago" suffix if it was in the past.
func Diff(earlier, now Time) string {
	var prefix, suffix, ret string
	if now.Unix() < earlier.Unix() {
		suffix = " ago"
	} else {
		prefix = "in "
	}

	duration := math.Abs(float64(now.Unix() - earlier.Unix()))
	if duration < 45 {
		ret = "seconds"
		goto addSuffix
	} else if duration < 90 {
		ret = "a minute"
		goto addSuffix
	}
	// compare duration in minutes
	duration /= 60
	if duration < 45 {
		ret = fmt.Sprint(math.Round(duration), " minutes")
		goto addSuffix
	} else if duration < 90 {
		ret = "an hour"
		goto addSuffix
	}
	//compare duration in hours
	duration /= 60
	if duration < 22 {
		ret = fmt.Sprint(math.Round(duration), " hours")
		goto addSuffix
	} else if duration < 36 {
		ret = "a day"
		goto addSuffix
	}
	// compare duration in days
	duration /= 24
	if duration < 30 {
		ret = fmt.Sprint(math.Round(duration), " days")
		goto addSuffix
	} else if duration < 46 {
		ret = "a month"
		goto addSuffix
	} else if duration < 320 {
		ret = fmt.Sprint(math.Round(duration/30), " months")
		goto addSuffix
	} else if duration < 548 {
		ret = "a year"
		goto addSuffix
	}

	ret = fmt.Sprint(math.Round(duration/365), " years")
	goto addSuffix

addSuffix:
	return prefix + ret + suffix
}

// ToNow is a wrapper over Diff calling it with the current time as comparison
func ToNow(other Time) string {
	return Diff(Now(), other)
}
