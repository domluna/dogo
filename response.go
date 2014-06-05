package dogo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Response encapsulates the entire DigitalOcean API Response.
// A Typical Response will have a status (either "OK", or "ERROR")
// and an error message, "" if no error is present.
//
// These fields will be followed with one of the other fields
// depending entirely on the query endpoint.
type Response struct {
	Status     string    `json:"status"`
	ErrMessage string    `json:"error_message,omitempty"`
	Droplets   []Droplet `json:"droplets,omitempty"`
	Droplet    *Droplet  `json:"droplet,omitempty"`
	Images     []Image   `json:"images,omitempty"`
	Regions    []Region  `json:"regions,omitempty"`
	Sizes      []Size    `json:"sizes,omitempty"`
	SSHKeys    []SSHKey  `json:"ssh_keys,omitempty"`
	SSHKey     *SSHKey   `json:"ssh_key,omitempty"`
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

	var resp Response
	// util function for reading a response for http get
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

// get sends a GET request to the endpoint and returns
// the response body.
func get(u string) ([]byte, error) {
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
