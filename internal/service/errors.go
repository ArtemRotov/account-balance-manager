package service

import (
	"errors"
)

var (
	ErrUserAlreadyExists        = errors.New("user already exists")
	ErrCannotCreateUser         = errors.New("cannot create user")
	ErrUserNotFound             = errors.New("wrong username or password")
	ErrAccountNotFound          = errors.New("wrong userId")
	ErrCannotSignToken          = errors.New("cannot sign token")
	ErrCannotParseToken         = errors.New("cannot parse token")
	ErrReservationAlreadyExists = errors.New("reservation already exists")
	ErrReservationNotFound      = errors.New("reservation not found")
	ErrNotEnoughMoney           = errors.New("not enough money")
)
