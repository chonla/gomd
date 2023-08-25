package doc_test

import (
	"gomd/doc"
	"gomd/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeBlock(t *testing.T) {
	codeblocks := []string{
		"\tfoo\tbaz\t\tbim",   // Example 1
		"  \tfoo\tbaz\t\tbim", // Example 2
		"    a\ta",            // Example 3
		"    ὐ\ta",            // Example 3
		"\t\tbar",             // Example 5
		"    # foo",           // Example 69
	}

	elements := []types.AnyElement{
		doc.CodeBlockElement{
			Value: "foo\tbaz\t\tbim",
		},
		doc.CodeBlockElement{
			Value: "foo\tbaz\t\tbim",
		},
		doc.CodeBlockElement{
			Value: "a\ta",
		},
		doc.CodeBlockElement{
			Value: "ὐ\ta",
		},
		doc.CodeBlockElement{
			Value: "bar",
		},
		doc.CodeBlockElement{
			Value: "# foo",
		},
	}

	for i, line := range codeblocks {
		elem, remainder, e := doc.TryCodeBlock([]types.Str{
			types.Str(line),
		})

		assert.NoError(t, e)
		assert.Equal(t, []types.Str{}, remainder)
		assert.Equal(t, elements[i], elem)
	}
}
