package dogo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func Test_ListKeys(t *testing.T) {
	setup(t)
	defer teardown()

	mux.HandleFunc("/account/keys", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")
		fmt.Fprint(w, listKeysExample)
	})

	want := &Key{
		ID:          1,
		FingerPrint: "2a:68:5c:e3:56:7e:a9:88:75:39:f4:f2:9d:2a:71:db",
		PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQDPeJTWZvNZiBxd9PBBS6W18wH6qjtP4FOu84ngCBszJlIF6NZoRPjtdOYz4gvgoS+HkfrlgjdVNupwxJT5uDwN example",
		Name:        "Example Key",
	}

	keys, err := client.ListKeys()
	assertEqual(t, err, nil)
	assertEqual(t, len(keys), 1)
	assertEqual(t, keys[0].ID, want.ID)
	assertEqual(t, keys[0].FingerPrint, want.FingerPrint)
	assertEqual(t, keys[0].PublicKey, want.PublicKey)
	assertEqual(t, keys[0].Name, want.Name)
}

func Test_GetKey_ByID(t *testing.T) {
	setup(t)
	defer teardown()

	want := &Key{
		ID:          3,
		FingerPrint: "a1:b2:c3",
		PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQDmfd/h4HxEHKQd6nsHYYYkn0mrfNE3QsxrLUD3vYnwb6dZIU6bNxPH4OHQ1lhevyUsw0WK4xi7dtNkJsb9lhtZ example",
		Name:        "Example Key",
	}

	mux.HandleFunc("/account/keys/3", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")
		fmt.Fprint(w, keyExample)
	})

	key, err := client.GetKey(want.ID)
	assertEqual(t, err, nil)
	assertEqual(t, key.ID, want.ID)
	assertEqual(t, key.FingerPrint, want.FingerPrint)
	assertEqual(t, key.PublicKey, want.PublicKey)
	assertEqual(t, key.Name, want.Name)
}

func Test_GetKey_ByFingerPrint(t *testing.T) {
	setup(t)
	defer teardown()

	want := &Key{
		ID:          3,
		FingerPrint: "a1:b2:c3",
		PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQDmfd/h4HxEHKQd6nsHYYYkn0mrfNE3QsxrLUD3vYnwb6dZIU6bNxPH4OHQ1lhevyUsw0WK4xi7dtNkJsb9lhtZ example",
		Name:        "Example Key",
	}

	mux.HandleFunc("/account/keys/a1:b2:c3", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")
		fmt.Fprint(w, keyExample)
	})

	key, err := client.GetKey(want.FingerPrint)
	assertEqual(t, err, nil)
	assertEqual(t, key.ID, want.ID)
	assertEqual(t, key.FingerPrint, want.FingerPrint)
	assertEqual(t, key.PublicKey, want.PublicKey)
	assertEqual(t, key.Name, want.Name)
}

func Test_CreateKey(t *testing.T) {
	setup(t)
	defer teardown()

	want := &Key{
		ID:          3,
		FingerPrint: "a1:b2:c3",
		PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQDmfd/h4HxEHKQd6nsHYYYkn0mrfNE3QsxrLUD3vYnwb6dZIU6bNxPH4OHQ1lhevyUsw0WK4xi7dtNkJsb9lhtZ example",
		Name:        "Example Key",
	}

	opts := &CreateKeyOpts{
		Name:      want.Name,
		PublicKey: want.PublicKey,
	}

	mux.HandleFunc("/account/keys", func(w http.ResponseWriter, r *http.Request) {
		v := new(CreateKeyOpts)
		json.NewDecoder(r.Body).Decode(v)

		assertEqual(t, r.Method, "POST")
		assertEqual(t, v, opts)

		fmt.Fprint(w, keyExample)
	})

	key, err := client.CreateKey(opts)
	assertEqual(t, err, nil)
	assertEqual(t, key.ID, want.ID)
	assertEqual(t, key.FingerPrint, want.FingerPrint)
	assertEqual(t, key.PublicKey, want.PublicKey)
	assertEqual(t, key.Name, want.Name)
}

