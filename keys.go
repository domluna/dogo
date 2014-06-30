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
	Client
}

// GetKeys retrieves all the users current ssh keys.
func (kc *KeyClient) GetKeys() (Keys, error) {
	var k Keys
	req, err := kc.Client.Get(KeysEndpoint)
	if err != nil {
		return k, err
	}

	err = kc.Client.DoRequest(req, &k)
	if err != nil {
		return k, err
	}
	return k, nil
}

// GetKey returns the public key, this includes the public key.
func (kc *KeyClient) GetKey(v interface{}) (Key, error) {
	u := fmt.Sprintf("%s/%v", KeysEndpoint, v)
	var k Key
	req, err := kc.Client.Get(u)
	if err != nil {
		return k, err
	}
	err = kc.Client.DoRequest(req, &k)
	if err != nil {
		return k, err
	}
	return k, nil

}

// CreateKey adds an ssh key to the user account.
func (kc *KeyClient) CreateKey(name string, pk []byte) (Key, error) {
	var k Key
	payload := map[string]interface{}{
		"name":       name,
		"public_key": string(pk),
	}
	req, err := kc.Client.Post(KeysEndpoint, payload)
	if err != nil {
		return k, err
	}
	err = kc.Client.DoRequest(req, &k)
	if err != nil {
		return k, err
	}
	return k, nil
}

// DestroyKey destroys the ssh key with
// passed id from user account.
func (kc *KeyClient) UpdateKey(v interface{}, name string) (Key, error) {
	u := fmt.Sprintf("%s/%v", KeysEndpoint, v)
	var k Key
	payload := map[string]interface{}{
		"name": name,
	}
	req, err := kc.Client.Put(u, payload)
	if err != nil {
		return k, err
	}
	err = kc.Client.DoRequest(req, &k)
	if err != nil {
		return k, err
	}
	return k, nil
}

// DestroyKey destroys the ssh key with
// passed id from user account.
func (kc *KeyClient) DestroyKey(v interface{}) error {
	u := fmt.Sprintf("%s/%v", KeysEndpoint, v)
	req, err := kc.Client.Del(u)
	if err != nil {
		return err
	}

	err = kc.Client.DoRequest(req, nil)
	if err != nil {
		return err
	}
	return nil
}
