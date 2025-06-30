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

	var filterQuery []string
	for column, each := range filters {
		if each.Value == nil || reflect.ValueOf(each.Value).IsZero() {
			continue
		}

		if each.IsRef {
			// products.user.id = users.id (reference)
			filterQuery = append(filterQuery, fmt.Sprintf("%s %s %v", column, each.Operator, each.Value))
		} else {
			// products.user.id = $1 (literal value with args)
			filterQuery = append(filterQuery, fmt.Sprintf("%s %s $%d", column, each.Operator, len(s.Args)+1))
			s.Args = append(s.Args, each.Value)
		}
	}

	s.Filters = fmt.Sprintf("WHERE %s", strings.Join(filterQuery, " AND "))
}
