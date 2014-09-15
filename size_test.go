package digitalocean

import (
        "testing"
        "net/http"
        "fmt"
        "reflect"
)

func Test_ListSizes(t *testing.T) {
        setup(t)
        defer teardown()

        mux.HandleFunc("/sizes", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintln(w, listSizesExample)
        })

        want := Sizes{
                &Size{
                        Slug: "512mb",
                        Disk: 20,
                        PriceHourly: 0.00744,
                        Regions: []string{
                                "nyc1",
                                "sfo1",
                                "ams1",
                        },
                },
                &Size{
                        Slug: "1gb",
                        Disk: 30,
                        PriceHourly: 0.01488,
                        Regions: []string{
                                "nyc1",
                                "sfo1",
                                "ams1",
                        },
                },
        }

        sizes, err := client.ListSizes()
        if err != nil {
                t.Errorf("Error retrieving sizes: %s", err)
        }


        if !reflect.DeepEqual(want[0].Slug, sizes[0].Slug) {
                t.Errorf("Expected %v, got %v", want[0].Slug, sizes[0].Slug)
        }
        if !reflect.DeepEqual(want[0].Disk, sizes[0].Disk) {
                t.Errorf("Expected %v, got %v", want[0].Disk, sizes[0].Disk)
        }
        if !reflect.DeepEqual(want[0].PriceHourly, sizes[0].PriceHourly) {
                t.Errorf("Expected %v, got %v", want[0].PriceHourly, sizes[0].PriceHourly)
        }
        if !reflect.DeepEqual(want[0].Regions, sizes[0].Regions) {
                t.Errorf("Expected %v, got %v", want[0].Regions, sizes[0].Regions)
        }
        if !reflect.DeepEqual(want[1].Slug, sizes[1].Slug) {
                t.Errorf("Expected %v, got %v", want[1].Slug, sizes[1].Slug)
        }
        if !reflect.DeepEqual(want[1].Disk, sizes[1].Disk) {
                t.Errorf("Expected %v, got %v", want[1].Disk, sizes[1].Disk)
        }
        if !reflect.DeepEqual(want[1].PriceHourly, sizes[1].PriceHourly) {
                t.Errorf("Expected %v, got %v", want[1].PriceHourly, sizes[1].PriceHourly)
        }
        if !reflect.DeepEqual(want[1].Regions, sizes[1].Regions) {
                t.Errorf("Expected %v, got %v", want[1].Regions, sizes[1].Regions)
        }
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
