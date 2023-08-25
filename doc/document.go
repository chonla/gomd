package doc

import (
	"errors"
	"gomd/types"
)

type Document struct {
	Elements []types.AnyElement
}

func NewDocument() *Document {
	return &Document{
		Elements: []types.AnyElement{},
	}
}

func (d *Document) AppendElement(elem types.AnyElement) {
	d.Elements = append(d.Elements, elem)
}

func (d *Document) ReplaceLastElement(elem types.AnyElement) error {
	if len(d.Elements) == 0 {
		return errors.New("document is empty")
	}
	d.Elements[len(d.Elements)-1] = elem
	return nil
}

func (d *Document) LastElement() (types.AnyElement, error) {
	if len(d.Elements) == 0 {
		return nil, errors.New("document is empty")
	}
	return d.Elements[len(d.Elements)-1], nil
}
