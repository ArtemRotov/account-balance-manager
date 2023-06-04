package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
	"github.com/ArtemRotov/account-balance-manager/internal/service"
	"github.com/gorilla/mux"
)

type authRoutes struct {
	service service.Auth
}

type signUpInput struct {
	Username string `json:"username" example:"example@mail.org"`
	Password string `json:"password" example:"password12345678"`
}

type signUpOutput struct {
	Id int `json:"id" example:"1"`
}

func NewAuthRoutes(router *mux.Router, s service.Auth) {
	r := &authRoutes{
		service: s,
	}

	router.HandleFunc("/sign-up", r.signUp()).Methods(http.MethodPost)
	router.HandleFunc("/sign-in", r.signIn()).Methods(http.MethodPost)
}

// @Summary Sign up
// @Description Sign up
// @Tags auth
// @Accept json
// @Produce json
// @Param req body signUpInput true "req"
// @Success 201 {object} signUpOutput
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /auth/sign-up [post]
func (route *authRoutes) signUp() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := &signUpInput{}

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

		respond(w, r, http.StatusCreated, &signUpOutput{Id: id})
	}
}

func (r *authRoutes) signIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("SignIn"))
	}
}
