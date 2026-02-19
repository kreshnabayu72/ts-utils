package sql_query

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

// NewSQLUpdateBuilder creates a new UpdateBuilder for the given table.
//
// Example:
//
//	builder := NewSQLUpdateBuilder("users","u")
//	builder.Update(user).Return("id").Build()
func NewSQLUpdateBuilder(tableName string, alias ...string) UpdateStarter {
	if len(alias) > 0 {
		tableName = fmt.Sprintf("%s %s", tableName, strings.TrimSpace(alias[0]))
	}

	return &UpdateBuilder{
		&SQLEloquentQuery{
			Table:       tableName,
			Filters:     []string{},
			OtherTables: []string{},
			Columns:     []string{},
			CustomQuery: "UPDATE",
			Args:        nil,
			Mode:        "update",
		},
	}
}

func (s *UpdateBuilder) Where(filters WhereQuery) Updater {
	s.SQLEloquentQuery.sharedWhereAndQuery(filters)
	return s
}

func (s *UpdateBuilder) WhereOr(filters ...WhereQuery) Updater {
	s.SQLEloquentQuery.sharedWhereOrQuery(filters...)
	return s
}

func (s *UpdateBuilder) Join(table string, onCondition string, additionalConditions ...WhereQuery) Updater {
	if table == "" {
		return s
	}

	var filterSb strings.Builder
	if len(additionalConditions) > 0 {
		var filters []string
		s.sharedWhereAndQuery(additionalConditions[0], &filters)

		for i := range filters {
			filterSb.WriteString(" AND ")
			filterSb.WriteString(filters[i])
		}
	}

	s.OtherTables = append(s.OtherTables, fmt.Sprintf("JOIN %s ON %s%s", table, onCondition, filterSb.String()))
	return s
}

func (s *UpdateBuilder) LeftJoin(table string, mainCondition string, additionalConditions ...WhereQuery) Updater {
	if table == "" {
		return s
	}

	var filterSb strings.Builder
	if len(additionalConditions) > 0 {
		var filters []string
		s.sharedWhereAndQuery(additionalConditions[0], &filters)

		for i := range filters {
			filterSb.WriteString(" AND ")
			filterSb.WriteString(filters[i])
		}
	}

	s.OtherTables = append(s.OtherTables, fmt.Sprintf("LEFT JOIN %s ON %s%s", table, mainCondition, filterSb.String()))
	return s
}

func (s *UpdateBuilder) WithCTEBuilder(cteName string, cteBuilder *SQLEloquentQuery) Updater {
	cteQuery, cteArgs, err := cteBuilder.Build()
	if err != nil {
		s.LastError = err
		return s
	}

	// Calculate the current argument offset
	offset := len(s.Args)

	// Shift the placeholders in the CTE query
	shiftedCTEQuery := shiftSQLPlaceholders(cteQuery, offset)

	s.WithClauses = append(s.WithClauses, fmt.Sprintf("%s AS (%s)", cteName, shiftedCTEQuery))
	s.Args = append(s.Args, cteArgs...)

	return s
}

func (s *UpdateBuilder) WithRecursiveCTEBuilder(cteName string, cteBuilder *SQLEloquentQuery) Updater {
	cteQuery, cteArgs, err := cteBuilder.Build()
	if err != nil {
		s.LastError = err
		return s
	}

	// Calculate the current argument offset
	offset := len(s.Args)

	// Shift the placeholders in the CTE query
	shiftedCTEQuery := shiftSQLPlaceholders(cteQuery, offset)

	s.WithClauses = append(s.WithClauses, fmt.Sprintf("%s AS (%s)", cteName, shiftedCTEQuery))
	s.Args = append(s.Args, cteArgs...)

	s.useWithRecursive = true

	return s
}

func (s *UpdateBuilder) ExcludeEmpty() Updater {
	s.excludeEmptyValue = true
	return s
}

func (s *UpdateBuilder) Return(column ...string) Updater {
	if len(column) > 0 {
		s.Columns = column
	} else {
		s.Columns = []string{"id"}
	}
	return s
}

func (s *UpdateBuilder) Conflict(constraint, do string) Updater {
	s.ConflictClause = fmt.Sprintf(" ON CONFLICT %s DO %s", constraint, do)
	return s
}

