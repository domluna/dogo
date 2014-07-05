package action

import (
	"fmt"
	"time"

	"github.com/domluna/dogo/digitalocean"
)

const (
	Endpoint = digitalocean.BaseURL + "/actions"
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
	client *digitalocean.Client
}

func (c *Client) GetAll() (Actions, error) {
	s := struct {
		Actions           `json:"actions,omitempty"`
		digitalocean.Meta `json:"meta,omitempty"`
	}{}
	err := c.client.Get(Endpoint, &s)
	if err != nil {
		return nil, err
	}
	return s.Actions, nil
}

func (c *Client) Get(id int) (Action, error) {
	u := fmt.Sprintf("%s/%d", Endpoint, id)
	s := struct {
		Action            `json:"action,omitempty"`
		digitalocean.Meta `json:"meta,omitempty"`
	}{}
	err := c.client.Get(u, &s)
	if err != nil {
		return s.Action, err
	}
	return s.Action, nil
}
