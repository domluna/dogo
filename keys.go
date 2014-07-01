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

type KeyClient struct {
	client Client
}

// GetKeys retrieves all the users current ssh keys.
func (kc *KeyClient) GetAll() (Keys, error) {
	s := struct {
		Keys `json:"ssh_keys,omitempty"`
		Meta `json:"meta,omitempty"`
	}{}
	err := kc.client.Get(KeysEndpoint, &s)
	if err != nil {
		return s.Keys, err
	}
	return s.Keys, nil
}

// GetKey returns the public key, this includes the public key.
func (kc *KeyClient) Get(v interface{}) (Key, error) {
	u := fmt.Sprintf("%s/%v", KeysEndpoint, v)
	s := struct {
		Key `json:"ssh_keys,omitempty"`
	}{}
	err := kc.client.Get(u, &s)
	if err != nil {
		return s.Key, err
	}
	return s.Key, nil

}

// CreateKey adds an ssh key to the user account.
func (kc *KeyClient) Create(name string, pk []byte) (Key, error) {
	s := struct {
		Key `json:"ssh_keys,omitempty"`
	}{}
	payload := map[string]interface{}{
		"name":       name,
		"public_key": string(pk),
	}
	err := kc.client.Post(KeysEndpoint, payload, &s)
	if err != nil {
		return s.Key, err
	}
	return s.Key, nil
}

// DestroyKey destroys the ssh key with
// passed id from user account.
func (kc *KeyClient) Update(v interface{}, name string) (Key, error) {
	u := fmt.Sprintf("%s/%v", KeysEndpoint, v)
	s := struct {
		Key `json:"ssh_keys,omitempty"`
	}{}
	payload := map[string]interface{}{
		"name": name,
	}
	err := kc.client.Post(u, payload, &s)
	if err != nil {
		return s.Key, err
	}
	return s.Key, nil
}

// DestroyKey destroys the ssh key with
// passed id from user account.
func (kc *KeyClient) Destroy(v interface{}) error {
	u := fmt.Sprintf("%s/%v", KeysEndpoint, v)
	err := kc.client.Del(u)
	if err != nil {
		return err
	}
	return nil
}
