package dogo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// SSHKey represents DigitalOcean ssh key.
type SSHKey struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	SSHPublicKey string `json:"ssh_pub_key"`
}

// GetSSHKeys retrieves all the users current ssh keys.
func (c *Client) GetSSHKeys() ([]SSHKey, error) {
	query := fmt.Sprintf("%s?client_id=%s&api_key=%s",
		KeysEndpoint,
		c.Auth.ClientID,
		c.Auth.APIKey)

	body, err := Request(query)
	if err != nil {
		return nil, err
	}

	resp := struct {
		Status  string   `json:"status"`
		SSHKeys []SSHKey `json:"ssh_keys"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Status != "OK" {
		return nil, errors.New("Error retrieving ssh keys")
	}

	return resp.SSHKeys, nil
}

// AddSSHKey adds an ssh key to the user account.
func (c *Client) AddSSHKey(name string, publicKey []byte) (SSHKey, error) {
	ks := url.QueryEscape(string(publicKey))
	query := fmt.Sprintf("%s/new/?name=%s&ssh_pub_key=%s&client_id=%s&api_key=%s",
		KeysEndpoint,
		name,
		ks,
		c.Auth.ClientID,
		c.Auth.APIKey)

	body, err := Request(query)
	if err != nil {
		return SSHKey{}, err
	}

	resp := struct {
		Status string `json:"status"`
		SSHKey SSHKey `json:"ssh_key"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return SSHKey{}, err
	}

	if resp.Status != "OK" {
		return SSHKey{}, errors.New("Error adding key, might be something wrong with the endpoint")
	}

	return resp.SSHKey, nil
}

// GetSSHKey returns the public key, this includes the public key.
func (c *Client) GetSSHKey(id int) (SSHKey, error) {
	query := fmt.Sprintf("%s/%d/?client_id=%s&api_key=%s",
		KeysEndpoint,
		id,
		c.Auth.ClientID,
		c.Auth.APIKey)

	body, err := Request(query)
	if err != nil {
		return SSHKey{}, err
	}

	resp := struct {
		Status string `json:"status"`
		SSHKey SSHKey `json:"ssh_key"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return SSHKey{}, err
	}

	if resp.Status != "OK" {
		return SSHKey{}, errors.New("Invalid ssh key id")
	}

	return resp.SSHKey, nil
}

// DestroySSHKey destroys the ssh key with
// passed id from user account.
func (c *Client) DestroySSHKey(id int) error {
	query := fmt.Sprintf("%s/%d/destroy/?client_id=%s&api_key=%s",
		KeysEndpoint,
		id,
		c.Auth.ClientID,
		c.Auth.APIKey)

	body, err := Request(query)
	if err != nil {
		return err
	}

	resp := struct {
		Status string `json:"status"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Status != "OK" {
		errors.New("Did not remove ssh key, are you sure the id is correct?")
	}

	return nil
}
