package sql_query

import (
	"fmt"
	"reflect"
	"strings"
)

func (s *SQLEloquentQuery) sharedWhereQuery(filters map[string]SQLCondition, isHavingClause bool) {
	fmt.Println("Sharedwhere")

	for column, each := range filters {
		var clause string
		if each.Value == nil || reflect.ValueOf(each.Value).IsZero() {
			continue
		}

		t := reflect.TypeOf(each.Value)

		if each.IsRef {
			// products.user.id = users.id (reference)
			clause = fmt.Sprintf("%s %s %v", column, each.Operator, each.Value)
		} else if t.Kind() == reflect.Slice {
			// products.user.id IN ($1,$2) (each element in array became args)
			v := reflect.ValueOf(each.Value)

			literals := ""
			for i := 0; i < v.Len(); i++ {
				if i > 0 {
					literals += ", "
				}
				literals += fmt.Sprintf("$%d", len(s.Args)+1)
				s.Args = append(s.Args, v.Index(i).Interface())
			}
			clause = fmt.Sprintf("%s %s (%s)", column, each.Operator, literals)
		} else {
			// products.user.id = $1 (literal value with args)

			clause = fmt.Sprintf("%s %s $%d", column, each.Operator, len(s.Args)+1)
			s.Args = append(s.Args, each.Value)
		}

		if isHavingClause {
			s.HavingClauses = append(s.HavingClauses, clause)
		} else {
			fmt.Println("APPEN CLAUSE TO FILTERS " + clause)
			s.Filters = append(s.Filters, clause)
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
