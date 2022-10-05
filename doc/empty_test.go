package doc_test

import (
	"gomd/doc"
	"gomd/err"
	"gomd/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	line := ""

	elem, remainder, e := doc.TryEmpty([]types.Str{
		types.Str(line),
	})

	assert.NoError(t, e)
	assert.Equal(t, []types.Str{}, remainder)
	assert.Equal(t, doc.EmptyElement{}, elem)
}

func TestNotEmpty(t *testing.T) {
	line := " "

	_, _, e := doc.TryEmpty([]types.Str{
		types.Str(line),
	})

	assert.ErrorIs(t, err.ErrElementNotMatch, e)
}
