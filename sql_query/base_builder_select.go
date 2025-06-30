package sql_query

import (
	"fmt"

	"strings"
)

// To ensure SQLSelectBuilder has its own initial methods
// e.g. SelectBuilder(...).Select()...Rest
type SQLSelectInitBuilder interface {
	Select(columns ...string) SQLSelectChainBuilder
}

// To ensure method .Select() has its own chaining methods
// e.g. .Select(...).Join(...).Where(...).Build()
type SQLSelectChainBuilder interface {
	Where(filters map[string]SQLCondition) SQLSelectChainBuilder
	Search(keyword string, fields []string) SQLSelectChainBuilder
	SetLimit(limit int) SQLSelectChainBuilder
	Join(table string, onCondition string) SQLSelectChainBuilder
	Paginate(page, limit int) SQLSelectChainBuilder
	OrderBy(sortBy string, asc bool) SQLSelectChainBuilder

	Build() (string, []interface{}, error)
}

type selectBuilder struct {
	*SQLEloquentQuery
}

// Where implements SQLSelectChainBuilder.
func (s *selectBuilder) Where(filters map[string]SQLCondition) SQLSelectChainBuilder {
	s.SQLEloquentQuery.sharedWhereQuery(filters)
	return s
}

// Search implementation
func (s *selectBuilder) Search(keyword string, fields []string) SQLSelectChainBuilder {
	if keyword != "" && len(fields) > 0 {
		searchFilters := map[string]SQLCondition{}

		for _, field := range fields {
			searchFilters[field] = SQLCondition{
				Operator: "ILIKE",
				Value:    "%" + keyword + "%",
			}
		}

		// Use sharedWhereQuery internally to handle argIndex, Args, etc.
		s.Where(searchFilters)
	}

	return s
}

// Select implements SQLSelectChainBuilder.
func (s *selectBuilder) Select(columns ...string) SQLSelectChainBuilder {
	if len(columns) < 1 {
		s.Columns = "*"
		return s
	}

	s.Columns = strings.Join(columns, ",")
	return s
}

// Join implements SQLSelectChainBuilder.
func (s *selectBuilder) Join(table string, onCondition string) SQLSelectChainBuilder {
	if table == "" {
		return s
	}

	s.OtherTables = append(s.OtherTables, fmt.Sprintf("JOIN %s ON %s", table, onCondition))
	return s
}

// OrderBy implements SQLSelectChainBuilder.
func (s *selectBuilder) OrderBy(sortBy string, asc bool) SQLSelectChainBuilder {
	panic("unimplemented")
}

// Paginate implements SQLSelectChainBuilder.
func (s *selectBuilder) Paginate(page int, limit int) SQLSelectChainBuilder {
	var normalizedPage int
	if page > 0 {
		normalizedPage = page - 1
	} else {
		normalizedPage = 0
	}

	s.UsePagination = true
	s.Limit = limit
	s.Offset = normalizedPage * limit

	return s
}
func (s *selectBuilder) SetLimit(limit int) SQLSelectChainBuilder {
	var normalizedPage int

	s.Limit = limit
	s.Offset = normalizedPage * limit

	return s
}

func NewSQLSelectBuilder(tableName string) SQLSelectInitBuilder {
	return &selectBuilder{
		&SQLEloquentQuery{
			Table:         tableName,
			Filters:       "",
			OtherTables:   []string{},
			Columns:       "",
			Limit:         0,
			Offset:        0,
			SortBy:        "created_at",
			CustomQuery:   "",
			Args:          nil,
			UsePagination: false,
			Mode:          "select",
		},
	}
}
