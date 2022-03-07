package str

import (
	"regexp"
	"strings"
)

type Str string

// Lines
// Break a string into lines (array of string), with autodetection of newline character.
func (s Str) Lines() []string {
	str := string(s)
	var sArray []string
	if strings.Contains(str, "\r\n") {
		sArray = strings.Split(str, "\r\n")
	} else {
		sArray = strings.Split(str, "\n")
	}
	return sArray
}

// First
// Get first n characters of string.
func (s Str) First(n int) string {
	str := string(s)
	if len(str) < n {
		return str
	}
	return str[0:n]
}

// From
// Get substring starting from (n + 1)th character.
func (s Str) From(n int) string {
	str := string(s)
	if len(str) < n {
		return ""
	}
	return str[n:]
}

// IsBlank
// Check if a string is blank or whitespace string.
func (s Str) IsBlank() bool {
	return s.Trim() == ""
}

// IsLike
// Pattern matching a string using regex.
func (s Str) IsLike(pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(string(s))
}

// String
// Get a string value
func (s Str) String() string {
	return string(s)
}

// Trim
// Remove surrounding whitespace from string
func (s Str) Trim() string {
	str := string(s)
	return strings.TrimSpace(str)
}

// Without
// Remove p from string.
func (s Str) Without(p string) string {
	str := string(s)
	return strings.ReplaceAll(str, p, "")
}

// Capture
// Capture a string that matches pattern. If not match, return empty string.
// If more than one string captured, return first captured.
func (s Str) Capture(pattern string) string {
	str := string(s)
	re := regexp.MustCompile(pattern)
	if s.IsLike(pattern) {
		m := re.FindStringSubmatch(str)
		if len(m) > 1 {
			return m[1]
		}
	}
	return ""
}
