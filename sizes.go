package dogo

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Representation for the size of a DigitalOcean droplet.
type Size struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (c *Client) GetSizes() ([]Size, error) {
	query := fmt.Sprintf("%s?client_id=%s&api_key=%s",
		SizesEndpoint,
		c.Auth.ClientID,
		c.Auth.APIKey)

	body, err := Request(query)
	if err != nil {
		return nil, err
	}

	resp := struct {
		Status string `json:"status"`
		Sizes  []Size `json:"sizes"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Status == "ERROR" {
		return nil, errors.New("Error retrieving droplet sizes")
	}

	return resp.Sizes, nil
}
