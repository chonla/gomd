package doc

import (
	"gomd/err"
	"gomd/types"
)

type H2Element struct {
	Value types.Str
}

func (e H2Element) SameType(elem types.AnyElement) bool {
	return e.TypeName() == elem.TypeName()
}

func (e H2Element) TypeName() string {
	return "h2"
}

func TryH2(lines []types.Str) (types.AnyElement, []types.Str, error) {
	if lines[0].StartsWith("## ") {
		return H2Element{
			Value: lines[0][3:],
		}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
