package sql_query

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// type string struct {
// 	Key       string
// 	Operator  SQLOperators // e.g., '=', '>', '<=', 'LIKE', 'IN', 'IS NULL'
// 	Value     interface{}  // could be a single value, slice, or nil
// 	IsRef     bool         // to determine whether WHERE is targeting literal value or reference
// 	ExtraArgs []interface{}
// }

type UpdateCases interface {
	Case(conditions MultiFilterCondition, value string, isRef bool)
	Else(value string, isRef bool)
}

// To ensure SQLUpdateBuilder has its own initial methods
// e.g. UpdateBuilder(...).Update()...Rest
type SQLUpdateInitBuilder interface {
	// Update builds an UPDATE query for a single struct using reflection.
	// Automatically sets updated_at = NOW() if not provided.
	Update(values interface{}) SQLUpdateChainBuilder
	// UpdateEach updates multiple rows at once using VALUES() with a slice of structs.
	// Matches rows using the given rowIdentifier (e.g., "id").
	//
	// Example:
	//
	//	builder.UpdateEach([]User{{ID: 1, Name: "A"}, {ID: 2, Name: "B"}}, "id")
	//
	// → UPDATE users SET name = v.name, updated_at = NOW()
	//
	//	FROM (VALUES ($1,$2),($3,$4)) AS v(id,name,updated_at)
	//	WHERE users.id = v.id
	UpdateEach(values interface{}, rowIdentifier string) SQLUpdateChainBuilder

	AddCase(setColumn string, fn func(b UpdateCases)) SQLUpdateChainBuilder

	// Increment is used to replace Update() for adding int value cause update cant handle that
	// Increment builds an UPDATE query that increases integer columns by a given value.
	// Automatically sets updated_at = NOW().
	//
	// Example:
	//
	//	builder.Increment(map[string]any{"count": 1})
	//
	// → UPDATE table SET "count" = "count" + $1, "updated_at" = NOW()
	Increment(values map[string]any) SQLUpdateChainBuilder

	// updateEachClausesGenerator looks at every struct in the slice and builds:
	//  1. The SET part of the query (e.g., "name = v.name"),
	//  2. The column names to use in the VALUES table,
	//  3. The placeholders ($1, $2, …) for each row.
	//
	// It's the core logic behind UpdateEach, letting you bulk-update rows
	// by matching them with a unique rowIdentifier (e.g. "id" or "name").
	//
	// This function loop through slice, returning generated clauses for everything in it
	// Expected result:
	// setClauses contains: [value = v.value, ...etc].
	// valueClauses contains: [value, ...etc].
	// valuePlaceholders contains: [($1, $2, $3, ...etc), ($4, $5, ...etc), ...etc]
	updateEachClausesGenerator(
		slice reflect.Value,
		rowIdentifier string,
	) ([]string, []string, []string)

	// extractUpdateFieldsStruct goes through a struct's fields and turns them into
	// SET clauses for an UPDATE query. It uses the "column" or "json" tag (or
	// the field name if tags are missing) to pick the column name.
	//
	// It also checks if an "updated_at" field exists, so the builder knows
	// whether to add "updated_at = NOW()" automatically.
	extractUpdateFieldsStruct(v reflect.Value) ([]string, bool)

	// extractUpdateFieldsMap takes a map[string]interface{} and returns a slice of
	// SET clauses for an UPDATE query. It uses the map keys as column names.
	// It also checks if an "updated_at" key exists, so the builder knows
	// whether to add "updated_at = NOW()" automatically.
	extractUpdateFieldsMap(v map[string]any) ([]string, bool)
}

