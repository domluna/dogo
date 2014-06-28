package dogo

import (
	"fmt"
	"time"
)

type Action struct {
	ID           int       `json:"id,omitempty"`
	Status       string    `json:"status,omitempty"`
	Type         string    `json:"type,omitempty"`
	StartedAt    time.Time `json:"started_at,omitempty"`
	CompletedAt  time.Time `json:"completed_at,omitempty"`
	ResourceID   int       `json:"resource_id,omitempty"`
	ResourceType string    `json:"resource_type,omitempty"`
	Region       string    `json:"region,omitempty"`
}

type Actions []Action

type ActionClient struct {
	Client
}

func (ac *ActionClient) GetActions() (Actions, error) {
	var a Actions
	err := ac.Client.Get(ActionsEndpoint, a)
	if err != nil {
		return a, err
	}
	return a, nil
}

func (ac *ActionClient) GetAction(id int) (Action, error) {
	u := fmt.Sprintf("%s/%d", ActionsEndpoint, id)
	var a Action
	err := ac.Client.Get(u, a)
	if err != nil {
		return a, err
	}
	return a, nil
}
