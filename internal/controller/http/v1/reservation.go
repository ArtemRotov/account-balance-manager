package v1

import (
	"net/http"

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
}

func (route *reservationRoutes) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
