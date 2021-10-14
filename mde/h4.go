package mde

import "gomd/adt"

type H4 struct {
	Value adt.String
}

func TryH4(s adt.String) (bool, Block) {
	if text, ok := s.Grep("^ {0,3}####\\s+(.*?)(\\s+#*)?\\s*$", 0); ok {
		return true, &H4{
			Value: text.Trim(),
		}
	}
	if s.EqualsTo("####") {
		return true, &H4{
			Value: "",
		}
	}

	return false, nil
}
