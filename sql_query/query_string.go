package sql_query

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func InterpolateQuery(query string, args ...any) string {
	re := regexp.MustCompile(`\$(\d+)`)

	result := re.ReplaceAllStringFunc(query, func(match string) string {
		// Extract the number from the matched placeholder
		indexStr := re.FindStringSubmatch(match)[1]
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			// Handle error if placeholder isn't a valid number
			return "NULL"
		}

		// Adjust for 0-based index
		i := index - 1
		if i < 0 || i >= len(args) {
			return "NULL" // safety fallback
		}

		arg := args[i]
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
