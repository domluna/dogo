package dogo

import (
	"encoding/json"
	"fmt"
)

// Representation for a DigitalOcean Image.
type Image struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Distribution string `json:"distribution"`
	Slug         string `json:"slug"`
	Public       bool   `json:"public"`
}

// GetImages returns DigitalOcean images, filter can either be
// "my_images" or "global".
//
// If filter is set to "my_images" user snapshots will be returned.
//
// If filter is set to "global" all default images will be returned.
func (c *Client) GetImages(filter string) ([]Image, error) {
	query := fmt.Sprintf("%s?client_id=%s&api_key=%s&filter=%s",
		ImagesEndpoint,
		c.Auth.ClientID,
		c.Auth.APIKey,
		filter)

	body, err := Request(query)
	if err != nil {
		return nil, err
	}

	resp := struct {
		Status     string  `json"status"`
		Images     []Image `json:"images"`
		ErrMessage string  `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Status == "ERROR" {
		return nil, fmt.Errorf("%s: %s", resp.Status, resp.ErrMessage)
	}

	return resp.Images, nil
}
