package dogo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func Test_CreateDroplet(t *testing.T) {
	setup(t)
	defer teardown()

	want := &Droplet{
		ID:     25,
		Name:   "My-Droplet",
		Memory: 512,
		VCPUS:  1,
		Disk:   20,
		Region: &Region{
			Slug: "nyc1",
			Name: "New York",
			Sizes: []string{
				"1gb",
				"512mb",
			},
			Available: true,
			Features: []string{
				"virtio",
				"private_networking",
				"backups",
				"ipv6",
			},
		},
		Image: &Image{
			ID:           449676389,
			Name:         "Ubuntu 13.04",
			Distribution: "ubuntu",
			Slug:         "",
			Public:       true,
			Regions: []string{
				"nyc1",
			},
			CreatedAt: "2014-09-05T02:02:07Z",
		},
		Size: &Size{
			Slug:         "512mb",
			Transfer:     1,
			PriceMonthly: 5.0,
			PriceHourly:  0.00744,
		},
		Locked: false,
		Status: "new",
		Networks: &Networks{
			[]*V4{
				&V4{
					IPAddress: "127.0.0.20",
					Netmask:   "255.255.255.0",
					Gateway:   "127.0.0.21",
					Type:      "public",
				},
				&V4{
					IPAddress: "10.130.0.0",
					Netmask:   "255.255.255.0",
					Gateway:   "127.0.0.21",
					Type:      "private",
				}},
			[]*V6{
				&V6{
					IPAddress: "2001::14",
					Cidr:      124,
					Gateway:   "2400:6180:0000:00D0:0000:0000:0009:7000",
					Type:      "public",
				},
			},
		},
		Kernel: &Kernel{
			ID:      485432972,
			Name:    "Ubuntu 14.04 x64 vmlinuz-3.13.0-24-generic (1221)",
			Version: "3.13.0-24-generic",
		},
		CreatedAt: "2014-09-05T02:02:07Z",
		Features: []string{
			"virtio",
		},
		BackupIDs:   []int{},
		SnapshotIDs: []int{},
	}

	opts := &CreateDropletOpts{
		Name:   "My-Droplet",
		Region: "nyc1",
		Size:   "512mb",
		Image:  449676389,
		Keys: []string{
			"123",
			"a1:b2:c3",
		},
		Backups:           true,
		IPV6:              false,
		PrivateNetworking: true,
		UserData:          "what user data?",
	}

	mux.HandleFunc("/droplets", func(w http.ResponseWriter, r *http.Request) {
		v := new(CreateDropletOpts)
		json.NewDecoder(r.Body).Decode(v)
		assertEqual(t, r.Method, "POST")
		assertEqual(t, v, opts)
		fmt.Fprint(w, createDropletExample)
	})

	droplet, err := client.CreateDroplet(opts)
	assertEqual(t, err, nil)
	assertEqual(t, droplet, want)
	assertEqual(t, droplet.IPV4Addr(), "127.0.0.20")
	assertEqual(t, droplet.IPV6Addr(), "2001::14")
	assertEqual(t, droplet.SizeSlug(), "512mb")
	assertEqual(t, droplet.ImageID(), 449676389)
	assertEqual(t, droplet.KernelName(), "Ubuntu 14.04 x64 vmlinuz-3.13.0-24-generic (1221)")
	assertEqual(t, droplet.RegionSlug(), "nyc1")
	assertEqual(t, droplet.ImageName(), "Ubuntu 13.04")
}
func Test_DeleteDroplet(t *testing.T) {
	setup(t)
	defer teardown()

	mux.HandleFunc("/droplets/20", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "DELETE")
		fmt.Fprint(w, "")
	})

	err := client.DeleteDroplet(20)
	assertEqual(t, err, nil)
}
func Test_GetDroplet(t *testing.T) {
	setup(t)
	defer teardown()

	want := &Droplet{
		ID: 20,
		Networks: &Networks{
			V4: []*V4{
				&V4{
					IPAddress: "127.0.0.20",
					Netmask:   "255.255.255.0",
					Gateway:   "127.0.0.21",
					Type:      "public",
				},
			},
			V6: []*V6{
				&V6{
					IPAddress: "2001::14",
					Cidr:      124,
					Gateway:   "2400:6180:0000:00D0:0000:0000:0009:7000",
					Type:      "public",
				},
			},
		},
		Features:    []string{},
		BackupIDs:   []int{11},
		SnapshotIDs: []int{20},
	}

	mux.HandleFunc("/droplets/20", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")
		fmt.Fprint(w, getDropletExample)
	})

	droplet, err := client.GetDroplet(20)
	assertEqual(t, err, nil)
	assertEqual(t, droplet, want)

}
func Test_ListDroplets(t *testing.T) {
	setup(t)
	defer teardown()

	want := Droplets{
		&Droplet{
			ID: 20,
		},
		&Droplet{
			ID: 30,
		},
		&Droplet{
			ID: 42,
		},
	}

	mux.HandleFunc("/droplets", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")
		fmt.Fprint(w, listDropletsExample)
	})

	droplets, err := client.ListDroplets()
	assertEqual(t, err, nil)
	assertEqual(t, droplets, want)
}
func Test_ListKernels(t *testing.T) {
	setup(t)
	defer teardown()
	want := Kernels{
		&Kernel{
			ID:      61833229,
			Name:    "Ubuntu 14.04 x32 vmlinuz-3.13.0-24-generic",
			Version: "3.13.0-24-generic",
		},
		&Kernel{
			ID:      485432972,
			Name:    "Ubuntu 14.04 x64 vmlinuz-3.13.0-24-generic (1221)",
			Version: "3.13.0-24-generic",
		},
	}

	mux.HandleFunc("/droplets/21/kernels", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")
		fmt.Fprint(w, listKernelsExample)
	})

	kernels, err := client.ListKernels(21)

	assertEqual(t, err, nil)
	assertEqual(t, kernels, want)
}

