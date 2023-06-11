package model

import "time"

type User struct {
	Id        int       `json:"id" validate:"-"`
	Username  string    `json:"username" validate:"email" example:"example@mail.org"`
	Password  string    `json:"password,omitempty" validate:"omitempty,min=6,max=30" example:"pass12345678"`
	CreatedAt time.Time `json:"-" validate:"-"`
}

func (u *User) Validate() error {
	return validate.Struct(u)
}
