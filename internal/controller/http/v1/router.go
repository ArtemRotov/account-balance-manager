package v1

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New(router *mux.Router) {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	}).Methods("GET")
}