// To ensure method .Update() has its own chaining methods
// e.g. .Update(...).From(...).Build()
type SQLUpdateChainBuilder interface {
	AddCase(setColumn string, fn func(b UpdateCases)) SQLUpdateChainBuilder

	// By default, the Update builder includes all fields, even if their values are zero or nil.
	// Calling this method tells the builder to skip fields with zero values when generating the SET clause.
	//
	// Note: This option only affects single-row Update operations.
	// It has no effect on bulk UPDATE, because all rows in those operations must have the same set of columns.
	ExcludeEmpty() SQLUpdateChainBuilder
	// Update implements SQLUpdateChainBuilder. (Only able to be called once, will override previous call).
	// Conflict adds an ON CONFLICT clause with the specified constraint and action.
	//
	// Example:
	//
	//	builder.Conflict("(id)", "NOTHING")
	//
	// → ON CONFLICT (id) DO NOTHING
	Conflict(constraint, do string) SQLUpdateChainBuilder
	// Where implements SQLUpdateChainBuilder. (Accumulates previous value if called again)
	Where(filters map[string]SQLCondition) SQLUpdateChainBuilder
	// WhereOr implements SQLUpdateChainBuilder. (Accumulates previous value if called again)
	WhereOr(filters ...map[string]SQLCondition) SQLUpdateChainBuilder
	// Return implements SQLUpdateChainBuilder. (Only able to be called once, overrides previous values if re-called).
	// Return sets the columns to return after the update.
	// Defaults to RETURNING id if no column is provided.
	Return(columns ...string) SQLUpdateChainBuilder

	// From implements SQLUpdateChainBuilder. (Overrides previous value if called again)
	// From adds a FROM clause to the UPDATE query, allowing joins with other tables.
	//
	// Example:
	//
	//	builder.From([]string{"users u", "roles r"})
	//
	// → FROM users u, roles r
	From(tables []string) SQLUpdateChainBuilder

	// buildUpdateQuery constructs the final UPDATE query string and its arguments.
	// Ensures that CustomQuery is set and that a WHERE clause exists for safety.
	Build() (string, []interface{}, error)
}

type UpdateBuilder struct {
	*SQLEloquentQuery
}

func (s *UpdateBuilder) Where(filters map[string]SQLCondition) SQLUpdateChainBuilder {
	s.SQLEloquentQuery.sharedWhereAndQuery(filters)
	return s
}

func (s *UpdateBuilder) WhereOr(filters ...map[string]SQLCondition) SQLUpdateChainBuilder {
	s.SQLEloquentQuery.sharedWhereOrQuery(filters...)
	return s
}

func (s *UpdateBuilder) ExcludeEmpty() SQLUpdateChainBuilder {
	s.excludeEmptyValue = true
	return s
}

func (s *UpdateBuilder) Return(column ...string) SQLUpdateChainBuilder {
	if len(column) > 0 {
		s.Columns = column
	} else {
		s.Columns = []string{"id"}
	}
	return s
}

func (s *UpdateBuilder) Conflict(constraint, do string) SQLUpdateChainBuilder {
	s.ConflictClause = fmt.Sprintf(" ON CONFLICT %s DO %s", constraint, do)
	return s
}

func (s *UpdateBuilder) AddCase(setColumn string, fn func(b UpdateCases)) SQLUpdateChainBuilder {
	s.UpdateCaseClauses[setColumn] = []UpdateCaseParam{}
	s.currentUpdateCaseKey = setColumn
	return s
}

