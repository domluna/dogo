package dogo

import (
	"net/url"
)

// SSHKey represents DigitalOcean ssh key.
type SSHKey struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	SSHPublicKey string `json:"ssh_pub_key,omitempty"`
}

// GetSSHKeys retrieves all the users current ssh keys.
func (c *Client) GetSSHKeys() ([]SSHKey, error) {
	resp, err := c.send(SSHKeysEndpoint, nil, nil)
	if err != nil {
		return resp.SSHKeys, err
	}
	return resp.SSHKeys, nil
}

// AddSSHKey adds an ssh key to the user account.
func (c *Client) AddSSHKey(name string, publicKey []byte) (SSHKey, error) {
	ks := url.QueryEscape(string(publicKey))
	resp, err := c.send(SSHKeysEndpoint, "new", Params{
		"name":        name,
		"ssh_pub_key": ks,
	})
	if err != nil {
		return *resp.SSHKey, err
	}
	return *resp.SSHKey, nil
}

// GetSSHKey returns the public key, this includes the public key.
func (c *Client) GetSSHKey(id int) (SSHKey, error) {
	resp, err := c.send(SSHKeysEndpoint, id, nil)
	if err != nil {
		return *resp.SSHKey, err
	}
	return *resp.SSHKey, nil
}

// DestroySSHKey destroys the ssh key with
// passed id from user account.
func (c *Client) DestroySSHKey(id int) error {
	_, err := c.send(SSHKeysEndpoint, id, nil)
	if err != nil {
		return err
	}
	return nil
}
