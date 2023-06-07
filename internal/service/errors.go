package service

import (
	"errors"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrCannotCreateUser  = errors.New("cannot create user")
	ErrUserNotFound      = errors.New("wrong username or password")
	ErrCannotSignToken   = errors.New("cannot sign token")
	ErrCannotParseToken  = errors.New("cannot parse token")
)
