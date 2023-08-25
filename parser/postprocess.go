package parser

import (
	"gomd/doc"
	"gomd/types"

	"github.com/samber/lo"
)

func Postprocess(document *doc.Document) *doc.Document {
	// Merge multiple consecutive code blocks
	newDoc := doc.NewDocument()

	lo.ForEach(document.Elements, func(element types.AnyElement, i int) {
		// fmt.Println("------->")
		if lastElement, err := newDoc.LastElement(); err == nil {
			// fmt.Println("Last element: ", lastElement.TypeName())
			// fmt.Println("Arg element: ", element.TypeName())
			if lastElement.TypeName() == "code_block" && lastElement.SameType(element) {
				newElement := lastElement.(doc.CodeBlockElement).Merge(element.(doc.CodeBlockElement))
				// fmt.Println("Replace last element with", newElement.TypeName())
				newDoc.ReplaceLastElement(types.AnyElement(newElement))
			} else {
				if lastElement.TypeName() == "code_block" && element.TypeName() == "empty_line" {
					newElement := lastElement.(doc.CodeBlockElement).Merge(doc.CodeBlockElement{})
					// fmt.Println("Replace last element with", newElement.TypeName())
					newDoc.ReplaceLastElement(types.AnyElement(newElement))
				} else {
					// fmt.Println("Append", element.TypeName(), "to doc")
					newDoc.AppendElement(element)
				}
			}
		} else {
			// fmt.Println("Append", element.TypeName(), "to doc")
			newDoc.AppendElement(element)
		}
		// fmt.Println("<-------")
	})
	return newDoc
}
