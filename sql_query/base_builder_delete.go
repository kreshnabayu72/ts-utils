package sql_query

import (
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
	Using(table string) SQLDeleteChainBuilder

	Build() (string, []interface{}, error)
}

// <---To wrap builder to its respective interface, used as pointer type in method of each builder--->
type deleteBuilder struct {
	*SQLEloquentQuery
}

// Where implements SQLDeleteChainBuilder.
func (s *deleteBuilder) Where(filters map[string]SQLCondition) SQLDeleteChainBuilder {
	s.SQLEloquentQuery.sharedWhereQuery(filters)
	return s
}

// Delete implements SQLDeleteChainBuilder.
func (s *deleteBuilder) Delete(returningColumns ...string) SQLDeleteChainBuilder {
	if len(returningColumns) > 0 {
		s.Columns = strings.Join(returningColumns, ",") // Columns for RETURNING (used in method Build)
	} else {
		s.Columns = "id"
	}

	s.CustomQuery = fmt.Sprintf(
		"DELETE FROM %s",
		s.Table,
	)
	return s
}

// Using implements SQLDeleteChainBuilder.
func (s *deleteBuilder) Using(table string) SQLDeleteChainBuilder {
	if table == "" {
		return s
	}

	s.OtherTables = append(s.OtherTables, fmt.Sprintf("USING %s", table))
	return s
}

func NewSQLDeleteBuilder(tableName string) SQLDeleteInitBuilder {
	return &deleteBuilder{
		&SQLEloquentQuery{
			Table:       tableName,
			Filters:     "",
			OtherTables: []string{},
			Columns:     "",
			CustomQuery: "",
			Args:        nil,
			Mode:        "delete",
		},
	}
}
