package functions

import (
	"fmt"
	"github.com/goccy/go-json"
)

// JSONStringify stringifies any representation of a json.
func JSONStringify(val any) (string, error) {
	jsonStr, err := json.Marshal(val)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	return string(jsonStr), nil
}

// JSONParse parses any value representation of json.
func JSONParse[T any](s string) (T, error) {
	var parsed T
	err := json.Unmarshal([]byte(s), &parsed)
	if err != nil {
		return parsed, fmt.Errorf("%v", err)
	}

	return parsed, nil
}
