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
	line := lines[0].RTrim()
	if line.LooksLike(`^ {0,3}#(\s+#*)?$`) {
		return H1Element{
			Value: "",
		}, lines[1:], nil
	}
	if captured, found := line.Capture(`^ {0,3}#\s(.+?)( #*)?$`); found {
		return H1Element{
			Value: captured[0].Trim(),
		}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
