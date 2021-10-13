package mde

import (
	"gomd/adt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryBreakWith3Asterisks(t *testing.T) {
	mdText := adt.String(`***`)
	expectedElement := &Break{}

	actualResult, actualElement := TryBreak(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryBreakWithMoreThan3Asterisks(t *testing.T) {
	mdText := adt.String(`****`)
	expectedElement := &Break{}

	actualResult, actualElement := TryBreak(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)

}

func TestTryBreakWith3Dashes(t *testing.T) {
	mdText := adt.String(`---`)
	expectedElement := &Break{}

	actualResult, actualElement := TryBreak(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryBreakWithMoreThan3Dashes(t *testing.T) {
	mdText := adt.String(`----`)
	expectedElement := &Break{}

	actualResult, actualElement := TryBreak(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryBreakWith3Lodashes(t *testing.T) {
	mdText := adt.String(`___`)
	expectedElement := &Break{}

	actualResult, actualElement := TryBreak(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryBreakWithMoreThan3Lodashes(t *testing.T) {
	mdText := adt.String(`____`)
	expectedElement := &Break{}

	actualResult, actualElement := TryBreak(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryBreakWithWrongCharacters(t *testing.T) {
	mdText := adt.String(`+++`)

	actualResult, actualElement := TryBreak(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryBreakWithNotEnoughAsterisks(t *testing.T) {
	mdText := adt.String(`**`)

	actualResult, actualElement := TryBreak(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryBreakWithNotEnoughDashes(t *testing.T) {
	mdText := adt.String(`--`)

	actualResult, actualElement := TryBreak(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryBreakWithNotEnoughLodashes(t *testing.T) {
	mdText := adt.String(`__`)

	actualResult, actualElement := TryBreak(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryBreakWith3AsterisksWithLessThan4Indent(t *testing.T) {
	mdText := adt.String(`   ***`)
	expectedElement := &Break{}

	actualResult, actualElement := TryBreak(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryBreakWithMoreThan3AsterisksWithLessThan4Indent(t *testing.T) {
	mdText := adt.String(`   ****`)
	expectedElement := &Break{}

	actualResult, actualElement := TryBreak(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)

}

func TestTryBreakWith3DashesWithLessThan4Indent(t *testing.T) {
	mdText := adt.String(`   ---`)
	expectedElement := &Break{}

	actualResult, actualElement := TryBreak(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryBreakWithMoreThan3DashesWithLessThan4Indent(t *testing.T) {
	mdText := adt.String(`   ----`)
	expectedElement := &Break{}

	actualResult, actualElement := TryBreak(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryBreakWith3LodashesWithLessThan4Indent(t *testing.T) {
	mdText := adt.String(`   ___`)
	expectedElement := &Break{}

	actualResult, actualElement := TryBreak(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryBreakWithMoreThan3LodashesWithLessThan4Indent(t *testing.T) {
	mdText := adt.String(`   ____`)
	expectedElement := &Break{}

	actualResult, actualElement := TryBreak(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryBreakWithSpaceAndTabBetweenAsterisks(t *testing.T) {
	mdText := adt.String(` **  * ** * ** * **`)
	expectedElement := &Break{}

	actualResult, actualElement := TryBreak(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryBreakWithSpaceAndTabBetweenDashes(t *testing.T) {
	mdText := adt.String(`-     -      -      -`)
	expectedElement := &Break{}

	actualResult, actualElement := TryBreak(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryBreakWithSpaceAndTabBetweenLodashes(t *testing.T) {
	mdText := adt.String(`_     _      _      _`)
	expectedElement := &Break{}

	actualResult, actualElement := TryBreak(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryBreakWithOtherCharactersBlended(t *testing.T) {
	mdText := adt.String(`_ _ _ _ a`)

	actualResult, actualElement := TryBreak(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryBreakWithMixBreakingCharacters(t *testing.T) {
	mdText := adt.String(` *-*`)

	actualResult, actualElement := TryBreak(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}
