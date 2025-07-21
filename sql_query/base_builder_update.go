package sql_query

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// To ensure SQLUpdateBuilder has its own initial methods
// e.g. UpdateBuilder(...).Update()...Rest
type SQLUpdateInitBuilder interface {
	Update(values interface{}, returningColumns ...string) SQLUpdateChainBuilder

	Increment(values map[string]any, returningColumns ...string) SQLUpdateChainBuilder
}

// To ensure method .Update() has its own chaining methods
// e.g. .Update(...).From(...).Build()
type SQLUpdateChainBuilder interface {
	Where(filters map[string]SQLCondition) SQLUpdateChainBuilder
	WhereOr(filters map[string]SQLCondition) SQLUpdateChainBuilder

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

// Update implements SQLUpdateChainBuilder using reflection like Insert(). (Only able to be called once)
func (s *updateBuilder) Update(
	values interface{},
	returningColumns ...string,
) SQLUpdateChainBuilder {
	v := reflect.ValueOf(values)
	if v.Kind() != reflect.Struct {
		s.LastError = fmt.Errorf("Update data must be a struct")
		return s
	}

	t := v.Type()
	var hasUpdatedAt bool
	var setClauses []string

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}

		snake := CamelToSnake(tag)
		if snake == "updated_at" {
			hasUpdatedAt = true
		}

		setClauses = append(setClauses, fmt.Sprintf(`"%s" = $%d`, snake, len(s.Args)+1))
		s.Args = append(s.Args, v.Field(i).Interface())
	}

	if len(returningColumns) > 0 {
		s.Columns = returningColumns // Columns for RETURNING (used in method Build)
	} else {
		s.Columns = []string{"id"}
	}

	if !hasUpdatedAt {
		setClauses = append(setClauses, fmt.Sprintf("updated_at = $%d", len(s.Args)+1))
		s.Args = append(s.Args, time.Now())
	}
	s.CustomQuery = fmt.Sprintf("UPDATE %s SET %s", s.Table, strings.Join(setClauses, ", "))

	return s
}

// Increment is used to replace Update() for adding int value cause update cant handle that
func (s *updateBuilder) Increment(
	values map[string]any,
	returningColumns ...string,
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

	if len(returningColumns) > 0 {
		s.Columns = returningColumns
	} else {
		s.Columns = []string{"id"}
	}

	setClauses = append(setClauses, fmt.Sprintf(`"updated_at" = $%d`, len(s.Args)+1))
	s.Args = append(s.Args, time.Now())

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
			Columns:     []string{"*"},
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
	}
	return query, s.Args, nil
}
