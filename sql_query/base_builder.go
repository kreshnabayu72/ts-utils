package sql_query

import (
	"errors"
)

type ArrayAggConfig struct {
	Expr      string
	SortBy    string
	SortOrder int
}

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
	Operator      SQLOperators  // e.g., '=', '>', '<=', 'LIKE', 'IN', 'IS NULL'
	Value         interface{}   // could be a single value, slice, or nil.
	IsRef         bool          // to determine whether WHERE is targeting literal value or reference, e.g. `"column_a" = value` vs `"column_a" = $2` based on given boolean.
	SourceIsValue bool          // to determine whether WHERE is sourcing from literal value or reference, e.g. `"column_a" = value` vs `$2 = value` based on given boolean.
	ExtraArgs     []interface{} // for Operator `SQLOperatorRaw`
}

type UpdateCaseParam struct {
	conditions []string
	value      string
	isElse     bool
}

// type UpdateCaseParam struct {
// 	column         string
// 	UpdateCaseExpr []UpdateCaseExpr
// }

type MultiFilterCondition struct {
	And map[string]SQLCondition
	Or  []map[string]SQLCondition
}

// Base struct that contains params used by methods
type SQLEloquentQuery struct {
	NestedAggregation []string
	WrapAggregation   bool

	ConflictClause       string
	WithClauses          []string
	WithRecursiveClauses []string
	Table                string
	Filters              []string
	OtherTables          []string
	UnionAllQueries      []string
	Columns              []string
	DistinctBy           []string
	DistinctAlias        string
	Offset               int
	Limit                int
	SortBy               []string
	Grouping             []string
	HavingClauses        []string

	CustomQuery       string
	UpdateCaseClauses map[string][]UpdateCaseParam

	Args          []interface{}
	UsePagination bool
	Mode          SQLMode
	LastError     error

	currentUpdateCaseKey string
	useUnionAll          bool
	useHaving            bool
	excludeEmptyValue    bool
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
