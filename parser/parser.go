package parser

import (
	"errors"
	"gomd/doc"
	"gomd/err"
	"gomd/types"
)

func Parse(data types.Str) (*doc.Document, error) {
	processed := Preprocess(data)
	factory := newFactory(processed.ToLines())
	mddoc := doc.NewDocument()

	for element, e := factory.Unshift(); !errors.Is(e, err.ErrEndOfFile); element, e = factory.Unshift() {
		mddoc.AppendElement(element)
	}

	cleanDoc := Postprocess(mddoc)

	return cleanDoc, nil
}
