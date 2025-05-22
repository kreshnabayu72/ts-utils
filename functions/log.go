package functions

import (
	"encoding/json"
	"fmt"
)

// PrintJS prints any Go object or array as a beautified JavaScript constant.
func PrintJS(obj interface{}) {
	// Marshal with indentation for readability
	jsonBytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Printf("%v", err.Error())
	}

	fmt.Printf("%v", string(jsonBytes))
}
