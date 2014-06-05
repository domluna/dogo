package dogo

import (
	"encoding/json"
	"testing"
)

func TestOmitFields(t *testing.T) {
	arr := []string{`{"status":"OK","sizes":[{"id":66,"name":"512MB","slug":"512mb"}]}`}
	var got Response
	for _, s := range arr {
		err := json.Unmarshal([]byte(s), &got)
		if err != nil {
			t.Fatal(err)
		}

		b, err := json.Marshal(got)
		if err != nil {
			t.Fatal(err)
		}
		if string(b) != s {
			t.Errorf("Expected %v, got %v", s, string(b))
		}
	}

}
