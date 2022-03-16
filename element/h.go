package element

// Header
type H1 struct {
	Value string
}

func (H1) Type() string {
	return "H1"
}

type H2 struct {
	Value string
}

func (H2) Type() string {
	return "H2"
}

type H3 struct {
	Value string
}

func (H3) Type() string {
	return "H3"
}

type H4 struct {
	Value string
}

func (H4) Type() string {
	return "H4"
}

type H5 struct {
	Value string
}

func (H5) Type() string {
	return "H5"
}

type H6 struct {
	Value string
}

func (H6) Type() string {
	return "H6"
}
