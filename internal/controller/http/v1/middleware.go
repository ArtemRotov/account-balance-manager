package v1

import (
	"context"
	"fmt"
	"github.com/ArtemRotov/account-balance-manager/pkg/responsewriter"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
	"net/http"
	"strings"
	"time"

	"github.com/ArtemRotov/account-balance-manager/internal/service"
)

type ctxKey int8

const (
	ctxKeyUser ctxKey = iota
	ctxRequestId
)

type authMiddleware struct {
	service service.Auth
	log     *slog.Logger
}

func NewAuthMiddleware(s service.Auth, log *slog.Logger) *authMiddleware {
	return &authMiddleware{
		service: s,
		log:     log,
	}
}

func (m *authMiddleware) verify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := bearerToken(r)
		if !ok {
			newErrorRespond(w, r, http.StatusUnauthorized, errInvalidAuthHeader)
			m.log.Error(fmt.Sprintf("authMiddleware.verify: bearerToken: %v", errInvalidAuthHeader.Error()))
			return
		}

		userId, err := m.service.ParseToken(token)
		if err != nil {
			newErrorRespond(w, r, http.StatusUnauthorized, errCannotParseToken)
			m.log.Error(fmt.Sprintf("authMiddleware.verify: m.service.ParseToken: %v", err.Error()))
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

type middleware struct {
	log *slog.Logger
}

func NewMiddleware(log *slog.Logger) *middleware {
	return &middleware{
		log: log,
	}
}

func (m *middleware) requestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxRequestId, id)))
	})
}

func (m *middleware) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := m.log.With(slog.String("ID", r.Context().Value(ctxRequestId).(string)))
		log.Debug(fmt.Sprintf("started %s %s", r.Method, r.RequestURI))
		startTime := time.Now()

		rw := &responsewriter.ResponseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		log.Debug(fmt.Sprintf("completed with code %d(%s) [%v]",
			rw.Code, http.StatusText(rw.Code), time.Since(startTime)))
	})
}
