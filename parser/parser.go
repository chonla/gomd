package parser

import (
	"gomd/element"
	"gomd/helper/str"
)

type Parser struct{}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(doc string) (*element.Doc, error) {
	lines := str.Str(doc).Lines()
	multilineProcess := false
	mdDoc := element.NewDoc()
	textCarriedOver := ""

	for _, line := range lines {
		if multilineProcess {

		} else {
			l := str.Str(line)

			if l.IsBlank() {
				if textCarriedOver != "" {
					mdDoc.Push(element.P{
						Value: textCarriedOver,
					})
					textCarriedOver = ""
				}
				continue
			}

			// Code Block
			if l.IsLike(`^ {0,3}\t`) || l.IsLike(`^ {4,}`) {
				mdDoc.Push(element.CodeBlock{
					Value: l.Trim(),
				})
				continue
			}

			// Setext H1
			if l.IsLike(`^=+$`) {
				if textCarriedOver != "" {
					mdDoc.Push(element.H1{
						Value: textCarriedOver,
					})
					textCarriedOver = ""
				} else {
					mdDoc.Push(element.P{
						Value: l.String(),
					})
				}
				continue
			}

			// Setext H2
			if l.IsLike(`^-+$`) {
				if textCarriedOver != "" {
					mdDoc.Push(element.H2{
						Value: textCarriedOver,
					})
					textCarriedOver = ""
				} else {
					if l.IsLike(`^-{3,}$`) {
						mdDoc.Push(element.Hr{})
						continue
					} else {
						mdDoc.Push(element.P{
							Value: l.String(),
						})
					}
				}
				continue
			}

			// Thematic Breaks
			hr := str.Str(l.Without(" ", "\t"))
			if hr.IsLike(`^\*{3,}$`) || hr.IsLike(`^_{3,}$`) || hr.IsLike(`^-{3,}$`) {
				if textCarriedOver != "" {
					mdDoc.Push(element.P{
						Value: textCarriedOver,
					})
					textCarriedOver = ""
				}
				mdDoc.Push(element.Hr{})
				continue
			}

			if l.IsLike(`^ {0,3}# `) {
				mdDoc.Push(element.H1{
					Value: l.Capture(`^ {0,3}# (.*)`),
				})
				continue
			}
			if l.IsLike(`^ {0,3}## `) {
				mdDoc.Push(element.H2{
					Value: l.Capture(`^ {0,3}## (.*)`),
				})
				continue
			}
			if l.IsLike(`^ {0,3}### `) {
				mdDoc.Push(element.H3{
					Value: l.Capture(`^ {0,3}### (.*)`),
				})
				continue
			}
			if l.IsLike(`^ {0,3}#### `) {
				mdDoc.Push(element.H4{
					Value: l.Capture(`^ {0,3}#### (.*)`),
				})
				continue
			}
			if l.IsLike(`^ {0,3}##### `) {
				mdDoc.Push(element.H5{
					Value: l.Capture(`^ {0,3}##### (.*)`),
				})
				continue
			}
			if l.IsLike(`^ {0,3}###### `) {
				mdDoc.Push(element.H6{
					Value: l.Capture(`^ {0,3}###### (.*)`),
				})
				continue
			}
			if textCarriedOver != "" {
				textCarriedOver += " "
			}
			textCarriedOver += line
		}
	}
	if textCarriedOver != "" {
		mdDoc.Push(element.P{
			Value: textCarriedOver,
		})
	}

	return mdDoc, nil
}
