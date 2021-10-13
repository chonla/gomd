package mde

import "gomd/adt"

type H1 struct {
	Value adt.String
}

func TryH1(s adt.String) (bool, Block) {
	if text, ok := s.Grep("^ {0,3}#\\s+(.*?)(\\s+#*)?\\s*$", 0); ok {
		return true, &H1{
			Value: text.Trim(),
		}
	}
	if s.EqualsTo("#") {
		return true, &H1{
			Value: "",
		}
	}

	return false, nil
}
