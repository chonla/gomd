package parser

import (
	"errors"
	"gomd/doc"
	"gomd/err"
	"gomd/types"
)

func Parse(lines []types.Str) (*doc.Document, error) {
	factory := newFactory(lines)
	mddoc := doc.NewDocument()

	for element, e := factory.Unshift(); !errors.Is(e, err.ErrEndOfFile); element, e = factory.Unshift() {
		mddoc.AppendElement(element)
	}

	return mddoc, nil
}
