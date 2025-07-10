package sql_query

import (
	"errors"
)

type SQLMode string

const (
	SQLDelete SQLMode = "delete"
	SQLInsert SQLMode = "insert"
	SQLSelect SQLMode = "select"
	SQLUpdate SQLMode = "update"
)

type QueryBuilder interface {
	buildSelectQuery() (string, []interface{}, error)
	buildInsertQuery() (string, []interface{}, error)
	buildUpdateQuery() (string, []interface{}, error)
	buildDeleteQuery() (string, []interface{}, error)
}

type SQLFilter map[string]SQLCondition

type SQLCondition struct {
	Operator  SQLOperators // e.g., '=', '>', '<=', 'LIKE', 'IN', 'IS NULL'
	Value     interface{}  // could be a single value, slice, or nil
	IsRef     bool         // to determine whether WHERE is targeting literal value or reference
	ExtraArgs []interface{}
}

// Base struct that contains params used by methods
type SQLEloquentQuery struct {
	AggregationAlias   string
	AggregationFields  string
	AggregationFilters []string

	Table         string
	Filters       []string
	OtherTables   []string
	Columns       []string
	Offset        int
	Limit         int
	SortBy        []string
	Grouping      []string
	HavingClauses []string

	CustomQuery string

	Args          []interface{}
	UsePagination bool
	Mode          SQLMode
	LastError     error

	useAggregation bool
	useHaving      bool
}

// Run respective build method based on given mode
func (s *SQLEloquentQuery) Build() (string, []interface{}, error) {
	switch s.Mode {
	case SQLDelete:
		return s.buildDeleteQuery()
	case SQLInsert:
		return s.buildInsertQuery()
	case SQLUpdate:
		return s.buildUpdateQuery()
	case SQLSelect:
		return s.buildSelectQuery()
	default:
		return "", nil, errors.New("unsupported query mode")
	}
}
