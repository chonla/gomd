package doc

import (
	"gomd/err"
	"gomd/types"
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
	if lines[0].StartsWith("    ") {
		return CodeBlockElement{
			Value: lines[0][4:],
		}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
