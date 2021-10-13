package gomd

import (
	"gomd/adt"
	"gomd/mde"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(data []byte) (*mde.Document, error) {
	// Reference: https://spec.commonmark.org/0.30

	str := adt.String(string(data))
	children := []mde.Block{}

	for _, line := range str.Lines() {
		// Blank line
		if len(adt.String(line).Trim()) == 0 {
			continue
		}

		// >> Leaf blocks parsing

		// ATX Headings
		if ok, element := mde.TryH1(line); ok {
			children = append(children, element)
			continue
		}
		if ok, element := mde.TryH2(line); ok {
			children = append(children, element)
			continue
		}
		if ok, element := mde.TryH3(line); ok {
			children = append(children, element)
			continue
		}
		if ok, element := mde.TryH4(line); ok {
			children = append(children, element)
			continue
		}
		if ok, element := mde.TryH5(line); ok {
			children = append(children, element)
			continue
		}
		if ok, element := mde.TryH6(line); ok {
			children = append(children, element)
			continue
		}
		if ok, element := mde.TryBreak(line); ok {
			children = append(children, element)
			continue
		}

		children = append(children, &mde.Text{Value: adt.String(line).Unescape()})
	}

	return &mde.Document{
		Children: children,
	}, nil
}
