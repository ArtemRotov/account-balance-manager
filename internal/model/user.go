package model

import "time"

type User struct {
	Id        int64     `json:"id" validate:"-"`
	Username  string    `json:"username" validate:"email"`
	Password  string    `json:"password,omitempty" validate:"omitempty,min=6,max=30"`
	CreatedAt time.Time `json:"-" validate:"-"`
}

func (u *User) Validate() error {
	return validate.Struct(u)
}
