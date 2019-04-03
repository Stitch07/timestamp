package timestamp

import (
	"fmt"
	"time"
)

var verbs = map[rune]func(t Time) string{
	'A': func(t Time) string {
		return t.Weekday().String()
	},
	'a': func(t Time) string {
		return t.Weekday().String()[0:3]
	},
	'w': func(t Time) string {
		return fmt.Sprint(int(t.Weekday()))
	},
	'd': func(t Time) string {
		_, _, day := t.Date()
		return fmt.Sprint(day)
	},
	'b': func(t Time) string {
		return t.Month().String()[0:3]
	},
	'B': func(t Time) string {
		return t.Month().String()
	},
	'm': func(t Time) string {
		return fmt.Sprint(int(t.Month()))
	},
	'y': func(t Time) string {
		return fmt.Sprint(int(t.Year()))[2:]
	},
	'Y': func(t Time) string {
		return fmt.Sprint(t.Year())
	},
	'H': func(t Time) string {
		return fmt.Sprint(t.Hour())
	},
	'I': func(t Time) string {
		if t.Hour() > 11 {
			return fmt.Sprint(t.Hour() - 12)
		}
		return fmt.Sprint(t.Hour())
	},
	'p': func(t Time) string {
		if t.Hour() >= 12 {
			return "PM"
		}
		return "AM"
	},
	'M': func(t Time) string {
		return fmt.Sprint(t.Minute())
	},
	'S': func(t Time) string {
		return fmt.Sprint(t.Second())
	},
	'f': func(t Time) string {
		return fmt.Sprint(t.Microsecond())
	},
	'j': func(t Time) string {
		return fmt.Sprint(t.YearDay())
	},
	'U': func(t Time) string {
		_, week := t.ISOWeek()
		return padZero(fmt.Sprint(week))
	},
	'W': func(t Time) string {
		_, week := t.ISOWeek()
		return fmt.Sprint(week)
	},
	'Z': func(t Time) string {
		return t.Location().String()
	},
	'z': func(t Time) string {
		_, offset := t.Zone()
		offsetTime := time.Duration(offset)
		if offset < 0 {
			return "-" + padZero(fmt.Sprint(offsetTime.Hours())) + padZero(fmt.Sprint(offsetTime.Minutes()))
		}
		return "+" + padZero(fmt.Sprint(offsetTime.Hours())) + padZero(fmt.Sprint(offsetTime.Minutes()))
	},
}

// padZero adds a leading zero to a date
func padZero(date string) string {
	if len(date) < 2 {
		return fmt.Sprint("0", date)
	}
	return date
}
