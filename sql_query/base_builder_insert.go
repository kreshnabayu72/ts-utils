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
	// Insert implements SQLInsertChainBuilder. (Only able to be called once)
	// Insert builds an INSERT query from either a single struct or a slice of structs.
	// - For a single struct → generates one row.
	// - For a slice → generates multiple rows.
	//
	// By default, it will RETURN the id::text column,
	// but you can pass custom returning columns:
	//
	//	Insert(user, "id", "name")
	//	Insert([]User{u1, u2})
	Insert(values interface{}, returningColumns ...string) SQLInsertChainBuilder

	// insertSingle handles the insert logic for a single struct.
	// It auto-generates a Snowflake ID and builds the VALUES list
	// based on struct field tags (using "column" or "json").
	insertSingle(v reflect.Value) SQLInsertChainBuilder
	// insertMany handles inserting multiple structs at once.
	// It generates a Snowflake ID for each row, creates placeholders
	// for every value, and builds the full VALUES (...), (...), ... list.
	//
	// Example Result:
	//
	//	INSERT INTO users (id, name, email)
	//	VALUES ($1, $2, $3), ($4, $5, $6)
	insertMany(slice reflect.Value) SQLInsertChainBuilder
	// preBuild is a helper that assembles the final
	// INSERT INTO ... (columns) VALUES ... statement
	// after the columns and value placeholders have been prepared.
	preBuild(columns, valuePlaceholders []string)
}

// To ensure method .Insert() has its own chaining methods
// e.g. .Insert(...).Build()
type SQLInsertChainBuilder interface {
	// By default, the Insert builder includes all fields, even if their values are zero or nil.
	// Calling this method tells the builder to skip fields with zero values when generating the SET clause.
	//
	// Note: This option only affects single-row Insert operations.
	// It has no effect on bulk INSERT, because all rows in those operations must have the same set of columns.
	ExcludeEmpty() SQLInsertChainBuilder
	// Insert implements SQLInsertChainBuilder. (Only able to be called once, will override previous call)
	// Conflict adds an ON CONFLICT clause to the insert statement.
	// Example:
	//
	//	.Conflict("(id)", "NOTHING")
	//	-> INSERT ... ON CONFLICT (id) DO NOTHING
	Conflict(constraint, do string) SQLInsertChainBuilder
	// buildInsertQuery finalizes the insert query into SQL string + args.
	// It prevents unsafe cases (like adding filters, joins, or pagination)
	// and appends RETURNING and ON CONFLICT if defined.
	Build() (string, []interface{}, error)
}

type InsertBuilder struct {
	*SQLEloquentQuery
}

func (s *InsertBuilder) ExcludeEmpty() SQLInsertChainBuilder {
	s.excludeEmptyValue = true
	return s
}

func (s *InsertBuilder) Conflict(constraint, do string) SQLInsertChainBuilder {
	s.ConflictClause = fmt.Sprintf(" ON CONFLICT %s DO %s", constraint, do)
	return s
}

func (s *InsertBuilder) Insert(values interface{}, returningColumns ...string) SQLInsertChainBuilder {
	if len(returningColumns) > 0 {
		s.Columns = returningColumns
	} else {
		s.Columns = []string{"id"}
	}

	v := reflect.ValueOf(values)

	if v.Kind() != reflect.Slice && v.Kind() != reflect.Struct {
		s.LastError = errors.New("insert values must be struct or slice of struct")
		return s
	}

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
	return s.insertSingle(v)
}

// NewSQLInsertBuilder creates a new insert builder for a given table.
// Example:
//
//	builder := NewSQLInsertBuilder("users")
func NewSQLInsertBuilder(tableName string) SQLInsertInitBuilder {
	return &InsertBuilder{
		&SQLEloquentQuery{
			Table:       tableName,
			Columns:     []string{},
			CustomQuery: "",
			Args:        nil,
			Mode:        "insert",
		},
	}
}

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

	if s.ConflictClause != "" {
		s.CustomQuery += s.ConflictClause
	}

	return s.CustomQuery, s.Args, nil
}

func (s *InsertBuilder) insertSingle(v reflect.Value) SQLInsertChainBuilder {
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
		val := v.Field(i)

		if val.IsZero() && s.excludeEmptyValue {
			continue
		}

		jsonTag := field.Tag.Get("json")
		columnTag := field.Tag.Get("column")

		if ArrayIncludes([]string{"_id", "id"}, jsonTag) || columnTag == "id" {
			continue
		}

		if ArrayIncludes([]string{"", "-"}, jsonTag) && ArrayIncludes([]string{"", "-"}, columnTag) {
			continue
		}

		setTag := CamelToSnake(jsonTag)
		if columnTag != "" {
			setTag = columnTag
		}

		columns = append(columns, `"`+setTag+`"`)

		if ArrayIncludes([]string{"updated_at", "created_at"}, setTag) {
			placeholders = append(placeholders, "NOW()")
			continue
		}

		args = append(args, val.Interface())
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(args)))
	}

	s.Args = args
	s.preBuild(columns, []string{fmt.Sprintf("(%s)", strings.Join(placeholders, ", "))})

	return s
}

func (s *InsertBuilder) insertMany(slice reflect.Value) SQLInsertChainBuilder {
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
				jsonTag := field.Tag.Get("json")
				columnTag := field.Tag.Get("column")

				if ArrayIncludes([]string{"_id", "id"}, jsonTag) || columnTag == "id" {
					continue
				}

				if ArrayIncludes([]string{"", "-"}, jsonTag) && ArrayIncludes([]string{"", "-"}, columnTag) {
					continue
				}

				setTag := CamelToSnake(jsonTag)
				if columnTag != "" {
					setTag = columnTag
				}
				columns = append(columns, `"`+setTag+`"`)
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
			jsonTag := field.Tag.Get("json")
			columnTag := field.Tag.Get("column")
			setTag := CamelToSnake(jsonTag)
			if columnTag != "" {
				setTag = columnTag
			}

			if ArrayIncludes([]string{"_id", "id"}, jsonTag) || columnTag == "id" {
				continue
			}

			if ArrayIncludes([]string{"", "-"}, jsonTag) && ArrayIncludes([]string{"", "-"}, columnTag) {
				continue
			}

			if ArrayIncludes([]string{"updated_at", "created_at"}, setTag) {
				rowPlaceholders = append(rowPlaceholders, "NOW()")
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

func (s *InsertBuilder) preBuild(columns, valuePlaceholders []string) {
	s.CustomQuery = fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES %s",
		s.Table,
		strings.Join(columns, ", "),
		strings.Join(valuePlaceholders, ", "),
	)
}
