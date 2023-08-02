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
	line := lines[0].RTrim()
	if line.LooksLike(`^ {0,3}###(\s+#*)?$`) {
		return H3Element{
			Value: "",
		}, lines[1:], nil
	}
	if captured, found := line.Capture(`^ {0,3}###\s(.+?)( #*)?$`); found {
		return H3Element{
			Value: captured[0].Trim(),
		}, lines[1:], nil
	}
	return nil, lines, err.ErrElementNotMatch
}
