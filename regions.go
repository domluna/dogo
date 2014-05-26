package dogo

import (
	"encoding/json"
	"fmt"
)

type Region struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func GetRegions() ([]Region, error) {
	query := fmt.Sprintf("%s?client_id=%s&api_key=%s",
		RegionsEndpoint,
		config.Conf.ClientID,
		config.Conf.APIKey,
	)

	body, err := sendQuery(query)
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
