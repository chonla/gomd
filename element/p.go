package element

// Paragraph
type P struct {
	Value string
}

func (P) Type() string {
	return "P"
}
