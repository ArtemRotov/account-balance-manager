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

	// Swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		// httpSwagger.DeepLinking(true),
		// httpSwagger.DocExpansion("none"),
		// httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	// authRoute := router.PathPrefix("/auth").Subrouter()
	// authRoute.HandleFunc("/h", authHanlder()).Methods("GET")
}
