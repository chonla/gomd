package doc

import "gomd/types"

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
