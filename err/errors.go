package err

import "errors"

var (
	ErrElementNotMatch = errors.New("element not match")
	ErrEndOfFile       = errors.New("end of file")
)
