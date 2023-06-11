package v1

import (
	"context"
	"net/http"
	"strings"

	"github.com/ArtemRotov/account-balance-manager/internal/service"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type ctxKey int8

const (
	ctxKeyUser ctxKey = iota
)

type authMiddleware struct {
	service service.Auth
}

func NewAuthMiddleware(router *mux.Router, s service.Auth) *authMiddleware {
	return &authMiddleware{
		service: s,
	}
}

func (m *authMiddleware) verify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := bearerToken(r)
		if !ok {
			newErrorRespond(w, r, http.StatusUnauthorized, errInvalidAuthHeader)
			log.Errorf("authMiddleware.verify: bearerToken: %v", errInvalidAuthHeader.Error())
			return
		}

		userId, err := m.service.ParseToken(token)
		if err != nil {
			newErrorRespond(w, r, http.StatusUnauthorized, errCannotParseToken)
			log.Errorf("authMiddleware.verify: m.service.ParseToken: %v", err.Error())
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, userId)))
	})
}

func bearerToken(r *http.Request) (string, bool) {
	const prefix = "Bearer "

	header := r.Header.Get("Authorization")
	if header == "" {
		return "", false
	}

	if len(header) > len(prefix) && strings.EqualFold(header[:len(prefix)], prefix) {
		return header[len(prefix):], true
	}

	return "", false
}
