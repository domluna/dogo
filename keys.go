package dogo

import (
	"fmt"
)

// Key represents DigitalOcean ssh key.
type Key struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Fingerprint  string `json:"fingerprint,omitempty"`
	SSHPublicKey string `json:"public_key,omitempty"`
}

type Keys []Key

// GetKeys retrieves all the users current ssh keys.
func (c *Client) GetKeys() (Keys, error) {
	var k Keys
	err := c.Get(KeysEndpoint, k)
	if err != nil {
		return k, err
	}
	return k, nil
}

// GetKey returns the public key, this includes the public key.
func (c *Client) GetKey(v interface{}) (Key, error) {
	u := fmt.Sprintf("%s/%v", KeysEndpoint, v)
	var k Key
	err := c.Get(u, k)
	if err != nil {
		return k, err
	}
	return k, nil

}

// CreateKey adds an ssh key to the user account.
func (c *Client) CreateKey(name string, pk []byte) (Key, error) {
	var k Key
	payload := map[string]interface{}{
		"name":       name,
		"public_key": string(pk),
	}
	err := c.Post(KeysEndpoint, payload, k)
	if err != nil {
		return k, err
	}
	return k, nil
}

// DestroyKey destroys the ssh key with
// passed id from user account.
func (c *Client) UpdateKey(v interface{}, name string) (Key, error) {
	u := fmt.Sprintf("%s/%v", KeysEndpoint, v)
	var k Key
	payload := map[string]interface{}{
		"name": name,
	}
	err := c.Put(u, payload, k)
	if err != nil {
		return k, err
	}
	return k, nil
}

// DestroyKey destroys the ssh key with
// passed id from user account.
func (c *Client) DestroyKey(v interface{}) error {
	u := fmt.Sprintf("%s/%v", KeysEndpoint, v)
	err := c.Del(u)
	if err != nil {
		return err
	}
	return nil
}
