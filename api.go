package dogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Params are to be used in conjuction with POST or PUT.
type Params map[string]interface{}

// Client is a client to the DigitalOcean API service
type Client struct {
	// DO Access Token
	Token string

	// Base DO API URL
	URL string
}

// NewClient creates a new Client.
func NewClient(token string) (*Client, error) {
	if token == "" {
		token = os.Getenv("DIGITALOCEAN_TOKEN")
	}
	if token == "" {
		return nil, EnvError
	}
	cl := &Client{
		Token: token,
		URL:   "https://api.digitalocean.com/v2",
	}
	return cl, nil
}

func (c *Client) get(endpoint string, v interface{}) error {
	endpoint = fmt.Sprintf("%s/%s", c.URL, endpoint)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	err = c.DoRequest(req, v)

	if err != nil {
		return err
	}
	return nil
}

func (c *Client) delete(endpoint string) error {
	endpoint = fmt.Sprintf("%s/%s", c.URL, endpoint)
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	err = c.DoRequest(req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) post(endpoint string, opts interface{}, v interface{}) error {
	endpoint = fmt.Sprintf("%s/%s", c.URL, endpoint)
	payload, err := json.Marshal(opts)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	c.DoRequest(req, v)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) put(endpoint string, opts interface{}, v interface{}) error {
	endpoint = fmt.Sprintf("%s/%s", c.URL, endpoint)
	payload, err := json.Marshal(opts)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", endpoint, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	c.DoRequest(req, v)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DoRequest(req *http.Request, v interface{}) error {
	cl := &http.Client{}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	resp, err := cl.Do(req)
	if err != nil {
		return fmt.Errorf("Error attemping request: %s", err)
	}
	err = decode(resp, v)
	if err != nil {
		return err
	}
	return nil
}

// Decode parses the response.
func decode(resp *http.Response, v interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading response: %s", err)
	}
	// create error
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		apiErr := &APIError{
			StatusCode: resp.StatusCode,
		}
		err := json.Unmarshal(body, apiErr)
		if err != nil {
			return fmt.Errorf("Error UnMarshaling JSON Response into error: %s", err)
		}
		return apiErr
	}

	if v != nil {
		err := json.Unmarshal(body, &v)
		if err != nil {
			return fmt.Errorf("Error UnMarshaling JSON Response into struct: %s", err)
		}
	}
	return nil
}

// DoAction performs an action on an endpoint resource with the passed id,
// type of action and its required params are described in the DigitalOcean API.
//
// For example in this case the image resource.
//
// https://developers.digitalocean.com/v2/#image-actions
//
// An example of some params:
//	params := digitalocean.Params{
//		"type": "transfer",
//		"region": "nyc2",
//	}
//
// The above example specifies the type of action, in this case resizing and
// the additional param in this case the size to resize to "1024mb".
//
// Params will sometimes only require the type of action and no additional params.
func (c *Client) DoAction(endpoint string, id int, params Params) error {
	u := fmt.Sprintf("%s/%d/actions", endpoint, id)
	err := c.post(u, params, nil)
	if err != nil {
		return err
	}
	return nil
}
