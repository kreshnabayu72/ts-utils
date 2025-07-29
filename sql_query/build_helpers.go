package sql_query

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// sharedWhereAndQuery builds SQL WHERE/HAVING clauses from filters.
// It supports operators like =, IN, BETWEEN, IS NULL, RAW, and column refs (IsRef).
// Generated SQL fragments are added to Filters or HavingClauses, and Args is updated.
//
// Example:
//
//	filters := map[string]sql_query.SQLCondition{
//	    "age": {Operator: ">", Value: 18},
//	    "status": {Operator: "IN", Value: []string{"active", "pending"}},
//	}
//	builder.sharedWhereAndQuery(filters)
//	// Produces: "age" > $1 AND "status" IN ($2, $3)
func (s *SQLEloquentQuery) sharedWhereAndQuery(
	filters map[string]SQLCondition,
) {
	for column, each := range filters {
		// Skip value nil except for IS NULL / IS NOT NULL
		if each.Value == nil &&
			each.Operator != SQLOperatorIsNull &&
			each.Operator != SQLOperatorIsNotNull {
			continue
		}

		var clause string

		switch each.Operator {
		/* ───────────────────── RAW ───────────────────── */
		case SQLOperatorRaw:
			// Value should be ready to use SQL query
			raw, ok := each.Value.(string)
			if !ok {
				continue // atau panic
			}
			clause = raw
			// If exists, ExtraArgs will be appended into main Args
			if len(each.ExtraArgs) > 0 {
				s.Args = append(s.Args, each.ExtraArgs...)
			}

		/* ───────────── IS NULL / IS NOT NULL ──────────── */
		case SQLOperatorIsNull, SQLOperatorIsNotNull:
			clause = fmt.Sprintf(`"%s" %s`, column, each.Operator)

		/* ─────────────────── BETWEEN ──────────────────── */
		case SQLOperatorBetween, SQLOperatorNotBetween:
			v := reflect.ValueOf(each.Value)
			if v.Kind() != reflect.Slice || v.Len() != 2 {
				continue // Skip if length is not 2
			}
			clause = fmt.Sprintf(
				`"%s" %s $%d AND $%d`,
				column, each.Operator,
				len(s.Args)+1, len(s.Args)+2,
			)
			s.Args = append(s.Args, v.Index(0).Interface(), v.Index(1).Interface())

		/* ────────────────── = ANY($n) ─────────────────── */
		case SQLOperatorAny:
			clause = fmt.Sprintf(`"%s" = ANY($%d)`, column, len(s.Args)+1)
			s.Args = append(s.Args, each.Value)

		/* ──────────────────── DEFAULT ─────────────────── */
		default:
			// Reference to other columns like users.id = user_assets.user_id
			if each.IsRef {
				clause = fmt.Sprintf(`"%s" %s %v`, column, each.Operator, each.Value)
				break
			}

			typ := reflect.TypeOf(each.Value)
			if typ.Kind() == reflect.Slice {
				// products.user.id IN ($1,$2) (each element in array became args)
				// To expand placeholder usage for operator IN / NOT IN
				v := reflect.ValueOf(each.Value)
				if v.Len() == 0 {
					continue
				}
				ph := make([]string, v.Len())
				for i := 0; i < v.Len(); i++ {
					ph[i] = fmt.Sprintf("$%d", len(s.Args)+1)
					s.Args = append(s.Args, v.Index(i).Interface())
				}

				parts := strings.Split(column, ".")
				for i, part := range parts {
					parts[i] = fmt.Sprintf(`"%s"`, part)
				}
				quotedColumn := column
				if !strings.Contains(column, "(") {
					quotedColumn = strings.Join(parts, ".")
				}

				clause = fmt.Sprintf(`%s %s (%s)`, quotedColumn, each.Operator, strings.Join(ph, ", "))
			} else {
				// Common operator, products.user.id = $1 (literal value with args)
				parts := strings.Split(column, ".")
				for i, part := range parts {
					parts[i] = fmt.Sprintf(`"%s"`, part)
				}
				quotedColumn := column
				if !strings.Contains(column, "(") {
					quotedColumn = strings.Join(parts, ".")
				}

				clause = fmt.Sprintf(`%s %s $%d`, quotedColumn, each.Operator, len(s.Args)+1)
				s.Args = append(s.Args, each.Value)
			}
		}

		// Insert to Where clause or Having clause based on given params
		if s.useHaving {
			s.HavingClauses = append(s.HavingClauses, clause)
		} else {
			s.Filters = append(s.Filters, clause)
		}
	}
}

// sharedWhereOrQuery builds OR-combined conditions from multiple filter maps.
// Each map is AND-combined internally, then OR-joined together. Args is updated.
//
// Example:
//
//	builder.sharedWhereOrQuery(
//	    map[string]SQLCondition{"role": {Operator: "=", Value: "admin"}},
//	    map[string]SQLCondition{"role": {Operator: "=", Value: "editor"}},
//	)
//	// Produces: (("role" = $1) OR ("role" = $2))
func (s *SQLEloquentQuery) sharedWhereOrQuery(filters ...map[string]SQLCondition) {
	orConditions := []string{}

	for _, filter := range filters {
		inner := &SQLEloquentQuery{Args: s.Args}
		inner.sharedWhereAndQuery(filter)
		orClause := fmt.Sprintf("(%s)", strings.Join(inner.Filters, " AND "))
		orConditions = append(orConditions, orClause)
		s.Args = inner.Args
	}

	s.Filters = append(s.Filters, fmt.Sprintf("(%s)", strings.Join(orConditions, " OR ")))
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
// func (b *SQLBuilder) flattenArgs(val interface{}) []interface{} {
// 	v := reflect.ValueOf(val)
// 	if v.Kind() == reflect.Slice {
// 		args := make([]interface{}, v.Len())
// 		for i := 0; i < v.Len(); i++ {
// 			args[i] = v.Index(i).Interface()
// 		}
// 		return args
// 	}
// 	// fallback for single values
// 	return []interface{}{val}
// }

// func (b *SQLBuilder) makePlaceholders(val interface{}) string {
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

// func (b *SQLBuilder) countArgs(val interface{}) int {
// 	v := reflect.ValueOf(val)
// 	if v.Kind() == reflect.Slice {
// 		return v.Len()
// 	}
// 	return 1
// }
