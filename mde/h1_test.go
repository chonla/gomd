package mde

import (
	"gomd/adt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryH1WithSpace(t *testing.T) {
	mdText := adt.String(`# foo`)
	expectedElement := &H1{
		Value: "foo",
	}

	actualResult, actualElement := TryH1(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH1WithTab(t *testing.T) {
	mdText := adt.String(`#	foo`)
	expectedElement := &H1{
		Value: "foo",
	}

	actualResult, actualElement := TryH1(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH1WithSpaceFollowedByEmptyText(t *testing.T) {
	mdText := adt.String(`# `)
	expectedElement := &H1{
		Value: "",
	}

	actualResult, actualElement := TryH1(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH1WithoutSpace(t *testing.T) {
	mdText := adt.String(`#5 bolt`)

	actualResult, actualElement := TryH1(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH1WithEscapedHash(t *testing.T) {
	mdText := adt.String(`\# foo`)

	actualResult, actualElement := TryH1(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH1WithoutSpaceButNewLine(t *testing.T) {
	mdText := adt.String(`#`)
	expectedElement := &H1{
		Value: "",
	}

	actualResult, actualElement := TryH1(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH1WithTrailingSpacesTitle(t *testing.T) {
	mdText := adt.String(`#      foo        `)
	expectedElement := &H1{
		Value: "foo",
	}

	actualResult, actualElement := TryH1(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH1WithTrailingTabsTitle(t *testing.T) {
	mdText := adt.String(`#      foo					`)
	expectedElement := &H1{
		Value: "foo",
	}

	actualResult, actualElement := TryH1(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH1With3Indentations(t *testing.T) {
	mdText := adt.String(`   # foo`)
	expectedElement := &H1{
		Value: "foo",
	}

	actualResult, actualElement := TryH1(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH1With4Indentations(t *testing.T) {
	mdText := adt.String(`    # foo`)

	actualResult, actualElement := TryH1(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH1WithClosingHashes(t *testing.T) {
	mdText := adt.String(`   # foo ########`)
	expectedElement := &H1{
		Value: "foo",
	}

	actualResult, actualElement := TryH1(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH1WithClosingHashesAndSpaces(t *testing.T) {
	mdText := adt.String(`   # foo ########     `)
	expectedElement := &H1{
		Value: "foo",
	}

	actualResult, actualElement := TryH1(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH1WithClosingHashesAndTabs(t *testing.T) {
	mdText := adt.String(`   # foo ########					`)
	expectedElement := &H1{
		Value: "foo",
	}

	actualResult, actualElement := TryH1(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH1WithClosingHashesAndTabsAndOtherCharacters(t *testing.T) {
	mdText := adt.String(`   # foo ######## b`)
	expectedElement := &H1{
		Value: "foo ######## b",
	}

	actualResult, actualElement := TryH1(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH1WithNoSpaceClosingHashes(t *testing.T) {
	mdText := adt.String(`   # foo########`)
	expectedElement := &H1{
		Value: "foo########",
	}

	actualResult, actualElement := TryH1(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}
