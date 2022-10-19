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
		"***",    // Example 43
		"---",    // Example 43
		"___",    // Example 43
		" ***",   // Example 47
		"  ***",  // Example 47
		"   ***", // Example 47
		"   *** ",
		"   ****",
		"   ***\t",
		"   ---",
		"   ----",
		"   ___",
		"   ____",
		"_____________________________________", // Example 50
		" - - -",                                // Example 51
		" **  * ** * ** * **",                   // Example 52
		"-     -      -      -",                 // Example 53
		"- - - -    ",                           // Example 54
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
		"+++",       // Example 44
		"===",       // Example 45
		"--",        // Example 46
		"**",        // Example 46
		"__",        // Example 46
		"    ***",   // Example 48
		"_ _ _ _ a", // Example 55
		"a------",   // Example 55
		"---a---",   // Example 55
		" *-*",      // Example 56
	}

	for _, line := range notABreaks {
		_, _, e := doc.TryBreak([]types.Str{
			types.Str(line),
		})

		assert.ErrorIs(t, err.ErrElementNotMatch, e)
	}
}
