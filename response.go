package dogo

import (
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

func (c *Client) Send(endpoint string, params map[string]interface{}) (Response, error) {
	return Response{}, nil
}
