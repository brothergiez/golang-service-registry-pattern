package registry

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	registry *ServiceRegistry
}

type HealthCheckResponse struct {
	Status     string `json:"status"`
	Database   string `json:"database"`
	ServerTime string `json:"server_time"`
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

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Cek koneksi database
	if err := h.registry.repo.PingDatabase(ctx); err != nil {
		response := HealthCheckResponse{
			Status:     "unhealthy",
			Database:   "disconnected",
			ServerTime: time.Now().Format(time.RFC3339),
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Jika semua sehat
	response := HealthCheckResponse{
		Status:     "healthy",
		Database:   "connected",
		ServerTime: time.Now().Format(time.RFC3339),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
