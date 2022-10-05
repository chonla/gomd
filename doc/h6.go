package doc

import (
	"gomd/err"
	"gomd/types"
)

type H6Element struct {
	Value types.Str
}

func (e H6Element) SameType(elem types.AnyElement) bool {
	return e.TypeName() == elem.TypeName()
}

func (e H6Element) TypeName() string {
	return "h6"
}

func TryH6(lines []types.Str) (types.AnyElement, []types.Str, error) {
	if lines[0].StartsWith("###### ") {
		return H6Element{
			Value: lines[0][7:],
		}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
