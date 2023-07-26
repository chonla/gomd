package doc_test

import (
	"gomd/doc"
	"gomd/err"
	"gomd/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestATXHeadings(t *testing.T) {
	headings := []string{
		"# foo",                 // Example 62
		"## foo",                // Example 62
		"### foo",               // Example 62
		"#### foo",              // Example 62
		"##### foo",             // Example 62
		"###### foo",            // Example 62
		"# foo *bar* \\*baz\\*", // Example 66
		"#                  foo                     ", // Example 67
		" ### foo",           // Example 68
		"  ## foo",           // Example 68
		"   # foo",           // Example 68
		"## foo ##",          // Example 71
		"  ###   bar    ###", // Example 71
		"# foo ##################################", // Example 72
		"##### foo ##",     // Example 72
		"### foo ###     ", // Example 73
		"### foo ### b",    // Example 74
		"# foo#",           // Example 75
	}

	elements := []types.AnyElement{
		doc.H1Element{
			Value: "foo",
		},
		doc.H2Element{
			Value: "foo",
		},
		doc.H3Element{
			Value: "foo",
		},
		doc.H4Element{
			Value: "foo",
		},
		doc.H5Element{
			Value: "foo",
		},
		doc.H6Element{
			Value: "foo",
		},
		doc.H1Element{
			Value: "foo *bar* \\*baz\\*",
		},
		doc.H1Element{
			Value: "foo",
		},
		doc.H3Element{
			Value: "foo",
		},
		doc.H2Element{
			Value: "foo",
		},
		doc.H1Element{
			Value: "foo",
		},
		doc.H2Element{
			Value: "foo",
		},
		doc.H3Element{
			Value: "bar",
		},
		doc.H1Element{
			Value: "foo",
		},
		doc.H5Element{
			Value: "foo",
		},
		doc.H3Element{
			Value: "foo",
		},
		doc.H3Element{
			Value: "foo ### b",
		},
		doc.H1Element{
			Value: "foo#",
		},
	}

	for i, line := range headings {
		elem, remainder, e := doc.TryATXHeading([]types.Str{
			types.Str(line),
		})

		assert.NoError(t, e)
		assert.Equal(t, []types.Str{}, remainder)
		assert.Equal(t, elements[i], elem)
	}
}

func TestNonATXHeading(t *testing.T) {
	headings := []string{
		"####### foo", // Example 63
		"#5 bolt",     // Example 64
		"#hashtag",    // Example 64
		"\\## foo",    // Example 65
	}

	for _, line := range headings {
		_, _, e := doc.TryATXHeading([]types.Str{
			types.Str(line),
		})

		assert.ErrorIs(t, err.ErrElementNotMatch, e)
	}
}
