package str

import (
	"regexp"
	"strings"
)

type Str string

// Lines
// Break a string into lines (array of Str), with autodetection of newline character.
func (s Str) Lines() []Str {
	str := string(s)
	var sArray []string
	if strings.Contains(str, "\r\n") {
		sArray = strings.Split(str, "\r\n")
	} else {
		sArray = strings.Split(str, "\n")
	}

	var result []Str
	for _, line := range sArray {
		result = append(result, Str(line))
	}

	return result
}

// First
// Get first n characters of string.
func (s Str) First(n int) Str {
	str := string(s)
	if len(str) < n {
		return s
	}
	return Str(str[0:n])
}

// From
// Get substring starting from (n + 1)th character.
func (s Str) From(n int) Str {
	str := string(s)
	if len(str) < n {
		return Str("")
	}
	return Str(str[n:])
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
func (s Str) Trim() Str {
	str := string(s)
	return Str(strings.TrimSpace(str))
}

// Without
// Remove p from string. If others is provided, remove others too.
func (s Str) Without(p string, others ...string) Str {
	str := string(s)
	result := strings.ReplaceAll(str, p, "")
	for _, other := range others {
		result = strings.ReplaceAll(result, other, "")
	}
	return Str(result)
}

// Capture
// Capture a string that matches pattern. If not match, return empty string.
// If more than one string captured, return first captured.
func (s Str) Capture(pattern string) Str {
	str := string(s)
	re := regexp.MustCompile(pattern)
	if s.IsLike(pattern) {
		m := re.FindStringSubmatch(str)
		if len(m) > 1 {
			return Str(m[1])
		}
	}
	return Str("")
}

// Append
// Append a string and return the result.
func (s Str) Append(text Str) Str {
	str := string(s) + string(text)
	return Str(str)
}

// IsEmpty
// Check whether string is empty.
func (s Str) IsEmpty() bool {
	return string(s) == ""
}

// Empty
// Empty string.
func (s Str) Empty() Str {
	return Str("")
}

// Replace
// Replace all occurences of old in s with new string.
func (s Str) Replace(old, new string) Str {
	return Str(strings.ReplaceAll(string(s), old, new))
}

// ReplaceRegex
// ReplaceRegex all occurences of old (regex pattern) in s with new string.
func (s Str) ReplaceRegex(pat, new string) Str {
	pattern := regexp.MustCompile(pat)
	return Str(pattern.ReplaceAllString(string(s), new))
}

// Len
// String length.
func (s Str) Len() int {
	return len(string(s))
}

// Repeat
// Repeat string n times. If n is zero or less, return empty string.
func (s Str) Repeat(n int) Str {
	if n <= 0 {
		return s.Empty()
	}
	return Str(strings.Repeat(string(s), n))
}
