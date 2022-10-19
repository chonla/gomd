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
	if captured, found := lines[0].Capture(`^ {0,3}###### (.+)$`); found {
		return H6Element{
			Value: captured[0].Trim(),
		}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
