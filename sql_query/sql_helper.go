package sql_query

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func EscapeQuoteColumns(column string) string {
	var hasSuffix bool
	var suffix string
	if strings.Contains(column, "::") {
		hasSuffix = true
		splittedByCast := strings.Split(column, "::")
		suffix = "::" + splittedByCast[len(splittedByCast)-1]
		column = splittedByCast[0]
	} else if strings.Contains(column, "->>") {
		hasSuffix = true
		splittedByOperator := strings.Split(column, "->>")
		suffix = "->>" + splittedByOperator[len(splittedByOperator)-1]
		column = splittedByOperator[0]
	}

	parts := strings.Split(column, ".")
	for i, part := range parts {
		parts[i] = fmt.Sprintf(`"%s"`, part)
	}
	quotedColumn := column
	if !strings.Contains(column, "(") {
		quotedColumn = strings.Join(parts, ".")
	}

	if hasSuffix {
		quotedColumn += suffix
	}
	return quotedColumn
}

func shiftSQLPlaceholders(query string, offset int) string {
	placeholderRegexp := regexp.MustCompile(`\$(\d+)`)
	return placeholderRegexp.ReplaceAllStringFunc(query, func(placeholder string) string {
		numStr := placeholder[1:]
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return placeholder // fallback to original if parse fails
		}
		return fmt.Sprintf("$%d", num+offset)
	})
}

// // flattenArgs handles []any or single value
// func (b *SQLBuilder) flattenArgs(val any) []any {
// 	v := reflect.ValueOf(val)
// 	if v.Kind() == reflect.Slice {
// 		args := make([]any, v.Len())
// 		for i := 0; i < v.Len(); i++ {
// 			args[i] = v.Index(i).Interface()
// 		}
// 		return args
// 	}
// 	// fallback for single values
// 	return []any{val}
// }

// func (b *SQLBuilder) makePlaceholders(val any) string {
// 	v := reflect.ValueOf(val)
// 	if v.Kind() != reflect.Slice {
// 		return fmt.Sprintf("$%d", b.argIndex)
// 	}

// 	placeholders := make([]string, v.Len())
// 	for i := 0; i < v.Len(); i++ {
// 		placeholders[i] = fmt.Sprintf("$%d", b.argIndex+i)
// 	}
// 	return strings.Join(placeholders, ", ")
// }

// func (b *SQLBuilder) countArgs(val any) int {
// 	v := reflect.ValueOf(val)
// 	if v.Kind() == reflect.Slice {
// 		return v.Len()
// 	}
// 	return 1
// }
