package v1

import (
	"net/http"

	_ "github.com/ArtemRotov/account-balance-manager/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func New(router *mux.Router) {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	}).Methods("GET")

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		// httpSwagger.DeepLinking(true),
		// httpSwagger.DocExpansion("none"),
		// httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	authRoute := router.PathPrefix("/auth").Subrouter()
	authRoute.HandleFunc("/h", authHanlder()).Methods("GET")
}

// @Summary authHanlder
// @Description authHanlder
// @Tags auth
// @Accept json
// @Produce plain
// @Success 200
// @Router /auth/h [get]

func authHanlder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("auth page!"))
	}
}