func Test_ListSnapshots(t *testing.T) {
	setup(t)
	defer teardown()

	want := Snapshots{
		&Snapshot{
			ID:           119192820,
			Name:         "Ubuntu 13.04",
			Distribution: "ubuntu",
			Slug:         "",
			Public:       false,
			Regions: []string{
				"nyc1",
			},
		},
	}

	mux.HandleFunc("/droplets/21/snapshots", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")
		fmt.Fprint(w, listSnapshotsExample)
	})

	snapshots, err := client.ListSnapshots(21)

	assertEqual(t, err, nil)
	assertEqual(t, snapshots, want)
}

func Test_ListBackups(t *testing.T) {
	setup(t)
	defer teardown()

	want := Backups{
		&Backup{
			ID:           119192820,
			Name:         "Ubuntu 13.04",
			Distribution: "ubuntu",
			Slug:         "ubuntu1304",
			Public:       false,
			Regions: []string{
				"nyc1",
			},
		},
	}

	mux.HandleFunc("/droplets/21/backups", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")
		fmt.Fprint(w, listBackupsExample)
	})

	backups, err := client.ListBackups(21)

	assertEqual(t, err, nil)
	assertEqual(t, backups, want)
}

var createDropletExample = `{
  "droplet": {
    "id": 25,
    "name": "My-Droplet",
    "memory": 512,
    "vcpus": 1,
    "disk": 20,
    "region": {
      "slug": "nyc1",
      "name": "New York",
      "sizes": [
        "1gb",
        "512mb"
      ],
      "available": true,
      "features": [
        "virtio",
        "private_networking",
        "backups",
        "ipv6"
      ]
    },
    "image": {
      "id": 449676389,
      "name": "Ubuntu 13.04",
      "distribution": "ubuntu",
      "slug": null,
      "public": true,
      "regions": [
        "nyc1"
      ],
      "created_at": "2014-09-05T02:02:07Z"
    },
    "size": {
      "slug": "512mb",
      "transfer": 1,
      "price_monthly": 5.0,
      "price_hourly": 0.00744
    },
    "locked": false,
    "status": "new",
    "networks": {
      "v4": [
        {
          "ip_address": "127.0.0.20",
          "netmask": "255.255.255.0",
          "gateway": "127.0.0.21",
          "type": "public"
        },
        {
          "ip_address": "10.130.0.0",
          "netmask": "255.255.255.0",
          "gateway": "127.0.0.21",
          "type": "private"
        }
      ],
      "v6": [
        {
          "ip_address": "2001::14",
          "cidr": 124,
          "gateway": "2400:6180:0000:00D0:0000:0000:0009:7000",
          "type": "public"
        }
      ]
    },
    "kernel": {
      "id": 485432972,
      "name": "Ubuntu 14.04 x64 vmlinuz-3.13.0-24-generic (1221)",
      "version": "3.13.0-24-generic"
    },
    "created_at": "2014-09-05T02:02:07Z",
    "features": [
      "virtio"
    ],
    "backup_ids": [

    ],
    "snapshot_ids": [

    ]
  },
  "links": {
    "actions": [
      {
        "id": 20,
        "rel": "create",
        "href": "http://example.org/v2/actions/20"
      }
    ]
  }
}`

var getDropletExample = `{
  "droplet": {
    "id": 20,
    "networks": {
      "v4": [
        {
          "ip_address": "127.0.0.20",
          "netmask": "255.255.255.0",
          "gateway": "127.0.0.21",
          "type": "public"
        }
      ],
      "v6": [
        {
          "ip_address": "2001::14",
          "cidr": 124,
          "gateway": "2400:6180:0000:00D0:0000:0000:0009:7000",
          "type": "public"
        }
      ]
    },
    "features": [
    ],
    "backup_ids": [
        11
    ],
    "snapshot_ids": [
        20
    ]
  }
}`

var listDropletsExample = `{
  "droplets": [
    {
      "id": 20
    },
    {
      "id": 30
    },
    {
      "id": 42 
    }
  ]
}`

var listKernelsExample = `{
  "kernels": [
    {
      "id": 61833229,
      "name": "Ubuntu 14.04 x32 vmlinuz-3.13.0-24-generic",
      "version": "3.13.0-24-generic"
    },
    {
      "id": 485432972,
      "name": "Ubuntu 14.04 x64 vmlinuz-3.13.0-24-generic (1221)",
      "version": "3.13.0-24-generic"
    }
  ],
  "meta": {
    "total": 2
  }
}`

var listSnapshotsExample = `{
  "snapshots": [
    {
      "id": 119192820,     
      "name": "Ubuntu 13.04",
      "distribution": "ubuntu",
      "slug": null,
      "public": false,
      "regions": [
        "nyc1"
      ]
   }
 ]
}`

var listBackupsExample = `{
  "backups": [
    {
      "id": 119192820,     
      "name": "Ubuntu 13.04",
      "distribution": "ubuntu",
      "slug": "ubuntu1304",
      "public": false,
      "regions": [
        "nyc1"
      ]
   }
 ]
}`
