package types

import (
	"regexp"
	"strings"

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
