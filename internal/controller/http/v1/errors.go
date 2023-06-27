package v1

import (
	"errors"
	"net/http"
)

var (
	errInvalidRequestBody        = errors.New("invalid request body")
	errInvalidUsernameOrPassword = errors.New("invalid username or password")
	errCannotParseToken          = errors.New("cannot parse token")
	errInvalidAuthHeader         = errors.New("invalid auth header")
	errInvalidUserId             = errors.New("invalid userId")
	errInvalidDepositValue       = errors.New("invalid deposit value. Value must be greater then 0")
	errNotEnoughMoney            = errors.New("not enough money")
	errReservationAlreadyExists  = errors.New("cannot create reservation. already exists")
	errReservationNotFound       = errors.New("reservation not found")
)

type ErrorOutput struct {
	Error string `json:"error" example:"example error"`
}

func newErrorRespond(w http.ResponseWriter, r *http.Request, code int, err error) {
	data := &ErrorOutput{Error: err.Error()}

	respond(w, r, code, data)
}
