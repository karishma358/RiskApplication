package store

import (
	"sync"

	"github.com/google/uuid"
	"riskapp/models"
)

// InMemoryStore holds risks in-memory
type InMemoryStore struct {
	mutex sync.Mutex
	risks map[string]models.Risk
}

// NewStore initializes a new in-memory store
func NewStore() *InMemoryStore {
	return &InMemoryStore{
		risks: make(map[string]models.Risk),
	}
}

// GetAllRisks retrieves all risks
func (s *InMemoryStore) GetAllRisks() []models.Risk {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var riskList []models.Risk
	for _, risk := range s.risks {
		riskList = append(riskList, risk)
	}
	return riskList
}

// GetRiskByID retrieves a specific risk by ID
func (s *InMemoryStore) GetRiskByID(id string) (models.Risk, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	risk, exists := s.risks[id]
	return risk, exists
}

// CreateRisk adds a new risk to the store
func (s *InMemoryStore) CreateRisk(state, title, description string) models.Risk {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	risk := models.Risk{
		ID:          uuid.New().String(),
		State:       state,
		Title:       title,
		Description: description,
	}

	s.risks[risk.ID] = risk
	return risk
}