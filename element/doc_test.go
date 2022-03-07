package element_test

import (
	"gomd/element"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopElementFromEmptyDoc(t *testing.T) {
	doc := element.NewDoc()

	result := doc.Pop()

	assert.Nil(t, result)
	assert.Equal(t, doc.Nodes, []element.Element{})
}

func TestPopElementFrom1ElementDoc(t *testing.T) {
	doc := element.NewDoc()
	doc.Push(element.H1{
		Value: "Test",
	})

	expected := element.H1{
		Value: "Test",
	}

	result := doc.Pop()

	assert.Equal(t, expected, result)
	assert.Equal(t, doc.Nodes, []element.Element{})
}

func TestPopElementFrom2ElementsDoc(t *testing.T) {
	doc := element.NewDoc()
	doc.Push(element.H1{
		Value: "Test",
	})
	doc.Push(element.H2{
		Value: "Test 2",
	})

	expected := element.H2{
		Value: "Test 2",
	}

	result := doc.Pop()

	assert.Equal(t, expected, result)
	assert.Equal(t, doc.Nodes, []element.Element{
		element.H1{
			Value: "Test",
		},
	})
}

func TestPopElementFrom3ElementsDoc(t *testing.T) {
	doc := element.NewDoc()
	doc.Push(element.H1{
		Value: "Test",
	})
	doc.Push(element.H2{
		Value: "Test 2",
	})
	doc.Push(element.H3{
		Value: "Test 3",
	})

	expected := element.H3{
		Value: "Test 3",
	}

	result := doc.Pop()

	assert.Equal(t, expected, result)
	assert.Equal(t, doc.Nodes, []element.Element{
		element.H1{
			Value: "Test",
		},
		element.H2{
			Value: "Test 2",
		},
	})
}
