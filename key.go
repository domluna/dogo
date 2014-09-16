package digitalocean

import (
	"fmt"
)

const KeyEndpoint = "account/keys"

// Key represents DigitalOcean ssh key.
type Key struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	FingerPrint string `json:"fingerprint,omitempty"`
	PublicKey   string `json:"public_key,omitempty"`
}

// Keys is a list of Key.
type Keys []*Key

type CreateKeyOpts struct {
        Name string `json:"name"`
        PublicKey string `json:"public_key"`
}

type UpdateKeyOpts struct {
        Name string `json:"name"`
}

// ListKeys retrieves all the users current ssh keys.
func (c *Client) ListKeys() (Keys, error) {
	s := struct {
		Keys `json:"ssh_keys,omitempty"`
	}{}
	err := c.get(KeyEndpoint, &s)
	if err != nil {
		return nil, err
	}
	return s.Keys, nil
}

// GetKey returns the public key, this includes the public key.
func (c *Client) GetKey(v interface{}) (*Key, error) {
	u := fmt.Sprintf("%s/%v", KeyEndpoint, v)
	s := struct {
		Key `json:"ssh_key,omitempty"`
	}{}
	err := c.get(u, &s)
	if err != nil {
		return nil, err
	}
	return &s.Key, nil
}

// CreateKey adds an ssh key to the user account.
func (c *Client) CreateKey(opts *CreateKeyOpts) (*Key, error) {
	s := struct {
		Key `json:"ssh_key,omitempty"`
	}{}
	err := c.post(KeyEndpoint, opts, &s)
	if err != nil {
		return nil, err
	}
	return &s.Key, nil
}

// UpdateKey updates an SSH Key. Can use the ID or FINGERPRINT of the key.
func (c *Client) UpdateKey(v interface{}, opts *UpdateKeyOpts) (*Key, error) {
	u := fmt.Sprintf("%s/%v", KeyEndpoint, v)
	s := struct {
		Key `json:"ssh_key,omitempty"`
	}{}
	err := c.put(u, opts, &s)
	if err != nil {
		return nil, err
	}
	return &s.Key, nil
}

// DestroyKey destroys an SSH Key. Can use the ID or FINGERPRINT of the key.
func (c *Client) DeleteKey(v interface{}) error {
	u := fmt.Sprintf("%s/%v", KeyEndpoint, v)
	err := c.delete(u)
	if err != nil {
		return err
	}
	return nil
}
