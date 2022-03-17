package str_test

import (
	"gomd/helper/str"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatingLinesFromLinuxBasedString(t *testing.T) {
	text := "simple text\nwith new line"
	expected := []str.Str{
		str.Str("simple text"),
		str.Str("with new line"),
	}

	result := str.Str(text).Lines()

	assert.Equal(t, expected, result)
}

func TestCreatingLinesFromWindowsBasedString(t *testing.T) {
	text := "simple text\r\nwith new line"
	expected := []str.Str{
		str.Str("simple text"),
		str.Str("with new line"),
	}

	result := str.Str(text).Lines()

	assert.Equal(t, expected, result)
}

func TestGetFirstNCharactersOfString(t *testing.T) {
	text := "simple text"
	expected := str.Str("simp")

	result := str.Str(text).First(4)

	assert.Equal(t, expected, result)
}

func TestGetFirstNCharactersOfStringWithNOfSameStringSize(t *testing.T) {
	text := "simple text"
	expected := str.Str("simple text")

	result := str.Str(text).First(11)

	assert.Equal(t, expected, result)
}

func TestGetFirstNCharactersOfStringWithNOfBiggerThanStringSize(t *testing.T) {
	text := "simple text"
	expected := str.Str("simple text")

	result := str.Str(text).First(12)

	assert.Equal(t, expected, result)
}

func TestGetFromNCharactersOfString(t *testing.T) {
	text := "simple text"
	expected := str.Str("le text")

	result := str.Str(text).From(4)

	assert.Equal(t, expected, result)
}

func TestGetFromNCharactersOfStringFromTheBegining(t *testing.T) {
	text := "simple text"
	expected := str.Str("simple text")

	result := str.Str(text).From(0)

	assert.Equal(t, expected, result)
}

func TestGetFromNCharactersOfStringFromSomewhereFarFromStringSize(t *testing.T) {
	text := "simple text"
	expected := str.Str("")

	result := str.Str(text).From(11)

	assert.Equal(t, expected, result)
}

func TestBlankLineWithEmptyString(t *testing.T) {
	text := ""

	result := str.Str(text).IsBlank()

	assert.True(t, result)
}

func TestBlankLineWithSpaceLine(t *testing.T) {
	text := "     "

	result := str.Str(text).IsBlank()

	assert.True(t, result)
}

func TestBlankLineWithTabLine(t *testing.T) {
	text := "\t"

	result := str.Str(text).IsBlank()

	assert.True(t, result)
}

func TestBlankLineWithNonBlankLine(t *testing.T) {
	text := "simple text"

	result := str.Str(text).IsBlank()

	assert.False(t, result)
}

func TestBlankLineWithNonBlankLineStartingWithSpace(t *testing.T) {
	text := " simple text"

	result := str.Str(text).IsBlank()

	assert.False(t, result)
}

func TestMatchPatternWithRegularExpression(t *testing.T) {
	text := "==="

	result := str.Str(text).IsLike(`^={3}$`)

	assert.True(t, result)
}

func TestStringValue(t *testing.T) {
	text := "simple text"
	expected := "simple text"

	result := str.Str(text).String()

	assert.Equal(t, expected, result)
}

func TestTrim(t *testing.T) {
	text := "  simple text  "
	expected := str.Str("simple text")

	result := str.Str(text).Trim()

	assert.Equal(t, expected, result)
}

func TestWithout(t *testing.T) {
	text := "simple text"
	expected := str.Str("simple ex")

	result := str.Str(text).Without("t")

	assert.Equal(t, expected, result)
}

func TestMultipleWithout(t *testing.T) {
	text := "simple text"
	expected := str.Str("sipl x")

	result := str.Str(text).Without("t", "m", "e")

	assert.Equal(t, expected, result)
}

func TestCaptureMatch(t *testing.T) {
	text := "simple text"
	expected := str.Str("text")

	result := str.Str(text).Capture("(t..t)")

	assert.Equal(t, expected, result)
}

func TestCaptureUnmatch(t *testing.T) {
	text := "simple text"
	expected := str.Str("")

	result := str.Str(text).Capture("(t.st)")

	assert.Equal(t, expected, result)
}

func TestMultipleCaptureMatch(t *testing.T) {
	text := "simple text"
	expected := str.Str("ple")

	result := str.Str(text).Capture("(.l.) (t..t)")

	assert.Equal(t, expected, result)
}

func TestAppendString(t *testing.T) {
	text := "simple"
	expected := str.Str("simple text")

	result := str.Str(text).Append(str.Str(" text"))

	assert.Equal(t, expected, result)
}

func TestEmpty(t *testing.T) {
	text := "simple"
	expected := str.Str("")

	result := str.Str(text).Empty()

	assert.Equal(t, expected, result)
}

func TestIsEmpty(t *testing.T) {
	text := str.Str("")

	assert.True(t, text.IsEmpty())
}

func TestNotIsEmpty(t *testing.T) {
	text := str.Str("simple text")

	assert.False(t, text.IsEmpty())
}

func TestReplace(t *testing.T) {
	text := str.Str("simple text")

	expected := str.Str("simplest test")

	result := text.Replace(" ", "st ").Replace("x", "s")

	assert.Equal(t, expected, result)
}

func TestReplaceRegex(t *testing.T) {
	text := str.Str("simple     text")

	expected := str.Str("simplest test")

	result := text.ReplaceRegex(`\s+`, "st ").ReplaceRegex(`x`, "s")

	assert.Equal(t, expected, result)
}

func TestLen(t *testing.T) {
	text := str.Str("simple text")

	expected := 11

	result := text.Len()

	assert.Equal(t, expected, result)
}

func TestRepeatWithPositive(t *testing.T) {
	text := str.Str("simple text")

	expected := str.Str("simple textsimple textsimple text")

	result := text.Repeat(3)

	assert.Equal(t, expected, result)
}

func TestRepeatWithZero(t *testing.T) {
	text := str.Str("simple text")

	expected := str.Str("")

	result := text.Repeat(0)

	assert.Equal(t, expected, result)
}

func TestRepeatWithNegative(t *testing.T) {
	text := str.Str("simple text")

	expected := str.Str("")

	result := text.Repeat(-1)

	assert.Equal(t, expected, result)
}
