package parser

import (
	"gomd/types"
)

func Preprocess(data types.Str) types.Str {
	// Setext H1 to ATX H1
	data = data.ReplaceLike(`(?s:(\n|^) {0,3}([^ \n][^\n]*)\n {0,3}={3,}(\n|$))`, "$1# $2$3")

	// Setext H2 to ATX H2
	data = data.ReplaceLike(`(?s:(\n|^) {0,3}([^ \n][^\n]*)\n {0,3}-{3,}(\n|$))`, "$1## $2$3")

	return data
}
