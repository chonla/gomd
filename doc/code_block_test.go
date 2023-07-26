package doc_test

import (
	"gomd/doc"
	"gomd/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeBlock(t *testing.T) {
	headings := []string{
		"    # foo", // Example 69
	}

	elements := []types.AnyElement{
		doc.CodeBlockElement{
			Value: "# foo",
		},
	}

	for i, line := range headings {
		elem, remainder, e := doc.TryCodeBlock([]types.Str{
			types.Str(line),
		})

		assert.NoError(t, e)
		assert.Equal(t, []types.Str{}, remainder)
		assert.Equal(t, elements[i], elem)
	}
}
