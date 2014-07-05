package action

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

type Client struct {
	client Client
}

func (c *Client) GetAll() (Actions, error) {
	s := struct {
		Actions `json:"actions,omitempty"`
		Meta    `json:"meta,omitempty"`
	}{}
	err := c.client.Get(ActionsEndpoint, &s)
	if err != nil {
		return nil, err
	}
	return s.Actions, nil
}

func (c *Client) Get(id int) (Action, error) {
	u := fmt.Sprintf("%s/%d", ActionsEndpoint, id)
	s := struct {
		Action `json:"action,omitempty"`
		Meta   `json:"meta,omitempty"`
	}{}
	err := c.client.Get(u, &s)
	if err != nil {
		return s.Action, err
	}
	return s.Action, nil
}
