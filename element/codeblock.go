package element

type CodeBlock struct {
	Value string
}

func (CodeBlock) Type() string {
	return "CodeBlock"
}
