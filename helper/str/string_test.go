package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatingLinesFromLinuxBasedString(t *testing.T) {
	text := "simple text\nwith new line"
	expected := []string{
		"simple text",
		"with new line",
	}

	result := Str(text).Lines()

	assert.Equal(t, expected, result)
}

func TestCreatingLinesFromWindowsBasedString(t *testing.T) {
	text := "simple text\r\nwith new line"
	expected := []string{
		"simple text",
		"with new line",
	}

	result := Str(text).Lines()

	assert.Equal(t, expected, result)
}

func TestGetFirstNCharactersOfString(t *testing.T) {
	text := "simple text"
	expected := "simp"

	result := Str(text).First(4)

	assert.Equal(t, expected, result)
}

func TestGetFirstNCharactersOfStringWithNOfSameStringSize(t *testing.T) {
	text := "simple text"
	expected := "simple text"

	result := Str(text).First(11)

	assert.Equal(t, expected, result)
}

func TestGetFirstNCharactersOfStringWithNOfBiggerThanStringSize(t *testing.T) {
	text := "simple text"
	expected := "simple text"

	result := Str(text).First(12)

	assert.Equal(t, expected, result)
}

func TestGetFromNCharactersOfString(t *testing.T) {
	text := "simple text"
	expected := "le text"

	result := Str(text).From(4)

	assert.Equal(t, expected, result)
}

func TestGetFromNCharactersOfStringFromTheBegining(t *testing.T) {
	text := "simple text"
	expected := "simple text"

	result := Str(text).From(0)

	assert.Equal(t, expected, result)
}

func TestGetFromNCharactersOfStringFromSomewhereFarFromStringSize(t *testing.T) {
	text := "simple text"
	expected := ""

	result := Str(text).From(11)

	assert.Equal(t, expected, result)
}

func TestBlankLineWithEmptyString(t *testing.T) {
	text := ""

	result := Str(text).IsBlank()

	assert.True(t, result)
}

func TestBlankLineWithSpaceLine(t *testing.T) {
	text := "     "

	result := Str(text).IsBlank()

	assert.True(t, result)
}

func TestBlankLineWithTabLine(t *testing.T) {
	text := "\t"

	result := Str(text).IsBlank()

	assert.True(t, result)
}

func TestBlankLineWithNonBlankLine(t *testing.T) {
	text := "simple text"

	result := Str(text).IsBlank()

	assert.False(t, result)
}

func TestBlankLineWithNonBlankLineStartingWithSpace(t *testing.T) {
	text := " simple text"

	result := Str(text).IsBlank()

	assert.False(t, result)
}

func TestMatchPatternWithRegularExpression(t *testing.T) {
	text := "==="

	result := Str(text).IsLike(`^={3}$`)

	assert.True(t, result)
}

func TestStringValue(t *testing.T) {
	text := "simple text"
	expected := "simple text"

	result := Str(text).String()

	assert.Equal(t, expected, result)
}

func TestTrim(t *testing.T) {
	text := "  simple text  "
	expected := "simple text"

	result := Str(text).Trim()

	assert.Equal(t, expected, result)
}

func TestCaptureMatch(t *testing.T) {
	text := "simple text"
	expected := "text"

	result := Str(text).Capture("(t..t)")

	assert.Equal(t, expected, result)
}

func TestCaptureUnmatch(t *testing.T) {
	text := "simple text"
	expected := ""

	result := Str(text).Capture("(t.st)")

	assert.Equal(t, expected, result)
}

func TestMultipleCaptureMatch(t *testing.T) {
	text := "simple text"
	expected := "ple"

	result := Str(text).Capture("(.l.) (t..t)")

	assert.Equal(t, expected, result)
}
