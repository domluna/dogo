package dogo

import (
	"fmt"
)

// Action is a representation of a DigitalOcean action.
type Action struct {
	ID           int    `json:"id,omitempty"`
	Status       string `json:"status,omitempty"`
	Type         string `json:"type,omitempty"`
	StartedAt    string `json:"started_at,omitempty"`
	CompletedAt  string `json:"completed_at,omitempty"`
	ResourceID   int    `json:"resource_id,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
	Region       string `json:"region,omitempty"`
}

// Actions is a list of Action.
type Actions []*Action

// ListActions retrieves all DigitalOcean actions.
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

// GetAction retrieves an the action related to the passed id.
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
