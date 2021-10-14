package mde

import (
	"gomd/adt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryH4WithSpace(t *testing.T) {
	mdText := adt.String(`#### foo`)
	expectedElement := &H4{
		Value: "foo",
	}

	actualResult, actualElement := TryH4(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH4WithTab(t *testing.T) {
	mdText := adt.String(`####	foo`)
	expectedElement := &H4{
		Value: "foo",
	}

	actualResult, actualElement := TryH4(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH4WithSpaceFollowedByEmptyText(t *testing.T) {
	mdText := adt.String(`#### `)
	expectedElement := &H4{
		Value: "",
	}

	actualResult, actualElement := TryH4(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH4WithoutSpace(t *testing.T) {
	mdText := adt.String(`####5 bolt`)

	actualResult, actualElement := TryH4(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH4WithEscapedHash(t *testing.T) {
	mdText := adt.String(`\#### foo`)

	actualResult, actualElement := TryH4(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH4WithoutSpaceButNewLine(t *testing.T) {
	mdText := adt.String(`####`)
	expectedElement := &H4{
		Value: "",
	}

	actualResult, actualElement := TryH4(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH4WithTrailingSpacesTitle(t *testing.T) {
	mdText := adt.String(`####      foo        `)
	expectedElement := &H4{
		Value: "foo",
	}

	actualResult, actualElement := TryH4(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH4WithTrailingTabsTitle(t *testing.T) {
	mdText := adt.String(`####      foo					`)
	expectedElement := &H4{
		Value: "foo",
	}

	actualResult, actualElement := TryH4(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH4With3Indentations(t *testing.T) {
	mdText := adt.String(`   #### foo`)
	expectedElement := &H4{
		Value: "foo",
	}

	actualResult, actualElement := TryH4(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH4With4Indentations(t *testing.T) {
	mdText := adt.String(`    #### foo`)

	actualResult, actualElement := TryH4(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH4WithClosingHashes(t *testing.T) {
	mdText := adt.String(`   #### foo ########`)
	expectedElement := &H4{
		Value: "foo",
	}

	actualResult, actualElement := TryH4(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH4WithClosingHashesAndSpaces(t *testing.T) {
	mdText := adt.String(`   #### foo ########     `)
	expectedElement := &H4{
		Value: "foo",
	}

	actualResult, actualElement := TryH4(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH4WithClosingHashesAndTabs(t *testing.T) {
	mdText := adt.String(`   #### foo ########					`)
	expectedElement := &H4{
		Value: "foo",
	}

	actualResult, actualElement := TryH4(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH4WithClosingHashesAndTabsAndOtherCharacters(t *testing.T) {
	mdText := adt.String(`   #### foo ######## b`)
	expectedElement := &H4{
		Value: "foo ######## b",
	}

	actualResult, actualElement := TryH4(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH4WithNoSpaceClosingHashes(t *testing.T) {
	mdText := adt.String(`   #### foo########`)
	expectedElement := &H4{
		Value: "foo########",
	}

	actualResult, actualElement := TryH4(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}
