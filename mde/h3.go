package mde

import "gomd/adt"

type H3 struct {
	Value adt.String
}

func TryH3(s adt.String) (bool, Block) {
	if text, ok := s.Grep("^ {0,3}###\\s+(.*?)(\\s+#*)?\\s*$", 0); ok {
		return true, &H3{
			Value: text.Trim(),
		}
	}
	if s.EqualsTo("###") {
		return true, &H3{
			Value: "",
		}
	}

	return false, nil
}
