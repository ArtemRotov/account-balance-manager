package repoerrors

import "errors"

var (
	ErrCannotStartTransaction = errors.New("entity already exists")
	ErrAlreadyExists          = errors.New("entity already exists")
	ErrNotFound               = errors.New("entity not found")
)
