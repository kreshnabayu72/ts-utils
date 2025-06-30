package sql_query

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type SQLFilter map[string]SQLCondition

type SQLCondition struct {
	Operator SQLOperators // e.g., '=', '>', '<=', 'LIKE', 'IN', 'IS NULL'
	Value    interface{}  // could be a single value, slice, or nil
	IsRef    bool         // to determine whether WHERE is targeting literal value or reference
}

type SQLMode string

const (
	SQLDelete SQLMode = "delete"
	SQLInsert SQLMode = "insert"
	SQLSelect SQLMode = "select"
	SQLUpdate SQLMode = "update"
)

type SQLBuilder interface {
	buildSelectQuery() (string, []interface{}, error)
	buildInsertQuery() (string, []interface{}, error)
	buildUpdateQuery() (string, []interface{}, error)
	buildDeleteQuery() (string, []interface{}, error)
}

// To make method Where has generic type (usable for chaining method in each builder)
type Whereable[T any] interface {
	Where(filters map[string]SQLCondition) T
}

// To ensure method .Delete() has its own chaining methods
// e.g. DeleteBuilder(...).Using(...).Where(...).Build()
type SQLDeleteBuilder interface {
	Whereable[SQLDeleteBuilder]
	Delete(returningColumns ...string) SQLDeleteBuilder
	Using(table string) SQLDeleteBuilder

	Build() (string, []interface{}, error)
}

// To ensure method .Insert() has its own chaining methods
// e.g. InsertBuilder(...).Build()
type SQLInsertBuilder interface {
	Insert(values interface{}, returningColumns ...string) SQLInsertBuilder
	Build() (string, []interface{}, error)
}

// To ensure method .Select() has its own chaining methods
// e.g. SelectBuilder(...).Join(...).Where(...).Build()
type SQLSelectBuilder interface {
	Whereable[SQLSelectBuilder]
	Select(columns ...string) SQLSelectBuilder
	Join(table string, onCondition string) SQLSelectBuilder
	Paginate(page, limit int64) SQLSelectBuilder
	OrderBy(sortBy string, asc bool) SQLSelectBuilder

	Build() (string, []interface{}, error)
}

// To ensure method .Update() has its own chaining methods
// e.g. UpdateBuilder(...).From(...).Build()
type SQLUpdateBuilder interface {
	Whereable[SQLUpdateBuilder]
	Update(values map[string]interface{}, returningColumns ...string) SQLUpdateBuilder
	From(table string) SQLUpdateBuilder

	Build() (string, []interface{}, error)
}

// <---To wrap builder to its respective interface, used as pointer type in method of each builder--->
type deleteBuilder struct {
	*SQLEloquentQuery
}

type insertBuilder struct {
	*SQLEloquentQuery
}

type selectBuilder struct {
	*SQLEloquentQuery
}

type updateBuilder struct {
	*SQLEloquentQuery
}

// Base struct that contains params used by methods
type SQLEloquentQuery struct {
	Table       string
	Filters     string
	OtherTables []string
	Columns     string
	Offset      int64
	Limit       int64
	SortBy      string

	CustomQuery string

	Args          []interface{}
	UsePagination bool
	Mode          SQLMode
	LastError     error
}

// Where implements SQLDeleteBuilder.
func (s *deleteBuilder) Where(filters map[string]SQLCondition) SQLDeleteBuilder {
	s.SQLEloquentQuery.sharedWhereQuery(filters)
	return s
}

// Where implements SQLSelectBuilder.
func (s *selectBuilder) Where(filters map[string]SQLCondition) SQLSelectBuilder {
	s.SQLEloquentQuery.sharedWhereQuery(filters)
	return s
}

// Where implements SQLUpdateBuilder.
func (s *updateBuilder) Where(filters map[string]SQLCondition) SQLUpdateBuilder {
	s.SQLEloquentQuery.sharedWhereQuery(filters)
	return s
}

// Delete implements SQLDeleteBuilder.
func (s *deleteBuilder) Delete(returningColumns ...string) SQLDeleteBuilder {
	if s.Mode != "" && s.Mode != "delete" {
		panic("Can't add Delete after Base Method other than Delete")
	}
	s.Mode = "delete"
	s.Args = nil

	s.Columns = strings.Join(returningColumns, ",") // Columns for RETURNING (used in method Build)

	s.CustomQuery = fmt.Sprintf(
		"DELETE FROM %s",
		s.Table,
	)
	return s
}

