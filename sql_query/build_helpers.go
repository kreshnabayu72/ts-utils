package sql_query

import (
	"fmt"
	"reflect"
	"strings"
)

func (s *SQLEloquentQuery) sharedWhereQuery(filters map[string]SQLCondition) {
	if len(filters) < 1 {
		return
	}

	for column, each := range filters {
		if each.Value == nil || reflect.ValueOf(each.Value).IsZero() {
			continue
		}

		if each.IsRef {
			// products.user.id = users.id (reference)
			s.Filters = append(s.Filters, fmt.Sprintf("%s %s %v", column, each.Operator, each.Value))
		} else {
			// products.user.id = $1 (literal value with args)

			s.Filters = append(s.Filters, fmt.Sprintf("%s %s $%d", column, each.Operator, len(s.Args)+1))
			s.Args = append(s.Args, each.Value)
		}
	}
}

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
