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
	line := lines[0].RTrim()
	if line.LooksLike(`^ {0,3}######(\s+#*)?$`) {
		return H6Element{
			Value: "",
		}, lines[1:], nil
	}
	if captured, found := line.Capture(`^ {0,3}######\s(.+?)( #*)?$`); found {
		return H6Element{
			Value: captured[0].Trim(),
		}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
