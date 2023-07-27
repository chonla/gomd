package doc

import (
	"gomd/err"
	"gomd/types"
	"strings"
)

type CodeBlockElement struct {
	Value types.Str
}

func (e CodeBlockElement) SameType(elem types.AnyElement) bool {
	return e.TypeName() == elem.TypeName()
}

func (e CodeBlockElement) TypeName() string {
	return "code_block"
}

func TryCodeBlock(lines []types.Str) (types.AnyElement, []types.Str, error) {
	line := lines[0]
	if captured, found := line.Capture(`^\t(\t+)`); found {
		// Expand leading 2 or more tabs into 3 spaces for each tabs
		line = line.ReplaceLike(`^\t\t+`, strings.Repeat("   ", len(captured[0])+1))
	}
	if captured, found := line.Capture(`^ *\t(.+)`); found {
		return CodeBlockElement{
			Value: captured[0].Trim(),
		}, lines[1:], nil
	}
	if line.StartsWith("    ") {
		return CodeBlockElement{
			Value: line[4:],
		}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
