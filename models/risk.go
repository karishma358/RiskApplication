package models

// Risk represents a risk entity
type Risk struct {
	ID          string `json:"id"`
	State       string `json:"state"`
	Title       string `json:"title"`
	Description string `json:"description"`
}