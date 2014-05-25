package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
		log.Fatal(err)
	}

	resp := struct {
		Status  string   `json"status"`
		Regions []Region `json:"regions"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Status == "ERROR" {
		return nil, errors.New("Error retrieving regions")
	}

	return resp.Regions, nil
}
