package model

import "time"

type Account struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"-"`
}
