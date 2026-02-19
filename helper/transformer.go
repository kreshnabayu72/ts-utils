package helper

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

var regexOnlyInt = regexp.MustCompile(`-?[0-9]+`)
var regexOnlyFloat = regexp.MustCompile(`-?[0-9]*\.?[0-9]+`)

// IntoString transforms anything into string.
func IntoString(v any) string {
	s, _ := v.(string)
	s = strings.TrimSpace(s)

	return s
}

// IntoInt transforms anything into int.
func IntoInt(v any) int {
	s := IntoString(v)

	intStr := regexOnlyInt.FindString(s)
	result, _ := strconv.Atoi(intStr)
	return result
}

// IntoFloat transforms anything into float (specifically float64).
func IntoFloat(v any) float64 {
	s := IntoString(v)

	floatStr := regexOnlyFloat.FindString(s)
	result, _ := strconv.ParseFloat(floatStr, 64)
	return result
}

// IntoBool transforms anything into boolean.
func IntoBool(v any) bool {
	s := IntoString(v)

	return slices.Contains([]string{"TRUE", "T", "YES", "Y", "1"}, strings.ToUpper(s))
}

// IntoTime transforms anything into time.Time.
//
// Currently only the date format "2006-01-02" is supported.
func IntoTime(v any) time.Time {
	s := IntoString(v)

	parsedTime, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return time.Time{}
	}

	return parsedTime
}
