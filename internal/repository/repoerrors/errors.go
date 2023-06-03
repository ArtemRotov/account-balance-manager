package repoerrors

import "errors"

var (
	ErrAlreadyExists = errors.New("an entity with these parameters already exists")
)
