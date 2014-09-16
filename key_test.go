package digitalocean

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_ListKeys(t *testing.T) {
	setup(t)
	defer teardown()

	mux.HandleFunc("/account/keys", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, listKeysExample)
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
		FingerPrint: "70:8a:81:98:9c:60:d9:d2:d4:82:c7:97:bf:95:4f:09",
		PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQDmfd/h4HxEHKQd6nsHYYYkn0mrfNE3QsxrLUD3vYnwb6dZIU6bNxPH4OHQ1lhevyUsw0WK4xi7dtNkJsb9lhtZ example",
		Name:        "Example Key",
	}

	u := fmt.Sprintf("/account/keys/%d", want.ID)
	mux.HandleFunc(u, func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, getKeyExample)
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
		FingerPrint: "70:8a:81:98:9c:60:d9:d2:d4:82:c7:97:bf:95:4f:09",
		PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQDmfd/h4HxEHKQd6nsHYYYkn0mrfNE3QsxrLUD3vYnwb6dZIU6bNxPH4OHQ1lhevyUsw0WK4xi7dtNkJsb9lhtZ example",
		Name:        "Example Key",
	}

	u := fmt.Sprintf("/account/keys/%s", want.FingerPrint)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, getKeyExample)
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
		FingerPrint: "70:8a:81:98:9c:60:d9:d2:d4:82:c7:97:bf:95:4f:09",
		PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQDmfd/h4HxEHKQd6nsHYYYkn0mrfNE3QsxrLUD3vYnwb6dZIU6bNxPH4OHQ1lhevyUsw0WK4xi7dtNkJsb9lhtZ example",
		Name:        "Example Key",
	}

	mux.HandleFunc("/account/keys", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, createKeyExample)
	})

	key, err := client.CreateKey(want.Name, want.PublicKey)
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
		FingerPrint: "70:8a:81:98:9c:60:d9:d2:d4:82:c7:97:bf:95:4f:09",
		PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQDmfd/h4HxEHKQd6nsHYYYkn0mrfNE3QsxrLUD3vYnwb6dZIU6bNxPH4OHQ1lhevyUsw0WK4xi7dtNkJsb9lhtZ example",
		Name:        "New Name",
	}

	u := fmt.Sprintf("/account/keys/%d", want.ID)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, updateKeyExample)
	})

	key, err := client.UpdateKey(want.ID, want.Name)
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
		FingerPrint: "70:8a:81:98:9c:60:d9:d2:d4:82:c7:97:bf:95:4f:09",
		PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQDmfd/h4HxEHKQd6nsHYYYkn0mrfNE3QsxrLUD3vYnwb6dZIU6bNxPH4OHQ1lhevyUsw0WK4xi7dtNkJsb9lhtZ example",
		Name:        "New Name",
	}

	u := fmt.Sprintf("/account/keys/%s", want.FingerPrint)
	mux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, updateKeyExample)
	})

	key, err := client.UpdateKey(want.FingerPrint, want.Name)
	assertEqual(t, err, nil)
	assertEqual(t, key.ID, want.ID)
	assertEqual(t, key.FingerPrint, want.FingerPrint)
	assertEqual(t, key.PublicKey, want.PublicKey)
	assertEqual(t, key.Name, want.Name)
}

func Test_DeleteKey(t *testing.T) {
	setup(t)
	defer teardown()

	want := &Key{
		ID:          3,
		FingerPrint: "70:8a:81:98:9c:60:d9:d2:d4:82:c7:97:bf:95:4f:09",
		PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQDmfd/h4HxEHKQd6nsHYYYkn0mrfNE3QsxrLUD3vYnwb6dZIU6bNxPH4OHQ1lhevyUsw0WK4xi7dtNkJsb9lhtZ example",
		Name:        "New Name",
	}

	u := fmt.Sprintf("/account/keys/%d", want.ID)
	mux.HandleFunc(u, func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(204)
	})

	err := client.DeleteKey(want.ID)
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

var getKeyExample = `{
  "ssh_key": {
    "id": 3,
    "fingerprint": "70:8a:81:98:9c:60:d9:d2:d4:82:c7:97:bf:95:4f:09",
    "public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQDmfd/h4HxEHKQd6nsHYYYkn0mrfNE3QsxrLUD3vYnwb6dZIU6bNxPH4OHQ1lhevyUsw0WK4xi7dtNkJsb9lhtZ example",
    "name": "Example Key"
  }
}`

var createKeyExample = getKeyExample
var updateKeyExample = `{
  "ssh_key": {
    "id": 3,
    "fingerprint": "70:8a:81:98:9c:60:d9:d2:d4:82:c7:97:bf:95:4f:09",
    "public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQDmfd/h4HxEHKQd6nsHYYYkn0mrfNE3QsxrLUD3vYnwb6dZIU6bNxPH4OHQ1lhevyUsw0WK4xi7dtNkJsb9lhtZ example",
    "name": "New Name"
  }
}`
