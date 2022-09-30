package reader

import (
	"bytes"
	"gomd/types"
	"gomd/types/str"
	"os"
)

type Reader struct{}

func New() *Reader {
	return &Reader{}
}

// ReadTextFile
// Reads given text file and returns lines as array of string
func (r *Reader) ReadTextFile(fileName string) ([]types.Str, error) {
	data, e := os.ReadFile(fileName)
	if e != nil {
		return nil, e
	}

	var dataWithoutBOM []byte
	if bytes.HasPrefix(data, []byte{0xef, 0xbb, 0xbf}) {
		dataWithoutBOM = data[3:]
	} else {
		dataWithoutBOM = data
	}

	dataAsString := str.FromBytes(dataWithoutBOM)

	return dataAsString.ToLines(), nil
}
