package dogo

import (
	"fmt"
	"strings"
	"net/url"
)

func createURL(endpoint string, id interface{}, params Params) string {
	if id == nil {
		return fmt.Sprintf("%s/?%s", endpoint, parseParams(params))
	}
	return fmt.Sprintf("%s/%v/?%s", endpoint, id, parseParams(params))
}

func parseParams(params Params) string {
	var s string
	for k, v := range params {
		s += fmt.Sprintf("%v=%v&", url.QueryEscape(k), v)
	}
	s = strings.TrimRight(s, "&")
	return s
}
