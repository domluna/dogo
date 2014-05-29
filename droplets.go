package dogo

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Droplet respresents a digitalocean droplet.
type Droplet struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	ImageID          int       `json:"image_id"`
	SizeID           int       `json:"size_id"`
	RegionID         int       `json:"region_id"`
	BackupsActive    bool      `json:"backups_active"`
	IPAddress        string    `json:"ip_address"`
	PrivateIPAddress string    `json:"private_ip_address,omitempty"`
	Snapshots        []Image   `json:"snapshots,omitempty"`
	Locked           bool      `json:"locked"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
}

// GetDroplets returns all users droplets, active or otherwise.
func (c *Client) GetDroplets() ([]Droplet, error) {
	query := fmt.Sprintf(
		"%s?client_id=%s&api_key=%s",
		DropletsEndpoint,
		c.Auth.ClientID,
		c.Auth.APIKey,
	)

	body, err := Request(query)
	if err != nil {
		return nil, err
	}

	resp := struct {
		Status     string    `json:"status"`
		Droplets   []Droplet `json:"droplets"`
		ErrMessage string    `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Status == "ERROR" {
		return nil, fmt.Errorf("%s: %s", resp.Status, resp.ErrMessage)
	}

	return resp.Droplets, nil

}

// GetDroplet return an individual droplet based on the passed id.
func (c *Client) GetDroplet(id int) (Droplet, error) {
	query := fmt.Sprintf(
		"%s/%d/?client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		c.Auth.ClientID,
		c.Auth.APIKey,
	)

	body, err := Request(query)
	if err != nil {
		return Droplet{}, err
	}

	fmt.Println(string(body))

	resp := struct {
		Status     string  `json:"status"`
		Droplet    Droplet `json:"droplet"`
		ErrMessage string  `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return Droplet{}, err
	}

	if resp.Status == "ERROR" {
		return Droplet{}, fmt.Errorf("%s: %s", resp.Status, resp.ErrMessage)
	}

	return resp.Droplet, nil
}

// CreateDroplet creates a droplet based on based specs.
func (c *Client) CreateDroplet(d Droplet, keys []int, privateNet bool) (Droplet, error) {

	// Create a string of the key ids
	var keyStr string
	for _, k := range keys {
		ks := strconv.Itoa(k)
		keyStr += ks + ","
	}

	query := fmt.Sprintf(
		"%s/new?client_id=%s&api_key=%s&name=%s&size_id=%d&image_id=%d&region_id=%d&ssh_key_ids=%s&private_networking=%t&backups_enabled=%t",
		DropletsEndpoint,
		c.Auth.ClientID,
		c.Auth.APIKey,
		d.Name,
		d.SizeID,
		d.ImageID,
		d.RegionID,
		keyStr,
		privateNet,
		d.BackupsActive,
	)

	body, err := Request(query)
	if err != nil {
		return Droplet{}, err
	}

	resp := struct {
		Status     string  `json:"status"`
		Droplet    Droplet `json:"droplet"`
		ErrMessage string  `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return Droplet{}, err
	}

	if resp.Status == "ERROR" {
		return Droplet{}, fmt.Errorf("%s: %s", resp.Status, resp.ErrMessage)
	}

	return resp.Droplet, nil
}

