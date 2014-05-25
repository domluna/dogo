package dogo

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
		Err string `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Status == "ERROR" {
		return nil, errors.New(resp.Err)
	}

	return resp.Images, nil
}