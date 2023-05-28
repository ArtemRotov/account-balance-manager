package v1

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New(router *mux.Router) {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	}).Methods("GET")

	authRoute := router.PathPrefix("/auth").Subrouter()
	authRoute.HandleFunc("/h", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("auth page!"))
	}).Methods("GET")
}
