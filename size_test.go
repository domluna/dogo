package digitalocean

import (
	"testing"
        "fmt"
        "net/http"
)

func Test_ListSizes(t *testing.T) {
	setup(t)
	defer teardown()

        mux.HandleFunc("/sizes", func(w http.ResponseWriter, r *http.Request){
                assertEqual(t, r.Method, "GET")
                fmt.Fprint(w, listSizesExample)
        })

	want := Sizes{
		&Size{
			Slug:        "512mb",
			Disk:        20,
			PriceHourly: 0.00744,
			Regions: []string{
				"nyc1",
				"sfo1",
				"ams1",
			},
		},
		&Size{
			Slug:        "1gb",
			Disk:        30,
			PriceHourly: 0.01488,
			Regions: []string{
				"nyc1",
				"sfo1",
				"ams1",
			},
		},
	}

	sizes, err := client.ListSizes()

	assertEqual(t, err, nil)
	assertEqual(t, len(sizes), 2)
	assertEqual(t, sizes[0].Slug, want[0].Slug)
	assertEqual(t, sizes[0].Disk, want[0].Disk)
	assertEqual(t, sizes[0].PriceHourly, want[0].PriceHourly)
	assertEqual(t, sizes[0].Regions, want[0].Regions)
	assertEqual(t, sizes[1].Slug, want[1].Slug)
	assertEqual(t, sizes[1].Disk, want[1].Disk)
	assertEqual(t, sizes[1].PriceHourly, want[1].PriceHourly)
	assertEqual(t, sizes[1].Regions, want[1].Regions)
}

var listSizesExample = `{ 
  "sizes": [
    {
      "slug": "512mb",
      "memory": 512,
      "vcpus": 1,
      "disk": 20,
      "transfer": 1,
      "price_monthly": 5.0,
      "price_hourly": 0.00744,
      "regions": [
        "nyc1",
        "sfo1",
        "ams1"
      ]
    },
    {
      "slug": "1gb",
      "memory": 1024,
      "vcpus": 2,
      "disk": 30,
      "transfer": 2,
      "price_monthly": 10.0,
      "price_hourly": 0.01488,
      "regions": [
        "nyc1",
        "sfo1",
        "ams1"
      ]
    }
  ],
  "meta": {
    "total": 2
  }
}`
