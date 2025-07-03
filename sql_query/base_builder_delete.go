package sql_query

import (
	"errors"
	"fmt"

	"strings"
)

// To ensure SQLDeleteBuilder has its own initial methods
// e.g. DeleteBuilder(...).Delete()...Rest
type SQLDeleteInitBuilder interface {
	Delete(returningColumns ...string) SQLDeleteChainBuilder
}

// To ensure method .Delete() has its own chaining methods
// e.g. .Delete(...).Using(...).Where(...).Build()
type SQLDeleteChainBuilder interface {
	Where(filters map[string]SQLCondition) SQLDeleteChainBuilder
	Using(tables []string) SQLDeleteChainBuilder

	Build() (string, []interface{}, error)
}

// <---To wrap builder to its respective interface, used as pointer type in method of each builder--->
type deleteBuilder struct {
	*SQLEloquentQuery
}

// Where implements SQLDeleteChainBuilder. (Accumulates previous value if called again)
func (s *deleteBuilder) Where(filters map[string]SQLCondition) SQLDeleteChainBuilder {
	s.SQLEloquentQuery.sharedWhereQuery(filters, false)
	return s
}

// Delete implements SQLDeleteChainBuilder. (Only able to be called once)
func (s *deleteBuilder) Delete(returningColumns ...string) SQLDeleteChainBuilder {
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

// Using implements SQLDeleteChainBuilder. (Overrides previous value if called again)
func (s *deleteBuilder) Using(tables []string) SQLDeleteChainBuilder {
	if len(tables) < 1 {
		return s
	}

	var otherTables []string
	otherTables = append(otherTables, fmt.Sprintf("USING %s", strings.Join(tables, ", ")))
	s.OtherTables = otherTables
	return s
}

func NewSQLDeleteBuilder(tableName string) SQLDeleteInitBuilder {
	return &deleteBuilder{
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

// buildDeleteQuery implements SQLBuilder.
func (s *SQLEloquentQuery) buildDeleteQuery() (string, []interface{}, error) {
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
		s.CustomQuery += " RETURNING " + strings.Join(s.Columns, ",")
	}

	return query, s.Args, nil
}
