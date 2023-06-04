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
// @Param user body model.User true "ID NO NEED"
// @Success 201 {object} signUpOutput
// @Failure 400 {object} ErrorOutput
// @Failure 422 {object} ErrorOutput
// @Failure 500 {object} ErrorOutput
// @Router /auth/sign-up [post]
func (route *authRoutes) signUp() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		user := &model.User{}

		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			newErrorRespond(w, r, http.StatusBadRequest, errInvalidRequestBody)
			return
		}

		if err := user.Validate(); err != nil {
			newErrorRespond(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		id, err := route.service.CreateUser(r.Context(), user.Username, user.Password)
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

type signInOutput struct {
	Token string `json:"Token" example:"eyJhbGc..."`
}

// @Summary Sign in
// @Description Sign in
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.User true "ID NO NEED"
// @Success 200 {object} signInOutput
// @Failure 400 {object} ErrorOutput
// @Failure 422 {object} ErrorOutput
// @Failure 500 {object} ErrorOutput
// @Router /auth/sign-in [post]
func (route *authRoutes) signIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &model.User{}

		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			newErrorRespond(w, r, http.StatusBadRequest, errInvalidRequestBody)
			return
		}

		if err := user.Validate(); err != nil {
			newErrorRespond(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		token, err := route.service.GenerateToken(r.Context(), user.Username, user.Password)
		if err != nil {
			if errors.Is(err, service.ErrUserNotFound) {
				newErrorRespond(w, r, http.StatusBadRequest, errInvalidUsernameOrPassword)
				return
			}
			newErrorRespond(w, r, http.StatusInternalServerError, err)
			return
		}

		respond(w, r, http.StatusOK, &signInOutput{Token: token})
	}
}
