package v1

import (
	"encoding/json"
	"errors"
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

func GetImages(filter string) ([]Image, error) {
	query := fmt.Sprintf("%s?client_id=%s&api_key=%s&filter=%s",
		ImagesEndpoint,
		config.Conf.ClientID,
		config.Conf.APIKey,
		filter)

	body, err := sendQuery(query)
	if err != nil {
		return nil, err
	}

	resp := struct {
		Status string  `json"status"`
		Images []Image `json:"images"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Status == "ERROR" {
		return nil, errors.New("Error retrieving images")
	}

	return resp.Images, nil
}
