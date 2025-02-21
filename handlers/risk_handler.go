package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"riskapp/store"

	"github.com/gorilla/mux"
)

// RiskHandler handles API requests for risks
type RiskHandler struct {
	Store *store.InMemoryStore
}

// NewRiskHandler initializes a new risk handler
func NewRiskHandler(s *store.InMemoryStore) *RiskHandler {
	return &RiskHandler{Store: s}
}

// respondWithError sends a structured error response
func respondWithError(w http.ResponseWriter, status int, message string) {
	log.Printf("Error %d: %s", status, message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// GetAllRisks handles GET /v1/risks
func (h *RiskHandler) GetAllRisks(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching all risks")

	risks := h.Store.GetAllRisks()

	if len(risks) == 0 {
		respondWithError(w, http.StatusOK, "No risks are currently present")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(risks)
}

// GetRiskByID handles GET /v1/risks/{id}
func (h *RiskHandler) GetRiskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Printf("Fetching risk with ID: %s", id)

	risk, exists := h.Store.GetRiskByID(id)
	if !exists {
		respondWithError(w, http.StatusNotFound, "Risk not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(risk)
}

// CreateRisk handles POST /v1/risks
func (h *RiskHandler) CreateRisk(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating a new risk")

	var input struct {
		State       string `json:"state"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if input.State == "" || input.Title == "" || input.Description == "" {
		respondWithError(w, http.StatusBadRequest, "State, Title, and Description are required fields")
		return
	}

	risk := h.Store.CreateRisk(input.State, input.Title, input.Description)
	log.Printf("Risk created with ID: %s", risk.ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(risk)
}
