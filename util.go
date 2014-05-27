package dogo

import (
	"io/ioutil"
	"net/http"
)

// Request sends a GET request to the endpoint and returns
// the response body.
func Request(endpoint string) ([]byte, error) {
	resp, err := http.Get(endpoint)
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
