package sql_query

import (
	"errors"
	"fmt"
	"strings"
)

// NewSQLDeleteBuilder creates a new delete builder for a given table.
//
// Example:
//
//	builder := NewSQLDeleteBuilder("users","u")
func NewSQLDeleteBuilder(tableName string, alias ...string) DeleteStarter {
	if len(alias) > 0 {
		tableName = fmt.Sprintf("%s %s", tableName, strings.TrimSpace(alias[0]))
	}

	return &DeleteBuilder{
		&SQLEloquentQuery{
			Table:       tableName,
			Filters:     []string{},
			OtherTables: []string{},
			Columns:     []string{},
			CustomQuery: "",
			Args:        nil,
			Mode:        "delete",
		},
	}
}

func (s *DeleteBuilder) Where(filters map[string]SQLCondition) Deleter {
	s.SQLEloquentQuery.sharedWhereAndQuery(filters)
	return s
}

func (s *DeleteBuilder) WhereOr(filters ...map[string]SQLCondition) Deleter {
	s.SQLEloquentQuery.sharedWhereOrQuery(filters...)
	return s
}

func (s *DeleteBuilder) Delete(returningColumns ...string) Deleter {
	if len(returningColumns) > 0 {
		s.Columns = returningColumns
	} else {
		s.Columns = []string{"id"}
	}

	s.CustomQuery = fmt.Sprintf(
		"DELETE FROM %s",
		s.Table,
	)
	return s
}

func (s *DeleteBuilder) Using(tables []string) Deleter {
	if len(tables) < 1 {
		return s
	}

	var otherTables []string
	otherTables = append(otherTables, fmt.Sprintf("USING %s", strings.Join(tables, ", ")))
	s.OtherTables = otherTables
	return s
}

func (s *SQLEloquentQuery) buildDeleteQuery() (string, []any, error) {
	if s.LastError != nil {
		return "", nil, errors.New(s.LastError.Error())
	}

	if s.CustomQuery == "" {
		return "", nil, errors.New("invalid update query: CustomQuery not set")
	}

	query := s.CustomQuery

	if len(s.OtherTables) > 0 {
		usingClause := strings.Join(s.OtherTables, " ")
		query += " " + usingClause
	}

	if len(s.Filters) > 0 {
		query += " " + fmt.Sprintf("WHERE %s", strings.Join(s.Filters, " AND "))
	}

	if len(s.Columns) > 0 {
		query += " RETURNING " + strings.Join(s.Columns, ",")
	}

	return query, s.Args, nil
}
