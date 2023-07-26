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
	line := lines[0].RTrim()
	if line.LooksLike(`^ {0,3}####( +#*)?$`) {
		return H4Element{
			Value: "",
		}, lines[1:], nil
	}
	if captured, found := line.Capture(`^ {0,3}#### (.+?)( #*)?$`); found {
		return H4Element{
			Value: captured[0].Trim(),
		}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
