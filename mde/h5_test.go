package mde

import (
	"gomd/adt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryH5WithSpace(t *testing.T) {
	mdText := adt.String(`##### foo`)
	expectedElement := &H5{
		Value: "foo",
	}

	actualResult, actualElement := TryH5(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH5WithTab(t *testing.T) {
	mdText := adt.String(`#####	foo`)
	expectedElement := &H5{
		Value: "foo",
	}

	actualResult, actualElement := TryH5(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH5WithSpaceFollowedByEmptyText(t *testing.T) {
	mdText := adt.String(`##### `)
	expectedElement := &H5{
		Value: "",
	}

	actualResult, actualElement := TryH5(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH5WithoutSpace(t *testing.T) {
	mdText := adt.String(`#####5 bolt`)

	actualResult, actualElement := TryH5(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH5WithEscapedHash(t *testing.T) {
	mdText := adt.String(`\##### foo`)

	actualResult, actualElement := TryH5(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH5WithoutSpaceButNewLine(t *testing.T) {
	mdText := adt.String(`#####`)
	expectedElement := &H5{
		Value: "",
	}

	actualResult, actualElement := TryH5(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH5WithTrailingSpacesTitle(t *testing.T) {
	mdText := adt.String(`#####      foo        `)
	expectedElement := &H5{
		Value: "foo",
	}

	actualResult, actualElement := TryH5(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH5WithTrailingTabsTitle(t *testing.T) {
	mdText := adt.String(`#####      foo					`)
	expectedElement := &H5{
		Value: "foo",
	}

	actualResult, actualElement := TryH5(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH5With3Indentations(t *testing.T) {
	mdText := adt.String(`   ##### foo`)
	expectedElement := &H5{
		Value: "foo",
	}

	actualResult, actualElement := TryH5(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH5With4Indentations(t *testing.T) {
	mdText := adt.String(`    ##### foo`)

	actualResult, actualElement := TryH5(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH5WithClosingHashes(t *testing.T) {
	mdText := adt.String(`   ##### foo ########`)
	expectedElement := &H5{
		Value: "foo",
	}

	actualResult, actualElement := TryH5(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH5WithClosingHashesAndSpaces(t *testing.T) {
	mdText := adt.String(`   ##### foo ########     `)
	expectedElement := &H5{
		Value: "foo",
	}

	actualResult, actualElement := TryH5(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH5WithClosingHashesAndTabs(t *testing.T) {
	mdText := adt.String(`   ##### foo ########					`)
	expectedElement := &H5{
		Value: "foo",
	}

	actualResult, actualElement := TryH5(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH5WithClosingHashesAndTabsAndOtherCharacters(t *testing.T) {
	mdText := adt.String(`   ##### foo ######## b`)
	expectedElement := &H5{
		Value: "foo ######## b",
	}

	actualResult, actualElement := TryH5(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH5WithNoSpaceClosingHashes(t *testing.T) {
	mdText := adt.String(`   ##### foo########`)
	expectedElement := &H5{
		Value: "foo########",
	}

	actualResult, actualElement := TryH5(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}
