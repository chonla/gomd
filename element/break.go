package element

type BlankLine struct {
	Value string
}

func (BlankLine) Type() string {
	return "BlankLine"
}
