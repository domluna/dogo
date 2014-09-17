package dogo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func Test_ListImages(t *testing.T) {
	setup(t)
	defer teardown()

	want := struct {
		Images `json:"images"`
	}{}
	json.NewDecoder(strings.NewReader(listImagesExample)).Decode(&want)

	mux.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")
		fmt.Fprint(w, listImagesExample)
	})

	images, err := client.ListImages()

	assertEqual(t, err, nil)
	assertEqual(t, len(images), 2)
	assertEqual(t, images, want.Images)
}

func Test_GetImage_ByID(t *testing.T) {
	setup(t)
	defer teardown()

	want := struct {
		*Image `json:"image"`
	}{}
	json.NewDecoder(strings.NewReader(getImageExample)).Decode(&want)

	mux.HandleFunc("/images/1", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")
		fmt.Fprint(w, getImageExample)
	})

	image, err := client.GetImage(1)

	assertEqual(t, err, nil)
	assertEqual(t, image, want.Image)
}

func Test_GetImage_BySlug(t *testing.T) {
	setup(t)
	defer teardown()

	want := struct {
		*Image `json:"image"`
	}{}
	json.NewDecoder(strings.NewReader(getImageExample)).Decode(&want)

	mux.HandleFunc("/images/ubuntu1404", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")
		fmt.Fprint(w, getImageExample)
	})

	image, err := client.GetImage("ubuntu1404")

	assertEqual(t, err, nil)
	assertEqual(t, image, want.Image)
}

func Test_DeleteImage(t *testing.T) {
	setup(t)
	defer teardown()

	mux.HandleFunc("/images/2", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "DELETE")
		fmt.Fprint(w, "")
	})

	err := client.DeleteImage(2)

	assertEqual(t, err, nil)
}

func Test_UpdateImage(t *testing.T) {
	setup(t)
	defer teardown()

	opts := &UpdateImageOpts{
		Name: "New Image Name",
	}

	want := struct {
		*Image `json:"image"`
	}{}

	json.NewDecoder(strings.NewReader(updateImageExample)).Decode(&want)

	mux.HandleFunc("/images/2", func(w http.ResponseWriter, r *http.Request) {
		v := new(UpdateImageOpts)
		json.NewDecoder(r.Body).Decode(v)
		assertEqual(t, r.Method, "PUT")
		assertEqual(t, v, opts)
		fmt.Fprint(w, updateImageExample)
	})

	image, err := client.UpdateImage(2, opts)

	assertEqual(t, err, nil)
	assertEqual(t, image, want.Image)
}

var listImagesExample = `{
  "images": [
    {
      "id": 119192817,
      "name": "Ubuntu 13.04",
      "distribution": "ubuntu",
      "slug": "ubuntu1304",
      "public": true,
      "regions": [
        "nyc1"
      ],
      "created_at": "2014-09-05T02:02:08Z"
    },
    {
      "id": 449676376,
      "name": "Ubuntu 13.04",
      "distribution": "ubuntu",
      "slug": "ubuntu1404",
      "public": true,
      "regions": [
        "nyc1"
      ],
      "created_at": "2014-09-05T02:02:08Z"
    }
  ]
}`

var getImageExample = `{
  "image": {
    "id": 1,
    "name": "Ubuntu 13.04",
    "distribution": "ubuntu",
    "slug": "ubuntu1404",
    "public": false,
    "regions": [
      "region--1"
    ],
    "created_at": "2014-09-05T02:02:08Z"
  }
}`

var updateImageExample = `{
  "image": {
    "id": 2,
    "name": "New Image Name",
    "distribution": null,
    "slug": null,
    "public": false,
    "regions": [
      "region--3"
    ],
    "created_at": "2014-09-05T02:02:09Z"
  }
}`
