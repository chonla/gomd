package doc_test

import (
	"gomd/doc"
	"gomd/err"
	"gomd/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBreak(t *testing.T) {
	breaks := []string{
		"***",
		"---",
		"___",
		" ***",
		"  ***",
		"   ***",
		"   *** ",
		"   ****",
		"   ***\t",
		"   ---",
		"   ----",
		"   ___",
		"   ____",
		"_____________________________________",
		" - - -",
		" **  * ** * ** * **",
		"-     -      -      -",
		"- - - -    ",
	}

	for _, line := range breaks {
		elem, remainder, e := doc.TryBreak([]types.Str{
			types.Str(line),
		})

		assert.NoError(t, e)
		assert.Equal(t, []types.Str{}, remainder)
		assert.Equal(t, doc.BreakElement{}, elem)
	}
}

func TestNotABreak(t *testing.T) {
	notABreaks := []string{
		"+++",
		"===",
		"--",
		"**",
		"__",
		"    ***",
		"_ _ _ _ a",
		"a------",
		"---a---",
		" *-*",
	}

	for _, line := range notABreaks {
		_, _, e := doc.TryBreak([]types.Str{
			types.Str(line),
		})

		assert.ErrorIs(t, err.ErrElementNotMatch, e)
	}
}
