package dogo

import (
	"encoding/json"
	"fmt"
)

// Representation for the size of a DigitalOcean droplet.
type Size struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// SizesMap is a mapping between the slug
// representation of a droplet size to it's
// id.
var SizesMap = map[string]int{
	"512MB": 66,
	"1GB":   63,
	"2GB":   62,
	"4GB":   64,
	"8GB":   65,
	"16GB":  61,
	"32GB":  60,
	"48GB":  70,
	"64GB":  69,
}

// GetSizes returns all currently available droplet sizes.
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
		Status     string `json:"status"`
		Sizes      []Size `json:"sizes"`
		ErrMessage string `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Status == "ERROR" {
		return nil, fmt.Errorf("%s: %s", resp.Status, resp.ErrMessage)
	}

	return resp.Sizes, nil
}
