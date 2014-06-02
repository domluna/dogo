package dogo

import (
	"io/ioutil"
	"net/http"
)

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
