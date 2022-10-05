package doc

import (
	"gomd/err"
	"gomd/types"
)

type LiElement struct {
	Value types.Str
}

func (e LiElement) SameType(elem types.AnyElement) bool {
	return e.TypeName() == elem.TypeName()
}

func (e LiElement) TypeName() string {
	return "li"
}

func TryLi(lines []types.Str) (types.AnyElement, []types.Str, error) {
	if lines[0].StartsWith("* ") {
		return LiElement{
			Value: lines[0][2:],
		}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
