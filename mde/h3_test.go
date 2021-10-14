package mde

import (
	"gomd/adt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryH3WithSpace(t *testing.T) {
	mdText := adt.String(`### foo`)
	expectedElement := &H3{
		Value: "foo",
	}

	actualResult, actualElement := TryH3(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH3WithTab(t *testing.T) {
	mdText := adt.String(`###	foo`)
	expectedElement := &H3{
		Value: "foo",
	}

	actualResult, actualElement := TryH3(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH3WithSpaceFollowedByEmptyText(t *testing.T) {
	mdText := adt.String(`### `)
	expectedElement := &H3{
		Value: "",
	}

	actualResult, actualElement := TryH3(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH3WithoutSpace(t *testing.T) {
	mdText := adt.String(`###5 bolt`)

	actualResult, actualElement := TryH3(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH3WithEscapedHash(t *testing.T) {
	mdText := adt.String(`\### foo`)

	actualResult, actualElement := TryH3(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH3WithoutSpaceButNewLine(t *testing.T) {
	mdText := adt.String(`###`)
	expectedElement := &H3{
		Value: "",
	}

	actualResult, actualElement := TryH3(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH3WithTrailingSpacesTitle(t *testing.T) {
	mdText := adt.String(`###      foo        `)
	expectedElement := &H3{
		Value: "foo",
	}

	actualResult, actualElement := TryH3(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH3WithTrailingTabsTitle(t *testing.T) {
	mdText := adt.String(`###      foo					`)
	expectedElement := &H3{
		Value: "foo",
	}

	actualResult, actualElement := TryH3(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH3With3Indentations(t *testing.T) {
	mdText := adt.String(`   ### foo`)
	expectedElement := &H3{
		Value: "foo",
	}

	actualResult, actualElement := TryH3(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH3With4Indentations(t *testing.T) {
	mdText := adt.String(`    ### foo`)

	actualResult, actualElement := TryH3(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH3WithClosingHashes(t *testing.T) {
	mdText := adt.String(`   ### foo ########`)
	expectedElement := &H3{
		Value: "foo",
	}

	actualResult, actualElement := TryH3(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH3WithClosingHashesAndSpaces(t *testing.T) {
	mdText := adt.String(`   ### foo ########     `)
	expectedElement := &H3{
		Value: "foo",
	}

	actualResult, actualElement := TryH3(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH3WithClosingHashesAndTabs(t *testing.T) {
	mdText := adt.String(`   ### foo ########					`)
	expectedElement := &H3{
		Value: "foo",
	}

	actualResult, actualElement := TryH3(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH3WithClosingHashesAndTabsAndOtherCharacters(t *testing.T) {
	mdText := adt.String(`   ### foo ######## b`)
	expectedElement := &H3{
		Value: "foo ######## b",
	}

	actualResult, actualElement := TryH3(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH3WithNoSpaceClosingHashes(t *testing.T) {
	mdText := adt.String(`   ### foo########`)
	expectedElement := &H3{
		Value: "foo########",
	}

	actualResult, actualElement := TryH3(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}
