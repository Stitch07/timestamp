package timestamp

import (
	"strings"
)

// Format formats a given time with the provided pattern. the pattern follows Python strftime(), meaning %A is converted to Monday.
// for a list of formats, check http://strftime.org/. For a regular percentage sign, use %%.
// All numeric constants are zero padded by default. To disable zero padding, add a - before the verb. (%-d).
func (t Time) Format(pattern string) string {
	return format(t, pattern)
}

// internal
func format(t Time, pattern string) string {
	w := strings.Builder{}
	r := strings.NewReader(pattern)
	for {
		char, _, err := r.ReadRune()
		// eof or some weird encoding issue. either way, we can stop parsing
		if err != nil {
			break
		}
		// check for % format
		if char == '%' {
			next, _, err := r.ReadRune()
			if err != nil {
				break
			}
			var s string
			if next == '-' {
				next, _, err := r.ReadRune()
				if err != nil {
					w.WriteRune('-')
					break
				}
				s = fmtVerb(next, t, false)
			} else {
				s = fmtVerb(next, t, true)
			}
			w.WriteString(s)
		} else {
			w.WriteRune(char)
		}
	}

	return w.String()
}

func fmtVerb(char rune, t Time, padZeroes bool) string {
	fn, ok := verbs[char]
	if ok {
		// we've found a format verb
		// check for zero padding some specific verbs
		s := fn(t)
		if padZeroes {
			s = padZero(s)
		}
		return s
	}
	return string(char)
}
