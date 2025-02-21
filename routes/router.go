package routes

import (
	"encoding/json"
	"net/http"
	"riskapp/handlers"
	"riskapp/store"

	"github.com/gorilla/mux"
)

// SetupRouter initializes the router and registers endpoints
func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	s := store.NewStore()
	h := handlers.NewRiskHandler(s)

	// Risk API Routes
	r.HandleFunc("/v1/risks", h.GetAllRisks).Methods(http.MethodGet)
	r.HandleFunc("/v1/risks", h.CreateRisk).Methods(http.MethodPost)
	r.HandleFunc("/v1/risks/{id}", h.GetRiskByID).Methods(http.MethodGet)

	// Health Check Route
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}).Methods(http.MethodGet)

	return r
}
