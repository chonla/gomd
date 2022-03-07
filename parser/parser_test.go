package parser_test

import (
	"gomd/element"
	"gomd/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDocumentContainingOnlyH1Element(t *testing.T) {
	doc := `# Header`
	expected := &element.Doc{
		Nodes: []element.Element{
			element.H1{
				Value: "Header",
			},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestParseDocumentContainingOnlyH2Element(t *testing.T) {
	doc := `## Header`
	expected := &element.Doc{
		Nodes: []element.Element{
			element.H2{
				Value: "Header",
			},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestParseDocumentContainingOnlyH3Element(t *testing.T) {
	doc := `### Header`
	expected := &element.Doc{
		Nodes: []element.Element{
			element.H3{
				Value: "Header",
			},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestParseDocumentContainingOnlyH4Element(t *testing.T) {
	doc := `#### Header`
	expected := &element.Doc{
		Nodes: []element.Element{
			element.H4{
				Value: "Header",
			},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestParseDocumentContainingOnlyH5Element(t *testing.T) {
	doc := `##### Header`
	expected := &element.Doc{
		Nodes: []element.Element{
			element.H5{
				Value: "Header",
			},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestParseDocumentContainingOnlyH6Element(t *testing.T) {
	doc := `###### Header`
	expected := &element.Doc{
		Nodes: []element.Element{
			element.H6{
				Value: "Header",
			},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestParseDocumentContainingOnlyBlankLine(t *testing.T) {
	doc := ``
	expected := &element.Doc{
		Nodes: []element.Element{},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestParseDocumentContainingBlankLine(t *testing.T) {
	doc := "# Header 1\n\n## Header 2"
	expected := &element.Doc{
		Nodes: []element.Element{
			element.H1{
				Value: "Header 1",
			},
			element.H2{
				Value: "Header 2",
			},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestParseDocumentContainingMultipleBlankLines(t *testing.T) {
	doc := "# Header 1\n\n\n\n## Header 2\n\n## Another header 2"
	expected := &element.Doc{
		Nodes: []element.Element{
			element.H1{
				Value: "Header 1",
			},
			element.H2{
				Value: "Header 2",
			},
			element.H2{
				Value: "Another header 2",
			},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestParseDocumentContainingH1ElementWithLeadingSpace(t *testing.T) {
	doc := `   # Header`
	expected := &element.Doc{
		Nodes: []element.Element{
			element.H1{
				Value: "Header",
			},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestParseSingleLineParagraph(t *testing.T) {
	doc := "Lorem Ipsum"
	expected := &element.Doc{
		Nodes: []element.Element{
			element.P{
				Value: "Lorem Ipsum",
			},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestParseMultipleLinesParagraph(t *testing.T) {
	doc := "Lorem Ipsum\ndolor sit amet"
	expected := &element.Doc{
		Nodes: []element.Element{
			element.P{
				Value: "Lorem Ipsum dolor sit amet",
			},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestParseMultipleParagraph(t *testing.T) {
	doc := "Lorem Ipsum\n\ndolor sit amet"
	expected := &element.Doc{
		Nodes: []element.Element{
			element.P{
				Value: "Lorem Ipsum",
			},
			element.P{
				Value: "dolor sit amet",
			},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestParseDocumentContainingSetextH1Element(t *testing.T) {
	doc := "Header\n="
	expected := &element.Doc{
		Nodes: []element.Element{
			element.H1{
				Value: "Header",
			},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestParseDocumentContainingJustEqualSign(t *testing.T) {
	doc := "="
	expected := &element.Doc{
		Nodes: []element.Element{
			element.P{
				Value: "=",
			},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestParseDocumentContainingSetextH2Element(t *testing.T) {
	doc := "Header\n-"
	expected := &element.Doc{
		Nodes: []element.Element{
			element.H2{
				Value: "Header",
			},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestParseDocumentContainingJustDashSign(t *testing.T) {
	doc := "-"
	expected := &element.Doc{
		Nodes: []element.Element{
			element.P{
				Value: "-",
			},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestParseDocumentContainingSetextH1MixWithDashSignElement(t *testing.T) {
	doc := "Header\n=\n-"
	expected := &element.Doc{
		Nodes: []element.Element{
			element.H1{
				Value: "Header",
			},
			element.P{
				Value: "-",
			},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestThematicBreakWith3Asterisks(t *testing.T) {
	doc := "***"
	expected := &element.Doc{
		Nodes: []element.Element{
			element.Hr{},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestThematicBreakWithLongerThan3Asterisks(t *testing.T) {
	doc := "****"
	expected := &element.Doc{
		Nodes: []element.Element{
			element.Hr{},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestThematicBreakWith3AsterisksAndInnerSpaces(t *testing.T) {
	doc := " *   *        *"
	expected := &element.Doc{
		Nodes: []element.Element{
			element.Hr{},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestThematicBreakWith3Underscores(t *testing.T) {
	doc := "___"
	expected := &element.Doc{
		Nodes: []element.Element{
			element.Hr{},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestThematicBreakWithLongerThan3Underscores(t *testing.T) {
	doc := "____"
	expected := &element.Doc{
		Nodes: []element.Element{
			element.Hr{},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestThematicBreakWith3UnderscoresAndInnerSpaces(t *testing.T) {
	doc := " _  _       _       "
	expected := &element.Doc{
		Nodes: []element.Element{
			element.Hr{},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestThematicBreakWith3Dashes(t *testing.T) {
	doc := "---"
	expected := &element.Doc{
		Nodes: []element.Element{
			element.Hr{},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestThematicBreakWithLongerThan3Dashes(t *testing.T) {
	doc := "----"
	expected := &element.Doc{
		Nodes: []element.Element{
			element.Hr{},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestThematicBreakWith3DashesAndInnerSpaces(t *testing.T) {
	doc := " -  -      -    "
	expected := &element.Doc{
		Nodes: []element.Element{
			element.Hr{},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}

func TestThematicBreakWithCarriedOverText(t *testing.T) {
	doc := "Header\n-  -      -"
	expected := &element.Doc{
		Nodes: []element.Element{
			element.P{
				Value: "Header",
			},
			element.Hr{},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}
