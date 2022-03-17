package parser

import (
	"fmt"
	"gomd/element"
	"gomd/helper/str"
)

const (
	TabSize             = 4
	CodeBlockIndentSize = 4
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
					mdDoc.Push(&element.P{
						Value: textCarriedOver.String(),
					})
					textCarriedOver = textCarriedOver.Empty()
				}
				mdDoc.Push(&element.BlankLine{})
				continue
			}

			// Indented Code Block
			if l.IsLike(`^ {0,3}\t`) || l.IsLike(`^ {4}`) || l.IsLike(`^\t`) {
				l1 := l.ReplaceRegex(`^( {0,3}\t| {4}|\t)`, "    ")
				if mdDoc.Last() != nil && mdDoc.Last().Type() == "CodeBlock" {
					// merge to previous code block
					codeBlock := mdDoc.Pop().(*element.CodeBlock)
					codeBlock.Value = fmt.Sprintf("%s\n%s", codeBlock.Value, l1.From(4).String())

					mdDoc.Push(codeBlock)
				} else {
					mdDoc.Push(&element.CodeBlock{
						Value: l1.From(4).String(),
					})
				}
				continue
			}

			// Setext H1
			if l.IsLike(`^=+$`) {
				if !textCarriedOver.IsEmpty() {
					mdDoc.Push(&element.H1{
						Value: textCarriedOver.String(),
					})
					textCarriedOver = textCarriedOver.Empty()
				} else {
					mdDoc.Push(&element.P{
						Value: l.String(),
					})
				}
				continue
			}

			// Setext H2
			if l.IsLike(`^-+$`) {
				if !textCarriedOver.IsEmpty() {
					mdDoc.Push(&element.H2{
						Value: textCarriedOver.String(),
					})
					textCarriedOver = textCarriedOver.Empty()
				} else {
					if l.IsLike(`^-{3,}$`) {
						mdDoc.Push(&element.Hr{})
						continue
					} else {
						mdDoc.Push(&element.P{
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
					mdDoc.Push(&element.P{
						Value: textCarriedOver.String(),
					})
					textCarriedOver = textCarriedOver.Empty()
				}
				mdDoc.Push(&element.Hr{})
				continue
			}

			if l.IsLike(`^ {0,3}#{1} `) {
				mdDoc.Push(&element.H1{
					Value: l.Capture(`^ {0,3}#{1} (.*)$`).Trim().String(),
				})
				continue
			}
			if l.IsLike(`^ {0,3}#{2} `) {
				mdDoc.Push(&element.H2{
					Value: l.Capture(`^ {0,3}#{2} (.*)$`).Trim().String(),
				})
				continue
			}
			if l.IsLike(`^ {0,3}#{3} `) {
				mdDoc.Push(&element.H3{
					Value: l.Capture(`^ {0,3}#{3} (.*)$`).Trim().String(),
				})
				continue
			}
			if l.IsLike(`^ {0,3}#{4} `) {
				mdDoc.Push(&element.H4{
					Value: l.Capture(`^ {0,3}#{4} (.*)$`).Trim().String(),
				})
				continue
			}
			if l.IsLike(`^ {0,3}#{5} `) {
				mdDoc.Push(&element.H5{
					Value: l.Capture(`^ {0,3}#{5} (.*)$`).Trim().String(),
				})
				continue
			}
			if l.IsLike(`^ {0,3}#{6} `) {
				mdDoc.Push(&element.H6{
					Value: l.Capture(`^ {0,3}#{6} (.*)$`).Trim().String(),
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
		mdDoc.Push(&element.P{
			Value: textCarriedOver.String(),
		})
	}

	return mdDoc, nil
}
