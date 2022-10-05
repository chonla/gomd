package doc

import (
	"gomd/err"
	"gomd/types"
)

type H5Element struct {
	Value types.Str
}

func (e H5Element) SameType(elem types.AnyElement) bool {
	return e.TypeName() == elem.TypeName()
}

func (e H5Element) TypeName() string {
	return "h5"
}

func TryH5(lines []types.Str) (types.AnyElement, []types.Str, error) {
	if lines[0].StartsWith("##### ") {
		return H5Element{
			Value: lines[0][6:],
		}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
