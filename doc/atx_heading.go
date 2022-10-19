package doc

import (
	"gomd/err"
	"gomd/types"
)

func TryATXHeading(lines []types.Str) (types.AnyElement, []types.Str, error) {
	if elem, remainder, err := TryH1(lines); err == nil {
		return elem, remainder, err
	}
	if elem, remainder, err := TryH2(lines); err == nil {
		return elem, remainder, err
	}
	if elem, remainder, err := TryH3(lines); err == nil {
		return elem, remainder, err
	}
	if elem, remainder, err := TryH4(lines); err == nil {
		return elem, remainder, err
	}
	if elem, remainder, err := TryH5(lines); err == nil {
		return elem, remainder, err
	}
	if elem, remainder, err := TryH6(lines); err == nil {
		return elem, remainder, err
	}
	return nil, lines, err.ErrElementNotMatch
}
