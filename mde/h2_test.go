package mde

import (
	"gomd/adt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryH2WithSpace(t *testing.T) {
	mdText := adt.String(`## foo`)
	expectedElement := &H2{
		Value: "foo",
	}

	actualResult, actualElement := TryH2(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH2WithTab(t *testing.T) {
	mdText := adt.String(`##	foo`)
	expectedElement := &H2{
		Value: "foo",
	}

	actualResult, actualElement := TryH2(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH2WithSpaceFollowedByEmptyText(t *testing.T) {
	mdText := adt.String(`## `)
	expectedElement := &H2{
		Value: "",
	}

	actualResult, actualElement := TryH2(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH2WithoutSpace(t *testing.T) {
	mdText := adt.String(`##5 bolt`)

	actualResult, actualElement := TryH2(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH2WithEscapedHash(t *testing.T) {
	mdText := adt.String(`\## foo`)

	actualResult, actualElement := TryH2(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH2WithoutSpaceButNewLine(t *testing.T) {
	mdText := adt.String(`##`)
	expectedElement := &H2{
		Value: "",
	}

	actualResult, actualElement := TryH2(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH2WithTrailingSpacesTitle(t *testing.T) {
	mdText := adt.String(`##      foo        `)
	expectedElement := &H2{
		Value: "foo",
	}

	actualResult, actualElement := TryH2(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH2WithTrailingTabsTitle(t *testing.T) {
	mdText := adt.String(`##      foo					`)
	expectedElement := &H2{
		Value: "foo",
	}

	actualResult, actualElement := TryH2(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH2With3Indentations(t *testing.T) {
	mdText := adt.String(`   ## foo`)
	expectedElement := &H2{
		Value: "foo",
	}

	actualResult, actualElement := TryH2(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH2With4Indentations(t *testing.T) {
	mdText := adt.String(`    ## foo`)

	actualResult, actualElement := TryH2(mdText)

	assert.False(t, actualResult)
	assert.Nil(t, actualElement)
}

func TestTryH2WithClosingHashes(t *testing.T) {
	mdText := adt.String(`   ## foo ########`)
	expectedElement := &H2{
		Value: "foo",
	}

	actualResult, actualElement := TryH2(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH2WithClosingHashesAndSpaces(t *testing.T) {
	mdText := adt.String(`   ## foo ########     `)
	expectedElement := &H2{
		Value: "foo",
	}

	actualResult, actualElement := TryH2(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH2WithClosingHashesAndTabs(t *testing.T) {
	mdText := adt.String(`   ## foo ########					`)
	expectedElement := &H2{
		Value: "foo",
	}

	actualResult, actualElement := TryH2(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH2WithClosingHashesAndTabsAndOtherCharacters(t *testing.T) {
	mdText := adt.String(`   ## foo ######## b`)
	expectedElement := &H2{
		Value: "foo ######## b",
	}

	actualResult, actualElement := TryH2(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}

func TestTryH2WithNoSpaceClosingHashes(t *testing.T) {
	mdText := adt.String(`   ## foo########`)
	expectedElement := &H2{
		Value: "foo########",
	}

	actualResult, actualElement := TryH2(mdText)

	assert.True(t, actualResult)
	assert.Equal(t, expectedElement, actualElement)
}