// DestroyDroplet destroys a droplet. CAUTION - this is irreversible.
// There may be more appropriate options.
func (c *Client) DestroyDroplet(id int) error {
	query := fmt.Sprintf(
		"%s/%d/destroy/?client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		c.Auth.ClientID,
		c.Auth.APIKey,
	)

	body, err := Request(query)
	if err != nil {
		return err
	}

	resp := struct {
		Status     string `json:"status"`
		ErrMessage string `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Status == "ERROR" {
		return fmt.Errorf("%s: %s", resp.Status, resp.ErrMessage)
	}

	return nil
}

// ResizeDroplet droplet resizes a droplet. Sizes are based on
// the digitalocean sizes api.
func (c *Client) ResizeDroplet(id, sizeID int) error {
	query := fmt.Sprintf(
		"%s/%d/resize/?size_id=%d&client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		sizeID,
		c.Auth.ClientID,
		c.Auth.APIKey,
	)

	body, err := Request(query)
	if err != nil {
		return err
	}

	resp := struct {
		Status     string `json:"status"`
		ErrMessage string `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Status == "ERROR" {
		return fmt.Errorf("%s: %s", resp.Status, resp.ErrMessage)
	}

	return nil
}

// RebootDroplet reboots the a droplet. This is the preferred method
// to use if a server is not responding.
func (c *Client) RebootDroplet(id int) error {
	query := fmt.Sprintf(
		"%s/%d/reboot/?client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		c.Auth.ClientID,
		c.Auth.APIKey,
	)

	body, err := Request(query)
	if err != nil {
		return err
	}

	resp := struct {
		Status     string `json:"status"`
		ErrMessage string `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Status == "ERROR" {
		return fmt.Errorf("%s: %s", resp.Status, resp.ErrMessage)
	}

	return nil
}

// RebootDroplet rebuilds a droplet with a default image. This can be
// useful if you want to use a different image but keep the ip address
// of the droplet.
func (c *Client) RebuildDroplet(id, imageID int) error {
	query := fmt.Sprintf(
		"%s/%d/rebuild/?image_id=%d&client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		imageID,
		c.Auth.ClientID,
		c.Auth.APIKey,
	)

	body, err := Request(query)
	if err != nil {
		return err
	}

	resp := struct {
		Status     string `json:"status"`
		ErrMessage string `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Status == "ERROR" {
		return fmt.Errorf("%s: %s", resp.Status, resp.ErrMessage)
	}

	return nil
}

// StopDroplet powers off a running droplet, the droplet will remain
// in your account.
func (c *Client) StopDroplet(id int) error {
	query := fmt.Sprintf(
		"%s/%d/power_off/?client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		c.Auth.ClientID,
		c.Auth.APIKey,
	)

	body, err := Request(query)
	if err != nil {
		return err
	}

	resp := struct {
		Status     string `json:"status"`
		ErrMessage string `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Status == "ERROR" {
		return fmt.Errorf("%s: %s", resp.Status, resp.ErrMessage)
	}

	return nil
}

// StartDroplet powers on a powered off droplet.
func (c *Client) StartDroplet(id int) error {
	query := fmt.Sprintf(
		"%s/%d/power_on/?client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		c.Auth.ClientID,
		c.Auth.APIKey,
	)

	body, err := Request(query)
	if err != nil {
		return err
	}

	resp := struct {
		Status     string `json:"status"`
		ErrMessage string `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Status == "ERROR" {
		return fmt.Errorf("%s: %s", resp.Status, resp.ErrMessage)
	}

	return nil
}

// SnapshotDroplet allows you to take a snapshot of a droplet once it is
// powered off. Be aware this may reboot the droplet.
func (c *Client) SnapshotDroplet(id int, name string) error {
	query := fmt.Sprintf(
		"%s/%d/snapshot/?name=%s&client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		name,
		c.Auth.ClientID,
		c.Auth.APIKey,
	)

	body, err := Request(query)
	if err != nil {
		return err
	}

	resp := struct {
		Status     string `json:"status"`
		ErrMessage string `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Status == "ERROR" {
		return fmt.Errorf("%s: %s", resp.Status, resp.ErrMessage)
	}

	return nil
}

// RestoreDroplet allows you to restore a droplet to a previous image
// or snapshot. This will be a mirror copy of the image or snapshot to
// your droplet.
func (c *Client) RestoreDroplet(id, imageID int) error {
	query := fmt.Sprintf(
		"%s/%d/restore/?image_id=%d&client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		imageID,
		c.Auth.ClientID,
		c.Auth.APIKey,
	)

	body, err := Request(query)
	if err != nil {
		return err
	}

	resp := struct {
		Status     string `json:"status"`
		ErrMessage string `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Status == "ERROR" {
		return fmt.Errorf("%s: %s", resp.Status, resp.ErrMessage)
	}

	return nil
}
