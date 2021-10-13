package mde

import "gomd/adt"

type Break struct{}

func TryBreak(s adt.String) (bool, Block) {
	if s.Like("^ {0,3}(\\*\\s*){3,}\\s*$") {
		return true, &Break{}
	}
	if s.Like("^ {0,3}(-\\s*){3,}\\s*$") {
		return true, &Break{}
	}
	if s.Like("^ {0,3}(_\\s*){3,}\\s*$") {
		return true, &Break{}
	}
	return false, nil
}
