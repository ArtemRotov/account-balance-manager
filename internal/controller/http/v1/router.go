package v1

import (
	"encoding/json"
	"net/http"

	_ "github.com/ArtemRotov/account-balance-manager/docs"
	"github.com/ArtemRotov/account-balance-manager/internal/service"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func New(router *mux.Router, services *service.Services) {
	// Swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		// httpSwagger.DeepLinking(true),
		// httpSwagger.DocExpansion("none"),
		// httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	authRoute := router.PathPrefix("/auth").Subrouter()
	NewAuthRoutes(authRoute, services.Auth)
}

func respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
