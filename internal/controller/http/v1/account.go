package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ArtemRotov/account-balance-manager/internal/service"
	"github.com/gorilla/mux"
)

type accountRoutes struct {
	service service.Account
}

func NewAccountRoutes(router *mux.Router, s service.Account) {
	r := &accountRoutes{
		service: s,
	}

	router.HandleFunc("/", r.Balance()).Methods(http.MethodGet)
}

type BalanceInput struct {
	UserId int `json:"user_id"`
}

type BalanceOutput struct {
	UserId  int `json:"user_id"`
	Balance int `json:"balance"`
}

func (router *accountRoutes) Balance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := &BalanceInput{}

		if err := json.NewDecoder(r.Body).Decode(input); err != nil {
			newErrorRespond(w, r, http.StatusBadRequest, errInvalidRequestBody)
			return
		}

		if input.UserId <= 0 {
			newErrorRespond(w, r, http.StatusBadRequest, errInvalidUserId)
			return
		}

		a, err := router.service.AccountByUserId(r.Context(), input.UserId)
		if err != nil {
			if errors.Is(err, service.ErrAccountNotFound) {
				newErrorRespond(w, r, http.StatusBadRequest, errInvalidUserId)
				return
			}
			newErrorRespond(w, r, http.StatusInternalServerError, err)
			return
		}

		respond(w, r, http.StatusOK, &BalanceOutput{UserId: a.UserId, Balance: a.Balance})
	}
}
