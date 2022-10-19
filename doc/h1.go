package doc

import (
	"gomd/err"
	"gomd/types"
)

type H1Element struct {
	Value types.Str
}

func (e H1Element) SameType(elem types.AnyElement) bool {
	return e.TypeName() == elem.TypeName()
}

func (e H1Element) TypeName() string {
	return "h1"
}

func TryH1(lines []types.Str) (types.AnyElement, []types.Str, error) {
	if captured, found := lines[0].Capture(`^ {0,3}# (.+)$`); found {
		return H1Element{
			Value: captured[0].Trim(),
		}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
