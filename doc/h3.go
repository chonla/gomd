package doc

import (
	"gomd/err"
	"gomd/types"
)

type H3Element struct {
	Value types.Str
}

func (e H3Element) SameType(elem types.AnyElement) bool {
	return e.TypeName() == elem.TypeName()
}

func (e H3Element) TypeName() string {
	return "h3"
}

func TryH3(lines []types.Str) (types.AnyElement, []types.Str, error) {
	if lines[0].StartsWith("### ") {
		return H3Element{
			Value: lines[0][4:],
		}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
