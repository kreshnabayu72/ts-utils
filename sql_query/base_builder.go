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

// Base struct that contains params used by methods
type SQLEloquentQuery struct {
	Table       string
	Filters     []string
	OtherTables []string
	Columns     []string
	Offset      int
	Limit       int
	SortBy      []string

	CustomQuery string

	Args          []interface{}
	argIndex      int
	WhereParts    []string
	UsePagination bool
	Mode          SQLMode
	LastError     error
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
