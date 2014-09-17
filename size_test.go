package digitalocean

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_ListSizes(t *testing.T) {
	setup(t)
	defer teardown()

	want := Sizes{
		&Size{
			Slug:        "512mb",
			Disk:        20,
                        Memory: 512,
                        VCPUS: 1,
                        Transfer: 1,
			PriceHourly: 0.00744,
                        PriceMonthly: 5.0,
			Regions: []string{
				"nyc1",
				"sfo1",
				"ams1",
			},
		},
		&Size{
			Slug:        "1gb",
			Disk:        30,
                        Memory: 1024,
                        VCPUS: 2,
                        Transfer: 2,
			PriceHourly: 0.01488,
                        PriceMonthly: 10.0,
			Regions: []string{
				"nyc1",
				"sfo1",
				"ams1",
			},
		},
	}

	mux.HandleFunc("/sizes", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")
		fmt.Fprint(w, listSizesExample)
	})


	sizes, err := client.ListSizes()

	assertEqual(t, err, nil)
	assertEqual(t, len(sizes), 2)
        assertEqual(t, sizes, want)
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
