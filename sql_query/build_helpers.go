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
// WHERE filters for table & HAVING clauses won't be generated if you provide destination in params (pointer of slice string), but this function will assign value to the destination instead.
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
	v ...*[]string,
) {
	var dest []string
	useDestination := len(v) > 0

	for column, each := range filters {
		// Skip value nil except for IS NULL / IS NOT NULL
		if each.Value == nil &&
			each.Operator != SQLOperatorIsNull &&
			each.Operator != SQLOperatorIsNotNull {
			continue
		}

		var clause string

		// Needed for filtering with subquery, e.g: WHERE category_id IN (SELECT id FROM category_recursive)
		if each.IsSubQuery {
			clause = fmt.Sprintf(`%s %s %v`, column, each.Operator, each.Value)
			s.Filters = append(s.Filters, clause)
			break
		}

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
			quotedColumn := escapeQuoteColumns(column)
			clause = fmt.Sprintf(`%s %s`, quotedColumn, each.Operator)

		/* ─────────────────── BETWEEN ──────────────────── */
		case SQLOperatorBetween, SQLOperatorNotBetween:
			v := reflect.ValueOf(each.Value)
			if v.Kind() != reflect.Slice || v.Len() != 2 {
				continue // Skip if length is not 2
			}
			clause = fmt.Sprintf(
				`%s %s $%d AND $%d`,
				escapeQuoteColumns(column), each.Operator,
				len(s.Args)+1, len(s.Args)+2,
			)
			s.Args = append(s.Args, v.Index(0).Interface(), v.Index(1).Interface())

		/* ────────────────── = ANY($n) ─────────────────── */
		// products.user.id = ANY ($1) (array became args)
		case SQLOperatorAny, SQLOperatorIn, SQLOperatorNotIn:
			quotedColumn := escapeQuoteColumns(column)

			v := reflect.ValueOf(each.Value)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}

			if !v.IsValid() || (v.Kind() == reflect.Slice && v.IsNil()) {
				continue
			}

			if v.Len() == 0 {
				if each.Operator != SQLOperatorNotIn {
					// Empty slice/array means show nothing
					clause = "FALSE"
				} else {
					// Empty slice/array means show all
					clause = "TRUE"
				}
				break
			}

			switch each.Operator {
			case SQLOperatorIn, SQLOperatorNotIn:
				ph := make([]string, v.Len())
				for i := 0; i < v.Len(); i++ {
					ph[i] = fmt.Sprintf("$%d", len(s.Args)+1)
					s.Args = append(s.Args, v.Index(i).Interface())
				}

				clause = fmt.Sprintf(`%s %s (%s)`, escapeQuoteColumns(column), each.Operator, strings.Join(ph, ", "))
			case SQLOperatorAny:
				clause = fmt.Sprintf(`%s = ANY($%d)`, quotedColumn, len(s.Args)+1)
				s.Args = append(s.Args, each.Value)
			}

		/* ──────────────────── DEFAULT ─────────────────── */
		default:
			// Reference to other columns like users.id = user_assets.user_id
			if each.IsRef {
				quotedColumn := escapeQuoteColumns(column)
				clause = fmt.Sprintf(`%s %s %v`, quotedColumn, each.Operator, each.Value)
				break
			}

			if each.SourceIsValue {
				s.Args = append(s.Args, column)
				clause = fmt.Sprintf(`$%d %s $%d`, len(s.Args), each.Operator, len(s.Args)+1)
				s.Args = append(s.Args, each.Value)
				break
			}

			// Common operator, products.user.id = $1 (literal value with args)=
			clause = fmt.Sprintf(`%s %s $%d`, escapeQuoteColumns(column), each.Operator, len(s.Args)+1)
			s.Args = append(s.Args, each.Value)
		}

		// If there's destination in param, then HavingClauses and Filters wont be used, instead assign value to destination
		if useDestination {
			dest = append(dest, clause)
			continue
		}

		// Insert to Where clause or Having clause based on given params
		if s.useHaving {
			s.HavingClauses = append(s.HavingClauses, clause)
		} else {
			s.Filters = append(s.Filters, clause)
		}
	}

	if useDestination {
		if v[0] == nil {
			v[0] = &[]string{}
		}
		*v[0] = dest
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
	s.whereOrDestination(filters)
}

// Core function of `Where Or` query.
// WHERE filters for table won't be generated if you provide destination in params (pointer of slice string), but this function will assign value to the destination instead.
func (s *SQLEloquentQuery) whereOrDestination(
	filters []map[string]SQLCondition,
	v ...*[]string,
) {
	var dest []string
	useDestination := len(v) > 0
	orConditions := []string{}

	for _, filter := range filters {
		inner := &SQLEloquentQuery{Args: s.Args}
		inner.sharedWhereAndQuery(filter)
		orClause := fmt.Sprintf("(%s)", strings.Join(inner.Filters, " AND "))
		orConditions = append(orConditions, orClause)
		s.Args = inner.Args
	}

	if !useDestination {
		s.Filters = append(s.Filters, fmt.Sprintf("(%s)", strings.Join(orConditions, " OR ")))
	} else {
		if v[0] == nil {
			v[0] = &[]string{}
		}
		*v[0] = dest
	}
}

func escapeQuoteColumns(column string) string {
	var hasSuffix bool
	var suffix string
	if strings.Contains(column, "::") {
		hasSuffix = true
		splittedByCast := strings.Split(column, "::")
		suffix = "::" + splittedByCast[len(splittedByCast)-1]
		column = splittedByCast[0]
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
