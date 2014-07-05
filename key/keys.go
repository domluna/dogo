package key

import (
	"fmt"

	"github.com/domluna/dogo/digitalocean"
)

const (
	Endpoint = digitalocean.BaseURL + "/accounts/keys"
)

// Key represents DigitalOcean ssh key.
type Key struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Fingerprint  string `json:"fingerprint,omitempty"`
	SSHPublicKey string `json:"public_key,omitempty"`
}

type Keys []Key

type Client struct {
	client *digitalocean.Client
}

// GetKeys retrieves all the users current ssh keys.
func (c *Client) GetAll() (Keys, error) {
	s := struct {
		Keys              `json:"ssh_keys,omitempty"`
		digitalocean.Meta `json:"meta,omitempty"`
	}{}
	err := c.client.Get(Endpoint, &s)
	if err != nil {
		return s.Keys, err
	}
	return s.Keys, nil
}

// GetKey returns the public key, this includes the public key.
func (c *Client) Get(v interface{}) (Key, error) {
	u := fmt.Sprintf("%s/%v", Endpoint, v)
	s := struct {
		Key `json:"ssh_keys,omitempty"`
	}{}
	err := c.client.Get(u, &s)
	if err != nil {
		return s.Key, err
	}
	return s.Key, nil

}

// CreateKey adds an ssh key to the user account.
func (c *Client) Create(name string, pk []byte) (Key, error) {
	s := struct {
		Key `json:"ssh_keys,omitempty"`
	}{}
	payload := map[string]interface{}{
		"name":       name,
		"public_key": string(pk),
	}
	err := c.client.Post(Endpoint, payload, &s)
	if err != nil {
		return s.Key, err
	}
	return s.Key, nil
}

// DestroyKey destroys the ssh key with
// passed id from user account.
func (c *Client) Update(v interface{}, name string) (Key, error) {
	u := fmt.Sprintf("%s/%v", Endpoint, v)
	s := struct {
		Key `json:"ssh_keys,omitempty"`
	}{}
	payload := map[string]interface{}{
		"name": name,
	}
	err := c.client.Post(u, payload, &s)
	if err != nil {
		return s.Key, err
	}
	return s.Key, nil
}

// Destroy destroys the ssh key with
// passed id from user account.
func (c *Client) Destroy(v interface{}) error {
	u := fmt.Sprintf("%s/%v", Endpoint, v)
	err := c.client.Del(u)
	if err != nil {
		return err
	}
	return nil
}
