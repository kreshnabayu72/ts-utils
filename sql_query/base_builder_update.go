package sql_query

import (
	"fmt"
	"reflect"
	"strings"
)

// To ensure SQLUpdateBuilder has its own initial methods
// e.g. UpdateBuilder(...).Update()...Rest
type SQLUpdateInitBuilder interface {
	Update(values interface{}, returningColumns ...string) SQLUpdateChainBuilder
}

// To ensure method .Update() has its own chaining methods
// e.g. .Update(...).From(...).Build()
type SQLUpdateChainBuilder interface {
	Where(filters map[string]SQLCondition) SQLUpdateChainBuilder
	From(table string) SQLUpdateChainBuilder

	Build() (string, []interface{}, error)
}

type updateBuilder struct {
	*SQLEloquentQuery
}

// Where implements SQLUpdateChainBuilder.
func (s *updateBuilder) Where(filters map[string]SQLCondition) SQLUpdateChainBuilder {
	s.SQLEloquentQuery.sharedWhereQuery(filters)
	return s
}

// Update implements SQLUpdateChainBuilder using reflection like Insert().
func (s *updateBuilder) Update(values interface{}, returningColumns ...string) SQLUpdateChainBuilder {
	v := reflect.ValueOf(values)
	if v.Kind() != reflect.Struct {
		s.LastError = fmt.Errorf("Update data must be a struct")
		return s
	}

	t := v.Type()
	var setClauses []string

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", tag, len(s.Args)+1))
		s.Args = append(s.Args, v.Field(i).Interface())
	}

	if len(returningColumns) > 0 {
		s.Columns = strings.Join(returningColumns, ",")
	} else {
		s.Columns = "id"
	}

	s.CustomQuery = fmt.Sprintf("UPDATE %s SET %s", s.Table, strings.Join(setClauses, ", "))

	return s
}

// From implements SQLUpdateChainBuilder.
func (s *updateBuilder) From(table string) SQLUpdateChainBuilder {
	if table == "" {
		return s
	}

	s.OtherTables = append(s.OtherTables, fmt.Sprintf("FROM %s", table))
	return s
}

func NewSQLUpdateBuilder(tableName string) SQLUpdateInitBuilder {
	return &updateBuilder{
		&SQLEloquentQuery{
			Table:       tableName,
			Filters:     "",
			OtherTables: []string{},
			Columns:     "",
			CustomQuery: "UPDATE",
			Args:        nil,
			Mode:        "update",
		},
	}
}
