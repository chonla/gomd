package mde

import "gomd/adt"

type H6 struct {
	Value adt.String
}

func TryH6(s adt.String) (bool, Block) {
	if text, ok := s.Grep("^ {0,3}######\\s+(.*?)(\\s+#*)?\\s*$", 0); ok {
		return true, &H6{
			Value: text.Trim(),
		}
	}
	if s.EqualsTo("######") {
		return true, &H6{
			Value: "",
		}
	}

	return false, nil
}
