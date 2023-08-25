package types

import (
	"strings"

	"github.com/samber/lo"
)

type StrArray []Str

func (sa StrArray) Join(sep string) Str {
	return Str(strings.Join(lo.Map(sa, func(s Str, index int) string { return string(s) }), sep))
}

func (sa StrArray) Merge() Str {
	return sa.Join("\n")
}
