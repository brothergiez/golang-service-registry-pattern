package registry

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	registry *ServiceRegistry
}

func NewHandler(registry *ServiceRegistry) *Handler {
	return &Handler{registry: registry}
}

func (h *Handler) RegisterService(w http.ResponseWriter, r *http.Request) {
	var service Service
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.registry.RegisterService(r.Context(), &service); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetServices(w http.ResponseWriter, r *http.Request) {
	services, err := h.registry.GetServices(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func (h *Handler) DeregisterService(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.registry.DeregisterService(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
