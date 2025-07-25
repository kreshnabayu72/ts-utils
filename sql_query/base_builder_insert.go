package sql_query

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
)

// To ensure SQLInsertBuilder has its own initial methods
// e.g. InsertBuilder(...).Insert()...Rest
type SQLInsertInitBuilder interface {
	Insert(values interface{}, returningColumns ...string) SQLInsertChainBuilder

	insertSingle(v reflect.Value) SQLInsertChainBuilder
	insertMany(slice reflect.Value) SQLInsertChainBuilder
	preBuild(columns, valuePlaceholders []string)
}

// To ensure method .Insert() has its own chaining methods
// e.g. .Insert(...).Build()
type SQLInsertChainBuilder interface {
	Build() (string, []interface{}, error)
}

type insertBuilder struct {
	*SQLEloquentQuery
}

// Insert implements SQLInsertChainBuilder. (Only able to be called once)
func (s *insertBuilder) Insert(values interface{}, returningColumns ...string) SQLInsertChainBuilder {
	if len(returningColumns) > 0 {
		s.Columns = returningColumns
	} else {
		s.Columns = []string{"id"}
	}

	v := reflect.ValueOf(values)

	// Slice case
	if v.Kind() == reflect.Slice {
		if v.Len() == 0 {
			s.LastError = errors.New("cannot insert with empty slice")
			return s
		}

		firstElem := v.Index(0)
		if firstElem.Kind() != reflect.Struct {
			s.LastError = errors.New("insert slice must contain structs")
			return s
		}

		return s.insertMany(v)
	}

	// Single struct case
	if v.Kind() == reflect.Struct {
		return s.insertSingle(v)
	}

	s.LastError = errors.New("insert values must be struct or slice of struct")
	return s
}

func NewSQLInsertBuilder(tableName string) SQLInsertInitBuilder {
	return &insertBuilder{
		&SQLEloquentQuery{
			Table:       tableName,
			Columns:     []string{},
			CustomQuery: "",
			Args:        nil,
			Mode:        "insert",
		},
	}
}

// buildInsertQuery implements SQLBuilder.
func (s *SQLEloquentQuery) buildInsertQuery() (string, []interface{}, error) {
	if s.LastError != nil {
		return "", nil, errors.New(s.LastError.Error())
	}

	if len(s.Filters) > 0 || s.UsePagination || len(s.OtherTables) > 0 {
		return "", nil, errors.New("invalid insert query: cannot include filters, joins, or pagination")
	}

	if len(s.Columns) > 0 {
		s.CustomQuery += " RETURNING " + strings.Join(s.Columns, ",")
	}

	return s.CustomQuery, s.Args, nil
}

func (s *insertBuilder) insertSingle(v reflect.Value) SQLInsertChainBuilder {
	var args []interface{}
	var columns []string
	var placeholders []string

	db.InitSnowflake() // prevents error in unit test (correct this if better method found) (Maybe mock snowflake idk)
	id := db.Node.Generate().Int64()
	columns = append(columns, "id")
	placeholders = append(placeholders, "$1")
	args = append(args, id)

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if ArrayIncludes([]string{"", "-", "_id", "id"}, tag) {
			continue
		}

		columns = append(columns, `"`+CamelToSnake(tag)+`"`)
		args = append(args, v.Field(i).Interface())
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(args)))
	}

	s.Args = args
	s.preBuild(columns, placeholders)

	return s
}

func (s *insertBuilder) insertMany(slice reflect.Value) SQLInsertChainBuilder {
	var args []interface{}
	var columns []string
	var valuePlaceholders []string

	db.InitSnowflake() // prevents error in unit test (correct this if better method found) (Maybe mock snowflake idk)

	for i := 0; i < slice.Len(); i++ {
		v := slice.Index(i)
		t := v.Type()

		if i == 0 {
			// Setup column list once
			columns = append(columns, "id")
			for j := 0; j < t.NumField(); j++ {
				field := t.Field(j)
				tag := field.Tag.Get("json")
				if ArrayIncludes([]string{"", "-", "_id", "id"}, tag) {
					continue
				}
				columns = append(columns, `"`+CamelToSnake(tag)+`"`)
			}
		}

		var rowPlaceholders []string

		// Add ID
		id := db.Node.Generate().Int64()
		args = append(args, id)
		startIndex := len(args)
		rowPlaceholders = append(rowPlaceholders, fmt.Sprintf("$%d", startIndex))

		// Add fields
		for j := 0; j < t.NumField(); j++ {
			field := t.Field(j)
			tag := field.Tag.Get("json")
			if ArrayIncludes([]string{"", "-", "_id", "id"}, tag) {
				continue
			}
			args = append(args, v.Field(j).Interface())
			rowPlaceholders = append(rowPlaceholders, fmt.Sprintf("$%d", len(args)))
		}

		valuePlaceholders = append(valuePlaceholders, fmt.Sprintf("(%s)", strings.Join(rowPlaceholders, ", ")))
	}

	s.Args = args
	s.preBuild(columns, valuePlaceholders)

	return s
}

func (s *insertBuilder) preBuild(columns, valuePlaceholders []string) {
	s.CustomQuery = fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		s.Table,
		strings.Join(columns, ", "),
		strings.Join(valuePlaceholders, ", "),
	)
}
