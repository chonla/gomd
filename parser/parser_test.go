package parser_test

import (
	"gomd/element"
	"gomd/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingThematicBreaks(t *testing.T) {
	docs := []string{
		`***`,
		`---`,
		`___`,
		` ***`,                                  // one-leading space
		`  ***`,                                 // two-leading spaces
		`   ***`,                                // three-leading spaces
		`_____________________________________`, // More than 3 characters break
		` - - -`,                                // Inner spaces
		"-\t-\t-",                               // Inner tabs
		" **  * ** * ** * **",                   // Inner spaces
		"-       -       -      -",              // Inner spaces
		"- - - -     ",                          // Spaces at the end
		"- - - -\t\t\t\t",                       // Tabs at the end
	}
	expected := &element.Doc{
		Nodes: []element.Element{
			&element.Hr{},
		},
	}
	p := parser.New()

	for _, doc := range docs {
		result, err := p.Parse(doc)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	}
}

func TestParsingNonThematicBreaks(t *testing.T) {
	docs := []string{
		`+++`,
		`===`,
		`--`,
		`**`,
		`__`,
		// `    ---`,  // TODO: Code block
		// "Foo\n    ***", // TODO: More-than-3-spaces thematic break
	}
	expected := []*element.Doc{
		{
			Nodes: []element.Element{
				&element.P{
					Value: "+++",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.P{
					Value: "===",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.P{
					Value: "--",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.P{
					Value: "**",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.P{
					Value: "__",
				},
			},
		},
		// TODO: Codeblock expectation
		// {Codeblock}
		// TODO: More-than-3-spaces thematic break
		// {
		// 	Nodes: []element.Element{
		// 		&element.P{
		// 			Value: "Foo\n***",
		// 		},
		// 	},
		// },
	}
	p := parser.New()

	for index, doc := range docs {
		result, err := p.Parse(doc)

		assert.NoError(t, err)
		assert.Equal(t, expected[index], result)
	}
}

func TestSingleLineCodeBlock(t *testing.T) {
	docs := []string{
		"\tfoo\tbaz\t\tbim",
		"  \tfoo\tbaz\t\tbim",
		"    foo\tbaz\t\tbim",
	}
	expected := &element.Doc{
		Nodes: []element.Element{
			&element.CodeBlock{
				Value: "foo\tbaz\t\tbim",
			},
		},
	}

	p := parser.New()

	for _, doc := range docs {
		result, err := p.Parse(doc)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	}
}

func TestMultipleLinesCodeBlock(t *testing.T) {
	docs := []string{
		"    a simple\n      indented code block",
		"\ta simple\n      indented code block",
		"      a simple\n    indented code block",
		// "    chunk1\n\nchunk2\n  \n \n \n    chunk3",
	}
	expected := []*element.Doc{
		{
			Nodes: []element.Element{
				&element.CodeBlock{
					Value: "a simple\n  indented code block",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.CodeBlock{
					Value: "a simple\n  indented code block",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.CodeBlock{
					Value: "  a simple\nindented code block",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.CodeBlock{
					Value: "chunk1\n\nchunk2\n\n\n\nchunk3",
				},
			},
		},
	}

	p := parser.New()

	for index, doc := range docs {
		result, err := p.Parse(doc)

		assert.NoError(t, err)
		assert.Equal(t, expected[index], result)
	}
}

func TestATXHeadings(t *testing.T) {
	docs := []string{
		`# foo`,
		`## foo`,
		`### foo`,
		`#### foo`,
		`##### foo`,
		`###### foo`,
		`#             foo            `,
		` ### foo`,
		`  ## foo`,
		`   # foo`,
	}
	expected := []*element.Doc{
		{
			Nodes: []element.Element{
				&element.H1{
					Value: "foo",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.H2{
					Value: "foo",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.H3{
					Value: "foo",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.H4{
					Value: "foo",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.H5{
					Value: "foo",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.H6{
					Value: "foo",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.H1{
					Value: "foo",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.H3{
					Value: "foo",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.H2{
					Value: "foo",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.H1{
					Value: "foo",
				},
			},
		},
	}
	p := parser.New()

	for index, doc := range docs {
		result, err := p.Parse(doc)

		assert.NoError(t, err)
		assert.Equal(t, expected[index], result)
	}
}

func TestNonATXHeadings(t *testing.T) {
	docs := []string{
		`####### foo`,
		`#5 bolt`,
		`#hashtag`,
		`    # foo`,
	}
	expected := []*element.Doc{
		{
			Nodes: []element.Element{
				&element.P{
					Value: "####### foo",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.P{
					Value: "#5 bolt",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.P{
					Value: "#hashtag",
				},
			},
		},
		{
			Nodes: []element.Element{
				&element.CodeBlock{
					Value: "# foo",
				},
			},
		},
	}
	p := parser.New()

	for index, doc := range docs {
		result, err := p.Parse(doc)

		assert.NoError(t, err)
		assert.Equal(t, expected[index], result)
	}
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
			&element.H1{
				Value: "Header 1",
			},
			&element.H2{
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
			&element.H1{
				Value: "Header 1",
			},
			&element.H2{
				Value: "Header 2",
			},
			&element.H2{
				Value: "Another header 2",
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
			&element.P{
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
			&element.P{
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
			&element.P{
				Value: "Lorem Ipsum",
			},
			&element.P{
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
			&element.H1{
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
			&element.P{
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
			&element.H2{
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
			&element.P{
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
			&element.H1{
				Value: "Header",
			},
			&element.P{
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
			&element.Hr{},
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
			&element.Hr{},
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
			&element.Hr{},
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
			&element.Hr{},
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
			&element.Hr{},
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
			&element.Hr{},
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
			&element.Hr{},
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
			&element.Hr{},
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
			&element.Hr{},
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
			&element.P{
				Value: "Header",
			},
			&element.Hr{},
		},
	}

	p := parser.New()

	result, e := p.Parse(doc)

	assert.NoError(t, e)
	assert.Equal(t, expected, result)
}
