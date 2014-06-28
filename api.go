package dogo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

// To make a new Client call NewClient.
type Client struct {
	Token string
}

// NewClient creates a new Client.
func NewClient(token string) *Client {
	return &Client{token}
}

// EnvAuth tries to get the api token from the environment
// variable DIGITALOCEAN_TOKEN.
func EnvAuth() (Client, error) {
	var cli Client
	cli.Token = os.Getenv("DIGITALOCEAN_TOKEN")
	if cli.Token == "" {
		return cli, errors.New("DIGITALOCEAN_TOKEN not found in environment")
	}
	return cli, nil
}

func (c *Client) Get(u string, obj interface{}) error {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}

	err = c.DoRequest(req, "application/json", obj)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Del(u string) error {
	req, err := http.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	err = c.DoRequest(req, "application/x-www-form-urlencoded", nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Put(u string, v map[string]interface{}, obj interface{}) error {
	payload, err := json.Marshal(v)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", u, bytes.NewReader(payload))
	if err != nil {
		return err
	}

	err = c.DoRequest(req, "application/json", obj)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Post(u string, v map[string]interface{}, obj interface{}) error {
	payload, err := json.Marshal(v)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", u, bytes.NewReader(payload))
	if err != nil {
		return err
	}

	err = c.DoRequest(req, "application/json", obj)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DoRequest(req *http.Request, ct string, v interface{}) error {
	cl := &http.Client{}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", ct)
	resp, err := cl.Do(req)
	if err != nil {
		return err
	}
	err = c.Decode(resp, v)
	if err != nil {
		return err
	}
	return nil
}

// Decode parses the response.
func (c *Client) Decode(resp *http.Response, v interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// error code
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return errors.New(string(body))
	}

	if v != nil {
		err := json.Unmarshal(body, &v)
		if err != nil {
			return err
		}
	}
	return nil
}
