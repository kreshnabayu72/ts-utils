package sql_query

import (
	"fmt"
	"regexp"
	"strings"
)

func InterpolateQuery(query string, args ...any) string {
	// Regex to find $1, $2, ...
	re := regexp.MustCompile(`\$\d+`)

	i := 0
	result := re.ReplaceAllStringFunc(query, func(_ string) string {
		if i >= len(args) {
			return "NULL" // safety fallback
		}
		arg := args[i]
		i++

		switch v := arg.(type) {
		case string:
			// wrap strings in quotes, escape single quotes
			return "'" + strings.ReplaceAll(v, "'", "''") + "'"
		default:
			// numbers, bools, etc → fmt
			return fmt.Sprint(v)
		}
	})

	return result
}
