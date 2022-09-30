package parser

import (
	"gomd/doc"
	"gomd/err"
	"gomd/types"
)

var elementTesters = []types.ElementTesterFunc{
	doc.TryH1,
	doc.TryH2,
	doc.TryH3,
	doc.TryParagraph,
	doc.TryEmpty,
}

type Factory struct {
	Lines []types.Str
}

func newFactory(lines []types.Str) *Factory {
	return &Factory{
		Lines: lines,
	}
}

func (f *Factory) Load(lines []types.Str) {
	f.Lines = lines
}

func (f *Factory) Unshift() (types.AnyElement, error) {
	if len(f.Lines) == 0 {
		return nil, err.ErrEndOfFile
	}

	for _, elementTest := range elementTesters {
		elem, remainder, e := elementTest(f.Lines)
		if e == nil {
			f.Lines = remainder
			return elem, nil
		}
	}

	return doc.EmptyElement{}, nil
}
