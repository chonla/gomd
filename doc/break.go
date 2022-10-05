package doc

import (
	"gomd/err"
	"gomd/types"
)

type BreakElement struct {
	Value types.Str
}

func (e BreakElement) SameType(elem types.AnyElement) bool {
	return e.TypeName() == elem.TypeName()
}

func (e BreakElement) TypeName() string {
	return "break"
}

func TryBreak(lines []types.Str) (types.AnyElement, []types.Str, error) {
	if lines[0].LooksLike(`^ {0,3}(\*[\t ]*){3,}[ \t]*$`) {
		return BreakElement{}, lines[1:], nil
	}
	if lines[0].LooksLike(`^ {0,3}(-[\t ]*){3,}[ \t]*$`) {
		return BreakElement{}, lines[1:], nil
	}
	if lines[0].LooksLike(`^ {0,3}(_[\t ]*){3,}[ \t]*$`) {
		return BreakElement{}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