func (s *UpdateBuilder) Case(conditions MultiFilterCondition, value interface{}, isRef bool) {
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

func (s *UpdateBuilder) Else(value interface{}, isRef bool) {
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

func (s *UpdateBuilder) Update(values interface{}) SQLUpdateChainBuilder {
	v := reflect.ValueOf(values)

	setClauses := []string{}
	hasUpdatedAt := false
	if v.Kind() == reflect.Struct {
		setClauses, hasUpdatedAt = s.extractUpdateFieldsStruct(v)
	} else if v.Kind() == reflect.Map {
		setClauses, hasUpdatedAt = s.extractUpdateFieldsMap(values.(map[string]any))
	} else {
		s.LastError = fmt.Errorf("invalid update values: expected struct or map, got %T", values)
		return s
	}

	if !hasUpdatedAt {
		setClauses = append(setClauses, `"updated_at" = NOW()`)
	}

	s.CustomQuery = fmt.Sprintf(`UPDATE %s SET %s`, s.Table, strings.Join(setClauses, ", "))

	return s
}

func (s *UpdateBuilder) UpdateEach(values interface{}, rowIdentifier string) SQLUpdateChainBuilder {
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
) SQLUpdateChainBuilder {
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

func (s *UpdateBuilder) From(tables []string) SQLUpdateChainBuilder {
	if len(tables) < 1 {
		return s
	}

	var otherTables []string
	otherTables = append(otherTables, fmt.Sprintf("FROM %s", strings.Join(tables, ", ")))
	s.OtherTables = otherTables
	return s
}

// NewSQLUpdateBuilder creates a new UpdateBuilder for the given table.
//
// Example:
//
//	builder := NewSQLUpdateBuilder("users")
//	builder.Update(user).Return("id").Build()
func NewSQLUpdateBuilder(tableName string) SQLUpdateInitBuilder {
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

func (s *SQLEloquentQuery) buildUpdateQuery() (string, []interface{}, error) {
	if s.CustomQuery == "" {
		return "", nil, errors.New("invalid update query: CustomQuery not set")
	}

	var query string
	if len(s.UpdateCaseClauses) > 0 {
		query = buildUpdateCase(s.UpdateCaseClauses, s.Table)
	} else {
		query = s.CustomQuery
	}

	if len(s.OtherTables) > 0 {
		fromClause := strings.Join(s.OtherTables, " ")
		query += " " + fromClause
	}

	if len(s.Filters) < 1 {
		return "", nil, errors.New("unsafe query: DELETE/UPDATE must have WHERE clause")
	}

	query += " " + fmt.Sprintf("WHERE %s", strings.Join(s.Filters, " AND "))

	if len(s.Columns) > 0 {
		query += " RETURNING " + strings.Join(s.Columns, ",")
	} else {
		query += " RETURNING id::text"
	}

	return query, s.Args, nil
}

func (s *UpdateBuilder) updateEachClausesGenerator(
	slice reflect.Value,
	rowIdentifier string,
) ([]string, []string, []string) {
	var args []interface{}
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

			if ok := ArrayIncludes(valueClauses, "updated_at"); !ok {
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

func (s *UpdateBuilder) extractUpdateFieldsStruct(v reflect.Value) ([]string, bool) {
	var setClauses []string
	var hasUpdatedAt bool

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		val := v.Field(i)

		if val.IsZero() && s.excludeEmptyValue {
			continue
		}

		// Skip unexported or invalid fields
		if !val.IsValid() || !val.CanInterface() {
			continue
		}

		// Handle ignored fields
		jsonTag := field.Tag.Get("json")
		if jsonTag == "-" {
			continue
		}

		// Use column tag > json tag > field name
		col := field.Tag.Get("column")
		if col == "" {
			if jsonTag != "" {
				col = CamelToSnake(jsonTag)
			} else {
				col = CamelToSnake(field.Name)
			}
		}

		// Recurse into nested structs (excluding time.Time)
		if val.Kind() == reflect.Struct && field.Type != reflect.TypeOf(time.Time{}) {
			childClauses, childHasUpdated := s.extractUpdateFieldsStruct(val)
			setClauses = append(setClauses, childClauses...)
			if childHasUpdated {
				hasUpdatedAt = true
			}
			continue
		}

		// If value is a map, recurse into it
		if val.Kind() == reflect.Map {
			childClauses, childHasUpdated := s.extractUpdateFieldsMap(
				val.Interface().(map[string]any),
			)
			setClauses = append(setClauses, childClauses...)
			if childHasUpdated {
				hasUpdatedAt = true
			}
			continue
		}

		setClauses = append(setClauses, fmt.Sprintf(`"%s" = $%d`, col, len(s.Args)+1))
		s.Args = append(s.Args, val.Interface())

		if col == "updated_at" {
			hasUpdatedAt = true
		}
	}

	return setClauses, hasUpdatedAt
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
		// if value is a struct, recurse into it
		if v.Kind() == reflect.Struct && v.Type() != reflect.TypeOf(time.Time{}) {
			childClauses, childHasUpdated := s.extractUpdateFieldsStruct(v)
			setClauses = append(setClauses, childClauses...)
			if childHasUpdated {
				hasUpdatedAt = true
			}
			continue
		}

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
		setClauses = append(setClauses, fmt.Sprintf(`"%s" = $%d`, col, len(s.Args)+1))
		s.Args = append(s.Args, value)
	}

	return setClauses, hasUpdatedAt
}

func buildUpdateCase(updateCaseClauses map[string][]UpdateCaseParam, tableName string) string {
	var updateExpr string
	updateExpr = "UPDATE " + tableName + "\n"
	updateExpr += "SET\n"

	for key, each := range updateCaseClauses {
		updateExpr += key + " CASE\n"

		for _, param := range each {
			updateExpr += "WHEN " + strings.Join(param.conditions, " AND ") + " THEN" + param.value
			updateExpr += "\n"
		}
	}

	return updateExpr
}

func isZeroValue(v interface{}) bool {
	if v == nil {
		return true
	}

	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}
