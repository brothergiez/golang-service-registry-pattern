package registry

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, handler *Handler) {
	router.HandleFunc("/health", handler.HealthCheck).Methods(http.MethodGet)

	router.HandleFunc("/services", handler.RegisterService).Methods(http.MethodPost)
	router.HandleFunc("/services", handler.GetServices).Methods(http.MethodGet)
	router.HandleFunc("/services/{id}", handler.DeregisterService).Methods((http.MethodDelete))
}
