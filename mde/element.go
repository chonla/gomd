package mde

import "gomd/adt"

type H3 struct {
	Value adt.String
}

type H4 struct {
	Value adt.String
}

type H5 struct {
	Value adt.String
}

type H6 struct {
	Value adt.String
}

type Text struct {
	Value adt.String
}

func TryH3(s adt.String) (bool, Block) {
	if s.FirstN(4) == "### " {
		return true, &H3{
			Value: s.From(4),
		}
	}
	return false, nil
}

func TryH4(s adt.String) (bool, Block) {
	if s.FirstN(5) == "#### " {
		return true, &H4{
			Value: s.From(5),
		}
	}
	return false, nil
}

func TryH5(s adt.String) (bool, Block) {
	if s.FirstN(6) == "##### " {
		return true, &H5{
			Value: s.From(6),
		}
	}
	return false, nil
}

func TryH6(s adt.String) (bool, Block) {
	if s.FirstN(7) == "###### " {
		return true, &H6{
			Value: s.From(7),
		}
	}
	return false, nil
}
