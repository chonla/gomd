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
	if captured, found := lines[0].RTrim().Capture(`^ {0,3}##### (.+?)( #*)?$`); found {
		return H5Element{
			Value: captured[0].Trim(),
		}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
