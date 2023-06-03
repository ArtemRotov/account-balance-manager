package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
	"github.com/ArtemRotov/account-balance-manager/internal/service"
	"github.com/gorilla/mux"
)

type AuthRoutes struct {
	service service.Auth
}

func NewAuthRoutes(router *mux.Router, s service.Auth) {
	r := &AuthRoutes{
		service: s,
	}

	router.HandleFunc("/sign-up", r.SignUp()).Methods(http.MethodPost)
	router.HandleFunc("/sign-in", r.SignIn()).Methods(http.MethodPost)
}

func (route *AuthRoutes) SignUp() http.HandlerFunc {

	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			newErrorRespond(w, r, http.StatusBadRequest, errInvalidRequestBody)
			return
		}

		u := &model.User{
			Username: req.Username,
			Password: req.Password,
		}

		if err := u.Validate(); err != nil {
			newErrorRespond(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		id, err := route.service.CreateUser(r.Context(), u.Username, u.Password)
		if err != nil {
			if errors.Is(err, service.ErrUserAlreadyExists) {
				newErrorRespond(w, r, http.StatusBadRequest, err)
				return
			}
			newErrorRespond(w, r, http.StatusInternalServerError, err)
			return
		}

		type response struct {
			Id int `json:"id"`
		}

		respond(w, r, http.StatusOK, &response{Id: id})
	}
}

func (r *AuthRoutes) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("SignIn"))
	}
}
