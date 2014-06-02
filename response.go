package dogo

import (
	"encoding/json"
	"fmt"
)

// Response encapsulates the entire DigitalOcean API Response.
// A Typical Response will have a status (either "OK", or "ERROR")
// and an error message, "" if no error is present.
//
// These fields will be followed with one of the other fields
// depending entirely on the query endpoint.
type Response struct {
	Status     string    `json:"status"`
	ErrMessage string    `json:"error_message"`
	Droplets   []Droplet `json:"droplets, omitempty"`
	Droplet    Droplet   `json:"droplet, omitempty"`
	Images     []Image   `json:"images, omitempty"`
	Regions    []Region  `json:"regions, omitempty"`
	Sizes      []Size    `json:"sizes, omitempty"`
	SSHKeys    []SSHKey  `json:"ssh_keys, omitempty"`
	SSHKey     SSHKey    `json:"ssh_key, omitempty"`
}

type Params map[string]interface{}

func (c *Client) Send(endpoint string, id interface{}, params Params) (Response, error) {

	if params == nil {
		params = Params{}
	}
	// Add credentials
	params["client_id"] = c.Auth.ClientID
	params["api_key"] = c.Auth.APIKey
	u := createURL(endpoint, id, params)

	fmt.Println("URL:", u)

	var resp Response
	body, err := get(u)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}

	if resp.Status == "ERROR" {
		return resp, fmt.Errorf("%s: %s", resp.Status, resp.ErrMessage)
	}

	return resp, nil
}
