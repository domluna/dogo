package dogo

import (
	"time"
)

type Action struct {
	ID           int       `json:"id"`
	Status       string    `json:"status"`
	Type         string    `json:"type"`
	StartedAt    time.Time `json:"started_at"`
	CompletedAt  time.Time `json:"completed_at"`
	ResourceID   int       `json:"resource_id"`
	ResourceType string    `json:"resource_type"`
	Region       string    `json:"region"`
}
