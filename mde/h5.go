package mde

import "gomd/adt"

type H5 struct {
	Value adt.String
}

func TryH5(s adt.String) (bool, Block) {
	if text, ok := s.Grep("^ {0,3}#####\\s+(.*?)(\\s+#*)?\\s*$", 0); ok {
		return true, &H5{
			Value: text.Trim(),
		}
	}
	if s.EqualsTo("#####") {
		return true, &H5{
			Value: "",
		}
	}

	return false, nil
}
