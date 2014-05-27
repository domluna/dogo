package dogo

import (
	"encoding/json"
	"fmt"
)

// Region represents a DigitalOcean region.
type Region struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// RegionsMap is a mapping between the slug
// representation of the region and it's id.
//
// Note that some regions listed may actually not be
// currently available.
var RegionsMap = map[string]int{
	"nyc1": 1,
	"ams1": 2,
	"sfo1": 3,
	"nyc2": 4,
	"ams2": 5,
	"sgp1": 6,
}

// GetRegions gets all current available regions a droplet may be created in.
func (c *Client) GetRegions() ([]Region, error) {
	query := fmt.Sprintf("%s?client_id=%s&api_key=%s",
		RegionsEndpoint,
		c.Auth.ClientID,
		c.Auth.APIKey,
	)

	body, err := Request(query)
	if err != nil {
		return nil, err
	}

	resp := struct {
		Status     string   `json"status"`
		Regions    []Region `json:"regions"`
		ErrMessage string   `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Status == "ERROR" {
		return nil, fmt.Errorf("%s: %s", resp.Status, resp.ErrMessage)
	}

	return resp.Regions, nil
}
