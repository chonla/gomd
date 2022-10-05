package doc

import (
	"gomd/err"
	"gomd/types"
)

type H4Element struct {
	Value types.Str
}

func (e H4Element) SameType(elem types.AnyElement) bool {
	return e.TypeName() == elem.TypeName()
}

func (e H4Element) TypeName() string {
	return "h4"
}

func TryH4(lines []types.Str) (types.AnyElement, []types.Str, error) {
	if lines[0].StartsWith("#### ") {
		return H4Element{
			Value: lines[0][5:],
		}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
