package responsewriter

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	Code int
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.Code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
