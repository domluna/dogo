package dogo

import (
	"fmt"
	"time"
)

// Kernel is a DigitalOcean Kernel.
type Kernel struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
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
	Regions      []string  `json:"regions,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

type Images []Image

type ImageClient struct {
	client Client
}

// GetMyImages gets all custom images/snapshots.
func (ic *ImageClient) GetAll() (Images, error) {
	s := struct {
		Images `json:"images,omitempty"`
		Meta   `json:"meta,omitempty"`
	}{}
	err := ic.client.Get(ImagesEndpoint, &s)
	if err != nil {
		return s.Images, err
	}
	return s.Images, nil
}

// GetMyImages gets all custom images/snapshots.
func (ic *ImageClient) Get(v interface{}) (Image, error) {
	u := fmt.Sprintf("%s/%v", ImagesEndpoint, v)
	s := struct {
		Image `json:"images,omitempty"`
		Meta  `json:"meta,omitempty"`
	}{}
	err := ic.client.Get(u, &s)
	if err != nil {
		return s.Image, err
	}
	return s.Image, nil
}

// GetMyImages gets all custom images/snapshots.
func (ic *ImageClient) Delete(id int) error {
	u := fmt.Sprintf("%s/%d", ImagesEndpoint, id)
	err := ic.client.Del(u)
	if err != nil {
		return err
	}
	return nil
}

// GetMyImages gets all custom images/snapshots.
func (ic *ImageClient) Update(id int, name string) (Image, error) {
	u := fmt.Sprintf("%s/%d", ImagesEndpoint, id)
	s := struct {
		Image `json:"image,omitempty"`
	}{}
	payload := map[string]interface{}{
		"name": name,
	}
	err := ic.client.Put(u, payload, &s)
	if err != nil {
		return s.Image, err
	}
	return s.Image, nil
}

func (ic *ImageClient) DoAction(id int, params map[string]interface{}) error {
	u := fmt.Sprintf("%s/%d", ImagesEndpoint, id)
	err := ic.client.Post(u, params, nil)
	if err != nil {
		return err
	}
	return nil
}
