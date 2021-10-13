package adt

import (
	"regexp"
	"strings"
)

type String string

func (s String) Lines() StringArray {
	str := string(s)
	var sArray StringArray
	if strings.Contains(str, "\r\n") {
		sArray = NewStringArray(strings.Split(str, "\r\n"))
	} else {
		sArray = NewStringArray(strings.Split(str, "\n"))
	}
	return sArray
}

func (s String) Len() int {
	return len(s)
}

func (s String) FirstN(n int) String {
	if s.Len() < n {
		return s
	}
	return String(string(s)[0:n])
}

func (s String) From(n int) String {
	if s.Len() < n {
		return String("")
	}
	return String(string(s)[n:])
}

func (s String) Trim() String {
	return String(strings.TrimSpace(string(s)))
}

func (s String) EqualsTo(other string) bool {
	return string(s) == other
}

func (s String) Like(pattern string) bool {
	matcher := regexp.MustCompile(pattern)
	return matcher.Match([]byte(string(s)))
}

func (s String) Grep(pattern string, index int) (String, bool) {
	matcher := regexp.MustCompile(pattern)
	matches := matcher.FindStringSubmatch(string(s))
	if len(matches) > 1 {
		return String(matches[index+1]), true
	}
	return String(""), false
}

func (s String) Unescape() String {
	pattern := `\\([!\"#$%&'\(\)\*\+,-\./:;<=>?@\[\\\]^_` + "`" + `\{\|\}~])`
	re := regexp.MustCompile(pattern)
	return String(string(re.ReplaceAll([]byte(string(s)), []byte("$1"))))
}
