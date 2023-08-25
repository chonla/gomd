package parser_test

import (
	"gomd/doc"
	"gomd/parser"
	"gomd/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationCommonmark1(t *testing.T) {
	markdown := types.Str(`	foo	baz		bim
`)
	expected := &doc.Document{
		Elements: []types.AnyElement{
			doc.CodeBlockElement{
				Value: `foo	baz		bim
`,
			},
		},
	}

	doc, err := parser.Parse(markdown)

	assert.NoError(t, err)
	assert.Equal(t, expected, doc)
}

func TestIntegrationCommonmark2(t *testing.T) {
	markdown :=
		types.Str(`  	foo	baz		bim
`)
	expected := &doc.Document{
		Elements: []types.AnyElement{
			doc.CodeBlockElement{
				Value: `foo	baz		bim
`,
			},
		},
	}

	doc, err := parser.Parse(markdown)

	assert.NoError(t, err)
	assert.Equal(t, expected, doc)
}

func TestIntegrationCommonmark3(t *testing.T) {
	markdown :=
		types.Str(`    a	a
	ὐ	a`)
	expected := &doc.Document{
		Elements: []types.AnyElement{
			doc.CodeBlockElement{
				Value: `a	a
ὐ	a`,
			},
		},
	}

	doc, err := parser.Parse(markdown)

	assert.NoError(t, err)
	assert.Equal(t, expected, doc)
}