func (s *UpdateBuilder) AddCase(setColumn string, fn func(b UpdateCases)) Updater {
	s.UpdateCaseClauses = make(map[string][]UpdateCaseParam, 0)
	s.UpdateCaseClauses[setColumn] = []UpdateCaseParam{}
	s.currentUpdateCaseKey = setColumn
	fn(s)
	return s
}

func (s *UpdateBuilder) Case(conditions MultiFilterCondition, value any, isRef bool) {
	filters := []string{}

	if len(conditions.And) > 0 {
		andFilters := []string{}
		s.sharedWhereAndQuery(conditions.And, &andFilters)
		filters = append(filters, andFilters...)
	}

	if len(conditions.Or) > 0 {
		orFilters := []string{}
		s.whereOrDestination(conditions.Or, &orFilters)
		filters = append(filters, orFilters...)
	}

	var valueSb strings.Builder
	if isRef {
		valueSb.WriteString(value.(string))
	} else {
		valueSb.WriteByte('$')
		s.Args = append(s.Args, value)
		valueSb.WriteString(strconv.Itoa(len(s.Args)))
	}

	s.UpdateCaseClauses[s.currentUpdateCaseKey] = append(
		s.UpdateCaseClauses[s.currentUpdateCaseKey],
		UpdateCaseParam{conditions: filters, value: valueSb.String()},
	)
}

func (s *UpdateBuilder) Else(value any, isRef bool) {
	var valueSb strings.Builder
	if isRef {
		valueSb.WriteString(value.(string))
	} else {
		valueSb.WriteByte('$')
		s.Args = append(s.Args, value)
		valueSb.WriteString(strconv.Itoa(len(s.Args)))
	}

	s.UpdateCaseClauses[s.currentUpdateCaseKey] = append(
		s.UpdateCaseClauses[s.currentUpdateCaseKey],
		UpdateCaseParam{conditions: []string{}, value: valueSb.String(), isElse: true},
	)
}

func (s *UpdateBuilder) Update(values UpdateQuery, skipUpdatedAt ...bool) Updater {
	// Directly call extractUpdateFieldsMap since UpdateQuery is map[string]any
	setClauses, hasUpdatedAt := s.extractUpdateFieldsMap(values)

	if !hasUpdatedAt && (len(skipUpdatedAt) == 0 || !skipUpdatedAt[0]) {
		setClauses = append(setClauses, `"updated_at" = NOW()`)
	}

	s.CustomQuery = fmt.Sprintf(`UPDATE %s SET %s`, s.Table, strings.Join(setClauses, ", "))

	return s
}

func (s *UpdateBuilder) UpdateEach(values any, rowIdentifier string) Updater {
	v := reflect.ValueOf(values)

	// Slice data, plus length checking
	if v.Kind() != reflect.Slice || v.Len() == 0 {
		s.LastError = errors.New("update many values must be non-empty slice of struct")
		return s
	}

	// Check type of slice, should be slice of struct (only check the first index)
	firstElem := v.Index(0)
	if firstElem.Kind() != reflect.Struct {
		s.LastError = errors.New("update slice must contain structs")
		return s
	}

	// Main condition to make sure each row assigned by their respective value by matching row identifier
	// Example input key is "id"/"name", value is given in slice data with column tag same as input key (rowIdentifier), operator is equal
	// Assume v is aliased VALUES, example output id = v.id or could be name = v.name if this column is unique
	s.Filters = append(
		s.Filters,
		fmt.Sprintf(`%s."%s" = v."%s"`, s.Table, rowIdentifier, rowIdentifier),
	)

	// Generate clauses for everything in slice, then stored in main array
	var setClauses, valueClauses, placeholders []string
	mappedSet, mappedValue, mappedPlaceholders := s.updateEachClausesGenerator(v, rowIdentifier)
	setClauses = append(setClauses, mappedSet...)
	valueClauses = append(valueClauses, mappedValue...)
	placeholders = append(placeholders, mappedPlaceholders...)

	// Build everything into one query dedicated for update many
	s.CustomQuery = fmt.Sprintf(
		`UPDATE %s SET %s FROM (VALUES %s) as v(%s)`,
		s.Table,
		strings.Join(setClauses, ", "),
		strings.Join(placeholders, ", "),
		strings.Join(valueClauses, ","),
	)

	return s
}

