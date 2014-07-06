package digitalocean

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	BaseURL = "https://api.digitalocean.com/v2"
)

// To make a new Client call NewClient.
type Client struct {
	Token string
}

// NewClient creates a new Client.
func NewClient(token string) Client {
	return Client{token}
}

// EnvAuth tries to get the api token from the environment
// variable DIGITALOCEAN_TOKEN.
func EnvAuth() (string, error) {
	token := os.Getenv("DIGITALOCEAN_TOKEN")
	if token == "" {
		return "", EnvError
	}
	return token, nil
}

func (c Client) Get(u string, v interface{}) error {
	req, err := http.NewRequest("GET", u, nil)
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

func (c Client) Del(u string) error {
	req, err := http.NewRequest("DELETE", u, nil)
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

func (c Client) Put(u string, params map[string]interface{}, v interface{}) error {
	payload, err := json.Marshal(params)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", u, bytes.NewReader(payload))
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

func (c Client) Post(u string, params map[string]interface{}, v interface{}) error {
	payload, err := json.Marshal(params)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", u, bytes.NewReader(payload))
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

func (c Client) DoRequest(req *http.Request, v interface{}) error {
	cl := &http.Client{}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	resp, err := cl.Do(req)
	if err != nil {
		return err
	}
	err = Decode(resp, v)
	if err != nil {
		return err
	}
	return nil
}

// Decode parses the response.
func Decode(resp *http.Response, v interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("Parsing response", string(body))
	// create error
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		apiErr := &APIError{
			StatusCode: resp.StatusCode,
		}
		err := json.Unmarshal(body, apiErr)
		if err != nil {
			return err
		}
		return apiErr
	}

	if v != nil {
		err := json.Unmarshal(body, &v)
		fmt.Printf("%v\n", v)
		if err != nil {
			return err
		}
	}
	return nil
}
