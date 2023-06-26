package model

import "time"

type Reservation struct {
	Id        int       `json:"id"`
	AccountId int       `json:"account_id"`
	ServiceId int       `json:"service_id"`
	OrderId   int       `json:"order_id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"-"`
}
