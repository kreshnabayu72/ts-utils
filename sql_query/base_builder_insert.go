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
	var args []interface{}

	v := reflect.ValueOf(values)
	if v.Kind() != reflect.Struct {
		s.LastError = fmt.Errorf("Insert data must be a struct")
	}

	t := v.Type()
	var columns []string
	var placeholders []string

	db.InitSnowflake() //prevents error in unit test (correct this if better method found) (Maybe mock snowflake idk)

	id := db.Node.Generate().Int64()
	columns = append(columns, "id")
	placeholders = append(placeholders, "$1")
	args = append(args, id)

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" || tag == "_id" || tag == "id" {
			continue
		}

		columns = append(columns, `"`+CamelToSnake(tag)+`"`)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(args)+1))
		args = append(args, v.Field(i).Interface())
	}

	if len(returningColumns) > 0 {
		s.Columns = returningColumns // Columns for RETURNING (used in method Build)
	} else {
		s.Columns = []string{"id"}
	}

	s.Args = args

	s.CustomQuery = fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		s.Table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)
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
