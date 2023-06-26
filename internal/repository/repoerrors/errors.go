package repoerrors

import "errors"

var (
	ErrCannotStartTransaction = errors.New("tx_error - entity already exists")
	ErrAlreadyExists          = errors.New("entity already exists")
	ErrNotFound               = errors.New("entity not found")
	ErrInsufficientBalance    = errors.New("insufficient balance")
)
