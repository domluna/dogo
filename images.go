package dogo

import (
	"fmt"
	"time"
)

// Kernel is a DigitalOcean Kernel.
type Kernel struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name",omitempty`
	Version string `json:"version,omitempty"`
}

// Snapshot is a DigitalOcean Snapshot/Backup.
type Snapshot struct {
	ID        int      `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Dist      string   `json:"distribution,omitempty"`
	Slug      string   `json:"slug,omitempty,omitempty"`
	Public    bool     `json:"public,omitempty"`
	Regions   []string `json:"regions,omitempty"`
	ActionIDs []int    `json:"action_ids,omitempty"`
}

// Representation for a DigitalOcean Image.
type Image struct {
	ID           int       `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Distribution string    `json:"distribution,omitempty"`
	Slug         string    `json:"slug,omitempty,omitempty"`
	Public       bool      `json:"public,omitempty"`
	Regions      []string  `json"regions,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

type Images []Image

type ImageClient struct {
	Client
}

// GetMyImages gets all custom images/snapshots.
func (ic *ImageClient) GetImages() (Images, error) {
	var i Images
	req, err := ic.Client.Get(ImagesEndpoint)
	if err != nil {
		return i, err
	}

	err = ic.Client.DoRequest(req, &i)
	if err != nil {
		return i, err
	}
	return i, nil
}

// GetMyImages gets all custom images/snapshots.
func (ic *ImageClient) GetImage(v interface{}) (Image, error) {
	u := fmt.Sprintf("%s/%v", ImagesEndpoint, v)
	var i Image
	req, err := ic.Client.Get(u)
	if err != nil {
		return i, err
	}

	err = ic.Client.DoRequest(req, &i)
	if err != nil {
		return i, err
	}
	return i, nil
}

// GetMyImages gets all custom images/snapshots.
func (ic *ImageClient) DelImage(id int) error {
	u := fmt.Sprintf("%s/%d", ImagesEndpoint, id)
	req, err := ic.Client.Del(u)
	if err != nil {
		return err
	}
	err = ic.Client.DoRequest(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// GetMyImages gets all custom images/snapshots.
func (ic *ImageClient) UpdateImage(id int, name string) (Image, error) {
	u := fmt.Sprintf("%s/%d", ImagesEndpoint, id)
	var i Image
	payload := map[string]interface{}{
		"name": name,
	}
	req, err := ic.Client.Put(u, payload)
	if err != nil {
		return i, err
	}

	err = ic.Client.DoRequest(req, &i)
	if err != nil {
		return i, err
	}
	return i, nil
}
