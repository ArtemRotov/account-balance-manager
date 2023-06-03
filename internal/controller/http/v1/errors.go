package v1

import (
	"fmt"
	"net/http"
)

var (
	errInvalidRequestBody = fmt.Errorf("invalid request body")
	errInvalidAuthHeader  = fmt.Errorf("invalid auth header")
	errCannotParseToken   = fmt.Errorf("cannot parse token")
)

func newErrorRespond(w http.ResponseWriter, r *http.Request, code int, err error) {
	w.WriteHeader(code)
	data := map[string]string{"error": err.Error()}

	respond(w, r, code, data)
}
