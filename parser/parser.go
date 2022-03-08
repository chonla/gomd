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
	textCarriedOver := str.Str("")

	for _, line := range lines {
		if multilineProcess {

		} else {
			l := str.Str(line)

			if l.IsBlank() {
				if !textCarriedOver.IsEmpty() {
					mdDoc.Push(element.P{
						Value: textCarriedOver.String(),
					})
					textCarriedOver = textCarriedOver.Empty()
				}
				continue
			}

			// Code Block
			if l.IsLike(`^ {0,3}\t`) || l.IsLike(`^ {4,}`) {
				mdDoc.Push(element.CodeBlock{
					Value: l.Trim().String(),
				})
				continue
			}

			// Setext H1
			if l.IsLike(`^=+$`) {
				if !textCarriedOver.IsEmpty() {
					mdDoc.Push(element.H1{
						Value: textCarriedOver.String(),
					})
					textCarriedOver = textCarriedOver.Empty()
				} else {
					mdDoc.Push(element.P{
						Value: l.String(),
					})
				}
				continue
			}

			// Setext H2
			if l.IsLike(`^-+$`) {
				if !textCarriedOver.IsEmpty() {
					mdDoc.Push(element.H2{
						Value: textCarriedOver.String(),
					})
					textCarriedOver = textCarriedOver.Empty()
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
				if !textCarriedOver.IsEmpty() {
					mdDoc.Push(element.P{
						Value: textCarriedOver.String(),
					})
					textCarriedOver = textCarriedOver.Empty()
				}
				mdDoc.Push(element.Hr{})
				continue
			}

			if l.IsLike(`^ {0,3}# `) {
				mdDoc.Push(element.H1{
					Value: l.Capture(`^ {0,3}# (.*)$`).Trim().String(),
				})
				continue
			}
			if l.IsLike(`^ {0,3}## `) {
				mdDoc.Push(element.H2{
					Value: l.Capture(`^ {0,3}## (.*)$`).Trim().String(),
				})
				continue
			}
			if l.IsLike(`^ {0,3}### `) {
				mdDoc.Push(element.H3{
					Value: l.Capture(`^ {0,3}### (.*)$`).Trim().String(),
				})
				continue
			}
			if l.IsLike(`^ {0,3}#### `) {
				mdDoc.Push(element.H4{
					Value: l.Capture(`^ {0,3}#### (.*)$`).Trim().String(),
				})
				continue
			}
			if l.IsLike(`^ {0,3}##### `) {
				mdDoc.Push(element.H5{
					Value: l.Capture(`^ {0,3}##### (.*)$`).Trim().String(),
				})
				continue
			}
			if l.IsLike(`^ {0,3}###### `) {
				mdDoc.Push(element.H6{
					Value: l.Capture(`^ {0,3}###### (.*)$`).Trim().String(),
				})
				continue
			}
			if !textCarriedOver.IsEmpty() {
				textCarriedOver = textCarriedOver.Append(str.Str(" "))
			}
			textCarriedOver = textCarriedOver.Append(line)
		}
	}
	if !textCarriedOver.IsEmpty() {
		mdDoc.Push(element.P{
			Value: textCarriedOver.String(),
		})
	}

	return mdDoc, nil
}
