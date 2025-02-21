package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"riskapp/handlers"
	"riskapp/store"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *mux.Router {
	s := store.NewStore()
	h := handlers.NewRiskHandler(s)
	r := mux.NewRouter()

	r.HandleFunc("/v1/risks", h.GetAllRisks).Methods(http.MethodGet)
	r.HandleFunc("/v1/risks", h.CreateRisk).Methods(http.MethodPost)
	r.HandleFunc("/v1/risks/{id}", h.GetRiskByID).Methods(http.MethodGet)

	return r
}

func TestCreateRisk(t *testing.T) {
	router := setupTestRouter()

	requestBody := `{"state":"open", "title":"Test Risk", "description":"This is a test risk"}`
	req, _ := http.NewRequest(http.MethodPost, "/v1/risks", bytes.NewBuffer([]byte(requestBody)))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestGetAllRisks(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest(http.MethodGet, "/v1/risks", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	// Ensure response is a valid JSON array (even if empty)
	var risks []map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &risks)
	assert.NoError(t, err, "Response should be valid JSON")
	assert.Equal(t, 0, len(risks), "Expected empty array when no risks exist")
}

func TestGetRiskByID(t *testing.T) {
	router := setupTestRouter()

	// Create a new risk first
	requestBody := `{"state":"investigating", "title":"Sample Risk", "description":"Risk details"}`
	req, _ := http.NewRequest(http.MethodPost, "/v1/risks", bytes.NewBuffer([]byte(requestBody)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	// Extract risk ID
	var createdRisk map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &createdRisk)
	riskID := createdRisk["id"].(string)

	// Fetch the created risk by ID
	req, _ = http.NewRequest(http.MethodGet, "/v1/risks/"+riskID, nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
