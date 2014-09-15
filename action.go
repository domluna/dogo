package digitalocean

import (
	"fmt"
	"time"
)

const ActionEndpoint = "actions"

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

type Actions []*Action

func (c *Client) ListActions() (Actions, error) {
	s := struct {
		Actions `json:"actions,omitempty"`
	}{}
	err := c.get(ActionEndpoint, &s)
	if err != nil {
		return nil, err
	}
	return s.Actions, nil
}

func (c *Client) GetAction(id int) (*Action, error) {
	u := fmt.Sprintf("%s/%d", ActionEndpoint, id)
	s := struct {
		Action `json:"action,omitempty"`
	}{}
	err := c.get(u, &s)
	if err != nil {
		return nil, err
	}
	return &s.Action, nil
}
