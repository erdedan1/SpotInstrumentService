package errs

import "errors"

var (
	ErrNotFound        = errors.New("Not Found")
	ErrInvalidArgument = errors.New("Invalid Argument")
)
