package doc

import (
	"gomd/err"
	"gomd/types"
)

type ParagrahElement struct {
	Value types.Str
}

func (e ParagrahElement) SameType(elem types.AnyElement) bool {
	return e.TypeName() == elem.TypeName()
}

func (e ParagrahElement) TypeName() string {
	return "paragrah"
}

func TryParagraph(lines []types.Str) (types.AnyElement, []types.Str, error) {
	if !lines[0].IsEmpty() {
		return ParagrahElement{
			Value: lines[0],
		}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
