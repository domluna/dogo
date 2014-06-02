package dogo

import (
	"fmt"
	"net/url"
	"strings"
)

// Creates a url required to interact with DigitalOcean API.
//
// url is of the form https://api.digitalocean.com/v1/{service}/{id}/?{params}.
//
// endpoint makes up everything up to and not including {id}, in some requests
// id might not be required, in this case nil should be passed.
//
// params are the query paramaters required for the specific action.
func createURL(endpoint string, id interface{}, params Params) string {
	if id == nil {
		return fmt.Sprintf("%s/?%s", endpoint, parseParams(params))
	}
	return fmt.Sprintf("%s/%v/?%s", endpoint, id, parseParams(params))
}

// Returns a queryescaped string of the passed params
//
// Example:
//
// Params{
//  "foo": "bar",
//  "blah": 5,
// }
//
// would return "foo=bar&blah=5".
func parseParams(params Params) string {
	var s string
	for k, v := range params {
		s += fmt.Sprintf("%v=%v&", url.QueryEscape(k), v)
	}
	s = strings.TrimRight(s, "&")
	return s
}