func (s *UpdateBuilder) Increment(
	values map[string]any,
) Updater {
	var setClauses []string

	for key, val := range values {
		snake := CamelToSnake(key)

		setClauses = append(
			setClauses,
			fmt.Sprintf(`"%s" = "%s" + $%d`, snake, snake, len(s.Args)+1),
		)
		s.Args = append(s.Args, val)
	}

	setClauses = append(setClauses, `"updated_at" = NOW()`)

	s.CustomQuery = fmt.Sprintf(`UPDATE %s SET %s`, s.Table, strings.Join(setClauses, ", "))
	return s
}

func (s *UpdateBuilder) From(tables []string) Updater {
	if len(tables) < 1 {
		return s
	}

	var otherTables []string
	otherTables = append(otherTables, fmt.Sprintf("FROM %s", strings.Join(tables, ", ")))
	s.OtherTables = otherTables
	return s
}

func (s *SQLEloquentQuery) buildUpdateQuery() (string, []any, error) {
	if s.LastError != nil {
		return "", nil, errors.New(s.LastError.Error())
	}

	if s.CustomQuery == "" {
		return "", nil, errors.New("invalid update query: CustomQuery not set")
	}

	var initSb strings.Builder
	var withSb strings.Builder
	var whereSb strings.Builder
	var joinSb strings.Builder
	var fromSb strings.Builder
	var returningSb strings.Builder

	initSb.Grow(256)      // preallocate ~256 bytes
	withSb.Grow(256)      // preallocate ~256 bytes
	whereSb.Grow(256)     // preallocate ~256 bytes
	joinSb.Grow(256)      // preallocate ~256 bytes
	fromSb.Grow(256)      // preallocate ~256 bytes
	returningSb.Grow(256) // preallocate ~256 bytes

	initSb.WriteByte('\n')
	if len(s.UpdateCaseClauses) > 0 {
		initSb.WriteString(buildUpdateCase(s.UpdateCaseClauses, s.Table))
	} else {
		initSb.WriteString(s.CustomQuery)
	}

	// WITH
	if len(s.WithClauses) > 0 {
		withSb.WriteByte('\n')

		// if one of WITH are recursive, add 'RECURSIVE' string
		if s.useWithRecursive {
			withSb.WriteString("WITH RECURSIVE ")
		} else {
			withSb.WriteString("WITH ")
		}

		for i, c := range s.WithClauses {
			if i > 0 {
				withSb.WriteString(", ")
			}
			withSb.WriteString(c)
		}
	}

	if len(s.OtherTables) > 0 {
		fromSb.WriteByte('\n')
		fromSb.WriteString(strings.Join(s.OtherTables, " "))
	}

	if len(s.Filters) < 1 {
		return "", nil, errors.New("unsafe query: DELETE/UPDATE must have WHERE clause")
	}

	// WHERE
	whereSb.WriteByte('\n')
	whereSb.WriteString("WHERE ")
	for i, f := range s.Filters {
		if i > 0 {
			whereSb.WriteString(" AND ")
		}
		whereSb.WriteString(f)
	}
	whereSb.WriteByte('\n')

	if len(s.Columns) > 0 {
		returningSb.WriteByte('\n')
		returningSb.WriteString(" RETURNING " + strings.Join(s.Columns, ","))
	} else {
		returningSb.WriteByte('\n')
		returningSb.WriteString(" RETURNING id")
	}

	query := withSb.String() + initSb.String() + fromSb.String() + whereSb.String() + returningSb.String()

	return query, s.Args, nil
}

