package types

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/samber/lo"
)

// Str
// String wrapper.
type Str string

// ToLines
// Splits Str to lines ([]Str) using either LF or CRLF as separator.
func (s Str) ToLines() []Str {
	maybeUnixText := string(s)
	unixText := strings.ReplaceAll(maybeUnixText, "\r\n", "\n")
	return lo.Map(strings.Split(unixText, "\n"), func(line string, _ int) Str {
		return Str(line)
	})
}

// StartsWith
// Returns true if this string starts with otherString, otherwise return false.
func (s Str) StartsWith(otherString string) bool {
	return strings.HasPrefix(string(s), otherString)
}

// IsEmpty
// Returns true if this string is empty, otherwise return false.
func (s Str) IsEmpty() bool {
	return string(s) == ""
}

// LooksLike
// Returns true if this string matches with otherString, otherwise return false.
func (s Str) LooksLike(pattern string) bool {
	rx := regexp.MustCompile(pattern)
	return rx.MatchString(string(s))
}

func (s Str) ReplaceLike(pattern string, withText string) Str {
	rx := regexp.MustCompile(pattern)
	return Str(rx.ReplaceAllString(string(s), withText))
}

// Capture
// Returns captured element from regular expression and true if substring can be captured from pattern.
// Otherwise returns nil with false.
func (s Str) Capture(pattern string) ([]Str, bool) {
	rx := regexp.MustCompile(pattern)
	matches := rx.FindStringSubmatch(string(s))
	if len(matches) == 0 {
		return nil, false
	}
	captured := lo.Map(matches[1:], func(elem string, _ int) Str {
		return Str(elem)
	})
	return captured, true
}

// Trim
// Trims surrounding whitespaces
func (s Str) Trim() Str {
	return Str(strings.TrimSpace(string(s)))
}

// RTrim
// Trims ending whitespaces
func (s Str) RTrim() Str {
	return Str(strings.TrimRightFunc(string(s), unicode.IsSpace))
}
