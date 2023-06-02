package v1

import (
	"net/http"

	"github.com/ArtemRotov/account-balance-manager/internal/service"
	"github.com/gorilla/mux"
)

type AuthRoutes struct {
	s service.Auth
}

func NewAuthRoutes(route *mux.Router, s service.Auth) {
	r := &AuthRoutes{
		s: s,
	}

	route.HandleFunc("/sign-up", r.SignUp()).Methods(http.MethodPost)
	route.HandleFunc("/sign-in", r.SignIn()).Methods(http.MethodPost)
}

func (r *AuthRoutes) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("SignUp"))
	}
}

func (r *AuthRoutes) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("SignIn"))
	}
}
