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
	if captured, found := lines[0].Capture(`^ {0,3}### (.+)$`); found {
		return H3Element{
			Value: captured[0].Trim(),
		}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
