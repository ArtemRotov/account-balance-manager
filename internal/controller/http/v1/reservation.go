package v1

import (
	"encoding/json"
	"errors"
	"golang.org/x/net/context"
	"net/http"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
	"github.com/ArtemRotov/account-balance-manager/internal/service"
	"github.com/gorilla/mux"
)

type ReservationService interface {
	CreateReservation(ctx context.Context, account_id, service_id, order_id, amount int) (*model.Reservation, error)
	Revenue(ctx context.Context, account_id, service_id, order_id, amount int) error
	Refund(ctx context.Context, account_id, service_id, order_id, amount int) error
}

type reservationRoutes struct {
	reservationService ReservationService
}

func NewReservationRoutes(router *mux.Router, s ReservationService) {
	r := &reservationRoutes{
		reservationService: s,
	}

	router.HandleFunc("/create", r.create()).Methods(http.MethodPost)
	router.HandleFunc("/revenue", r.revenue()).Methods(http.MethodPost)
	router.HandleFunc("/refund", r.refund()).Methods(http.MethodPost)
}

type createOutput struct {
	ReservationId int `json:"id" example:"1"`
}

// @Summary create
// @Description create new reservation
// @Tags api/v1/reservation/create
// @Accept json
// @Produce json
// @Param reservation body model.Reservation true  "ID NO NEED"
// @Success 201 {object} createOutput
// @Failure 400 {object} ErrorOutput
// @Failure 422 {object} ErrorOutput
// @Failure 500 {object} ErrorOutput
// @Security JWT
// @Router /api/v1/reservation/create [post]
func (route *reservationRoutes) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := &model.Reservation{}
		if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
			newErrorRespond(w, r, http.StatusBadRequest, errInvalidRequestBody)
			return
		}

		if err := res.Validate(); err != nil {
			newErrorRespond(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		res, err := route.reservationService.CreateReservation(r.Context(), res.AccountId, res.ServiceId, res.OrderId, res.Amount)
		if err != nil {
			if errors.Is(err, service.ErrNotEnoughMoney) {
				newErrorRespond(w, r, http.StatusBadRequest, errNotEnoughMoney)
				return
			} else if errors.Is(err, service.ErrReservationAlreadyExists) {
				newErrorRespond(w, r, http.StatusBadRequest, errReservationAlreadyExists)
				return
			} else {
				newErrorRespond(w, r, http.StatusInternalServerError, err)
				return
			}
		}
		respond(w, r, http.StatusCreated, createOutput{ReservationId: res.Id})
	}
}

type revenueOutput struct {
	Msg string `json:"msg" example:"OK"`
}

// @Summary revenue
// @Description recognizes revenue
// @Tags api/v1/reservation/revenue
// @Accept json
// @Produce json
// @Param reservation body model.Reservation true  "ID NO NEED"
// @Success 200 {object} revenueOutput
// @Failure 400 {object} ErrorOutput
// @Failure 404 {object} ErrorOutput
// @Failure 422 {object} ErrorOutput
// @Failure 500 {object} ErrorOutput
// @Security JWT
// @Router /api/v1/reservation/revenue [post]
func (route *reservationRoutes) revenue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := &model.Reservation{}
		if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
			newErrorRespond(w, r, http.StatusBadRequest, errInvalidRequestBody)
			return
		}

		if err := res.Validate(); err != nil {
			newErrorRespond(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		err := route.reservationService.Revenue(r.Context(), res.AccountId, res.ServiceId, res.OrderId, res.Amount)
		if err != nil {
			if errors.Is(err, service.ErrReservationNotFound) {
				newErrorRespond(w, r, http.StatusNotFound, errReservationNotFound)
				return
			}
			newErrorRespond(w, r, http.StatusInternalServerError, err)
			return
		}

		respond(w, r, http.StatusOK, revenueOutput{Msg: "OK"})
	}
}

type refundOutput struct {
	Msg string `json:"msg" example:"OK"`
}

// @Summary refund
// @Description refund
// @Tags api/v1/reservation/refund
// @Accept json
// @Produce json
// @Param reservation body model.Reservation true  "ID NO NEED"
// @Success 200 {object} refundOutput
// @Failure 400 {object} ErrorOutput
// @Failure 404 {object} ErrorOutput
// @Failure 422 {object} ErrorOutput
// @Failure 500 {object} ErrorOutput
// @Security JWT
// @Router /api/v1/reservation/refund [post]
func (route *reservationRoutes) refund() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := &model.Reservation{}
		if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
			newErrorRespond(w, r, http.StatusBadRequest, errInvalidRequestBody)
			return
		}

		if err := res.Validate(); err != nil {
			newErrorRespond(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		err := route.reservationService.Refund(r.Context(), res.AccountId, res.ServiceId, res.OrderId, res.Amount)
		if err != nil {
			if errors.Is(err, service.ErrReservationNotFound) {
				newErrorRespond(w, r, http.StatusNotFound, errReservationNotFound)
				return
			}
			newErrorRespond(w, r, http.StatusInternalServerError, err)
			return
		}

		respond(w, r, http.StatusOK, refundOutput{Msg: "OK"})
	}
}
