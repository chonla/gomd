package mde

import (
	"gomd/adt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryH6WithSpace(t *testing.T) {
	mdText := adt.String(`###### foo`)
	expectedElement := &H6{
		Value: "foo",
	}

	actualResult, actualElement := TryH6(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH6WithTab(t *testing.T) {
	mdText := adt.String(`######	foo`)
	expectedElement := &H6{
		Value: "foo",
	}

	actualResult, actualElement := TryH6(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH6WithSpaceFollowedByEmptyText(t *testing.T) {
	mdText := adt.String(`###### `)
	expectedElement := &H6{
		Value: "",
	}

	actualResult, actualElement := TryH6(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH6WithoutSpace(t *testing.T) {
	mdText := adt.String(`######5 bolt`)

	actualResult, actualElement := TryH6(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH6WithEscapedHash(t *testing.T) {
	mdText := adt.String(`\###### foo`)

	actualResult, actualElement := TryH6(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH6WithoutSpaceButNewLine(t *testing.T) {
	mdText := adt.String(`######`)
	expectedElement := &H6{
		Value: "",
	}

	actualResult, actualElement := TryH6(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH6WithTrailingSpacesTitle(t *testing.T) {
	mdText := adt.String(`######      foo        `)
	expectedElement := &H6{
		Value: "foo",
	}

	actualResult, actualElement := TryH6(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH6WithTrailingTabsTitle(t *testing.T) {
	mdText := adt.String(`######      foo					`)
	expectedElement := &H6{
		Value: "foo",
	}

	actualResult, actualElement := TryH6(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH6With3Indentations(t *testing.T) {
	mdText := adt.String(`   ###### foo`)
	expectedElement := &H6{
		Value: "foo",
	}

	actualResult, actualElement := TryH6(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH6With4Indentations(t *testing.T) {
	mdText := adt.String(`    ###### foo`)

	actualResult, actualElement := TryH6(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH6WithClosingHashes(t *testing.T) {
	mdText := adt.String(`   ###### foo ########`)
	expectedElement := &H6{
		Value: "foo",
	}

	actualResult, actualElement := TryH6(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH6WithClosingHashesAndSpaces(t *testing.T) {
	mdText := adt.String(`   ###### foo ########     `)
	expectedElement := &H6{
		Value: "foo",
	}

	actualResult, actualElement := TryH6(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH6WithClosingHashesAndTabs(t *testing.T) {
	mdText := adt.String(`   ###### foo ########					`)
	expectedElement := &H6{
		Value: "foo",
	}

	actualResult, actualElement := TryH6(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH6WithClosingHashesAndTabsAndOtherCharacters(t *testing.T) {
	mdText := adt.String(`   ###### foo ######## b`)
	expectedElement := &H6{
		Value: "foo ######## b",
	}

	actualResult, actualElement := TryH6(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH6WithNoSpaceClosingHashes(t *testing.T) {
	mdText := adt.String(`   ###### foo########`)
	expectedElement := &H6{
		Value: "foo########",
	}

	actualResult, actualElement := TryH6(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}
