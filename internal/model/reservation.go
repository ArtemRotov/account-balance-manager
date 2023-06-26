package model

import "time"

type Reservation struct {
	Id        int       `json:"id" validate:"-"`
	AccountId int       `json:"account_id" validate:"gt=0" example:"12"`
	ServiceId int       `json:"service_id" validate:"gt=0" example:"134"`
	OrderId   int       `json:"order_id" validate:"gt=0" example:"11231"`
	Amount    int       `json:"amount" validate:"gt=-1" example:"12441"`
	CreatedAt time.Time `json:"-" validate:"-"`
}

func (r *Reservation) Validate() error {
	return validate.Struct(r)
}