func (s *UpdateBuilder) updateEachClausesGenerator(
	slice reflect.Value,
	rowIdentifier string,
) ([]string, []string, []string) {
	var args []any
	var setClauses, valueClauses, valuePlaceholders []string

	// Loop through slice
	for i := 0; i < slice.Len(); i++ {
		v := slice.Index(i)
		t := v.Type()

		if i == 0 {
			// Setup column list once
			// Append set clauses and value clauses, also insert updated_at if not exists
			for j := 0; j < t.NumField(); j++ {
				field := t.Field(j)
				columnTag := field.Tag.Get("column")

				// If column tag is equal with rowIdentifier given by param, then it should not append into set clauses, only append to value clauses for WHERE condition
				if columnTag == rowIdentifier {
					valueClauses = append(valueClauses, fmt.Sprintf(`"%s"`, columnTag))
				} else {
					setClauses = append(setClauses, fmt.Sprintf(`"%s" = v."%s"`, columnTag, columnTag))
					valueClauses = append(valueClauses, fmt.Sprintf(`"%s"`, columnTag))
				}
			}

			if ok := slices.Contains(valueClauses, "updated_at"); !ok {
				setClauses = append(setClauses, `"updated_at" = NOW()`)
			}
		}

		var rowPlaceholders []string

		// Add fields
		for j := 0; j < t.NumField(); j++ {
			field := t.Field(j)
			transformTag := field.Tag.Get("transform")
			args = append(args, v.Field(j).Interface())
			rowPlaceholders = append(
				rowPlaceholders,
				fmt.Sprintf("$%d::%s", len(args), transformTag),
			)
		}

		// Append every placeholders to main array that will be returned
		valuePlaceholders = append(
			valuePlaceholders,
			fmt.Sprintf("(%s)", strings.Join(rowPlaceholders, ", ")),
		)
	}

	s.Args = args
	return setClauses, valueClauses, valuePlaceholders
}

func (s *UpdateBuilder) extractUpdateFieldsMap(v map[string]any) ([]string, bool) {
	setClauses := []string{}
	hasUpdatedAt := false

	for key, value := range v {
		isZeroValue := isZeroValue(value)
		if isZeroValue && s.excludeEmptyValue {
			continue
		}

		v := reflect.ValueOf(value)

		// if value is map, recurse into it
		if v.Kind() == reflect.Map {
			childClauses, childHasUpdated := s.extractUpdateFieldsMap(value.(map[string]any))
			setClauses = append(setClauses, childClauses...)
			if childHasUpdated {
				hasUpdatedAt = true
			}
			continue
		}

		// Use column name as key, or field name if not provided
		col := key
		if col == "updated_at" {
			hasUpdatedAt = true
		}

		switch v := value.(type) {
		case UpdateRawSQL:
			fmt.Println("v := value.(type) UpdateRawSQL", v)
			expr := v.Expr

			// replace ? with correct $n placeholders
			for _, arg := range v.Args {
				expr = strings.Replace(expr, "?", fmt.Sprintf("$%d", len(s.Args)+1), 1)
				s.Args = append(s.Args, arg)
			}
			setClauses = append(setClauses, fmt.Sprintf(`"%s" = %s`, col, expr))
		default:
			setClauses = append(setClauses, fmt.Sprintf(`"%s" = $%d`, col, len(s.Args)+1))
			s.Args = append(s.Args, v)
		}
	}

	return setClauses, hasUpdatedAt
}

// buildUpdateCase constructs a SQL UPDATE statement with CASE expressions.
//
// It takes a map of column names to slices of UpdateCaseParam, where each slice
// represents the WHEN/THEN/ELSE clauses for that column's CASE statement.
//
// Parameters:
//   - updateCaseClauses: A map where each key is a column name, and each value is a
//     slice of UpdateCaseParam defining the conditional logic for updating that column.
//   - tableName: The name of the table to update.
//
// Returns:
//   - A string containing the full SQL UPDATE query with conditional CASE logic.
//
// Example output:
//
//	UPDATE table_name
//	SET
//	  column1 = CASE
//	    WHEN condition1 AND condition2 THEN value1
//	    ELSE default_value
//	  END,
//	  updated_at = NOW()
//
// Notes:
//   - The function appends "updated_at = NOW()" to the final SET clause.
//   - It assumes all values and conditions are properly escaped/formatted.
func buildUpdateCase(updateCaseClauses map[string][]UpdateCaseParam, tableName string) string {
	var updateExpr string
	updateExpr = "UPDATE " + tableName + "\n"
	updateExpr += "SET\n"

	for key, each := range updateCaseClauses {
		updateExpr += key + " = CASE\n"
		fmt.Println("each", each)

		for _, param := range each {
			if param.isElse {
				updateExpr += "ELSE " + param.value
			} else {
				updateExpr += "WHEN " + strings.Join(param.conditions, " AND ") + " THEN " + param.value
			}
			updateExpr += "\n"
		}

		updateExpr += "END,\n"
	}

	updateExpr += "updated_at = NOW()"
	return updateExpr
}

func isZeroValue(v any) bool {
	if v == nil {
		return true
	}

	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}