func Test_UpdateKey_ByID(t *testing.T) {
	setup(t)
	defer teardown()

	want := &Key{
		ID:          3,
		FingerPrint: "a1:b2:c3",
		PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQDmfd/h4HxEHKQd6nsHYYYkn0mrfNE3QsxrLUD3vYnwb6dZIU6bNxPH4OHQ1lhevyUsw0WK4xi7dtNkJsb9lhtZ example",
		Name:        "New Name",
	}

	opts := &UpdateKeyOpts{
		Name: "New Name",
	}

	mux.HandleFunc("/account/keys/3", func(w http.ResponseWriter, r *http.Request) {
		v := new(UpdateKeyOpts)
		json.NewDecoder(r.Body).Decode(v)

		assertEqual(t, r.Method, "PUT")
		assertEqual(t, v, opts)
		fmt.Fprint(w, updateKeyExample)
	})

	key, err := client.UpdateKey(want.ID, opts)
	assertEqual(t, err, nil)
	assertEqual(t, key.ID, want.ID)
	assertEqual(t, key.FingerPrint, want.FingerPrint)
	assertEqual(t, key.PublicKey, want.PublicKey)
	assertEqual(t, key.Name, want.Name)
}

func Test_UpdateKey_ByFingerPrint(t *testing.T) {
	setup(t)
	defer teardown()

	want := &Key{
		ID:          3,
		FingerPrint: "a1:b2:c3",
		PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQDmfd/h4HxEHKQd6nsHYYYkn0mrfNE3QsxrLUD3vYnwb6dZIU6bNxPH4OHQ1lhevyUsw0WK4xi7dtNkJsb9lhtZ example",
		Name:        "New Name",
	}

	opts := &UpdateKeyOpts{
		Name: "New Name",
	}

	mux.HandleFunc("/account/keys/a1:b2:c3", func(w http.ResponseWriter, r *http.Request) {
		v := new(UpdateKeyOpts)
		json.NewDecoder(r.Body).Decode(v)

		assertEqual(t, r.Method, "PUT")
		assertEqual(t, v, opts)
		fmt.Fprint(w, updateKeyExample)
	})

	key, err := client.UpdateKey(want.FingerPrint, opts)
	assertEqual(t, err, nil)
	assertEqual(t, key.ID, want.ID)
	assertEqual(t, key.FingerPrint, want.FingerPrint)
	assertEqual(t, key.PublicKey, want.PublicKey)
	assertEqual(t, key.Name, want.Name)
}

func Test_DeleteKey_ByID(t *testing.T) {
	setup(t)
	defer teardown()

	id := 3

	mux.HandleFunc("/account/keys/3", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "DELETE")
		fmt.Fprint(w, "")

	})

	err := client.DeleteKey(id)
	assertEqual(t, err, nil)
}

func Test_DeleteKey_ByFingerPrint(t *testing.T) {
	setup(t)
	defer teardown()

	fingerprint := "a1:b2:c3"

	mux.HandleFunc("/account/keys/a1:b2:c3", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "DELETE")
		fmt.Fprint(w, "")

	})

	err := client.DeleteKey(fingerprint)
	assertEqual(t, err, nil)
}

var listKeysExample = `{
  "ssh_keys": [
    {
      "id": 1,
      "fingerprint": "2a:68:5c:e3:56:7e:a9:88:75:39:f4:f2:9d:2a:71:db",
      "public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQDPeJTWZvNZiBxd9PBBS6W18wH6qjtP4FOu84ngCBszJlIF6NZoRPjtdOYz4gvgoS+HkfrlgjdVNupwxJT5uDwN example",
      "name": "Example Key"
    }
  ],
  "meta": {
    "total": 1
  }
}
`

var keyExample = `{
  "ssh_key": {
    "id": 3,
    "fingerprint": "a1:b2:c3",
    "public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQDmfd/h4HxEHKQd6nsHYYYkn0mrfNE3QsxrLUD3vYnwb6dZIU6bNxPH4OHQ1lhevyUsw0WK4xi7dtNkJsb9lhtZ example",
    "name": "Example Key"
  }
}`

var updateKeyExample = `{
  "ssh_key": {
    "id": 3,
    "fingerprint": "a1:b2:c3",
    "public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQDmfd/h4HxEHKQd6nsHYYYkn0mrfNE3QsxrLUD3vYnwb6dZIU6bNxPH4OHQ1lhevyUsw0WK4xi7dtNkJsb9lhtZ example",
    "name": "New Name"
  }
}`
