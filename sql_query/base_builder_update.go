package sql_query

import (
	"errors"
	"fmt"
	"reflect"
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

// To ensure SQLUpdateBuilder has its own initial methods
// e.g. UpdateBuilder(...).Update()...Rest
type SQLUpdateInitBuilder interface {
	Update(values interface{}) SQLUpdateChainBuilder
	UpdateEach(values interface{}, rowIdentifier string) SQLUpdateChainBuilder

	Increment(values map[string]any) SQLUpdateChainBuilder
}

// To ensure method .Update() has its own chaining methods
// e.g. .Update(...).From(...).Build()
type SQLUpdateChainBuilder interface {
	Where(filters map[string]SQLCondition) SQLUpdateChainBuilder
	WhereOr(filters map[string]SQLCondition) SQLUpdateChainBuilder
	Return(columns ...string) SQLUpdateChainBuilder

	From(tables []string) SQLUpdateChainBuilder

	Build() (string, []interface{}, error)
}

type updateBuilder struct {
	*SQLEloquentQuery
}

// Where implements SQLUpdateChainBuilder. (Accumulates previous value if called again)
func (s *updateBuilder) Where(filters map[string]SQLCondition) SQLUpdateChainBuilder {
	s.SQLEloquentQuery.sharedWhereAndQuery(filters)
	return s
}

// WhereOr implements SQLUpdateChainBuilder. (Accumulates previous value if called again)
func (s *updateBuilder) WhereOr(filters map[string]SQLCondition) SQLUpdateChainBuilder {
	s.SQLEloquentQuery.sharedWhereOrQuery(filters)
	return s
}

// Return implements SQLUpdateChainBuilder. (Only able to be called once, overrides previous values if re-called)
func (s *updateBuilder) Return(column ...string) SQLUpdateChainBuilder {
	if len(column) > 0 {
		s.Columns = column
	} else {
		s.Columns = []string{"id"}
	}
	return s
}

// Update implements SQLUpdateChainBuilder using reflection like Insert(). (Only able to be called once)
func (s *updateBuilder) Update(values interface{}) SQLUpdateChainBuilder {
	v := reflect.ValueOf(values)
	if v.Kind() != reflect.Struct {
		s.LastError = fmt.Errorf("update one data must be a struct")
		return s
	}

	setClauses, hasUpdatedAt := s.extractUpdateFields(v)

	if !hasUpdatedAt {
		setClauses = append(setClauses, `"updated_at" = NOW()`)
	}

	s.CustomQuery = fmt.Sprintf(`UPDATE %s SET %s`, s.Table, strings.Join(setClauses, ", "))

	return s
}

// UpdateEach implements SQLUpdateChainBuilder using reflection like Insert(). (Only able to be called once)
// Example of expected result:
// UPDATE alias_codes
// SET is_checked = v.is_checked, value = v.value, updated_at = v.updated_at
// FROM (VALUES ($1, $2, $3), ($4, $5, $6), ...more row) as v(name, is_checked, value, updated_at)
// WHERE name = v.name
func (s *updateBuilder) UpdateEach(values interface{}, rowIdentifier string) SQLUpdateChainBuilder {
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
	s.Filters = append(s.Filters, fmt.Sprintf(`%s."%s" = v."%s"`, s.Table, rowIdentifier, rowIdentifier))

	// Generate clauses for everything in slice, then stored in main array
	var setClauses, valueClauses, placeholders []string
	mappedSet, mappedValue, mappedPlaceholders := s.updateEachClausesGenerator(v, rowIdentifier)
	setClauses = append(setClauses, mappedSet...)
	valueClauses = append(valueClauses, mappedValue...)
	placeholders = append(placeholders, mappedPlaceholders...)

	// Build everything into one query dedicated for update many
	s.CustomQuery = fmt.Sprintf(`UPDATE %s SET %s FROM (VALUES %s) as v(%s)`, s.Table, strings.Join(setClauses, ", "), strings.Join(placeholders, ", "), strings.Join(valueClauses, ","))

	return s
}

// Increment is used to replace Update() for adding int value cause update cant handle that
func (s *updateBuilder) Increment(
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

// From implements SQLUpdateChainBuilder. (Overrides previous value if called again)
func (s *updateBuilder) From(tables []string) SQLUpdateChainBuilder {
	if len(tables) < 1 {
		return s
	}

	var otherTables []string
	otherTables = append(otherTables, fmt.Sprintf("FROM %s", strings.Join(tables, ", ")))
	s.OtherTables = otherTables
	return s
}

func NewSQLUpdateBuilder(tableName string) SQLUpdateInitBuilder {
	return &updateBuilder{
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

// buildUpdateQuery implements SQLBuilder.
func (s *SQLEloquentQuery) buildUpdateQuery() (string, []interface{}, error) {
	if s.CustomQuery == "" {
		return "", nil, errors.New("invalid update query: CustomQuery not set")
	}

	query := s.CustomQuery

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
		query += " RETURNING id"
	}

	return query, s.Args, nil
}

// This function loop through slice, returning generated clauses for everything in it
// Expected result:
// setClauses contains: [value = v.value, ...etc].
// valueClauses contains: [value, ...etc].
// valuePlaceholders contains: [($1, $2, $3, ...etc), ($4, $5, ...etc), ...etc]
func (s *updateBuilder) updateEachClausesGenerator(slice reflect.Value, rowIdentifier string) ([]string, []string, []string) {
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
					valueClauses = append(valueClauses, columnTag)
				} else {
					setClauses = append(setClauses, fmt.Sprintf(`"%s" = v."%s"`, columnTag, columnTag))
					valueClauses = append(valueClauses, columnTag)
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
			rowPlaceholders = append(rowPlaceholders, fmt.Sprintf("$%d::%s", len(args), transformTag))
		}

		// Append every placeholders to main array that will be returned
		valuePlaceholders = append(valuePlaceholders, fmt.Sprintf("(%s)", strings.Join(rowPlaceholders, ", ")))
	}

	s.Args = args
	return setClauses, valueClauses, valuePlaceholders
}

func (s *updateBuilder) extractUpdateFields(v reflect.Value) ([]string, bool) {
	var setClauses []string
	var hasUpdatedAt bool

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		val := v.Field(i)

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
			childClauses, childHasUpdated := s.extractUpdateFields(val)
			setClauses = append(setClauses, childClauses...)
			if childHasUpdated {
				hasUpdatedAt = true
			}
			continue
		}

		setClauses = append(setClauses, fmt.Sprintf(`%s = $%d`, col, len(s.Args)+1))
		s.Args = append(s.Args, val.Interface())

		if col == "updated_at" {
			hasUpdatedAt = true
		}
	}

	return setClauses, hasUpdatedAt
}