// Insert implements SQLInsertBuilder.
func (s *insertBuilder) Insert(values interface{}, returningColumns ...string) SQLInsertBuilder {
	var args []interface{}

	if s.Mode != "" && s.Mode != "insert" {
		s.LastError = fmt.Errorf("Can't add Insert after Base Method other than Insert")
	}
	s.Mode = "insert"
	s.Args = nil

	v := reflect.ValueOf(values)
	if v.Kind() != reflect.Struct {
		s.LastError = fmt.Errorf("Insert data must be a struct")
	}

	t := v.Type()
	var columns []string
	var placeholders []string

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		columns = append(columns, tag)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(args)+1))
		args = append(args, v.Field(i).Interface())
	}

	s.Args = args
	s.Columns = strings.Join(returningColumns, ",") // Columns for RETURNING (used in method Build)

	s.CustomQuery = fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		s.Table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)
	return s
}

// Update implements SQLUpdateBuilder.
func (s *updateBuilder) Update(values map[string]interface{}, returningColumns ...string) SQLUpdateBuilder {
	if s.Mode != "" && s.Mode != "update" {
		panic("Can't add Update after Base Method other than Update")
	}
	s.Mode = "update"
	s.Args = nil

	var setClauses []string
	i := 1
	for key, value := range values {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", key, i))
		s.Args = append(s.Args, value)
		i++
	}

	s.Columns = strings.Join(returningColumns, ",") // Columns for RETURNING (used in method Build)
	s.CustomQuery = fmt.Sprintf("UPDATE %s SET %s", s.Table, strings.Join(setClauses, ", "))

	return s
}

// Select implements SQLSelectBuilder.
func (s *selectBuilder) Select(columns ...string) SQLSelectBuilder {
	if s.Mode != "" && s.Mode != "select" {
		panic("Can't add Select after Base Method other than Select")
	}
	s.Mode = "select"
	s.Args = nil

	if len(columns) < 1 {
		s.Columns = "*"
		return s
	}

	s.Columns = strings.Join(columns, ",")
	return s
}

// Using implements SQLDeleteBuilder.
func (s *deleteBuilder) Using(table string) SQLDeleteBuilder {
	if table == "" {
		return s
	}

	s.OtherTables = append(s.OtherTables, fmt.Sprintf("USING %s", table))
	return s
}

// From implements SQLUpdateBuilder.
func (s *updateBuilder) From(table string) SQLUpdateBuilder {
	if table == "" {
		return s
	}

	s.OtherTables = append(s.OtherTables, fmt.Sprintf("FROM %s", table))
	return s
}

// Join implements SQLSelectBuilder.
func (s *selectBuilder) Join(table string, onCondition string) SQLSelectBuilder {
	if table == "" {
		return s
	}

	s.OtherTables = append(s.OtherTables, fmt.Sprintf("JOIN %s ON %s", table, onCondition))
	return s
}

// OrderBy implements SQLSelectBuilder.
func (s *selectBuilder) OrderBy(sortBy string, asc bool) SQLSelectBuilder {
	panic("unimplemented")
}

// Paginate implements SQLSelectBuilder.
func (s *selectBuilder) Paginate(page int64, limit int64) SQLSelectBuilder {
	var normalizedPage int64
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

func NewSQLDeleteBuilder(tableName string) SQLDeleteBuilder {
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

func NewSQLInsertBuilder(tableName string) SQLInsertBuilder {
	return &insertBuilder{
		&SQLEloquentQuery{
			Table:       tableName,
			Columns:     "",
			CustomQuery: "",
			Args:        nil,
			Mode:        "insert",
		},
	}
}

func NewSQLSelectBuilder(tableName string) SQLSelectBuilder {
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

func NewSQLUpdateBuilder(tableName string) SQLUpdateBuilder {
	return &updateBuilder{
		&SQLEloquentQuery{
			Table:       tableName,
			Filters:     "",
			OtherTables: []string{},
			Columns:     "",
			CustomQuery: "",
			Args:        nil,
			Mode:        "update",
		},
	}
}
