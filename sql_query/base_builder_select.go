package sql_query

import (
	"errors"
	"fmt"

	"strings"
)

// To ensure method .Select() has its own chaining methods
// e.g. .Select(...)....).Where(...).Build()
type SQLSelectChainBuilder interface {
	Select(columns ...string) SQLSelectChainBuilder

	Where(filters map[string]SQLCondition) SQLSelectChainBuilder
	Search(keyword string, fields []string) SQLSelectChainBuilder
	SetLimit(limit int) SQLSelectChainBuilder
	Join(table string, onCondition string) SQLSelectChainBuilder
	LeftJoin(table string, onCondition string) SQLSelectChainBuilder
	Paginate(page, limit int) SQLSelectChainBuilder
	OrderBy(sortBy []string, asc bool) SQLSelectChainBuilder

	Build() (string, []interface{}, error)
}

type selectBuilder struct {
	*SQLEloquentQuery
}

// Where implements SQLSelectChainBuilder. (Accumulates previous value if called again)
func (s *selectBuilder) Where(filters map[string]SQLCondition) SQLSelectChainBuilder {
	s.SQLEloquentQuery.sharedWhereQuery(filters)
	return s
}

// Search implements SQLSelectChainBuilder. (Accumulates previous value if called again)
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

func extractAlias(column string) string {
	parts := strings.Split(strings.ToLower(column), " as ")
	if len(parts) < 2 {
		return "" // No alias
	}
	return strings.TrimSpace(parts[len(parts)-1])
}

// Select implements SQLSelectChainBuilder. (Only able to be called once)
func (s *selectBuilder) Select(columns ...string) SQLSelectChainBuilder {

	for _, newCol := range columns {
		newAlias := extractAlias(newCol)

		// Check if alias exists in current list
		replaced := false
		for i, existing := range s.Columns {
			if extractAlias(existing) == newAlias {
				s.Columns[i] = newCol // Overwrite
				replaced = true
				break
			}
		}

		if !replaced {
			s.Columns = append(s.Columns, newCol)
		}
	}
	return s
}

// Join implements SQLSelectChainBuilder. (Accumulates previous value if called again)
func (s *selectBuilder) Join(table string, onCondition string) SQLSelectChainBuilder {
	if table == "" {
		return s
	}

	s.OtherTables = append(s.OtherTables, fmt.Sprintf("JOIN %s ON %s", table, onCondition))
	return s
}

// Join implements SQLSelectChainBuilder. (Accumulates previous value if called again)
func (s *selectBuilder) LeftJoin(table string, onCondition string) SQLSelectChainBuilder {
	if table == "" {
		return s
	}

	s.OtherTables = append(s.OtherTables, fmt.Sprintf("LEFT JOIN %s ON %s", table, onCondition))
	return s
}

// OrderBy implements SQLSelectChainBuilder. (Accumulates with previous value if called again)
func (s *selectBuilder) OrderBy(sortBy []string, asc bool) SQLSelectChainBuilder {
	direction := "ASC"

	if !asc {
		direction = "DESC"
	}

	sortingRule := fmt.Sprintf("%s %s", strings.Join(sortBy, ", "), direction)
	s.SortBy = append(s.SortBy, sortingRule)
	return s
}

// Paginate implements SQLSelectChainBuilder. (Overrides previous value if called again)
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

// SetLimit implements SQLSelectChainBuilder. (Overrides previous value if called again)
func (s *selectBuilder) SetLimit(limit int) SQLSelectChainBuilder {
	var normalizedPage int

	s.Limit = limit
	s.Offset = normalizedPage * limit

	return s
}

func NewSQLSelectBuilder[T any](tableName string) SQLSelectChainBuilder {
	columns := ExtractJSONTags[T]()

	defaultColumns := []string{"*"}
	if len(columns) > 0 {
		defaultColumns = columns
	}
	return &selectBuilder{
		&SQLEloquentQuery{
			Table:         tableName,
			Filters:       []string{},
			OtherTables:   []string{},
			Columns:       defaultColumns,
			Limit:         0,
			Offset:        0,
			SortBy:        []string{},
			CustomQuery:   "",
			Args:          nil,
			UsePagination: false,
			Mode:          "select",
		},
	}
}

// buildSelectQuery implements SQLBuilder.
func (s *SQLEloquentQuery) buildSelectQuery() (string, []interface{}, error) {
	if len(s.Columns) == 0 {
		return "", nil, errors.New("select query must have columns")
	}

	var dataQuery []string
	var countQuery []string

	dataQuery = append(dataQuery, fmt.Sprintf("SELECT %s FROM %s", strings.Join(s.Columns, ","), s.Table))
	countQuery = append(countQuery, fmt.Sprintf("SELECT COUNT(DISTINCT %s.id) FROM %s", s.Table, s.Table))

	if len(s.OtherTables) > 0 {
		joinClause := strings.Join(s.OtherTables, " ")
		dataQuery = append(dataQuery, joinClause)
		countQuery = append(countQuery, joinClause)
	}

	if len(s.Filters) > 0 {
		whereClause := fmt.Sprintf("WHERE %s", strings.Join(s.Filters, " AND "))
		dataQuery = append(dataQuery, whereClause)
		countQuery = append(countQuery, whereClause)
	}

	if len(s.SortBy) > 0 {
		sortingClause := fmt.Sprintf("ORDER BY %s", strings.Join(s.SortBy, ", "))
		dataQuery = append(dataQuery, sortingClause)
	}

	if s.UsePagination {
		if s.Limit > 0 {
			dataQuery = append(dataQuery, fmt.Sprintf("LIMIT %d", s.Limit))
		}
		dataQuery = append(dataQuery, fmt.Sprintf("OFFSET %d", s.Offset))

		return PaginationQuery(strings.Join(dataQuery, " "), strings.Join(countQuery, " ")), s.Args, nil
	}

	result := strings.Join(dataQuery, " ")
	return result, s.Args, nil
}
