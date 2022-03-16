package element

type Doc struct {
	Nodes []Element
}

func NewDoc() *Doc {
	return &Doc{
		Nodes: []Element{},
	}
}

// Push
// Append an element into doc node list.
func (d *Doc) Push(el Element) {
	d.Nodes = append(d.Nodes, el)
}

// Pop
// Remove an element from doc node list and return the removed element.
func (d *Doc) Pop() Element {
	var poppedElement Element
	nodeCount := len(d.Nodes)
	if nodeCount > 0 {
		poppedElement = d.Nodes[nodeCount-1]
		d.Nodes = d.Nodes[0 : nodeCount-1]
	}
	return poppedElement
}

// Last
// Return the reference of last element in node list.
func (d *Doc) Last() Element {
	var refElement Element
	nodeCount := len(d.Nodes)
	if nodeCount > 0 {
		refElement = d.Nodes[nodeCount-1]
	}
	return refElement

}
