package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
	"github.com/ArtemRotov/account-balance-manager/internal/service"
	"github.com/gorilla/mux"
)

type reservationRoutes struct {
	reservationService service.Reservation
}

func NewReservationRoutes(router *mux.Router, s service.Reservation) {
	r := &reservationRoutes{
		reservationService: s,
	}

	router.HandleFunc("/create", r.create()).Methods(http.MethodPost)
	router.HandleFunc("/revenue", r.revenue()).Methods(http.MethodPut)
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
// @Router /api/v1/reservation/revenue [put]
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
