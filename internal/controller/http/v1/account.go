package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ArtemRotov/account-balance-manager/internal/service"
	"github.com/gorilla/mux"
)

type accountRoutes struct {
	accountService service.Account
}

func NewAccountRoutes(router *mux.Router, s service.Account) {
	r := &accountRoutes{
		accountService: s,
	}

	router.HandleFunc("/", r.balance()).Methods(http.MethodGet)
	router.HandleFunc("/deposit", r.deposit()).Methods(http.MethodPost)
}

type balanceInput struct {
	UserId int `json:"user_id" example:"1"`
}

type balanceOutput struct {
	UserId  int `json:"user_id" example:"1"`
	Balance int `json:"balance" example:"10000"`
}

// @Summary balance
// @Description User balance
// @Tags api/v1/account
// @Accept json
// @Produce json
// @Param balanceInput body balanceInput true "user_id"
// @Success 200 {object} balanceOutput
// @Failure 400 {object} ErrorOutput
// @Failure 500 {object} ErrorOutput
// @Security JWT
// @Router /api/v1/account/ [get]
func (router *accountRoutes) balance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := &balanceInput{}

		if err := json.NewDecoder(r.Body).Decode(input); err != nil {
			newErrorRespond(w, r, http.StatusBadRequest, errInvalidRequestBody)
			return
		}

		if input.UserId <= 0 {
			newErrorRespond(w, r, http.StatusBadRequest, errInvalidUserId)
			return
		}

		a, err := router.accountService.AccountByUserId(r.Context(), input.UserId)
		if err != nil {
			if errors.Is(err, service.ErrAccountNotFound) {
				newErrorRespond(w, r, http.StatusBadRequest, errInvalidUserId)
				return
			}
			newErrorRespond(w, r, http.StatusInternalServerError, err)
			return
		}

		respond(w, r, http.StatusOK, &balanceOutput{UserId: a.UserId, Balance: a.Balance})
	}
}

type depositInput struct {
	UserId int `json:"user_id" example:"1"`
	Amount int `json:"amount" example:"1000"`
}

type depositOutput struct {
	UserId  int `json:"user_id" example:"1"`
	Balance int `json:"balance" example:"1000"`
}

// @Summary deposit
// @Description deposit by userId
// @Tags api/v1/account/deposit
// @Accept json
// @Produce json
// @Param depositInput body depositInput true "user_id, amount"
// @Success 200 {object} depositOutput
// @Failure 400 {object} ErrorOutput
// @Failure 500 {object} ErrorOutput
// @Security JWT
// @Router /api/v1/account/deposit/ [post]
func (router *accountRoutes) deposit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := &depositInput{}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			newErrorRespond(w, r, http.StatusBadRequest, errInvalidRequestBody)
			return
		}

		if input.Amount <= 0 {
			newErrorRespond(w, r, http.StatusBadRequest, errInvalidDepositValue)
			return
		}

		a, err := router.accountService.DepositByUserId(r.Context(), input.UserId, input.Amount)
		if err != nil {
			if errors.Is(err, service.ErrAccountNotFound) {
				newErrorRespond(w, r, http.StatusBadRequest, errInvalidUserId)
				return
			}
			newErrorRespond(w, r, http.StatusInternalServerError, err)
			return
		}
		respond(w, r, http.StatusOK, &depositOutput{UserId: a.UserId, Balance: a.Balance})
	}
}
