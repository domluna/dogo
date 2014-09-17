package digitalocean

import (
        "testing"
        "net/http"
        "fmt"
)

func Test_ListActions(t *testing.T) {
        setup(t)
        defer teardown()
        want := Actions{
                &Action{
                        ID: 1,
                        Status: "in-progress",
                        Type: "test",
                        StartedAt: "2014-09-05T02:01:57Z",
                        CompletedAt: "",
                        ResourceID: 0,
                        ResourceType: "backend",
                        Region: "nyc1",
                },
        }
        mux.HandleFunc("/actions", func(w http.ResponseWriter, r *http.Request){
                assertEqual(t, r.Method, "GET")
                fmt.Fprint(w, listActionsExample)
        })

        actions, err := client.ListActions()

        assertEqual(t, err, nil)
        assertEqual(t, len(actions), 1)
        assertEqual(t, actions, want)
}

func Test_GetAction(t *testing.T) {
        setup(t)
        defer teardown()
        want := &Action{
                ID: 1,
                Status: "in-progress",
                Type: "test",
                StartedAt: "2014-09-05T02:01:57Z",
                CompletedAt: "",
                ResourceID: 0,
                ResourceType: "backend",
                Region: "nyc1",
        }
        mux.HandleFunc("/actions/1", func(w http.ResponseWriter, r *http.Request){
                assertEqual(t, r.Method, "GET")
                fmt.Fprint(w, getActionExample)
        })

        action, err := client.GetAction(1)

        assertEqual(t, err, nil)
        assertEqual(t, action, want)
}
var listActionsExample = `{
  "actions": [
    {
      "id": 1,
      "status": "in-progress",
      "type": "test",
      "started_at": "2014-09-05T02:01:57Z",
      "completed_at": null,
      "resource_id": null,
      "resource_type": "backend",
      "region": "nyc1"
    }
  ],
  "meta": {
    "total": 1
  }
}`

var getActionExample = `{
 "action": {
    "id": 1,
    "status": "in-progress",
    "type": "test",
    "started_at": "2014-09-05T02:01:57Z",
    "completed_at": null,
    "resource_id": null,
    "resource_type": "backend",
    "region": "nyc1"
  }
}`
