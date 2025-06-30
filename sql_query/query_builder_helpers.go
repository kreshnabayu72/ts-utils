package sql_query

import (
	"fmt"
	"reflect"
	"strings"
)

// flattenArgs handles []any or single value
func (b *SQLBuilder) flattenArgs(val interface{}) []interface{} {
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Slice {
		args := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			args[i] = v.Index(i).Interface()
		}
		return args
	}
	// fallback for single values
	return []interface{}{val}
}

func (b *SQLBuilder) makePlaceholders(val interface{}) string {
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Slice {
		return fmt.Sprintf("$%d", b.argIndex)
	}

	placeholders := make([]string, v.Len())
	for i := 0; i < v.Len(); i++ {
		placeholders[i] = fmt.Sprintf("$%d", b.argIndex+i)
	}
	return strings.Join(placeholders, ", ")
}

func (b *SQLBuilder) countArgs(val interface{}) int {
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Slice {
		return v.Len()
	}
	return 1
}
