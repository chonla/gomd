package str

import (
	"gomd/types"
)

// FromBytes
// Converts byes array to Str.
func FromBytes(data []byte) types.Str {
	return types.Str(data)
}
