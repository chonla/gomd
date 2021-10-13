package mde

import "gomd/adt"

type H2 struct {
	Value adt.String
}

func TryH2(s adt.String) (bool, Block) {
	if text, ok := s.Grep("^ {0,3}##\\s+(.*?)(\\s+#*)?\\s*$", 0); ok {
		return true, &H2{
			Value: text.Trim(),
		}
	}
	if s.EqualsTo("##") {
		return true, &H2{
			Value: "",
		}
	}

	return false, nil
}
