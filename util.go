package dogo

import (
	"io/ioutil"
	"net/http"
	"os"
)

// Sends a GET request to the query url and returns
// the response or an error.
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
