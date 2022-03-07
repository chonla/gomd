package element

type Doc struct {
	Nodes []Element
}

func NewDoc() *Doc {
	return &Doc{
		Nodes: []Element{},
	}
}

func (d *Doc) Push(el Element) {
	d.Nodes = append(d.Nodes, el)
}

func (d *Doc) Pop() Element {
	var poppedElement Element
	nodeCount := len(d.Nodes)
	if nodeCount > 0 {
		poppedElement = d.Nodes[nodeCount-1]
		d.Nodes = d.Nodes[0 : nodeCount-1]
	}
	return poppedElement
}
