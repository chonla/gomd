package doc

import (
	"gomd/err"
	"gomd/types"
)

type EmptyElement struct {
	Value string
}

func (e EmptyElement) SameType(elem types.AnyElement) bool {
	return e.TypeName() == elem.TypeName()
}

func (e EmptyElement) TypeName() string {
	return "empty_line"
}

func TryEmpty(lines []types.Str) (types.AnyElement, []types.Str, error) {
	if lines[0].IsEmpty() {
		return EmptyElement{}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
