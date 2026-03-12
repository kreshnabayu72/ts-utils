package helper

import (
	"fmt"
	"math/big"
)

// Increment takes a string of number, adds 1, and returns the new string.
func Increment(numStr string) (string, error) {
	// 1. Create a new big integer
	n := new(big.Int)

	// 2. Parse the string as base 10
	// SetString returns the *Int and a boolean indicating success
	_, success := n.SetString(numStr, 10)
	if !success {
		return "", fmt.Errorf("invalid number string provided: %s", numStr)
	}

	// 3. Add 1 to the integer
	n.Add(n, big.NewInt(1))

	// 4. Convert back to Hex string
	// .Text(10) returns the string representation in base 10
	result := n.Text(10)

	return result, nil
}
