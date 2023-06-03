package v1

import (
	"errors"
	"net/http"
)

var (
	errInvalidRequestBody = errors.New("invalid request body")
)

func newErrorRespond(w http.ResponseWriter, r *http.Request, code int, err error) {
	w.WriteHeader(code)
	data := map[string]string{"error": err.Error()}

	respond(w, r, code, data)
}
